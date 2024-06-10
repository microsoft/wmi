// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPortSecuritySettingData struct
type Msvm_EthernetSwitchPortSecuritySettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	AllowIeeePriorityTag bool

	//
	AllowMacSpoofing bool

	//
	AllowTeaming bool

	//
	DynamicIPAddressLimit uint32

	//
	EnableDhcpGuard bool

	//
	EnableFixSpeed10G bool

	//
	EnableRouterGuard bool

	//
	MonitorMode uint8

	//
	MonitorSession uint8

	//
	Reserved bool

	//
	StormLimit uint32

	//
	TeamName string

	//
	TeamNumber uint32

	//
	VirtualSubnetId uint32
}

func NewMsvm_EthernetSwitchPortSecuritySettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortSecuritySettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortSecuritySettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortSecuritySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortSecuritySettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortSecuritySettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetAllowIeeePriorityTag sets the value of AllowIeeePriorityTag for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowIeeePriorityTag(value bool) (err error) {
	return instance.SetProperty("AllowIeeePriorityTag", (value))
}

// GetAllowIeeePriorityTag gets the value of AllowIeeePriorityTag for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowIeeePriorityTag() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowIeeePriorityTag")
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

// SetAllowMacSpoofing sets the value of AllowMacSpoofing for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowMacSpoofing(value bool) (err error) {
	return instance.SetProperty("AllowMacSpoofing", (value))
}

// GetAllowMacSpoofing gets the value of AllowMacSpoofing for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowMacSpoofing() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowMacSpoofing")
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

// SetAllowTeaming sets the value of AllowTeaming for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowTeaming(value bool) (err error) {
	return instance.SetProperty("AllowTeaming", (value))
}

// GetAllowTeaming gets the value of AllowTeaming for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowTeaming() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowTeaming")
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

// SetDynamicIPAddressLimit sets the value of DynamicIPAddressLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyDynamicIPAddressLimit(value uint32) (err error) {
	return instance.SetProperty("DynamicIPAddressLimit", (value))
}

// GetDynamicIPAddressLimit gets the value of DynamicIPAddressLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyDynamicIPAddressLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicIPAddressLimit")
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

// SetEnableDhcpGuard sets the value of EnableDhcpGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableDhcpGuard(value bool) (err error) {
	return instance.SetProperty("EnableDhcpGuard", (value))
}

// GetEnableDhcpGuard gets the value of EnableDhcpGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableDhcpGuard() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDhcpGuard")
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

// SetEnableFixSpeed10G sets the value of EnableFixSpeed10G for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableFixSpeed10G(value bool) (err error) {
	return instance.SetProperty("EnableFixSpeed10G", (value))
}

// GetEnableFixSpeed10G gets the value of EnableFixSpeed10G for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableFixSpeed10G() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFixSpeed10G")
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

// SetEnableRouterGuard sets the value of EnableRouterGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableRouterGuard(value bool) (err error) {
	return instance.SetProperty("EnableRouterGuard", (value))
}

// GetEnableRouterGuard gets the value of EnableRouterGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableRouterGuard() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableRouterGuard")
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

// SetMonitorMode sets the value of MonitorMode for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyMonitorMode(value uint8) (err error) {
	return instance.SetProperty("MonitorMode", (value))
}

// GetMonitorMode gets the value of MonitorMode for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyMonitorMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("MonitorMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMonitorSession sets the value of MonitorSession for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyMonitorSession(value uint8) (err error) {
	return instance.SetProperty("MonitorSession", (value))
}

// GetMonitorSession gets the value of MonitorSession for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyMonitorSession() (value uint8, err error) {
	retValue, err := instance.GetProperty("MonitorSession")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyReserved(value bool) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyReserved() (value bool, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetStormLimit sets the value of StormLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyStormLimit(value uint32) (err error) {
	return instance.SetProperty("StormLimit", (value))
}

// GetStormLimit gets the value of StormLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyStormLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("StormLimit")
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

// SetTeamName sets the value of TeamName for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyTeamName(value string) (err error) {
	return instance.SetProperty("TeamName", (value))
}

// GetTeamName gets the value of TeamName for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyTeamName() (value string, err error) {
	retValue, err := instance.GetProperty("TeamName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTeamNumber sets the value of TeamNumber for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyTeamNumber(value uint32) (err error) {
	return instance.SetProperty("TeamNumber", (value))
}

// GetTeamNumber gets the value of TeamNumber for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyTeamNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("TeamNumber")
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

// SetVirtualSubnetId sets the value of VirtualSubnetId for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyVirtualSubnetId(value uint32) (err error) {
	return instance.SetProperty("VirtualSubnetId", (value))
}

// GetVirtualSubnetId gets the value of VirtualSubnetId for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyVirtualSubnetId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSubnetId")
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
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
