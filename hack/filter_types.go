// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Import all Envoy and Gloo filter types so they are registered and deserialization does
// not fail when using them in the "typed_config" attributes.
// The filter types are autogenerated by looking at all packages in go-control-plane and gloo
// api extensions. As a result, this will need to be re-run when updating go-control-plane
// or gloo extensions if new packages are added.

//go:generate sh filter_types_gen.sh

package main
