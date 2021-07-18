/*
Copyright 2021.

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

// GCPRoleBindingSpec defines the desired state of GCPRoleBinding
type GCPRoleBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Project is the project that resources exists relating to the role.
	// It Cannot be set with folder at the same time.
	// +optional
	Project string `json:"project,omitempty"`

	// Folder is the folder that projects exists relating to the role.
	// It Cannot be set with project at the same time.
	// +optional
	Folder string `json:"folder,omitempty"`

	// Roles specify the roles that should be applied.
	// Note that custom roles must be of the format [projects|organizations]/{parent-name}/roles/{role-name}.
	// +required
	Roles []string `json:"roles"`

	// TTL specifies time to live since the GCPRoleBindingSpec was applied or modified.
	// After exceeding ttl, the role binding created by this comtroller will be deleted.
	// Acceptable units are "d", "h", "m".
	// +required
	TTL string `json:"ttl"`

	// Members specifies identities that will be granted the privilege in role.
	// user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	// domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// +required
	Members []string `json:"members"`
}

// GCPRoleBindingStatus defines the observed state of GCPRoleBinding
type GCPRoleBindingStatus struct {
	// LastSpecModifiedTimestamp specifies when the GCPRoleBindingSpec was modified.
	// +required
	LastSpecModifiedTimestamp metav1.Time `json:"lastSpecModifiedTimestamp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GCPRoleBinding is the Schema for the iams API
type GCPRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPRoleBindingSpec   `json:"spec,omitempty"`
	Status GCPRoleBindingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GCPRoleBindingList contains a list of GCPRoleBinding
type GCPRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GCPRoleBinding{}, &GCPRoleBindingList{})
}
