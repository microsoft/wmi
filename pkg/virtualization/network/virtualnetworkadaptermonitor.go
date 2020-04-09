// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package network

import (
	"github.com/microsoft/wmi/pkg/base/monitor"
	"github.com/microsoft/wmi/pkg/constant"
)

type VirtualNetworkAdapterMonitor struct {
	*monitor.Monitor
}

// CreateVirtualNetworkAdapterMonitor createa a new VirtualNetworkAdapterMonitor
func CreateVirtualNetworkAdapterMonitor(callbackContext interface{},
	callbackFunction func(interface{}, string),
	propertyToQuery string) *VirtualNetworkAdapterMonitor {

	// FIXME the query string
	queryString := "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_ComputerSystem' "
	monitorBase := monitor.CreateMonitor(
		string(constant.Virtualization),
		callbackContext,
		callbackFunction,
		queryString,
		propertyToQuery,
	)

	return &VirtualNetworkAdapterMonitor{monitorBase}
}
