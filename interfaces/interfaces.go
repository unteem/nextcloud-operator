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

package interfaces

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reconcile is the interface for Reconcile object structs . This
// interface can be used to pass around Reconcile structs commonly
// used in Operators.
//
// Note however that by default Reconcile structs generated using
// Operator SDK do not implement this interface. Add following
// functions to implement this interface.
//
//     func (r *ReconcileObject) GetClient() client.Client { return r.client }
//     func (r *ReconcileObject) GetScheme() *runtime.Scheme { return r.scheme }
//     func (r *ReconcileObject) GetScheme() *runtime.Recorder { return r.recorder }
//
// The Reconcile object structs must implement this interface to use
// Operatorlib functions.
type Reconcile interface {
	// Getter function for reconcile client
	GetClient() client.Client
	// Getter function for reconcile Scheme
	GetScheme() *runtime.Scheme
	// Getter function for reconcile Scheme
	GetRecorder() record.EventRecorder
	// Getter function for reconcile Scheme
	//GetLogger() logr.Logger
}

// Object is the interface which all Kubernetes objects
// implements. This interface can be used to pass around any
// Kubernetes Object. This helps keep the functions more generic and
// less tied to the specific Objects.
type Object interface {
	// The object needs to implement Meta Object interface from API
	// machinery. This interface is used for various Client operations
	// on Kubernetes objects.
	metav1.Object
	// The object needs to implement Runtime Object interface from API
	// machinery.
	runtime.Object
}

type EnvSource interface {
	GetLocalObjectReference() corev1.LocalObjectReference
	GetValue() string
	GetKey() string
}
