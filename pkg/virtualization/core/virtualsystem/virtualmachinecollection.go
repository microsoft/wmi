// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

type VirtualMachineCollection []*VirtualMachine

func (vms *VirtualMachineCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *VirtualMachineCollection) String() string {
	return ""
}
