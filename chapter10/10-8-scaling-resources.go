package chapter10

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	types "k8s.io/apimachinery/pkg/types"
	kubevirtcorev1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ScaleResources(client client.Client) {
	vm := &kubevirtcorev1.VirtualMachine{}
	err := client.Get(context.Background(), types.NamespacedName{
		Name:      "myvm",
		Namespace: "my-namespace",
	}, vm)
	if err != nil {
		panic(err)
	}

	vm.Spec.Template.Spec.Domain.Resources.Requests[corev1.ResourceCPU] = resource.MustParse("500m")

	vm.Spec.Template.Spec.Domain.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("1Gi")

	err = client.Update(context.Background(), vm)
	if err != nil {
		panic(err)
	}
}
