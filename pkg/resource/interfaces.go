/*
Copyright 2019 The Crossplane Authors.

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

package resource

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
)

// A Bindable resource may be bound to another resource. Resources are bindable
// when they available for use.
type Bindable interface {
	SetBindingPhase(p v1alpha1.BindingPhase)
	GetBindingPhase() v1alpha1.BindingPhase
}

// A Conditioned may have conditions set or retrieved. Conditions are typically
// indicate the status of both a resource and its reconciliation process.
type Conditioned interface {
	SetConditions(c ...v1alpha1.Condition)
	GetCondition(v1alpha1.ConditionType) v1alpha1.Condition
}

// A ClaimReferencer may reference a resource claim.
type ClaimReferencer interface {
	SetClaimReference(r *corev1.ObjectReference)
	GetClaimReference() *corev1.ObjectReference
}

// A NonPortableClassReferencer may reference a non-portable resource class.
type NonPortableClassReferencer interface {
	SetNonPortableClassReference(r *corev1.ObjectReference)
	GetNonPortableClassReference() *corev1.ObjectReference
}

// A PortableClassReferencer may reference a local portable class.
type PortableClassReferencer interface {
	SetPortableClassReference(r *corev1.LocalObjectReference)
	GetPortableClassReference() *corev1.LocalObjectReference
}

// A ManagedResourceReferencer may reference a concrete managed resource.
type ManagedResourceReferencer interface {
	SetResourceReference(r *corev1.ObjectReference)
	GetResourceReference() *corev1.ObjectReference
}

// A ConnectionSecretWriterTo may write a connection secret.
type ConnectionSecretWriterTo interface {
	SetWriteConnectionSecretToReference(r corev1.LocalObjectReference)
	GetWriteConnectionSecretToReference() corev1.LocalObjectReference
}

// A Reclaimer may specify a ReclaimPolicy.
type Reclaimer interface {
	SetReclaimPolicy(p v1alpha1.ReclaimPolicy)
	GetReclaimPolicy() v1alpha1.ReclaimPolicy
}

// A PortableClassLister may contain a list of portable classes.
type PortableClassLister interface {
	SetPortableClassItems(i []PortableClass)
	GetPortableClassItems() []PortableClass
}

// A Claim is a Kubernetes object representing an abstract resource claim (e.g.
// an SQL database) that may be bound to a concrete managed resource (e.g. a
// CloudSQL instance).
type Claim interface {
	runtime.Object
	metav1.Object

	PortableClassReferencer
	ManagedResourceReferencer
	ConnectionSecretWriterTo

	Conditioned
	Bindable
}

// A NonPortableClass is a Kubernetes object representing configuration
// specifications for a manged resource.
type NonPortableClass interface {
	runtime.Object
	metav1.Object

	Reclaimer
}

// A Managed is a Kubernetes object representing a concrete managed
// resource (e.g. a CloudSQL instance).
type Managed interface {
	runtime.Object
	metav1.Object

	NonPortableClassReferencer
	ClaimReferencer
	ConnectionSecretWriterTo
	Reclaimer

	Conditioned
	Bindable
}

// A PortableClass is a Kubernetes object representing a default
// behavior for a given claim kind.
type PortableClass interface {
	runtime.Object
	metav1.Object

	NonPortableClassReferencer
}

// A PortableClassList is a Kubernetes object representing representing
// a list of portable classes.
type PortableClassList interface {
	runtime.Object
	metav1.ListInterface

	PortableClassLister
}
