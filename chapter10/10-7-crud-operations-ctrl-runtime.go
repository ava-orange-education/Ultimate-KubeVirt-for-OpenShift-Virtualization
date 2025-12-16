package chapter10

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/ptr"
	kubevirtcorev1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetVMs(c client.Client) {
	vms := &kubevirtcorev1.VirtualMachineList{}
	err := c.List(context.Background(), vms, &client.ListOptions{
		Namespace: "mynamespace",
	})
	if err != nil {
		panic(err)
	}
	for _, vm := range vms.Items {
		fmt.Println("VM Name: %s, Namespace: %s", vm.Name,
			vm.Namespace)
	}
}

func CreateVM(client client.Client) {
	vm := &kubevirtcorev1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "Name",
			Namespace: "project",
			Labels:    map[string]string{},
		},
		Spec: kubevirtcorev1.VirtualMachineSpec{
			Running: ptr.To(true),
			Template: &kubevirtcorev1.VirtualMachineInstanceTemplateSpec{

				Spec: kubevirtcorev1.VirtualMachineInstanceSpec{
					Domain: kubevirtcorev1.DomainSpec{
						Devices: kubevirtcorev1.Devices{
							Disks: []kubevirtcorev1.Disk{
								{
									Name: "emptydisk",
									DiskDevice: kubevirtcorev1.DiskDevice{
										Disk: &kubevirtcorev1.DiskTarget{
											Bus: kubevirtcorev1.DiskBusVirtio,
										},
									},
								},
							},
							Interfaces: []kubevirtcorev1.Interface{
								{
									Name: "default",
									InterfaceBindingMethod: kubevirtcorev1.InterfaceBindingMethod{

										Masquerade: &kubevirtcorev1.InterfaceMasquerade{},
									},
								},
							},
						},
						Resources: kubevirtcorev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceMemory: resource.MustParse("512m"),
								corev1.ResourceCPU:    resource.MustParse("200m"),
							},
						},
						Machine: &kubevirtcorev1.Machine{
							Type: "",
						},
					},
					Networks: []kubevirtcorev1.Network{
						{
							Name: "default",
							NetworkSource: kubevirtcorev1.NetworkSource{

								Pod: &kubevirtcorev1.PodNetwork{},
							},
						},
					},
					Volumes: []kubevirtcorev1.Volume{
						{
							Name: "emptydisk",
							VolumeSource: kubevirtcorev1.VolumeSource{
								EmptyDisk: &kubevirtcorev1.EmptyDiskSource{
									Capacity: resource.MustParse("10GB"),
								},
							},
						},
					},
				},
			},
		},
	}

	err := client.Create(context.Background(), vm)
	if err != nil {
		panic(err)
	}
}

func DeleteVM(client client.Client) {
	vm := &kubevirtcorev1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-vm-name",
			Namespace: "my-vm-namespace",
		},
	}

	err := client.Delete(context.Background(), vm)
	if err != nil {
		panic(err)
	}
}
