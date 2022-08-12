// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"log"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

// GetVirtualMachines gets all existing virtual machines
func GetVirtualMachines(whost *host.WmiHost) (vms virtualsystem.VirtualMachineCollection, err error) {
	query := query.NewWmiQuery("Msvm_ComputerSystem")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if err != nil {
		log.Printf("[WMI] Error getting Msvm_ComputerSystem instances - Error details [%+v]\n", err)
		return
	}
	defer rasdcollection.Close()

	vms, err = virtualsystem.NewVirtualMachineCollection(rasdcollection)
	if err != nil {
		log.Printf("[WMI] Error getting new VirtualMachineCollection for rasdcollection [%s] - Error details [%+v]\n", rasdcollection, err)
		return
	}

	log.Printf("[WMI] Got virtual machines [%s]", vms)
	return
}
