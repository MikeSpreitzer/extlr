# extlr
Demo of troubles building an external controller of Kubernetes

```
mjs10:extlr mspreitz$ go build
# github.com/MikeSpreitzer/extlr

./main.go:30: cannot use "k8s.io/apimachinery/pkg/apis/meta/v1".ObjectMeta literal (type "k8s.io/apimachinery/pkg/apis/meta/v1".ObjectMeta) as type "k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1".ObjectMeta in field value

./main.go:31: cannot use leaderElectionClient (type *kubernetes.Clientset) as type "k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/core/v1".EndpointsGetter in field value:
	*kubernetes.Clientset does not implement "k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/core/v1".EndpointsGetter (wrong type for Endpoints method)
		have Endpoints(string) "k8s.io/client-go/kubernetes/typed/core/v1".EndpointsInterface
		want Endpoints(string) "k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/core/v1".EndpointsInterface

./main.go:36: cannot use "github.com/MikeSpreitzer/extlr/systemb/defs".AStruct literal (type "github.com/MikeSpreitzer/extlr/systemb/defs".AStruct) as type "github.com/MikeSpreitzer/extlr/systema/defs".AStruct in assignment
```

The third error is a very direct demonstration of a feature of golang typing.  The first error is really the same thing, and easy enough to work around.

The second error is more troublesome.  Its consequence is that a `client-go` kind of `clientset` can not be used in leader election.
