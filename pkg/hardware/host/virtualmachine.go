// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
)

// GetVirtualMachines gets all existing virtual machines
func GetVirtualMachines(whost *host.WmiHost) (vms virtualsystem.VirtualMachineCollection, err error) {
	query := query.NewWmiQuery("Msvm_ComputerSystem")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Msvm_ComputerSystem instances")
	}
	defer func() {
		if err != nil {
			rasdcollection.Close()
		}
	}()

	vms, err = virtualsystem.NewVirtualMachineCollection(rasdcollection)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get new VirtualMachineCollection for rasdcollection [%+v]", rasdcollection)
	}

	return
}
