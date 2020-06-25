/*
Copyright 2019 The xridge kubestone contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package iperf2

import (
	"fmt"
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	perfv1alpha1 "github.com/xridge/kubestone/api/v1alpha1"

	"github.com/firepear/qsplit"
)

// Iperf2ServerPort is the TCP or UDP port where
// the iperf2 server deployment and service listens
const Iperf2ServerPort = 5201

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;create;delete;watch

func serverDeploymentName(cr *perfv1alpha1.Iperf2) string {
	return cr.Name
}

// NewServerDeployment create a iperf2 server deployment from the
// provided Iperf2 Benchmark Definition.
func NewServerDeployment(cr *perfv1alpha1.Iperf2) *appsv1.Deployment {
	replicas := int32(1)

	labels := map[string]string{
		"kubestone.xridge.io/app":     "iperf2",
		"kubestone.xridge.io/cr-name": cr.Name,
	}
	// Let's be nice and don't mutate CRs label field
	for k, v := range cr.Spec.ServerConfiguration.PodLabels {
		labels[k] = v
	}

	iperfCmdLineArgs := []string{
		"--server",
		"--port", strconv.Itoa(Iperf2ServerPort)}

	protocol := corev1.Protocol(corev1.ProtocolTCP)
	if cr.Spec.UDP {
		iperfCmdLineArgs = append(iperfCmdLineArgs, "--udp")
		protocol = corev1.Protocol(corev1.ProtocolUDP)
	}

	iperfCmdLineArgs = append(iperfCmdLineArgs,
		qsplit.ToStrings([]byte(cr.Spec.ServerConfiguration.CmdLineArgs))...)

	// Iperf2 Server does not like if probe connections are made to the port,
	// therefore we are checking if the port if open or not via shell script
	// the solution does not assume to have netstat installed in the container
	readinessAwkCmd := fmt.Sprintf("BEGIN{err=1}toupper($2)~/:%04X$/{err=0}END{exit err}", Iperf2ServerPort)

	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        serverDeploymentName(cr),
			Namespace:   cr.Namespace,
			Annotations: cr.Spec.ServerConfiguration.PodConfigurationSpec.Annotations,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Replicas: &replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: cr.Spec.ServerConfiguration.PodConfigurationSpec.Annotations,
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: cr.Spec.Image.PullSecret,
						},
					},
					Containers: []corev1.Container{
						{
							Name:            "server",
							Image:           cr.Spec.Image.Name,
							ImagePullPolicy: corev1.PullPolicy(cr.Spec.Image.PullPolicy),
							Command:         []string{"iperf2"},
							Args:            iperfCmdLineArgs,
							Ports: []corev1.ContainerPort{
								{
									Name:          "iperf-server",
									ContainerPort: Iperf2ServerPort,
									Protocol:      protocol,
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									Exec: &corev1.ExecAction{
										Command: []string{
											"awk",
											readinessAwkCmd,
											"/proc/1/net/tcp",
											"/proc/1/net/tcp6",
										},
									},
								},
								InitialDelaySeconds: 5,
								TimeoutSeconds:      2,
								PeriodSeconds:       2,
							},
							Resources: cr.Spec.ServerConfiguration.Resources,
						},
					},
					Affinity:     cr.Spec.ServerConfiguration.PodScheduling.Affinity,
					Tolerations:  cr.Spec.ServerConfiguration.PodScheduling.Tolerations,
					NodeSelector: cr.Spec.ServerConfiguration.PodScheduling.NodeSelector,
					NodeName:     cr.Spec.ServerConfiguration.PodScheduling.NodeName,
					HostNetwork:  cr.Spec.ServerConfiguration.HostNetwork,
				},
			},
		},
	}

	return &deployment
}
