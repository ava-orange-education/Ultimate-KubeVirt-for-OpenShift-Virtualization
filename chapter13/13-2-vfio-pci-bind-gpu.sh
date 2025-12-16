sudo modprobe vfio-pci

# Replace 0000:XX:YY.Z with the GPU's PCI address (e.g., 0000:01:00.0)
echo "0000:XX:YY.Z" > /sys/bus/pci/drivers/nvidia/unbind # Or other default driver
echo "vfio-pci" > /sys/bus/pci/devices/0000:XX:YY.Z/driver_override
echo "0000:XX:YY.Z" > /sys/bus/pci/drivers/vfio-pci/bind