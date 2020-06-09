module k8s.libre.sh/apps/nextcloud

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/hashicorp/go-version v1.2.0
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/presslabs/controller-util v0.2.2
	github.com/redhat-cop/operator-utils v0.2.4
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.libre.sh v0.0.0-20200512221307-33a77e4357d3
	sigs.k8s.io/controller-runtime v0.6.0
)

replace k8s.libre.sh/apps/nextcloud => ./

replace k8s.libre.sh => ./../application/

replace k8s.io/client-go => k8s.io/client-go v0.18.2
