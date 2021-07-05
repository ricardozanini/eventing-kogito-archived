/*
Copyright 2021 The Knative Authors

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
package reconciler

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	kogitofake "github.com/kiegroup/kogito-operator/client/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	"knative.dev/eventing-kogito/pkg/apis/kogito/v1alpha1"
	"knative.dev/eventing-kogito/pkg/reconciler/kogito/resources"
	v1 "knative.dev/eventing/pkg/apis/sources/v1"
	"knative.dev/pkg/reconciler"
)

func TestKogitoRuntimeReconciler_VerifyCreation(t *testing.T) {
	kogitoSourceBytes, err := ioutil.ReadFile("../../examples/order-processing-workflow.yaml")
	assert.NoError(t, err)
	src := &v1alpha1.KogitoSource{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(kogitoSourceBytes), 1000)
	err = decoder.Decode(src)
	assert.NoError(t, err)

	ctx := context.TODO()
	runtimeReconciler := KogitoRuntimeReconciler{KubeClientSet: k8sfake.NewSimpleClientset(), KogitoClientSet: kogitofake.NewSimpleClientset()}
	deployment, binder, event := runtimeReconciler.ReconcileKogitoRuntime(ctx, src,
		&v1.SinkBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      src.GetName(),
				Namespace: src.GetNamespace(),
			},
			Spec: v1.SinkBindingSpec{
				SourceSpec: src.Spec.SourceSpec,
			},
		},
		resources.MakeReceiveAdapter(&resources.ReceiveAdapterArgs{
			EventSource: src.Namespace + "/" + src.Name,
			Source:      src,
			Labels:      resources.Labels(src.Name),
		}))

	// we don't have Kogito Operator to handle the Deployment for us
	assert.Nil(t, deployment)
	assert.NotNil(t, binder)
	assert.True(t, reconciler.EventIs(event, newKogitoRuntimeCreated(src.Namespace, src.Name)))
}
