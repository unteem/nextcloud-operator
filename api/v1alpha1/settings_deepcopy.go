/*

Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.gnu.org/licenses/agpl-3.0.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.libre.sh/controller-utils/application/settings"
	"k8s.libre.sh/controller-utils/application/settings/parameters"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebSettings) DeepCopyInto(out *WebSettings) {
	*out = *in
	if in.CreateOptions != nil {
		in, out := &in.CreateOptions, &out.CreateOptions
		*out = new(settings.CreateOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = new(settings.Sources)
		if **in != nil {
			in, out := *in, *out
			*out = make([]settings.Source, len(*in))
			copy(*out, *in)
		}
	}
	if in.ConfTemplate != nil {
		in, out := &in.ConfTemplate, &out.ConfTemplate
		*out = new(parameters.Parameter)
		**out = **in
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSettings) DeepCopyInto(out *AppSettings) {
	*out = *in
	in.CreateOptions.DeepCopyInto(&out.CreateOptions)
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = new(settings.Sources)
		if **in != nil {
			in, out := *in, *out
			*out = make([]settings.Source, len(*in))
			copy(*out, *in)
		}
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.SMTP != nil {
		in, out := &in.SMTP, &out.SMTP
		*out = new(SMTP)
		(*in).DeepCopyInto(*out)
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(General)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectStore != nil {
		in, out := &in.ObjectStore, &out.ObjectStore
		*out = new(ObjectStore)
		(*in).DeepCopyInto(*out)
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(Redis)
		(*in).DeepCopyInto(*out)
	}
}
