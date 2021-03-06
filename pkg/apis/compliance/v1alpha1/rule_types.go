package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Rule is the Schema for the rules API
// +kubebuilder:resource:path=rules,scope=Namespaced
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The XCCDF ID
	ID string `json:"id"`
	// The title of the Rule
	Title string `json:"title"`
	// The description of the Rule
	Description string `json:"description,omitempty"`
	// The rationale of the Rule
	Rationale string `json:"rationale,omitempty"`
	// A discretionary warning about the of the Rule
	Warning string `json:"warning,omitempty"`
	// The severity level
	Severity string `json:"severity,omitempty"`
	// The Available fixes
	// +nullable
	// +optional
	AvailableFixes []FixDefinition `json:"availableFixes,omitempty"`
}

// FixDefinition Specifies a fix or remediation
// that applies to a rule
type FixDefinition struct {
	// The platform that the fix applies to
	Platform string `json:"platform,omitempty"`
	// An estimate of the potential disruption or operational
	// degradation that this fix will impose in the target system
	Disruption string `json:"disruption,omitempty"`
	// an object that should bring the rule into compliance
	FixObject *unstructured.Unstructured `json:"fixObject,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RuleList contains a list of Rule
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
