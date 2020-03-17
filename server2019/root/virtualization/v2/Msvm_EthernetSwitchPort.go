// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchPort struct
type Msvm_EthernetSwitchPort struct {
	CIM_EthernetPort

	//
	IOVOffloadUsage uint32

	//
	VMQOffloadUsage uint32
}

// SetIOVOffloadUsage sets the value of IOVOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) SetPropertyIOVOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("IOVOffloadUsage", value)
}

// GetIOVOffloadUsage gets the value of IOVOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) GetPropertyIOVOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IOVOffloadUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMQOffloadUsage sets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) SetPropertyVMQOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("VMQOffloadUsage", value)
}

// GetVMQOffloadUsage gets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) GetPropertyVMQOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQOffloadUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPort) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetSwitchPortOffloadData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPortOffloadData")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetSwitchPortBandwidthData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPortBandwidthData")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetPortAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetPortAllocationSettingData")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedDynamicForwardingEntry() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_DynamicForwardingEntry")
}
