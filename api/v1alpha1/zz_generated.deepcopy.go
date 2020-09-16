// +build !ignore_autogenerated

/*
Copyright 2020 GramLabs, Inc.

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
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Assignment) DeepCopyInto(out *Assignment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assignment.
func (in *Assignment) DeepCopy() *Assignment {
	if in == nil {
		return nil
	}
	out := new(Assignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapHelmValuesFromSource) DeepCopyInto(out *ConfigMapHelmValuesFromSource) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapHelmValuesFromSource.
func (in *ConfigMapHelmValuesFromSource) DeepCopy() *ConfigMapHelmValuesFromSource {
	if in == nil {
		return nil
	}
	out := new(ConfigMapHelmValuesFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraint) DeepCopyInto(out *Constraint) {
	*out = *in
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(OrderConstraint)
		**out = **in
	}
	if in.Sum != nil {
		in, out := &in.Sum, &out.Sum
		*out = new(SumConstraint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraint.
func (in *Constraint) DeepCopy() *Constraint {
	if in == nil {
		return nil
	}
	out := new(Constraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Experiment) DeepCopyInto(out *Experiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Experiment.
func (in *Experiment) DeepCopy() *Experiment {
	if in == nil {
		return nil
	}
	out := new(Experiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Experiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentList) DeepCopyInto(out *ExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Experiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentList.
func (in *ExperimentList) DeepCopy() *ExperimentList {
	if in == nil {
		return nil
	}
	out := new(ExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentSpec) DeepCopyInto(out *ExperimentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Optimization != nil {
		in, out := &in.Optimization, &out.Optimization
		*out = make([]Optimization, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]Constraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]Metric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]PatchTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceTemplate != nil {
		in, out := &in.NamespaceTemplate, &out.NamespaceTemplate
		*out = new(NamespaceTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentSpec.
func (in *ExperimentSpec) DeepCopy() *ExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(ExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatus) DeepCopyInto(out *ExperimentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatus.
func (in *ExperimentStatus) DeepCopy() *ExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValue) DeepCopyInto(out *HelmValue) {
	*out = *in
	out.Value = in.Value
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(HelmValueSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValue.
func (in *HelmValue) DeepCopy() *HelmValue {
	if in == nil {
		return nil
	}
	out := new(HelmValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValueSource) DeepCopyInto(out *HelmValueSource) {
	*out = *in
	if in.ParameterRef != nil {
		in, out := &in.ParameterRef, &out.ParameterRef
		*out = new(ParameterSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValueSource.
func (in *HelmValueSource) DeepCopy() *HelmValueSource {
	if in == nil {
		return nil
	}
	out := new(HelmValueSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValuesFromSource) DeepCopyInto(out *HelmValuesFromSource) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapHelmValuesFromSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValuesFromSource.
func (in *HelmValuesFromSource) DeepCopy() *HelmValuesFromSource {
	if in == nil {
		return nil
	}
	out := new(HelmValuesFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Port = in.Port
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceTemplateSpec) DeepCopyInto(out *NamespaceTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceTemplateSpec.
func (in *NamespaceTemplateSpec) DeepCopy() *NamespaceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Optimization) DeepCopyInto(out *Optimization) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Optimization.
func (in *Optimization) DeepCopy() *Optimization {
	if in == nil {
		return nil
	}
	out := new(Optimization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderConstraint) DeepCopyInto(out *OrderConstraint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderConstraint.
func (in *OrderConstraint) DeepCopy() *OrderConstraint {
	if in == nil {
		return nil
	}
	out := new(OrderConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSelector) DeepCopyInto(out *ParameterSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSelector.
func (in *ParameterSelector) DeepCopy() *ParameterSelector {
	if in == nil {
		return nil
	}
	out := new(ParameterSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchOperation) DeepCopyInto(out *PatchOperation) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchOperation.
func (in *PatchOperation) DeepCopy() *PatchOperation {
	if in == nil {
		return nil
	}
	out := new(PatchOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchReadinessGate) DeepCopyInto(out *PatchReadinessGate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchReadinessGate.
func (in *PatchReadinessGate) DeepCopy() *PatchReadinessGate {
	if in == nil {
		return nil
	}
	out := new(PatchReadinessGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchTemplate) DeepCopyInto(out *PatchTemplate) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.ReadinessGates != nil {
		in, out := &in.ReadinessGates, &out.ReadinessGates
		*out = make([]PatchReadinessGate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchTemplate.
func (in *PatchTemplate) DeepCopy() *PatchTemplate {
	if in == nil {
		return nil
	}
	out := new(PatchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheck) DeepCopyInto(out *ReadinessCheck) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionTypes != nil {
		in, out := &in.ConditionTypes, &out.ConditionTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastCheckTime != nil {
		in, out := &in.LastCheckTime, &out.LastCheckTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheck.
func (in *ReadinessCheck) DeepCopy() *ReadinessCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetupTask) DeepCopyInto(out *SetupTask) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelmValues != nil {
		in, out := &in.HelmValues, &out.HelmValues
		*out = make([]HelmValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelmValuesFrom != nil {
		in, out := &in.HelmValuesFrom, &out.HelmValuesFrom
		*out = make([]HelmValuesFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetupTask.
func (in *SetupTask) DeepCopy() *SetupTask {
	if in == nil {
		return nil
	}
	out := new(SetupTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumConstraint) DeepCopyInto(out *SumConstraint) {
	*out = *in
	out.Bound = in.Bound.DeepCopy()
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]SumConstraintParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumConstraint.
func (in *SumConstraint) DeepCopy() *SumConstraint {
	if in == nil {
		return nil
	}
	out := new(SumConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumConstraintParameter) DeepCopyInto(out *SumConstraintParameter) {
	*out = *in
	out.Weight = in.Weight.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumConstraintParameter.
func (in *SumConstraintParameter) DeepCopy() *SumConstraintParameter {
	if in == nil {
		return nil
	}
	out := new(SumConstraintParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trial) DeepCopyInto(out *Trial) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trial.
func (in *Trial) DeepCopy() *Trial {
	if in == nil {
		return nil
	}
	out := new(Trial)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trial) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialCondition) DeepCopyInto(out *TrialCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialCondition.
func (in *TrialCondition) DeepCopy() *TrialCondition {
	if in == nil {
		return nil
	}
	out := new(TrialCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialList) DeepCopyInto(out *TrialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trial, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialList.
func (in *TrialList) DeepCopy() *TrialList {
	if in == nil {
		return nil
	}
	out := new(TrialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialReadinessGate) DeepCopyInto(out *TrialReadinessGate) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionTypes != nil {
		in, out := &in.ConditionTypes, &out.ConditionTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialReadinessGate.
func (in *TrialReadinessGate) DeepCopy() *TrialReadinessGate {
	if in == nil {
		return nil
	}
	out := new(TrialReadinessGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialSpec) DeepCopyInto(out *TrialSpec) {
	*out = *in
	if in.ExperimentRef != nil {
		in, out := &in.ExperimentRef, &out.ExperimentRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Assignments != nil {
		in, out := &in.Assignments, &out.Assignments
		*out = make([]Assignment, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1beta1.JobTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StartTimeOffset != nil {
		in, out := &in.StartTimeOffset, &out.StartTimeOffset
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ApproximateRuntime != nil {
		in, out := &in.ApproximateRuntime, &out.ApproximateRuntime
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.TTLSecondsAfterFailure != nil {
		in, out := &in.TTLSecondsAfterFailure, &out.TTLSecondsAfterFailure
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessGates != nil {
		in, out := &in.ReadinessGates, &out.ReadinessGates
		*out = make([]TrialReadinessGate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PatchOperations != nil {
		in, out := &in.PatchOperations, &out.PatchOperations
		*out = make([]PatchOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadinessChecks != nil {
		in, out := &in.ReadinessChecks, &out.ReadinessChecks
		*out = make([]ReadinessCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]Value, len(*in))
		copy(*out, *in)
	}
	if in.SetupTasks != nil {
		in, out := &in.SetupTasks, &out.SetupTasks
		*out = make([]SetupTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SetupVolumes != nil {
		in, out := &in.SetupVolumes, &out.SetupVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SetupDefaultRules != nil {
		in, out := &in.SetupDefaultRules, &out.SetupDefaultRules
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialSpec.
func (in *TrialSpec) DeepCopy() *TrialSpec {
	if in == nil {
		return nil
	}
	out := new(TrialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialStatus) DeepCopyInto(out *TrialStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TrialCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialStatus.
func (in *TrialStatus) DeepCopy() *TrialStatus {
	if in == nil {
		return nil
	}
	out := new(TrialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialTemplateSpec) DeepCopyInto(out *TrialTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialTemplateSpec.
func (in *TrialTemplateSpec) DeepCopy() *TrialTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TrialTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Value) DeepCopyInto(out *Value) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Value.
func (in *Value) DeepCopy() *Value {
	if in == nil {
		return nil
	}
	out := new(Value)
	in.DeepCopyInto(out)
	return out
}
