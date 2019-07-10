package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Parameter represents the domain of a single component of the experiment search space
type Parameter struct {
	// The name of the parameter
	Name string `json:"name"`
	// The inclusive minimum value of the parameter
	Min int64 `json:"min,omitempty"`
	// The inclusive maximum value of the parameter
	Max int64 `json:"max,omitempty"`
}

// MetricType represents the allowable types of metrics
type MetricType string

const (
	// Local metrics are Go Templates evaluated against the trial itself. No external service is consulted, primarily
	// useful for extracting start and completion times.
	MetricLocal MetricType = "local"
	// Prometheus metrics issue PromQL queries to a matched service. Queries MUST evaluate to a scalar value.
	MetricPrometheus = "prometheus"
	// JSON path metrics fetch a JSON resource from the matched service. Queries are JSON path expression evaluated against the resource.
	MetricJSONPath = "jsonpath"
	// TODO "regex"?
)

// Metric represents an observable outcome from a trial run
type Metric struct {
	// The name of the metric
	Name string `json:"name"`
	// Indicator that the goal of the experiment is to minimize the value of this metric
	Minimize bool `json:"minimize,omitempty"`

	// The metric collection type, e.g. "prometheus"
	Type MetricType `json:"type,omitempty"`
	// Collection type specific query, e.g. PromQL or a JSON pointer expression
	Query string `json:"query"`

	// Selector matching services to collect this metric from, only the first matched service to provide a value is used
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// The port number or name on the matched service to collect the metric value from
	Port intstr.IntOrString `json:"port,omitempty"`
	// URL path component used to collect the metric value from an endpoint (used as a prefix for the Prometheus API)
	Path string `json:"path,omitempty"`

	// TODO ErrorQuery?
}

// PatchTemplate defines a target resource and a patch template to apply
type PatchTemplate struct {
	// The patch type, one of: json|merge|strategic, default: strategic
	Type string `json:"type,omitempty"`
	// A Go Template that evaluates to valid patch.
	Patch string `json:"patch"`
	// Direct reference to the object the patch should be applied to. The name can be omitted to match by label selector.
	TargetRef corev1.ObjectReference `json:"targetRef"`
	// A selector matching multiple labeled objects the patch should be applied to.
	// Used only if the target reference name is empty, the target reference API version and kind are used
	// to determine what type of object should be matched.
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

// TrialTemplateSpec is used as a template for creating new trials
type TrialTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrialSpec `json:"spec"`
}

// ExperimentSpec defines the desired state of Experiment
type ExperimentSpec struct {
	// Replicas is the number of trials to execute concurrently, defaults to 1
	Replicas *int32 `json:"replicas,omitempty"`
	// Parallelism is the total number of expected replicas across all clusters, defaults to the replica count
	Parallelism *int32 `json:"parallelism,omitempty"`
	// Parameters defines the search space for the experiment
	Parameters []Parameter `json:"parameters,omitempty"`
	// Metrics defines the outcomes for the experiment
	Metrics []Metric `json:"metrics,omitempty"`
	// Patches is a sequence of templates written against the experiment parameters that will be used to put the
	// cluster into the desired state
	Patches []PatchTemplate `json:"patches,omitempty"`
	// NamespaceSelector is used to determine which namespaces on a cluster can be used to create trials. Only a single
	// trial can be created in each namespace so if there are fewer matching namespaces then replicas, no trials will
	// be created
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	// Selector locates trial resources that are part of this experiment
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// Template for creating a new trial. The resulting trial must be matched by Selector. The template can provide an
	// initial namespace, however other namespaces (matched by NamespaceSelector) will be used if the effective
	// replica count is more then one
	Template TrialTemplateSpec `json:"template"`
}

// ExperimentStatus defines the observed state of Experiment
type ExperimentStatus struct {
	// TODO Number of trials: Active, Succeeded, Failed int32 (this is difficult, if not impossible, because we delete trials)
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Experiment is the Schema for the experiments API
// +k8s:openapi-gen=true
type Experiment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExperimentSpec   `json:"spec,omitempty"`
	Status ExperimentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExperimentList contains a list of Experiment
type ExperimentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Experiment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Experiment{}, &ExperimentList{})
}
