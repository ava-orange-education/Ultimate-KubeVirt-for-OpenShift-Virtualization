package chapter10

import (
	"k8s.io/client-go/rest"
	// kubernetes "k8s.io/client-go/kubernetes"
	// kubevirt "kubevirt.io/client-go/clientset/versioned"
)

func buildConfigInCluster() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
