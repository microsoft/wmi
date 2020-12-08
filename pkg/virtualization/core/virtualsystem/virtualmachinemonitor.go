// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"github.com/microsoft/wmi/pkg/base/monitor"
	"github.com/microsoft/wmi/pkg/base/query"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/constant"
)

const (
	// Monitor for State Changes
	queryString = "SELECT * FROM __InstanceOperationEvent WITHIN 5 WHERE TargetInstance isa 'Msvm_ComputerSystem' "
	// Monitor for Setting Changes
	settingQueryString = "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_VirtualSystemSettingData' "
)

type VirtualMachineMonitor struct {
	*monitor.Monitor
}

// CreateVirtualMachineMonitor createa a new VirtualMachineMonitor
func CreateVirtualMachineMonitor(callbackContext interface{},
	callbackFunction func(interface{}, string)) *VirtualMachineMonitor {

	monitorBase := monitor.CreateMonitor(
		string(constant.Virtualization),
		callbackContext,
		callbackFunction,
	)

	return &VirtualMachineMonitor{monitorBase}
}

func (m *VirtualMachineMonitor) AddEntity(vmID string) error {
	filters := query.WmiQueryFilterCollection{}
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.Name", vmID, query.Equals))
	return m.AddEntityWithFilter(vmID, queryString, filters)
}
