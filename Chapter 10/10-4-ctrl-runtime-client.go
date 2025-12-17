package chapter10

import (
	"log"

	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	kubevirtcorev1 "kubevirt.io/api/core/v1"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	if err := kubevirtcorev1.AddToScheme(scheme); err != nil {
		panic(err)
	}

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		log.Fatalf("Failed to add corev1 to scheme: %v", err)
	}
	if err := cdiv1.AddToScheme(scheme); err != nil {
		log.Fatalf("Failed to add cdiv1 to scheme: %v", err)
	}
}
func createControllerRuntimeClientWithScheme() {
	c, err := client.New(config.GetConfigOrDie(), client.
		Options{
		Scheme: scheme,
	})
	if err != nil {
		panic(err)
	}

	// use c
	_ = c
}
