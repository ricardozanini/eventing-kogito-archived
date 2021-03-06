module knative.dev/eventing-kogito

go 1.14

require (
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kiegroup/kogito-operator/api v1.8.0
	github.com/kiegroup/kogito-operator/client v1.8.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/eventing v0.23.1-0.20210623160544-0cb787308255
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210622173328-dd0db4b05c80
)

replace (
	github.com/kiegroup/kogito-operator/api => github.com/kiegroup/kogito-operator/api v0.0.0-20210702132500-6452df3eb8be
	github.com/kiegroup/kogito-operator/client => github.com/kiegroup/kogito-operator/client v0.0.0-20210702132500-6452df3eb8be
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
)
