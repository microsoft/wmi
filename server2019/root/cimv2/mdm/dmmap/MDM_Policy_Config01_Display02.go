// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Config01_Display02 struct
type MDM_Policy_Config01_Display02 struct {
	cim.WmiInstance

	//
	DisablePerProcessDpiForApps string

	//
	EnablePerProcessDpi int32

	//
	EnablePerProcessDpiForApps string

	//
	InstanceID string

	//
	ParentID string

	//
	TurnOffGdiDPIScalingForApps string

	//
	TurnOnGdiDPIScalingForApps string
}

// SetDisablePerProcessDpiForApps sets the value of DisablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyDisablePerProcessDpiForApps(value string) (err error) {
	return instance.SetProperty("DisablePerProcessDpiForApps", value)
}

// GetDisablePerProcessDpiForApps gets the value of DisablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyDisablePerProcessDpiForApps() (value string, err error) {
	retValue, err := instance.GetProperty("DisablePerProcessDpiForApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnablePerProcessDpi sets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyEnablePerProcessDpi(value int32) (err error) {
	return instance.SetProperty("EnablePerProcessDpi", value)
}

// GetEnablePerProcessDpi gets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyEnablePerProcessDpi() (value int32, err error) {
	retValue, err := instance.GetProperty("EnablePerProcessDpi")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnablePerProcessDpiForApps sets the value of EnablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyEnablePerProcessDpiForApps(value string) (err error) {
	return instance.SetProperty("EnablePerProcessDpiForApps", value)
}

// GetEnablePerProcessDpiForApps gets the value of EnablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyEnablePerProcessDpiForApps() (value string, err error) {
	retValue, err := instance.GetProperty("EnablePerProcessDpiForApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOffGdiDPIScalingForApps sets the value of TurnOffGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyTurnOffGdiDPIScalingForApps(value string) (err error) {
	return instance.SetProperty("TurnOffGdiDPIScalingForApps", value)
}

// GetTurnOffGdiDPIScalingForApps gets the value of TurnOffGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyTurnOffGdiDPIScalingForApps() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffGdiDPIScalingForApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOnGdiDPIScalingForApps sets the value of TurnOnGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyTurnOnGdiDPIScalingForApps(value string) (err error) {
	return instance.SetProperty("TurnOnGdiDPIScalingForApps", value)
}

// GetTurnOnGdiDPIScalingForApps gets the value of TurnOnGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyTurnOnGdiDPIScalingForApps() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnGdiDPIScalingForApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
