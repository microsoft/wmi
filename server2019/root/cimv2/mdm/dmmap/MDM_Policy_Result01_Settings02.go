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

// MDM_Policy_Result01_Settings02 struct
type MDM_Policy_Result01_Settings02 struct {
	cim.WmiInstance

	//
	AllowAutoPlay int32

	//
	AllowDataSense int32

	//
	AllowDateTime int32

	//
	AllowEditDeviceName int32

	//
	AllowLanguage int32

	//
	AllowOnlineTips int32

	//
	AllowPowerSleep int32

	//
	AllowRegion int32

	//
	AllowSignInOptions int32

	//
	AllowVPN int32

	//
	AllowWorkplace int32

	//
	AllowYourAccount int32

	//
	InstanceID string

	//
	PageVisibilityList string

	//
	ParentID string
}

// SetAllowAutoPlay sets the value of AllowAutoPlay for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowAutoPlay(value int32) (err error) {
	return instance.SetProperty("AllowAutoPlay", value)
}

// GetAllowAutoPlay gets the value of AllowAutoPlay for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowAutoPlay() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutoPlay")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowDataSense sets the value of AllowDataSense for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowDataSense(value int32) (err error) {
	return instance.SetProperty("AllowDataSense", value)
}

// GetAllowDataSense gets the value of AllowDataSense for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowDataSense() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDataSense")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowDateTime sets the value of AllowDateTime for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowDateTime(value int32) (err error) {
	return instance.SetProperty("AllowDateTime", value)
}

// GetAllowDateTime gets the value of AllowDateTime for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowDateTime() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDateTime")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEditDeviceName sets the value of AllowEditDeviceName for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowEditDeviceName(value int32) (err error) {
	return instance.SetProperty("AllowEditDeviceName", value)
}

// GetAllowEditDeviceName gets the value of AllowEditDeviceName for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowEditDeviceName() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowEditDeviceName")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLanguage sets the value of AllowLanguage for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowLanguage(value int32) (err error) {
	return instance.SetProperty("AllowLanguage", value)
}

// GetAllowLanguage gets the value of AllowLanguage for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowLanguage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowLanguage")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOnlineTips sets the value of AllowOnlineTips for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowOnlineTips(value int32) (err error) {
	return instance.SetProperty("AllowOnlineTips", value)
}

// GetAllowOnlineTips gets the value of AllowOnlineTips for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowOnlineTips() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowOnlineTips")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowPowerSleep sets the value of AllowPowerSleep for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowPowerSleep(value int32) (err error) {
	return instance.SetProperty("AllowPowerSleep", value)
}

// GetAllowPowerSleep gets the value of AllowPowerSleep for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowPowerSleep() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPowerSleep")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowRegion sets the value of AllowRegion for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowRegion(value int32) (err error) {
	return instance.SetProperty("AllowRegion", value)
}

// GetAllowRegion gets the value of AllowRegion for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowRegion() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowRegion")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSignInOptions sets the value of AllowSignInOptions for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowSignInOptions(value int32) (err error) {
	return instance.SetProperty("AllowSignInOptions", value)
}

// GetAllowSignInOptions gets the value of AllowSignInOptions for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowSignInOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSignInOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowVPN sets the value of AllowVPN for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowVPN(value int32) (err error) {
	return instance.SetProperty("AllowVPN", value)
}

// GetAllowVPN gets the value of AllowVPN for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowVPN() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVPN")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWorkplace sets the value of AllowWorkplace for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowWorkplace(value int32) (err error) {
	return instance.SetProperty("AllowWorkplace", value)
}

// GetAllowWorkplace gets the value of AllowWorkplace for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowWorkplace() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWorkplace")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowYourAccount sets the value of AllowYourAccount for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyAllowYourAccount(value int32) (err error) {
	return instance.SetProperty("AllowYourAccount", value)
}

// GetAllowYourAccount gets the value of AllowYourAccount for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyAllowYourAccount() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowYourAccount")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyInstanceID() (value string, err error) {
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

// SetPageVisibilityList sets the value of PageVisibilityList for the instance
func (instance *MDM_Policy_Result01_Settings02) SetPropertyPageVisibilityList(value string) (err error) {
	return instance.SetProperty("PageVisibilityList", value)
}

// GetPageVisibilityList gets the value of PageVisibilityList for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyPageVisibilityList() (value string, err error) {
	retValue, err := instance.GetProperty("PageVisibilityList")
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
func (instance *MDM_Policy_Result01_Settings02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Settings02) GetPropertyParentID() (value string, err error) {
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
