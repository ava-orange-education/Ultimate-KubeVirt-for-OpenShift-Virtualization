package chapter10

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtapiv1 "kubevirt.io/api/core/v1"
	kubevirtclientset "kubevirt.io/client-go/kubevirt/typed/core/v1"

	ptr "k8s.io/utils/ptr"
)

// createVM creates a new VirtualMachine in the specified namespace
func createVM(kvClient kubevirtclientset.KubevirtV1Client,
	namespace string, vmDef *kubevirtapiv1.VirtualMachine) (*kubevirtapiv1.VirtualMachine, error) {

	createdVM, err := kvClient.VirtualMachines(namespace).Create(context.TODO(), vmDef, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	log.Printf("Created VirtualMachine: %s", createdVM.Name)
	return createdVM, nil
}

// updateVM updates an existing VirtualMachine in the specified namespace
func startVM(kvClient kubevirtclientset.KubevirtV1Client, namespace, vmName string) (*kubevirtapiv1.VirtualMachine, error) {

	vm, err := kvClient.VirtualMachines(namespace).Get(context.TODO(), vmName, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	// Modify the spec to start the VM (if not using a runStrategy that keeps it running)

	// For VMs managed by a runStrategy, changing 'running' might not be the direct way.
	// This example assumes direct control over the 'running' state.
	if vm.Spec.Running != nil && !*vm.Spec.Running {
		newVm := vm.DeepCopy()
		newVm.Spec.Running = ptr.To(true) // ptr is a helper to get a pointer to a bool

		updatedVM, err := kvClient.VirtualMachines(namespace).Update(context.TODO(), newVm, metav1.UpdateOptions{})

		if err != nil {
			return nil, err
		}

		log.Printf("VirtualMachine %s updated to running state.", updatedVM.Name)
		return updatedVM, nil
	}
	log.Printf("VirtualMachine %s is already set to run or managed by a runStrategy.", vm.Name)
	return vm, nil
}
