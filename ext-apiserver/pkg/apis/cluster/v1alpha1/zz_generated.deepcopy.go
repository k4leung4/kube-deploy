// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	common "k8s.io/kube-deploy/ext-apiserver/pkg/apis/cluster/common"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*APIEndpoint).DeepCopyInto(out.(*APIEndpoint))
			return nil
		}, InType: reflect.TypeOf(&APIEndpoint{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Cluster).DeepCopyInto(out.(*Cluster))
			return nil
		}, InType: reflect.TypeOf(&Cluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterList).DeepCopyInto(out.(*ClusterList))
			return nil
		}, InType: reflect.TypeOf(&ClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterNetworkingConfig).DeepCopyInto(out.(*ClusterNetworkingConfig))
			return nil
		}, InType: reflect.TypeOf(&ClusterNetworkingConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterSpec).DeepCopyInto(out.(*ClusterSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterStatus).DeepCopyInto(out.(*ClusterStatus))
			return nil
		}, InType: reflect.TypeOf(&ClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ContainerRuntimeInfo).DeepCopyInto(out.(*ContainerRuntimeInfo))
			return nil
		}, InType: reflect.TypeOf(&ContainerRuntimeInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Machine).DeepCopyInto(out.(*Machine))
			return nil
		}, InType: reflect.TypeOf(&Machine{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineList).DeepCopyInto(out.(*MachineList))
			return nil
		}, InType: reflect.TypeOf(&MachineList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineSet).DeepCopyInto(out.(*MachineSet))
			return nil
		}, InType: reflect.TypeOf(&MachineSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineSetList).DeepCopyInto(out.(*MachineSetList))
			return nil
		}, InType: reflect.TypeOf(&MachineSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineSetSpec).DeepCopyInto(out.(*MachineSetSpec))
			return nil
		}, InType: reflect.TypeOf(&MachineSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineSetStatus).DeepCopyInto(out.(*MachineSetStatus))
			return nil
		}, InType: reflect.TypeOf(&MachineSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineSpec).DeepCopyInto(out.(*MachineSpec))
			return nil
		}, InType: reflect.TypeOf(&MachineSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineStatus).DeepCopyInto(out.(*MachineStatus))
			return nil
		}, InType: reflect.TypeOf(&MachineStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineTemplateSpec).DeepCopyInto(out.(*MachineTemplateSpec))
			return nil
		}, InType: reflect.TypeOf(&MachineTemplateSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MachineVersionInfo).DeepCopyInto(out.(*MachineVersionInfo))
			return nil
		}, InType: reflect.TypeOf(&MachineVersionInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NetworkRanges).DeepCopyInto(out.(*NetworkRanges))
			return nil
		}, InType: reflect.TypeOf(&NetworkRanges{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkingConfig) DeepCopyInto(out *ClusterNetworkingConfig) {
	*out = *in
	in.Services.DeepCopyInto(&out.Services)
	in.Pods.DeepCopyInto(&out.Pods)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkingConfig.
func (in *ClusterNetworkingConfig) DeepCopy() *ClusterNetworkingConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ClusterNetwork.DeepCopyInto(&out.ClusterNetwork)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make([]APIEndpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeInfo) DeepCopyInto(out *ContainerRuntimeInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeInfo.
func (in *ContainerRuntimeInfo) DeepCopy() *ContainerRuntimeInfo {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSet) DeepCopyInto(out *MachineSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSet.
func (in *MachineSet) DeepCopy() *MachineSet {
	if in == nil {
		return nil
	}
	out := new(MachineSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetList) DeepCopyInto(out *MachineSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetList.
func (in *MachineSetList) DeepCopy() *MachineSetList {
	if in == nil {
		return nil
	}
	out := new(MachineSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetSpec) DeepCopyInto(out *MachineSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetSpec.
func (in *MachineSetSpec) DeepCopy() *MachineSetSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetStatus) DeepCopyInto(out *MachineSetStatus) {
	*out = *in
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		if *in == nil {
			*out = nil
		} else {
			*out = new(common.MachineSetStatusError)
			**out = **in
		}
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetStatus.
func (in *MachineSetStatus) DeepCopy() *MachineSetStatus {
	if in == nil {
		return nil
	}
	out := new(MachineSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]common.MachineRole, len(*in))
		copy(*out, *in)
	}
	out.Versions = in.Versions
	if in.ConfigSource != nil {
		in, out := &in.ConfigSource, &out.ConfigSource
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.NodeConfigSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	if in.NodeRef != nil {
		in, out := &in.NodeRef, &out.NodeRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		if *in == nil {
			*out = nil
		} else {
			*out = new(MachineVersionInfo)
			**out = **in
		}
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		if *in == nil {
			*out = nil
		} else {
			*out = new(common.MachineStatusError)
			**out = **in
		}
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineTemplateSpec) DeepCopyInto(out *MachineTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineTemplateSpec.
func (in *MachineTemplateSpec) DeepCopy() *MachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineVersionInfo) DeepCopyInto(out *MachineVersionInfo) {
	*out = *in
	out.ContainerRuntime = in.ContainerRuntime
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineVersionInfo.
func (in *MachineVersionInfo) DeepCopy() *MachineVersionInfo {
	if in == nil {
		return nil
	}
	out := new(MachineVersionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRanges) DeepCopyInto(out *NetworkRanges) {
	*out = *in
	if in.CIDRBlocks != nil {
		in, out := &in.CIDRBlocks, &out.CIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRanges.
func (in *NetworkRanges) DeepCopy() *NetworkRanges {
	if in == nil {
		return nil
	}
	out := new(NetworkRanges)
	in.DeepCopyInto(out)
	return out
}
