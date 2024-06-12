// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_Display02 struct
type MDM_Policy_Config01_Display02 struct {
	*cim.WmiInstance

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

func NewMDM_Policy_Config01_Display02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Display02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Display02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Display02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Display02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Display02{
		WmiInstance: tmp,
	}
	return
}

// SetDisablePerProcessDpiForApps sets the value of DisablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyDisablePerProcessDpiForApps(value string) (err error) {
	return instance.SetProperty("DisablePerProcessDpiForApps", (value))
}

// GetDisablePerProcessDpiForApps gets the value of DisablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyDisablePerProcessDpiForApps() (value string, err error) {
	retValue, err := instance.GetProperty("DisablePerProcessDpiForApps")
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

// SetEnablePerProcessDpi sets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyEnablePerProcessDpi(value int32) (err error) {
	return instance.SetProperty("EnablePerProcessDpi", (value))
}

// GetEnablePerProcessDpi gets the value of EnablePerProcessDpi for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyEnablePerProcessDpi() (value int32, err error) {
	retValue, err := instance.GetProperty("EnablePerProcessDpi")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetEnablePerProcessDpiForApps sets the value of EnablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyEnablePerProcessDpiForApps(value string) (err error) {
	return instance.SetProperty("EnablePerProcessDpiForApps", (value))
}

// GetEnablePerProcessDpiForApps gets the value of EnablePerProcessDpiForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyEnablePerProcessDpiForApps() (value string, err error) {
	retValue, err := instance.GetProperty("EnablePerProcessDpiForApps")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetTurnOffGdiDPIScalingForApps sets the value of TurnOffGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyTurnOffGdiDPIScalingForApps(value string) (err error) {
	return instance.SetProperty("TurnOffGdiDPIScalingForApps", (value))
}

// GetTurnOffGdiDPIScalingForApps gets the value of TurnOffGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyTurnOffGdiDPIScalingForApps() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffGdiDPIScalingForApps")
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

// SetTurnOnGdiDPIScalingForApps sets the value of TurnOnGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) SetPropertyTurnOnGdiDPIScalingForApps(value string) (err error) {
	return instance.SetProperty("TurnOnGdiDPIScalingForApps", (value))
}

// GetTurnOnGdiDPIScalingForApps gets the value of TurnOnGdiDPIScalingForApps for the instance
func (instance *MDM_Policy_Config01_Display02) GetPropertyTurnOnGdiDPIScalingForApps() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnGdiDPIScalingForApps")
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
