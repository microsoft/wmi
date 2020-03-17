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

// Msvm_EthernetSwitchPortSecuritySettingData struct
type Msvm_EthernetSwitchPortSecuritySettingData struct {
	Msvm_EthernetSwitchPortFeatureSettingData

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

// SetAllowIeeePriorityTag sets the value of AllowIeeePriorityTag for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowIeeePriorityTag(value bool) (err error) {
	return instance.SetProperty("AllowIeeePriorityTag", value)
}

// GetAllowIeeePriorityTag gets the value of AllowIeeePriorityTag for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowIeeePriorityTag() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowIeeePriorityTag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowMacSpoofing sets the value of AllowMacSpoofing for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowMacSpoofing(value bool) (err error) {
	return instance.SetProperty("AllowMacSpoofing", value)
}

// GetAllowMacSpoofing gets the value of AllowMacSpoofing for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowMacSpoofing() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowMacSpoofing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowTeaming sets the value of AllowTeaming for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyAllowTeaming(value bool) (err error) {
	return instance.SetProperty("AllowTeaming", value)
}

// GetAllowTeaming gets the value of AllowTeaming for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyAllowTeaming() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowTeaming")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDynamicIPAddressLimit sets the value of DynamicIPAddressLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyDynamicIPAddressLimit(value uint32) (err error) {
	return instance.SetProperty("DynamicIPAddressLimit", value)
}

// GetDynamicIPAddressLimit gets the value of DynamicIPAddressLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyDynamicIPAddressLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicIPAddressLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableDhcpGuard sets the value of EnableDhcpGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableDhcpGuard(value bool) (err error) {
	return instance.SetProperty("EnableDhcpGuard", value)
}

// GetEnableDhcpGuard gets the value of EnableDhcpGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableDhcpGuard() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDhcpGuard")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableFixSpeed10G sets the value of EnableFixSpeed10G for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableFixSpeed10G(value bool) (err error) {
	return instance.SetProperty("EnableFixSpeed10G", value)
}

// GetEnableFixSpeed10G gets the value of EnableFixSpeed10G for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableFixSpeed10G() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFixSpeed10G")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableRouterGuard sets the value of EnableRouterGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyEnableRouterGuard(value bool) (err error) {
	return instance.SetProperty("EnableRouterGuard", value)
}

// GetEnableRouterGuard gets the value of EnableRouterGuard for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyEnableRouterGuard() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableRouterGuard")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitorMode sets the value of MonitorMode for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyMonitorMode(value uint8) (err error) {
	return instance.SetProperty("MonitorMode", value)
}

// GetMonitorMode gets the value of MonitorMode for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyMonitorMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("MonitorMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitorSession sets the value of MonitorSession for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyMonitorSession(value uint8) (err error) {
	return instance.SetProperty("MonitorSession", value)
}

// GetMonitorSession gets the value of MonitorSession for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyMonitorSession() (value uint8, err error) {
	retValue, err := instance.GetProperty("MonitorSession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyReserved(value bool) (err error) {
	return instance.SetProperty("Reserved", value)
}

// GetReserved gets the value of Reserved for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyReserved() (value bool, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStormLimit sets the value of StormLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyStormLimit(value uint32) (err error) {
	return instance.SetProperty("StormLimit", value)
}

// GetStormLimit gets the value of StormLimit for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyStormLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("StormLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeamName sets the value of TeamName for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyTeamName(value string) (err error) {
	return instance.SetProperty("TeamName", value)
}

// GetTeamName gets the value of TeamName for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyTeamName() (value string, err error) {
	retValue, err := instance.GetProperty("TeamName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeamNumber sets the value of TeamNumber for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyTeamNumber(value uint32) (err error) {
	return instance.SetProperty("TeamNumber", value)
}

// GetTeamNumber gets the value of TeamNumber for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyTeamNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("TeamNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSubnetId sets the value of VirtualSubnetId for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) SetPropertyVirtualSubnetId(value uint32) (err error) {
	return instance.SetProperty("VirtualSubnetId", value)
}

// GetVirtualSubnetId gets the value of VirtualSubnetId for the instance
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetPropertyVirtualSubnetId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSubnetId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortSecuritySettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
