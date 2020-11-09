// +build !ignore_autogenerated

/*


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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredCluster) DeepCopyInto(out *DiscoveredCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredCluster.
func (in *DiscoveredCluster) DeepCopy() *DiscoveredCluster {
	if in == nil {
		return nil
	}
	out := new(DiscoveredCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterList) DeepCopyInto(out *DiscoveredClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveredCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterList.
func (in *DiscoveredClusterList) DeepCopy() *DiscoveredClusterList {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterRefresh) DeepCopyInto(out *DiscoveredClusterRefresh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterRefresh.
func (in *DiscoveredClusterRefresh) DeepCopy() *DiscoveredClusterRefresh {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterRefresh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredClusterRefresh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterRefreshList) DeepCopyInto(out *DiscoveredClusterRefreshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveredClusterRefresh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterRefreshList.
func (in *DiscoveredClusterRefreshList) DeepCopy() *DiscoveredClusterRefreshList {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterRefreshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredClusterRefreshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterRefreshSpec) DeepCopyInto(out *DiscoveredClusterRefreshSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterRefreshSpec.
func (in *DiscoveredClusterRefreshSpec) DeepCopy() *DiscoveredClusterRefreshSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterRefreshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterRefreshStatus) DeepCopyInto(out *DiscoveredClusterRefreshStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterRefreshStatus.
func (in *DiscoveredClusterRefreshStatus) DeepCopy() *DiscoveredClusterRefreshStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterRefreshStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterSpec) DeepCopyInto(out *DiscoveredClusterSpec) {
	*out = *in
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ActivityTimestamp != nil {
		in, out := &in.ActivityTimestamp, &out.ActivityTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ProviderConnections != nil {
		in, out := &in.ProviderConnections, &out.ProviderConnections
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.Subscription = in.Subscription
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterSpec.
func (in *DiscoveredClusterSpec) DeepCopy() *DiscoveredClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredClusterStatus) DeepCopyInto(out *DiscoveredClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredClusterStatus.
func (in *DiscoveredClusterStatus) DeepCopy() *DiscoveredClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveredClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfig) DeepCopyInto(out *DiscoveryConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfig.
func (in *DiscoveryConfig) DeepCopy() *DiscoveryConfig {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfigList) DeepCopyInto(out *DiscoveryConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveryConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfigList.
func (in *DiscoveryConfigList) DeepCopy() *DiscoveryConfigList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfigSpec) DeepCopyInto(out *DiscoveryConfigSpec) {
	*out = *in
	if in.ProviderConnections != nil {
		in, out := &in.ProviderConnections, &out.ProviderConnections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Filters = in.Filters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfigSpec.
func (in *DiscoveryConfigSpec) DeepCopy() *DiscoveryConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfigStatus) DeepCopyInto(out *DiscoveryConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfigStatus.
func (in *DiscoveryConfigStatus) DeepCopy() *DiscoveryConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}
