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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServerType is the type of server to connect to.
type ServerType string

const (
	// KubernetesServer type.
	KubernetesServer ServerType = "kubernetes"

	// TanzuServer type.
	TanzuServer ServerType = "tanzu"
)

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	// Servers available.
	Servers []Server `json:"server,omitempty" yaml:"server"`

	// Current server.
	Current Server `json:"server,omitempty" yaml"server"`
}

// Server connection.
type Server struct {
	// Type of the endpoint.
	Type ServerType `json:"type,omitempty" yaml:"type"`

	// Path to the server config.
	Path string `json:"path,omitempty" yaml:"path"`

	// The context to use, defaults to current.
	Context string `json:"context,omitempty" yaml:"context"`
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct {
}

// +kubebuilder:object:root=true

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
