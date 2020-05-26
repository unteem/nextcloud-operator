module git.indie.host/operators/nextcloud-operator

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/gojp/goreportcard v0.0.0-20200415071653-59167b516f3f // indirect
	github.com/hashicorp/go-version v1.2.0
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/presslabs/controller-util v0.2.2
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.libre.sh v0.0.0-20200512221307-33a77e4357d3
	sigs.k8s.io/controller-runtime v0.5.0
	sigs.k8s.io/kustomize/kstatus v0.0.1 // indirect
)

replace git.indie.host/operators/nextcloud-operator => ./

replace k8s.libre.sh => ./../application/

replace k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible => k8s.io/client-go v0.17.2

replace k8s.io/apimachinery v0.18.1 => k8s.io/apimachinery v0.17.2
