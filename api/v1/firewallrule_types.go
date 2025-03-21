/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Chain string
type Protocol string
type Action string

const (
	InputChain   Chain = "input"
	ForwardChain Chain = "forward"
	OutputChain  Chain = "output"
)

const (
	IPV4 Protocol = "IPv4"
	IPV6 Protocol = "IPv6"
)

const (
	ActionAccept Action = "accept"
	ActionDrop   Action = "drop"
)

// FirewallRuleSpec defines the desired state of FirewallRule.
type FirewallRuleSpec struct {
	Protocol    Protocol `json:"protocol,omitempty"`
	Chain       Chain    `json:"chain,omitempty"`
	RuleID      *int32   `json:"ruleID,omitempty"`
	Action      Action   `json:"action,omitempty"`
	Source      *string  `json:"source,omitempty"`
	Destination *string  `json:"destination,omitempty"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	State         string `json:"state,omitempty"`
	Message       string `json:"message,omitempty"`
	LastUpdated   string `json:"lastUpdated,omitempty"`
	ErrorReason   string `json:"errorReason,omitempty"`
	AppliedRuleID *int32 `json:"appliedRuleID,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FirewallRule is the Schema for the firewallrules API.
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirewallRuleSpec   `json:"spec,omitempty"`
	Status FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRule.
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}
