/*
A demonstration of problems encountered when writing an external controller
that mostly uses client-go but needs leader election --- which is not
in client-go.

Are these problems only because of
https://github.com/kubernetes/kubernetes/issues/43246
or is there more involved?
*/

package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/client/leaderelection/resourcelock"

	adefs "github.com/MikeSpreitzer/extlr/systema/defs"
	bdefs "github.com/MikeSpreitzer/extlr/systemb/defs"
)

func main() {
	var config *rest.Config = nil
	leaderElectionClient := kubernetes.NewForConfigOrDie(rest.AddUserAgent(config, "leader-election"))
	rl := resourcelock.EndpointsLock{
		EndpointsMeta: metav1.ObjectMeta{
			Namespace: "kube-system",
			Name:      "extlr",
		},
		Client: leaderElectionClient,
		LockConfig: resourcelock.ResourceLockConfig{
			Identity:      "me",
		},
	}
	var _ adefs.AStruct = bdefs.AStruct{ID: 1,}
}
