// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualEthernetSwitchSettingData struct
type Msvm_VirtualEthernetSwitchSettingData struct {
	*CIM_VirtualEthernetSwitchSettingData

	//
	AllowNetLbfoTeams bool

	//
	BandwidthReservationMode uint32

	//
	BypassExtensionStack bool

	//
	ExtensionOrder []string

	//
	IOVPreferred bool

	//
	PacketDirectEnabled bool

	//
	RequiredExtensionIds []string

	//
	TeamingEnabled bool
}

func NewMsvm_VirtualEthernetSwitchSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualEthernetSwitchSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchSettingData{
		CIM_VirtualEthernetSwitchSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualEthernetSwitchSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchSettingData{
		CIM_VirtualEthernetSwitchSettingData: tmp,
	}
	return
}

// SetAllowNetLbfoTeams sets the value of AllowNetLbfoTeams for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyAllowNetLbfoTeams(value bool) (err error) {
	return instance.SetProperty("AllowNetLbfoTeams", (value))
}

// GetAllowNetLbfoTeams gets the value of AllowNetLbfoTeams for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyAllowNetLbfoTeams() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowNetLbfoTeams")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBandwidthReservationMode sets the value of BandwidthReservationMode for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyBandwidthReservationMode(value uint32) (err error) {
	return instance.SetProperty("BandwidthReservationMode", (value))
}

// GetBandwidthReservationMode gets the value of BandwidthReservationMode for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyBandwidthReservationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("BandwidthReservationMode")
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

// SetBypassExtensionStack sets the value of BypassExtensionStack for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyBypassExtensionStack(value bool) (err error) {
	return instance.SetProperty("BypassExtensionStack", (value))
}

// GetBypassExtensionStack gets the value of BypassExtensionStack for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyBypassExtensionStack() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassExtensionStack")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetExtensionOrder sets the value of ExtensionOrder for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyExtensionOrder(value []string) (err error) {
	return instance.SetProperty("ExtensionOrder", (value))
}

// GetExtensionOrder gets the value of ExtensionOrder for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyExtensionOrder() (value []string, err error) {
	retValue, err := instance.GetProperty("ExtensionOrder")
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

// SetIOVPreferred sets the value of IOVPreferred for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyIOVPreferred(value bool) (err error) {
	return instance.SetProperty("IOVPreferred", (value))
}

// GetIOVPreferred gets the value of IOVPreferred for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyIOVPreferred() (value bool, err error) {
	retValue, err := instance.GetProperty("IOVPreferred")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetPacketDirectEnabled sets the value of PacketDirectEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyPacketDirectEnabled(value bool) (err error) {
	return instance.SetProperty("PacketDirectEnabled", (value))
}

// GetPacketDirectEnabled gets the value of PacketDirectEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyPacketDirectEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PacketDirectEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRequiredExtensionIds sets the value of RequiredExtensionIds for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyRequiredExtensionIds(value []string) (err error) {
	return instance.SetProperty("RequiredExtensionIds", (value))
}

// GetRequiredExtensionIds gets the value of RequiredExtensionIds for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyRequiredExtensionIds() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredExtensionIds")
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

// SetTeamingEnabled sets the value of TeamingEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyTeamingEnabled(value bool) (err error) {
	return instance.SetProperty("TeamingEnabled", (value))
}

// GetTeamingEnabled gets the value of TeamingEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyTeamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TeamingEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetRelatedVirtualEthernetSwitchNicTeamingSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchNicTeamingSettingData")
}

func (instance *Msvm_VirtualEthernetSwitchSettingData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_VirtualEthernetSwitchSettingData) GetRelatedEthernetPortAllocationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_EthernetPortAllocationSettingData")
}
