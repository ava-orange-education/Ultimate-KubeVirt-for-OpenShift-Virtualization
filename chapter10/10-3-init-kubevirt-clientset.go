package chapter10

import (
	"log"

	"k8s.io/client-go/rest"

	// Import the KubeVirt versioned clientset
	kubevirtclientset "kubevirt.io/client-go/kubevirt/typed/core/v1"
	// For KubeVirt specific API group clients, if needed moregranularity:
	// kubevirtcorev1 "kubevirt.io/client-go/kubevirt/typed/core/v1"
)

func initializeKubeVirtClient(config *rest.Config) (*kubevirtclientset.KubevirtV1Client, error) {
	// Create a new KubeVirt clientset for the given config.
	clientset, err := kubevirtclientset.NewForConfig(config)
	if err != nil {
		log.Printf("Error creating KubeVirt clientset: %v", err)
		return nil, err
	}
	log.Println("KubeVirt clientset initialized successfully.")
	return clientset, nil
}
