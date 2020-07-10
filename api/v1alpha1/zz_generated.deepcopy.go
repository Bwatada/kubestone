// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BenchmarkStatus) DeepCopyInto(out *BenchmarkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BenchmarkStatus.
func (in *BenchmarkStatus) DeepCopy() *BenchmarkStatus {
	if in == nil {
		return nil
	}
	out := new(BenchmarkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Drill) DeepCopyInto(out *Drill) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Drill.
func (in *Drill) DeepCopy() *Drill {
	if in == nil {
		return nil
	}
	out := new(Drill)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Drill) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrillList) DeepCopyInto(out *DrillList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Drill, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrillList.
func (in *DrillList) DeepCopy() *DrillList {
	if in == nil {
		return nil
	}
	out := new(DrillList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DrillList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrillSpec) DeepCopyInto(out *DrillSpec) {
	*out = *in
	out.Image = in.Image
	if in.BenchmarksVolume != nil {
		in, out := &in.BenchmarksVolume, &out.BenchmarksVolume
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrillSpec.
func (in *DrillSpec) DeepCopy() *DrillSpec {
	if in == nil {
		return nil
	}
	out := new(DrillSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fio) DeepCopyInto(out *Fio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fio.
func (in *Fio) DeepCopy() *Fio {
	if in == nil {
		return nil
	}
	out := new(Fio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FioList) DeepCopyInto(out *FioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FioList.
func (in *FioList) DeepCopy() *FioList {
	if in == nil {
		return nil
	}
	out := new(FioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FioSpec) DeepCopyInto(out *FioSpec) {
	*out = *in
	out.Image = in.Image
	if in.BuiltinJobFiles != nil {
		in, out := &in.BuiltinJobFiles, &out.BuiltinJobFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomJobFiles != nil {
		in, out := &in.CustomJobFiles, &out.CustomJobFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PodConfig.DeepCopyInto(&out.PodConfig)
	in.Volume.DeepCopyInto(&out.Volume)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FioSpec.
func (in *FioSpec) DeepCopy() *FioSpec {
	if in == nil {
		return nil
	}
	out := new(FioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ioping) DeepCopyInto(out *Ioping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ioping.
func (in *Ioping) DeepCopy() *Ioping {
	if in == nil {
		return nil
	}
	out := new(Ioping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ioping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IopingList) DeepCopyInto(out *IopingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ioping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IopingList.
func (in *IopingList) DeepCopy() *IopingList {
	if in == nil {
		return nil
	}
	out := new(IopingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IopingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IopingSpec) DeepCopyInto(out *IopingSpec) {
	*out = *in
	out.Image = in.Image
	in.PodConfig.DeepCopyInto(&out.PodConfig)
	in.Volume.DeepCopyInto(&out.Volume)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IopingSpec.
func (in *IopingSpec) DeepCopy() *IopingSpec {
	if in == nil {
		return nil
	}
	out := new(IopingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf2) DeepCopyInto(out *Iperf2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf2.
func (in *Iperf2) DeepCopy() *Iperf2 {
	if in == nil {
		return nil
	}
	out := new(Iperf2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf2ConfigurationSpec) DeepCopyInto(out *Iperf2ConfigurationSpec) {
	*out = *in
	in.PodConfigurationSpec.DeepCopyInto(&out.PodConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf2ConfigurationSpec.
func (in *Iperf2ConfigurationSpec) DeepCopy() *Iperf2ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(Iperf2ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf2List) DeepCopyInto(out *Iperf2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Iperf2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf2List.
func (in *Iperf2List) DeepCopy() *Iperf2List {
	if in == nil {
		return nil
	}
	out := new(Iperf2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf2Spec) DeepCopyInto(out *Iperf2Spec) {
	*out = *in
	out.Image = in.Image
	in.ServerConfiguration.DeepCopyInto(&out.ServerConfiguration)
	in.ClientConfiguration.DeepCopyInto(&out.ClientConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf2Spec.
func (in *Iperf2Spec) DeepCopy() *Iperf2Spec {
	if in == nil {
		return nil
	}
	out := new(Iperf2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3) DeepCopyInto(out *Iperf3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3.
func (in *Iperf3) DeepCopy() *Iperf3 {
	if in == nil {
		return nil
	}
	out := new(Iperf3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3ConfigurationSpec) DeepCopyInto(out *Iperf3ConfigurationSpec) {
	*out = *in
	in.PodConfigurationSpec.DeepCopyInto(&out.PodConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3ConfigurationSpec.
func (in *Iperf3ConfigurationSpec) DeepCopy() *Iperf3ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(Iperf3ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3List) DeepCopyInto(out *Iperf3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Iperf3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3List.
func (in *Iperf3List) DeepCopy() *Iperf3List {
	if in == nil {
		return nil
	}
	out := new(Iperf3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3Spec) DeepCopyInto(out *Iperf3Spec) {
	*out = *in
	out.Image = in.Image
	in.ServerConfiguration.DeepCopyInto(&out.ServerConfiguration)
	in.ClientConfiguration.DeepCopyInto(&out.ClientConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3Spec.
func (in *Iperf3Spec) DeepCopy() *Iperf3Spec {
	if in == nil {
		return nil
	}
	out := new(Iperf3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaBench) DeepCopyInto(out *KafkaBench) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaBench.
func (in *KafkaBench) DeepCopy() *KafkaBench {
	if in == nil {
		return nil
	}
	out := new(KafkaBench)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaBench) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaBenchList) DeepCopyInto(out *KafkaBenchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaBench, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaBenchList.
func (in *KafkaBenchList) DeepCopy() *KafkaBenchList {
	if in == nil {
		return nil
	}
	out := new(KafkaBenchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaBenchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaBenchSpec) DeepCopyInto(out *KafkaBenchSpec) {
	*out = *in
	out.Image = in.Image
	in.PodConfig.DeepCopyInto(&out.PodConfig)
	in.KafkaClusterInfo.DeepCopyInto(&out.KafkaClusterInfo)
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]KafkaTestSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaBenchSpec.
func (in *KafkaBenchSpec) DeepCopy() *KafkaBenchSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaBenchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterInfo) DeepCopyInto(out *KafkaClusterInfo) {
	*out = *in
	if in.ZooKeepers != nil {
		in, out := &in.ZooKeepers, &out.ZooKeepers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Brokers != nil {
		in, out := &in.Brokers, &out.Brokers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterInfo.
func (in *KafkaClusterInfo) DeepCopy() *KafkaClusterInfo {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTestSpec) DeepCopyInto(out *KafkaTestSpec) {
	*out = *in
	if in.ConsumerSleep != nil {
		in, out := &in.ConsumerSleep, &out.ConsumerSleep
		*out = new(int32)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ExtraProducerOpts != nil {
		in, out := &in.ExtraProducerOpts, &out.ExtraProducerOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTestSpec.
func (in *KafkaTestSpec) DeepCopy() *KafkaTestSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MixedDistributionOptions) DeepCopyInto(out *MixedDistributionOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MixedDistributionOptions.
func (in *MixedDistributionOptions) DeepCopy() *MixedDistributionOptions {
	if in == nil {
		return nil
	}
	out := new(MixedDistributionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OcpLogtest) DeepCopyInto(out *OcpLogtest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OcpLogtest.
func (in *OcpLogtest) DeepCopy() *OcpLogtest {
	if in == nil {
		return nil
	}
	out := new(OcpLogtest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OcpLogtest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OcpLogtestList) DeepCopyInto(out *OcpLogtestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OcpLogtest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OcpLogtestList.
func (in *OcpLogtestList) DeepCopy() *OcpLogtestList {
	if in == nil {
		return nil
	}
	out := new(OcpLogtestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OcpLogtestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OcpLogtestSpec) DeepCopyInto(out *OcpLogtestSpec) {
	*out = *in
	out.Image = in.Image
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OcpLogtestSpec.
func (in *OcpLogtestSpec) DeepCopy() *OcpLogtestSpec {
	if in == nil {
		return nil
	}
	out := new(OcpLogtestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgbench) DeepCopyInto(out *Pgbench) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgbench.
func (in *Pgbench) DeepCopy() *Pgbench {
	if in == nil {
		return nil
	}
	out := new(Pgbench)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgbench) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbenchList) DeepCopyInto(out *PgbenchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgbench, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbenchList.
func (in *PgbenchList) DeepCopy() *PgbenchList {
	if in == nil {
		return nil
	}
	out := new(PgbenchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgbenchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbenchSpec) DeepCopyInto(out *PgbenchSpec) {
	*out = *in
	out.Image = in.Image
	out.Postgres = in.Postgres
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbenchSpec.
func (in *PgbenchSpec) DeepCopy() *PgbenchSpec {
	if in == nil {
		return nil
	}
	out := new(PgbenchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfigurationSpec) DeepCopyInto(out *PodConfigurationSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodScheduling.DeepCopyInto(&out.PodScheduling)
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfigurationSpec.
func (in *PodConfigurationSpec) DeepCopy() *PodConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(PodConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingSpec) DeepCopyInto(out *PodSchedulingSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingSpec.
func (in *PodSchedulingSpec) DeepCopy() *PodSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSpec) DeepCopyInto(out *PostgresSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSpec.
func (in *PostgresSpec) DeepCopy() *PostgresSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Qperf) DeepCopyInto(out *Qperf) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Qperf.
func (in *Qperf) DeepCopy() *Qperf {
	if in == nil {
		return nil
	}
	out := new(Qperf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Qperf) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QperfConfigurationSpec) DeepCopyInto(out *QperfConfigurationSpec) {
	*out = *in
	in.PodConfigurationSpec.DeepCopyInto(&out.PodConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QperfConfigurationSpec.
func (in *QperfConfigurationSpec) DeepCopy() *QperfConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(QperfConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QperfList) DeepCopyInto(out *QperfList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Qperf, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QperfList.
func (in *QperfList) DeepCopy() *QperfList {
	if in == nil {
		return nil
	}
	out := new(QperfList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QperfList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QperfSpec) DeepCopyInto(out *QperfSpec) {
	*out = *in
	out.Image = in.Image
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ServerConfiguration.DeepCopyInto(&out.ServerConfiguration)
	in.ClientConfiguration.DeepCopyInto(&out.ClientConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QperfSpec.
func (in *QperfSpec) DeepCopy() *QperfSpec {
	if in == nil {
		return nil
	}
	out := new(QperfSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AnalysisOptions) DeepCopyInto(out *S3AnalysisOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AnalysisOptions.
func (in *S3AnalysisOptions) DeepCopy() *S3AnalysisOptions {
	if in == nil {
		return nil
	}
	out := new(S3AnalysisOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AutoTermOptions) DeepCopyInto(out *S3AutoTermOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AutoTermOptions.
func (in *S3AutoTermOptions) DeepCopy() *S3AutoTermOptions {
	if in == nil {
		return nil
	}
	out := new(S3AutoTermOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bench) DeepCopyInto(out *S3Bench) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bench.
func (in *S3Bench) DeepCopy() *S3Bench {
	if in == nil {
		return nil
	}
	out := new(S3Bench)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Bench) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BenchList) DeepCopyInto(out *S3BenchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Bench, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BenchList.
func (in *S3BenchList) DeepCopy() *S3BenchList {
	if in == nil {
		return nil
	}
	out := new(S3BenchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BenchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BenchOptions) DeepCopyInto(out *S3BenchOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BenchOptions.
func (in *S3BenchOptions) DeepCopy() *S3BenchOptions {
	if in == nil {
		return nil
	}
	out := new(S3BenchOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BenchSpec) DeepCopyInto(out *S3BenchSpec) {
	*out = *in
	out.Image = in.Image
	in.PodConfig.DeepCopyInto(&out.PodConfig)
	out.S3BenchOptions = in.S3BenchOptions
	out.S3ObjectOptions = in.S3ObjectOptions
	out.S3AutoTermOptions = in.S3AutoTermOptions
	out.S3AnalysisOptions = in.S3AnalysisOptions
	out.MixedDistributionOptions = in.MixedDistributionOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BenchSpec.
func (in *S3BenchSpec) DeepCopy() *S3BenchSpec {
	if in == nil {
		return nil
	}
	out := new(S3BenchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectOptions) DeepCopyInto(out *S3ObjectOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectOptions.
func (in *S3ObjectOptions) DeepCopy() *S3ObjectOptions {
	if in == nil {
		return nil
	}
	out := new(S3ObjectOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sysbench) DeepCopyInto(out *Sysbench) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sysbench.
func (in *Sysbench) DeepCopy() *Sysbench {
	if in == nil {
		return nil
	}
	out := new(Sysbench)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sysbench) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SysbenchList) DeepCopyInto(out *SysbenchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sysbench, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SysbenchList.
func (in *SysbenchList) DeepCopy() *SysbenchList {
	if in == nil {
		return nil
	}
	out := new(SysbenchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SysbenchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SysbenchSpec) DeepCopyInto(out *SysbenchSpec) {
	*out = *in
	out.Image = in.Image
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SysbenchSpec.
func (in *SysbenchSpec) DeepCopy() *SysbenchSpec {
	if in == nil {
		return nil
	}
	out := new(SysbenchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YcsbBench) DeepCopyInto(out *YcsbBench) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YcsbBench.
func (in *YcsbBench) DeepCopy() *YcsbBench {
	if in == nil {
		return nil
	}
	out := new(YcsbBench)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YcsbBench) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YcsbBenchList) DeepCopyInto(out *YcsbBenchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]YcsbBench, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YcsbBenchList.
func (in *YcsbBenchList) DeepCopy() *YcsbBenchList {
	if in == nil {
		return nil
	}
	out := new(YcsbBenchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YcsbBenchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YcsbBenchOptions) DeepCopyInto(out *YcsbBenchOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YcsbBenchOptions.
func (in *YcsbBenchOptions) DeepCopy() *YcsbBenchOptions {
	if in == nil {
		return nil
	}
	out := new(YcsbBenchOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YcsbBenchSpec) DeepCopyInto(out *YcsbBenchSpec) {
	*out = *in
	out.Image = in.Image
	out.Options = in.Options
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YcsbBenchSpec.
func (in *YcsbBenchSpec) DeepCopy() *YcsbBenchSpec {
	if in == nil {
		return nil
	}
	out := new(YcsbBenchSpec)
	in.DeepCopyInto(out)
	return out
}
