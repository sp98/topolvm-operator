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

package v2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassState) DeepCopyInto(out *ClassState) {
	*out = *in
	if in.DeviceStates != nil {
		in, out := &in.DeviceStates, &out.DeviceStates
		*out = make([]DeviceState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassState.
func (in *ClassState) DeepCopy() *ClassState {
	if in == nil {
		return nil
	}
	out := new(ClassState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClass) DeepCopyInto(out *DeviceClass) {
	*out = *in
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = make([]Disk, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClass.
func (in *DeviceClass) DeepCopy() *DeviceClass {
	if in == nil {
		return nil
	}
	out := new(DeviceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceState) DeepCopyInto(out *DeviceState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceState.
func (in *DeviceState) DeepCopy() *DeviceState {
	if in == nil {
		return nil
	}
	out := new(DeviceState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopState) DeepCopyInto(out *LoopState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopState.
func (in *LoopState) DeepCopy() *LoopState {
	if in == nil {
		return nil
	}
	out := new(LoopState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDevices) DeepCopyInto(out *NodeDevices) {
	*out = *in
	if in.DeviceClasses != nil {
		in, out := &in.DeviceClasses, &out.DeviceClasses
		*out = make([]DeviceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDevices.
func (in *NodeDevices) DeepCopy() *NodeDevices {
	if in == nil {
		return nil
	}
	out := new(NodeDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStorageState) DeepCopyInto(out *NodeStorageState) {
	*out = *in
	if in.FailClasses != nil {
		in, out := &in.FailClasses, &out.FailClasses
		*out = make([]ClassState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuccessClasses != nil {
		in, out := &in.SuccessClasses, &out.SuccessClasses
		*out = make([]ClassState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Loops != nil {
		in, out := &in.Loops, &out.Loops
		*out = make([]LoopState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStorageState.
func (in *NodeStorageState) DeepCopy() *NodeStorageState {
	if in == nil {
		return nil
	}
	out := new(NodeStorageState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.DeviceClasses != nil {
		in, out := &in.DeviceClasses, &out.DeviceClasses
		*out = make([]NodeDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]Disk, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopolvmCluster) DeepCopyInto(out *TopolvmCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopolvmCluster.
func (in *TopolvmCluster) DeepCopy() *TopolvmCluster {
	if in == nil {
		return nil
	}
	out := new(TopolvmCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopolvmCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopolvmClusterList) DeepCopyInto(out *TopolvmClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TopolvmCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopolvmClusterList.
func (in *TopolvmClusterList) DeepCopy() *TopolvmClusterList {
	if in == nil {
		return nil
	}
	out := new(TopolvmClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopolvmClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopolvmClusterSpec) DeepCopyInto(out *TopolvmClusterSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopolvmClusterSpec.
func (in *TopolvmClusterSpec) DeepCopy() *TopolvmClusterSpec {
	if in == nil {
		return nil
	}
	out := new(TopolvmClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopolvmClusterStatus) DeepCopyInto(out *TopolvmClusterStatus) {
	*out = *in
	if in.NodeStorageStatus != nil {
		in, out := &in.NodeStorageStatus, &out.NodeStorageStatus
		*out = make([]NodeStorageState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopolvmClusterStatus.
func (in *TopolvmClusterStatus) DeepCopy() *TopolvmClusterStatus {
	if in == nil {
		return nil
	}
	out := new(TopolvmClusterStatus)
	in.DeepCopyInto(out)
	return out
}
