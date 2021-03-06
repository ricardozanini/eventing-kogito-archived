/*
Copyright 2020 The Knative Authors

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
	"context"

	"knative.dev/pkg/apis"
)

// Validate validates KogitoSource.
func (s *KogitoSource) Validate(ctx context.Context) *apis.FieldError {
	var errs *apis.FieldError

	//example: validation for "spec" field.
	errs = errs.Also(s.Spec.Validate(ctx).ViaField("spec"))

	//errs is nil if everything is fine.
	return errs
}

// Validate validates KogitoSourceSpec.
func (sspec *KogitoSourceSpec) Validate(ctx context.Context) *apis.FieldError {
	//Add code for validation webhook for KogitoSourceSpec.
	var errs *apis.FieldError

	if fe := sspec.Sink.Validate(ctx); fe != nil {
		errs = errs.Also(fe.ViaField("sink"))
	}

	if sspec.ServiceAccountName == "" {
		errs = errs.Also(apis.ErrMissingField("serviceAccountName"))
	}

	if sspec.Image == "" {
		errs = errs.Also(apis.ErrMissingField("image"))
	}

	return errs
}
