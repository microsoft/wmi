package virtualmachine

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	wmivm "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

func CreateVirtualMachine() (vm *wmivm.Msvm_ComputeSystem, err error) {
}

func GetVirtualMachine(hostname, vmName string) {
}
