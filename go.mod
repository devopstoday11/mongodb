module kubedb.dev/mongodb

go 1.12

require (
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/aws/aws-sdk-go v1.31.3
	github.com/codeskyblue/go-sh v0.0.0-20190412065543-76bd3d59ff27
	github.com/coreos/prometheus-operator v0.39.0
	github.com/fatih/structs v1.1.0
	github.com/jetstack/cert-manager v0.16.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/prom2json v1.3.0
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.2
	gomodules.xyz/cert v1.0.3
	gomodules.xyz/envsubst v0.1.0
	gomodules.xyz/version v0.1.0
	k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/apiserver v0.18.5
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-aggregator v0.18.5
	kmodules.xyz/client-go v0.0.0-20200818171030-24b2ce405feb
	kmodules.xyz/constants v0.0.0-20200506032633-a21e58ceec72
	kmodules.xyz/custom-resources v0.0.0-20200604135349-9e9f5c4fdba9
	kmodules.xyz/monitoring-agent-api v0.0.0-20200525002655-2aa50cb10ce9
	kmodules.xyz/objectstore-api v0.0.0-20200521103120-92080446e04d
	kmodules.xyz/offshoot-api v0.0.0-20200521035628-e135bf07b226
	kmodules.xyz/webhook-runtime v0.0.0-20200522123600-ca70a7e28ed0
	kubedb.dev/apimachinery v0.14.0-beta.1.0.20200726203932-93f0af4b3d70
	stash.appscode.dev/apimachinery v0.10.0-beta.1
)

replace bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d

replace git.apache.org/thrift.git => github.com/apache/thrift v0.13.0

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v35.0.0+incompatible

replace github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible

replace github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0

replace github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.5.0

replace github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.2.0

replace github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.1.0

replace github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.2.0

replace github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.2.0

replace github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0

replace github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.1.0

replace github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.5.0

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.0.0

replace go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace k8s.io/api => github.com/kmodules/api v0.18.4-0.20200524125823-c8bc107809b9

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625

replace k8s.io/client-go => k8s.io/client-go v0.18.3

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200521033432-49d3646051ad
