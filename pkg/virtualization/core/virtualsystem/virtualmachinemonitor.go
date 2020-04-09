// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"github.com/microsoft/wmi/pkg/base/monitor"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/constant"
)

type VirtualMachineMonitor struct {
	*monitor.Monitor
}

// CreateVirtualMachineMonitor createa a new VirtualMachineMonitor
func CreateVirtualMachineMonitor(callbackContext interface{},
	callbackFunction func(interface{}, string),
	propertyToQuery string) *VirtualMachineMonitor {

	queryString := "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_ComputerSystem' "
	monitorBase := monitor.CreateMonitor(
		string(constant.Virtualization),
		callbackContext,
		callbackFunction,
		queryString,
		propertyToQuery,
	)

	return &VirtualMachineMonitor{monitorBase}
}
