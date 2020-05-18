/*
Copyright © 2020 Kris Nóva <kris@nivenly.com>

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A SourceType represents an OBS source type.
type SourceType string

const (
	// TypeInput is an OBS input source type.
	TypeInput SourceType = "OBS_SOURCE_TYPE_INPUT"

	// TypeFilter is an OBS filter source type.
	TypeFilter SourceType = "OBS_SOURCE_TYPE_FILTER"

	// TypeTransition is an OBS transition source type.
	TypeTransition SourceType = "OBS_SOURCE_TYPE_TRANSITION"

	// TypeScene is an OBS scene source type.
	TypeScene SourceType = "OBS_SOURCE_TYPE_SCENE"
)

// SourceSpec defines the desired state of Source.
type SourceSpec struct {

	// ID is the plugin ID.
	// +immutable
	ID string `json:"id"`

	// Type is the type of OBS source.
	// +immutable
	// +kubebuilder:validation:Enum=OBS_SOURCE_TYPE_INPUT;OBS_SOURCE_TYPE_FILTER;OBS_SOURCE_TYPE_TRANSITION;OBS_SOURCE_TYPE_SCENE
	Type SourceType `json:"type"`
}

// SourceStatus defines the observed state of Source
type SourceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Source is the Schema for the sources API
type Source struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceSpec   `json:"spec,omitempty"`
	Status SourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourceList contains a list of Source
type SourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Source `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Source{}, &SourceList{})
}
