// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

import (
	"github.com/microsoft/wmi/pkg/base/monitor"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
)

type VirtualSwitchMonitor struct {
	*monitor.Monitor
}

const (
	queryString = "SELECT * FROM __InstanceOperationEvent WITHIN 10 WHERE TargetInstance isa 'Msvm_VirtualEthernetSwitch' "
)

// CreateVirtualSwitchMonitor createa a new VirtualSwitchMonitor
func CreateVirtualSwitchMonitor(callbackContext interface{},
	callbackFunction func(interface{}, string)) *VirtualSwitchMonitor {

	monitorBase := monitor.CreateMonitor(
		string(constant.Virtualization),
		callbackContext,
		callbackFunction,
	)

	return &VirtualSwitchMonitor{monitorBase}
}

func (m *VirtualSwitchMonitor) AddEntity(entityName string) error {
	filters := query.WmiQueryFilterCollection{}
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.ElementName", entityName, query.Equals))
	return m.AddEntityWithFilter(entityName, queryString, filters)
}
