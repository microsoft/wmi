// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_PciExpressSettingData struct
type Msvm_PciExpressSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	VirtualFunctions []uint16

	//
	VirtualSystemIdentifiers []string
}

func NewMsvm_PciExpressSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_PciExpressSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PciExpressSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_PciExpressSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PciExpressSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PciExpressSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetVirtualFunctions sets the value of VirtualFunctions for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyVirtualFunctions(value []uint16) (err error) {
	return instance.SetProperty("VirtualFunctions", value)
}

// GetVirtualFunctions gets the value of VirtualFunctions for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyVirtualFunctions() (value []uint16, err error) {
	retValue, err := instance.GetProperty("VirtualFunctions")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_PciExpressSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", value)
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_PciExpressSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_PciExpressSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
