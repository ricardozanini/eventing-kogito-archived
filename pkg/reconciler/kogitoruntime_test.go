package reconciler

import (
	"bytes"
	"context"
	kogitofake "github.com/kiegroup/kogito-operator/client/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	"knative.dev/eventing-kogito/pkg/apis/kogito/v1alpha1"
	"knative.dev/eventing-kogito/pkg/reconciler/kogito/resources"
	v1 "knative.dev/eventing/pkg/apis/sources/v1"
	"knative.dev/pkg/reconciler"
	"testing"
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


