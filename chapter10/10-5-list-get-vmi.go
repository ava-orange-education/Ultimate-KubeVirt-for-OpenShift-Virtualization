package chapter10

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtcorev1 "kubevirt.io/api/core/v1"
	kubevirtclientset "kubevirt.io/client-go/kubevirt/typed/core/v1"
)

// listVMIsInNamespace lists all VMIs in the specified namespace
func listVMIsInNamespace(kvClient kubevirtclientset.KubevirtV1Client,
	namespace string) error {

	vmiList, err := kvClient.VirtualMachineInstances(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return err
	}
	for _, vmi := range vmiList.Items {
		log.Printf("Found VMI: %s, Status: %s", vmi.Name,
			vmi.Status.Phase)
	}
	return nil
}

// getVMI retrieves a specific VMI by name in the given namespace
func getVMI(kvClient *kubevirtclientset.KubevirtV1Client, namespace,
	vmiName string) (*kubevirtcorev1.VirtualMachineInstance, error) {

	vmi, err := kvClient.VirtualMachineInstances(namespace).Get(context.TODO(), vmiName, v1.GetOptions{})

	if err != nil {
		return nil, err
	}
	log.Printf("VMI %s details: CPU Cores %d", vmi.Name, vmi.
		Spec.Domain.CPU.Cores)
	return vmi, nil
}
