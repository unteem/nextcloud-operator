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
