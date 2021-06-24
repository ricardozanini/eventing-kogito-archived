/*
Copyright 2019 The Knative Authors

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

package kogito

import (
	"context"

	reconcilersource "knative.dev/eventing/pkg/reconciler/source"

	"knative.dev/eventing-kogito/pkg/apis/kogito/v1alpha1"

	"github.com/kelseyhightower/envconfig"
	"k8s.io/client-go/tools/cache"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/resolver"

	"knative.dev/eventing-kogito/pkg/reconciler"

	kubeclient "knative.dev/pkg/client/injection/kube/client"
	deploymentinformer "knative.dev/pkg/client/injection/kube/informers/apps/v1/deployment"
	samplesourceinformer "knative.dev/eventing-kogito/pkg/client/injection/informers/kogito/v1alpha1/kogitosource"
	"knative.dev/eventing-kogito/pkg/client/injection/reconciler/kogito/v1alpha1/kogitosource"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	deploymentInformer := deploymentinformer.Get(ctx)
	sampleSourceInformer := samplesourceinformer.Get(ctx)

	r := &Reconciler{
		dr: &reconciler.DeploymentReconciler{KubeClientSet: kubeclient.Get(ctx)},
		// Config accessor takes care of tracing/config/logging config propagation to the receive adapter
		configAccessor: reconcilersource.WatchConfigurations(ctx, "kogito-source", cmw),
	}
	if err := envconfig.Process("", r); err != nil {
		logging.FromContext(ctx).Panicf("required environment variable is not defined: %v", err)
	}

	impl := kogitosource.NewImpl(ctx, r)

	r.sinkResolver = resolver.NewURIResolver(ctx, impl.EnqueueKey)

	logging.FromContext(ctx).Info("Setting up event handlers")

	sampleSourceInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	deploymentInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterControllerGK(v1alpha1.Kind("KogitoSource")),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
