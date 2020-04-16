// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	"strings"

	"github.com/microsoft/wmi/pkg/base/monitor"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
)

const (
	// Nic Properties
	syntheticEthernetPortSettingDataQueryString = "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_SyntheticEthernetPortSettingData' "
	// Port Properties
	ethernetPortAllocationSettingDataQueryString = "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_EthernetPortAllocationSettingData' "
	// Guest Network Propertiesa - Depends on Integration Component
	guestNetworkAdapterSettingDataQueryString = "SELECT * FROM __InstanceOperationEvent WITHIN 1 WHERE TargetInstance isa 'Msvm_GuestNetworkAdapterConfiguration' "
)

type VirtualNetworkAdapterMonitor struct {
	*monitor.Monitor
}

// CreateVirtualNetworkAdapterMonitor createa a new VirtualNetworkAdapterMonitor
func CreateVirtualNetworkAdapterMonitor(callbackContext interface{},
	callbackFunction func(interface{}, string)) *VirtualNetworkAdapterMonitor {

	monitorBase := monitor.CreateMonitor(
		string(constant.Virtualization),
		callbackContext,
		callbackFunction,
	)

	return &VirtualNetworkAdapterMonitor{monitorBase}
}

func (m *VirtualNetworkAdapterMonitor) AddEntityToMonitorForGuestNetworkChanges(entityName, instanceID string) error {
	if len(instanceID) == 0 {
		return errors.Wrapf(errors.InvalidFilter, "AddEntityToMonitorForGuestNetworkChanges Entity[%s]InstanceID[%s]", entityName, instanceID)
	}
	filters := query.WmiQueryFilterCollection{}
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.InstanceID", strings.ReplaceAll(instanceID, "\\", "\\\\"), query.Equals))
	return m.AddEntityWithFilter(entityName, guestNetworkAdapterSettingDataQueryString, filters)
}

func (m *VirtualNetworkAdapterMonitor) AddEntityToMonitorAdapterChanges(entityName, instanceID string) error {
	filters := query.WmiQueryFilterCollection{}
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.ElementName", entityName, query.Equals))
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.InstanceID", instanceID, query.Like))
	return m.AddEntityWithFilter(entityName, syntheticEthernetPortSettingDataQueryString, filters)
}

func (m *VirtualNetworkAdapterMonitor) AddEntityToMonitorPortChanges(entityName, instanceID string) error {
	filters := query.WmiQueryFilterCollection{}
	filters = append(filters, query.NewWmiQueryFilter("TargetInstance.InstanceID", instanceID, query.Like))
	return m.AddEntityWithFilter(entityName, ethernetPortAllocationSettingDataQueryString, filters)
}
