module k8s.libre.sh/apps/nextcloud

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/redhat-cop/operator-utils v0.2.5
	golang.org/x/tools v0.0.0-20200616195046-dc31b401abb5 // indirect
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.libre.sh/controller-utils v0.2.1
	sigs.k8s.io/controller-runtime v0.6.2
)

replace k8s.libre.sh/apps/nextcloud => ./

replace k8s.io/client-go => k8s.io/client-go v0.18.2
