module knative.dev/eventing-kogito

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kiegroup/kogito-operator v1.7.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v12.0.0+incompatible
	knative.dev/eventing v0.23.1-0.20210623160544-0cb787308255
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210622173328-dd0db4b05c80
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	k8s.io/api => k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.7
	// transitive dependencies from kogito-operator
	k8s.io/client-go => k8s.io/client-go v0.20.7
	// TODO: migrate Kogito Operator Knative dependency (or remove it)
	knative.dev/eventing => knative.dev/eventing v0.22.1
)

// Kogito Operator unnecessary dependencies
// TODO: create a package to export Kogito Data Structure
exclude (
	code.gitea.io/sdk/gitea v0.12.0
	github.com/Azure/azure-amqp-common-go/v2 v2.1.0
	github.com/RHsyseng/operator-utils v0.0.0-20200304191317-2425bf382482
	github.com/RHsyseng/operator-utils v0.0.0-20200506183821-e3b4a2ba9c30
	github.com/cucumber/gherkin-go/v11 v11.0.0
	github.com/cucumber/godog v0.11.0
	github.com/go-logr/logr v0.3.0
	github.com/go-logr/zapr v0.2.0
	github.com/go-openapi/spec v0.19.14
	github.com/google/uuid v1.1.2
	github.com/imdario/mergo v0.3.10
	github.com/imdario/mergo v0.3.6
	github.com/imdario/mergo v0.3.7
	github.com/imdario/mergo v0.3.8
	github.com/infinispan/infinispan-operator v0.0.0-20210106103300-03aa6d76d1b2
	github.com/integr8ly/grafana-operator/v3 v3.10.0
	github.com/integr8ly/grafana-operator/v3 v3.4.0
	github.com/keycloak/keycloak-operator v0.0.0-20200917060808-9858b19ca8bf
	github.com/machinebox/graphql v0.2.2
	github.com/mongodb/mongodb-kubernetes-operator v0.3.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/openshift/client-go v3.9.0+incompatible
	github.com/operator-framework/operator-lifecycle-manager v0.0.0-20200321030439-57b580e57e88
	github.com/operator-framework/operator-marketplace v0.0.0-20190919183128-4ef67b2f50e9
	github.com/operator-framework/operator-sdk v0.15.2
	github.com/operator-framework/operator-sdk v0.16.0
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.46.0
	github.com/spf13/cobra v1.4.2
	go.uber.org/zap v1.15.0
	software.sslmate.com/src/go-pkcs12 v0.0.0-20210415151418-c5206de65a78
)
