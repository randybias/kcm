//go:build !ignore_autogenerated

// Copyright 2024
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/fluxcd/helm-controller/api/v2"
	"k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRule) DeepCopyInto(out *AccessRule) {
	*out = *in
	in.TargetNamespaces.DeepCopyInto(&out.TargetNamespaces)
	if in.ClusterTemplateChains != nil {
		in, out := &in.ClusterTemplateChains, &out.ClusterTemplateChains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceTemplateChains != nil {
		in, out := &in.ServiceTemplateChains, &out.ServiceTemplateChains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRule.
func (in *AccessRule) DeepCopy() *AccessRule {
	if in == nil {
		return nil
	}
	out := new(AccessRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailableUpgrade) DeepCopyInto(out *AvailableUpgrade) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailableUpgrade.
func (in *AvailableUpgrade) DeepCopy() *AvailableUpgrade {
	if in == nil {
		return nil
	}
	out := new(AvailableUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplate) DeepCopyInto(out *ClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplate.
func (in *ClusterTemplate) DeepCopy() *ClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateChain) DeepCopyInto(out *ClusterTemplateChain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateChain.
func (in *ClusterTemplateChain) DeepCopy() *ClusterTemplateChain {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateChain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplateChain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateChainList) DeepCopyInto(out *ClusterTemplateChainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTemplateChain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateChainList.
func (in *ClusterTemplateChainList) DeepCopy() *ClusterTemplateChainList {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateChainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplateChainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateList) DeepCopyInto(out *ClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateList.
func (in *ClusterTemplateList) DeepCopy() *ClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateSpec) DeepCopyInto(out *ClusterTemplateSpec) {
	*out = *in
	in.TemplateSpecCommon.DeepCopyInto(&out.TemplateSpecCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateSpec.
func (in *ClusterTemplateSpec) DeepCopy() *ClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateStatus) DeepCopyInto(out *ClusterTemplateStatus) {
	*out = *in
	in.TemplateStatusCommon.DeepCopyInto(&out.TemplateStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateStatus.
func (in *ClusterTemplateStatus) DeepCopy() *ClusterTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Core) DeepCopyInto(out *Core) {
	*out = *in
	in.HMC.DeepCopyInto(&out.HMC)
	in.CAPI.DeepCopyInto(&out.CAPI)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Core.
func (in *Core) DeepCopy() *Core {
	if in == nil {
		return nil
	}
	out := new(Core)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreProviderTemplate) DeepCopyInto(out *CoreProviderTemplate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreProviderTemplate.
func (in *CoreProviderTemplate) DeepCopy() *CoreProviderTemplate {
	if in == nil {
		return nil
	}
	out := new(CoreProviderTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Credential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialList) DeepCopyInto(out *CredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Credential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialList.
func (in *CredentialList) DeepCopy() *CredentialList {
	if in == nil {
		return nil
	}
	out := new(CredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSpec) DeepCopyInto(out *CredentialSpec) {
	*out = *in
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSpec.
func (in *CredentialSpec) DeepCopy() *CredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialStatus) DeepCopyInto(out *CredentialStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialStatus.
func (in *CredentialStatus) DeepCopy() *CredentialStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmSpec) DeepCopyInto(out *HelmSpec) {
	*out = *in
	if in.ChartRef != nil {
		in, out := &in.ChartRef, &out.ChartRef
		*out = new(v2.CrossNamespaceSourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmSpec.
func (in *HelmSpec) DeepCopy() *HelmSpec {
	if in == nil {
		return nil
	}
	out := new(HelmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCluster) DeepCopyInto(out *ManagedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCluster.
func (in *ManagedCluster) DeepCopy() *ManagedCluster {
	if in == nil {
		return nil
	}
	out := new(ManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterList) DeepCopyInto(out *ManagedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterList.
func (in *ManagedClusterList) DeepCopy() *ManagedClusterList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSpec) DeepCopyInto(out *ManagedClusterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSpec.
func (in *ManagedClusterSpec) DeepCopy() *ManagedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterStatus) DeepCopyInto(out *ManagedClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterStatus.
func (in *ManagedClusterStatus) DeepCopy() *ManagedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Management) DeepCopyInto(out *Management) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Management.
func (in *Management) DeepCopy() *Management {
	if in == nil {
		return nil
	}
	out := new(Management)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Management) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementList) DeepCopyInto(out *ManagementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Management, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementList.
func (in *ManagementList) DeepCopy() *ManagementList {
	if in == nil {
		return nil
	}
	out := new(ManagementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementSpec) DeepCopyInto(out *ManagementSpec) {
	*out = *in
	if in.Core != nil {
		in, out := &in.Core, &out.Core
		*out = new(Core)
		(*in).DeepCopyInto(*out)
	}
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementSpec.
func (in *ManagementSpec) DeepCopy() *ManagementSpec {
	if in == nil {
		return nil
	}
	out := new(ManagementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementStatus) DeepCopyInto(out *ManagementStatus) {
	*out = *in
	in.AvailableProviders.DeepCopyInto(&out.AvailableProviders)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]ComponentStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementStatus.
func (in *ManagementStatus) DeepCopy() *ManagementStatus {
	if in == nil {
		return nil
	}
	out := new(ManagementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedProviderTemplate) DeepCopyInto(out *NamedProviderTemplate) {
	*out = *in
	out.CoreProviderTemplate = in.CoreProviderTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedProviderTemplate.
func (in *NamedProviderTemplate) DeepCopy() *NamedProviderTemplate {
	if in == nil {
		return nil
	}
	out := new(NamedProviderTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderTemplate) DeepCopyInto(out *ProviderTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderTemplate.
func (in *ProviderTemplate) DeepCopy() *ProviderTemplate {
	if in == nil {
		return nil
	}
	out := new(ProviderTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderTemplateList) DeepCopyInto(out *ProviderTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderTemplateList.
func (in *ProviderTemplateList) DeepCopy() *ProviderTemplateList {
	if in == nil {
		return nil
	}
	out := new(ProviderTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderTemplateSpec) DeepCopyInto(out *ProviderTemplateSpec) {
	*out = *in
	in.TemplateSpecCommon.DeepCopyInto(&out.TemplateSpecCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderTemplateSpec.
func (in *ProviderTemplateSpec) DeepCopy() *ProviderTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderTemplateStatus) DeepCopyInto(out *ProviderTemplateStatus) {
	*out = *in
	in.TemplateStatusCommon.DeepCopyInto(&out.TemplateStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderTemplateStatus.
func (in *ProviderTemplateStatus) DeepCopy() *ProviderTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Providers) DeepCopyInto(out *Providers) {
	*out = *in
	if in.InfrastructureProviders != nil {
		in, out := &in.InfrastructureProviders, &out.InfrastructureProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapProviders != nil {
		in, out := &in.BootstrapProviders, &out.BootstrapProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ControlPlaneProviders != nil {
		in, out := &in.ControlPlaneProviders, &out.ControlPlaneProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Providers.
func (in *Providers) DeepCopy() *Providers {
	if in == nil {
		return nil
	}
	out := new(Providers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	if in.UpgradeableVersions != nil {
		in, out := &in.UpgradeableVersions, &out.UpgradeableVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.HMC = in.HMC
	out.CAPI = in.CAPI
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]NamedProviderTemplate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	out.Templates = in.Templates
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplate) DeepCopyInto(out *ServiceTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplate.
func (in *ServiceTemplate) DeepCopy() *ServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplateChain) DeepCopyInto(out *ServiceTemplateChain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplateChain.
func (in *ServiceTemplateChain) DeepCopy() *ServiceTemplateChain {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplateChain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceTemplateChain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplateChainList) DeepCopyInto(out *ServiceTemplateChainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceTemplateChain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplateChainList.
func (in *ServiceTemplateChainList) DeepCopy() *ServiceTemplateChainList {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplateChainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceTemplateChainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplateList) DeepCopyInto(out *ServiceTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplateList.
func (in *ServiceTemplateList) DeepCopy() *ServiceTemplateList {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplateSpec) DeepCopyInto(out *ServiceTemplateSpec) {
	*out = *in
	in.TemplateSpecCommon.DeepCopyInto(&out.TemplateSpecCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplateSpec.
func (in *ServiceTemplateSpec) DeepCopy() *ServiceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplateStatus) DeepCopyInto(out *ServiceTemplateStatus) {
	*out = *in
	in.TemplateStatusCommon.DeepCopyInto(&out.TemplateStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplateStatus.
func (in *ServiceTemplateStatus) DeepCopy() *ServiceTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportedTemplate) DeepCopyInto(out *SupportedTemplate) {
	*out = *in
	if in.AvailableUpgrades != nil {
		in, out := &in.AvailableUpgrades, &out.AvailableUpgrades
		*out = make([]AvailableUpgrade, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportedTemplate.
func (in *SupportedTemplate) DeepCopy() *SupportedTemplate {
	if in == nil {
		return nil
	}
	out := new(SupportedTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNamespaces) DeepCopyInto(out *TargetNamespaces) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNamespaces.
func (in *TargetNamespaces) DeepCopy() *TargetNamespaces {
	if in == nil {
		return nil
	}
	out := new(TargetNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateChainSpec) DeepCopyInto(out *TemplateChainSpec) {
	*out = *in
	if in.SupportedTemplates != nil {
		in, out := &in.SupportedTemplates, &out.SupportedTemplates
		*out = make([]SupportedTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateChainSpec.
func (in *TemplateChainSpec) DeepCopy() *TemplateChainSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateChainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateManagement) DeepCopyInto(out *TemplateManagement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateManagement.
func (in *TemplateManagement) DeepCopy() *TemplateManagement {
	if in == nil {
		return nil
	}
	out := new(TemplateManagement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateManagement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateManagementList) DeepCopyInto(out *TemplateManagementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplateManagement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateManagementList.
func (in *TemplateManagementList) DeepCopy() *TemplateManagementList {
	if in == nil {
		return nil
	}
	out := new(TemplateManagementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateManagementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateManagementSpec) DeepCopyInto(out *TemplateManagementSpec) {
	*out = *in
	if in.AccessRules != nil {
		in, out := &in.AccessRules, &out.AccessRules
		*out = make([]AccessRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateManagementSpec.
func (in *TemplateManagementSpec) DeepCopy() *TemplateManagementSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateManagementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateManagementStatus) DeepCopyInto(out *TemplateManagementStatus) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = make([]AccessRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateManagementStatus.
func (in *TemplateManagementStatus) DeepCopy() *TemplateManagementStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateManagementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecCommon) DeepCopyInto(out *TemplateSpecCommon) {
	*out = *in
	in.Helm.DeepCopyInto(&out.Helm)
	in.Providers.DeepCopyInto(&out.Providers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecCommon.
func (in *TemplateSpecCommon) DeepCopy() *TemplateSpecCommon {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStatusCommon) DeepCopyInto(out *TemplateStatusCommon) {
	*out = *in
	out.TemplateValidationStatus = in.TemplateValidationStatus
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.ChartRef != nil {
		in, out := &in.ChartRef, &out.ChartRef
		*out = new(v2.CrossNamespaceSourceReference)
		**out = **in
	}
	in.Providers.DeepCopyInto(&out.Providers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStatusCommon.
func (in *TemplateStatusCommon) DeepCopy() *TemplateStatusCommon {
	if in == nil {
		return nil
	}
	out := new(TemplateStatusCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateValidationStatus) DeepCopyInto(out *TemplateValidationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateValidationStatus.
func (in *TemplateValidationStatus) DeepCopy() *TemplateValidationStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateValidationStatus)
	in.DeepCopyInto(out)
	return out
}
