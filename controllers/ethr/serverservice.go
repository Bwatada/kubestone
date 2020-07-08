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

package ethr

import (
	perfv1alpha1 "github.com/xridge/kubestone/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;create;delete;watch

func serverServiceName(cr *perfv1alpha1.Ethr) string {
	return cr.Name
}

// NewServerService creates k8s headless service (which targets the server deployment)
// from the Ethr Benchmark Definition
func NewServerService(cr *perfv1alpha1.Ethr) *corev1.Service {
	labels := map[string]string{
		"kubestone.xridge.io/app":     "ethr",
		"kubestone.xridge.io/cr-name": cr.Name,
	}
	protocol := corev1.Protocol(corev1.ProtocolTCP)
	if cr.Spec.UDP {
		protocol = corev1.Protocol(corev1.ProtocolUDP)
	}
	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        serverServiceName(cr),
			Namespace:   cr.Namespace,
			Annotations: cr.Spec.ClientConfiguration.PodConfigurationSpec.Annotations,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:          "ethr-server-tb",
					Port: perfv1alpha1.TCPBandwidthPort,
					Protocol:      protocol,
				},
				{
					Name:          "ethr-server-tc",
					Port: perfv1alpha1.TCPConnectionsPort,
					Protocol:      protocol,
				},
				{
					Name:          "ethr-server-tl",
					Port: perfv1alpha1.TCPLatencyPort,
					Protocol:      protocol,
				},
				{
					Name:          "ethr-server-hb",
					Port: perfv1alpha1.HTTPBandwidthPort,
					Protocol:      protocol,
				},
				{
					Name:          "ethr-server-hsb",
					Port: perfv1alpha1.HTTPSBandwidthPort,
					Protocol:      protocol,
				},
				{
					Name:          "ethr-server-hl",
					Port: perfv1alpha1.HTTPLatencyPort,
					Protocol:      protocol,
				},
			},
			Selector:  labels,
			ClusterIP: "None", // Headless service!
		},
	}

	return &service
}
