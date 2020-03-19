// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSRemoteDesktop struct
type Win32_TSRemoteDesktop struct {
	*CIM_LogicalElement

	// Alias of the Desktop
	Alias string

	// Contents of the icon corresponding to the application
	IconContents []uint8

	// Index of the icon
	IconIndex int32

	// Path to the Desktop icon
	IconPath string

	// Whether this Remote Desktop is meant for a virtual machine farm
	IsVmFarm bool

	// Contents of the RDP file corresponding to the desktop
	RDPFileContents string

	// Security Descriptor controlling access to the application, in SDDL Format. Empty string implies allow all access. Does not support DENY ACEs, or ACEs referring to non-domain users or groups.
	SecurityDescriptor string

	// Whether this application should be shown in the TS Web Access
	ShowInPortal bool

	// Virtual machine farm settigns corresponding to the desktop
	VmFarmSettings string
}

func NewWin32_TSRemoteDesktopEx1(instance *cim.WmiInstance) (newInstance *Win32_TSRemoteDesktop, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSRemoteDesktop{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSRemoteDesktopEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSRemoteDesktop, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSRemoteDesktop{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetAlias sets the value of Alias for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyAlias(value string) (err error) {
	return instance.SetProperty("Alias", value)
}

// GetAlias gets the value of Alias for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyAlias() (value string, err error) {
	retValue, err := instance.GetProperty("Alias")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIconContents sets the value of IconContents for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyIconContents(value []uint8) (err error) {
	return instance.SetProperty("IconContents", value)
}

// GetIconContents gets the value of IconContents for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyIconContents() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IconContents")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIconIndex sets the value of IconIndex for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyIconIndex(value int32) (err error) {
	return instance.SetProperty("IconIndex", value)
}

// GetIconIndex gets the value of IconIndex for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyIconIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("IconIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIconPath sets the value of IconPath for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyIconPath(value string) (err error) {
	return instance.SetProperty("IconPath", value)
}

// GetIconPath gets the value of IconPath for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyIconPath() (value string, err error) {
	retValue, err := instance.GetProperty("IconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsVmFarm sets the value of IsVmFarm for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyIsVmFarm(value bool) (err error) {
	return instance.SetProperty("IsVmFarm", value)
}

// GetIsVmFarm gets the value of IsVmFarm for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyIsVmFarm() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVmFarm")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDPFileContents sets the value of RDPFileContents for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyRDPFileContents(value string) (err error) {
	return instance.SetProperty("RDPFileContents", value)
}

// GetRDPFileContents gets the value of RDPFileContents for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyRDPFileContents() (value string, err error) {
	retValue, err := instance.GetProperty("RDPFileContents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertySecurityDescriptor(value string) (err error) {
	return instance.SetProperty("SecurityDescriptor", value)
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertySecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShowInPortal sets the value of ShowInPortal for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyShowInPortal(value bool) (err error) {
	return instance.SetProperty("ShowInPortal", value)
}

// GetShowInPortal gets the value of ShowInPortal for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyShowInPortal() (value bool, err error) {
	retValue, err := instance.GetProperty("ShowInPortal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmFarmSettings sets the value of VmFarmSettings for the instance
func (instance *Win32_TSRemoteDesktop) SetPropertyVmFarmSettings(value string) (err error) {
	return instance.SetProperty("VmFarmSettings", value)
}

// GetVmFarmSettings gets the value of VmFarmSettings for the instance
func (instance *Win32_TSRemoteDesktop) GetPropertyVmFarmSettings() (value string, err error) {
	retValue, err := instance.GetProperty("VmFarmSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
