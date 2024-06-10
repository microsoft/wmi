// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VirtualEthernetSwitchSettingData struct
type CIM_VirtualEthernetSwitchSettingData struct {
	*CIM_VirtualSystemSettingData

	// A list of host resource pools to be associated or that are currently associated with the Ethernet Switch for the purpose of the allocation of Ethernet connections between a virtual machine and an Ethernet switch. Each non-Null value of the AssociatedResourcePool property shall conform to the production WBEM_URI_UntypedInstancePath as defined in DSP0207.
	AssociatedResourcePool []string

	// This property specifies the number of unique MAC addresses that can be learned by the switch to support MAC Address Learning, as defined in the IEEE 802.1 standard.
	MaxNumMACAddress uint32

	// A list of VLAN Ids that this switch can access.
	VLANConnection []string
}

func NewCIM_VirtualEthernetSwitchSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_VirtualEthernetSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualEthernetSwitchSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

func NewCIM_VirtualEthernetSwitchSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VirtualEthernetSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualEthernetSwitchSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

// SetAssociatedResourcePool sets the value of AssociatedResourcePool for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyAssociatedResourcePool(value []string) (err error) {
	return instance.SetProperty("AssociatedResourcePool", (value))
}

// GetAssociatedResourcePool gets the value of AssociatedResourcePool for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyAssociatedResourcePool() (value []string, err error) {
	retValue, err := instance.GetProperty("AssociatedResourcePool")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMaxNumMACAddress sets the value of MaxNumMACAddress for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyMaxNumMACAddress(value uint32) (err error) {
	return instance.SetProperty("MaxNumMACAddress", (value))
}

// GetMaxNumMACAddress gets the value of MaxNumMACAddress for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyMaxNumMACAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumMACAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVLANConnection sets the value of VLANConnection for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyVLANConnection(value []string) (err error) {
	return instance.SetProperty("VLANConnection", (value))
}

// GetVLANConnection gets the value of VLANConnection for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyVLANConnection() (value []string, err error) {
	retValue, err := instance.GetProperty("VLANConnection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
