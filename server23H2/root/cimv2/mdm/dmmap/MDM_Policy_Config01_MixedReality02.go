// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_MixedReality02 struct
type MDM_Policy_Config01_MixedReality02 struct {
	*cim.WmiInstance

	//
	AADGroupMembershipCacheValidityInDays int32

	//
	AllowLaunchUriInSingleAppKiosk int32

	//
	AutomaticDisplayAdjustment int32

	//
	BrightnessButtonDisabled int32

	//
	ConfigureMovingPlatform int32

	//
	EyeTrackingCalibrationPrompt int32

	//
	FallbackDiagnostics int32

	//
	HeadTrackingMode int32

	//
	InstanceID string

	//
	ManualDownDirectionDisabled int32

	//
	MicrophoneDisabled int32

	//
	ParentID string

	//
	VisitorAutoLogon int32

	//
	VolumeButtonDisabled int32
}

func NewMDM_Policy_Config01_MixedReality02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_MixedReality02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_MixedReality02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_MixedReality02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_MixedReality02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_MixedReality02{
		WmiInstance: tmp,
	}
	return
}

// SetAADGroupMembershipCacheValidityInDays sets the value of AADGroupMembershipCacheValidityInDays for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyAADGroupMembershipCacheValidityInDays(value int32) (err error) {
	return instance.SetProperty("AADGroupMembershipCacheValidityInDays", (value))
}

// GetAADGroupMembershipCacheValidityInDays gets the value of AADGroupMembershipCacheValidityInDays for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyAADGroupMembershipCacheValidityInDays() (value int32, err error) {
	retValue, err := instance.GetProperty("AADGroupMembershipCacheValidityInDays")
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

// SetAllowLaunchUriInSingleAppKiosk sets the value of AllowLaunchUriInSingleAppKiosk for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyAllowLaunchUriInSingleAppKiosk(value int32) (err error) {
	return instance.SetProperty("AllowLaunchUriInSingleAppKiosk", (value))
}

// GetAllowLaunchUriInSingleAppKiosk gets the value of AllowLaunchUriInSingleAppKiosk for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyAllowLaunchUriInSingleAppKiosk() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowLaunchUriInSingleAppKiosk")
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

// SetAutomaticDisplayAdjustment sets the value of AutomaticDisplayAdjustment for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyAutomaticDisplayAdjustment(value int32) (err error) {
	return instance.SetProperty("AutomaticDisplayAdjustment", (value))
}

// GetAutomaticDisplayAdjustment gets the value of AutomaticDisplayAdjustment for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyAutomaticDisplayAdjustment() (value int32, err error) {
	retValue, err := instance.GetProperty("AutomaticDisplayAdjustment")
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

// SetBrightnessButtonDisabled sets the value of BrightnessButtonDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyBrightnessButtonDisabled(value int32) (err error) {
	return instance.SetProperty("BrightnessButtonDisabled", (value))
}

// GetBrightnessButtonDisabled gets the value of BrightnessButtonDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyBrightnessButtonDisabled() (value int32, err error) {
	retValue, err := instance.GetProperty("BrightnessButtonDisabled")
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

// SetConfigureMovingPlatform sets the value of ConfigureMovingPlatform for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyConfigureMovingPlatform(value int32) (err error) {
	return instance.SetProperty("ConfigureMovingPlatform", (value))
}

// GetConfigureMovingPlatform gets the value of ConfigureMovingPlatform for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyConfigureMovingPlatform() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureMovingPlatform")
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

// SetEyeTrackingCalibrationPrompt sets the value of EyeTrackingCalibrationPrompt for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyEyeTrackingCalibrationPrompt(value int32) (err error) {
	return instance.SetProperty("EyeTrackingCalibrationPrompt", (value))
}

// GetEyeTrackingCalibrationPrompt gets the value of EyeTrackingCalibrationPrompt for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyEyeTrackingCalibrationPrompt() (value int32, err error) {
	retValue, err := instance.GetProperty("EyeTrackingCalibrationPrompt")
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

// SetFallbackDiagnostics sets the value of FallbackDiagnostics for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyFallbackDiagnostics(value int32) (err error) {
	return instance.SetProperty("FallbackDiagnostics", (value))
}

// GetFallbackDiagnostics gets the value of FallbackDiagnostics for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyFallbackDiagnostics() (value int32, err error) {
	retValue, err := instance.GetProperty("FallbackDiagnostics")
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

// SetHeadTrackingMode sets the value of HeadTrackingMode for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyHeadTrackingMode(value int32) (err error) {
	return instance.SetProperty("HeadTrackingMode", (value))
}

// GetHeadTrackingMode gets the value of HeadTrackingMode for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyHeadTrackingMode() (value int32, err error) {
	retValue, err := instance.GetProperty("HeadTrackingMode")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyInstanceID() (value string, err error) {
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

// SetManualDownDirectionDisabled sets the value of ManualDownDirectionDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyManualDownDirectionDisabled(value int32) (err error) {
	return instance.SetProperty("ManualDownDirectionDisabled", (value))
}

// GetManualDownDirectionDisabled gets the value of ManualDownDirectionDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyManualDownDirectionDisabled() (value int32, err error) {
	retValue, err := instance.GetProperty("ManualDownDirectionDisabled")
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

// SetMicrophoneDisabled sets the value of MicrophoneDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyMicrophoneDisabled(value int32) (err error) {
	return instance.SetProperty("MicrophoneDisabled", (value))
}

// GetMicrophoneDisabled gets the value of MicrophoneDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyMicrophoneDisabled() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrophoneDisabled")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyParentID() (value string, err error) {
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

// SetVisitorAutoLogon sets the value of VisitorAutoLogon for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyVisitorAutoLogon(value int32) (err error) {
	return instance.SetProperty("VisitorAutoLogon", (value))
}

// GetVisitorAutoLogon gets the value of VisitorAutoLogon for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyVisitorAutoLogon() (value int32, err error) {
	retValue, err := instance.GetProperty("VisitorAutoLogon")
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

// SetVolumeButtonDisabled sets the value of VolumeButtonDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) SetPropertyVolumeButtonDisabled(value int32) (err error) {
	return instance.SetProperty("VolumeButtonDisabled", (value))
}

// GetVolumeButtonDisabled gets the value of VolumeButtonDisabled for the instance
func (instance *MDM_Policy_Config01_MixedReality02) GetPropertyVolumeButtonDisabled() (value int32, err error) {
	retValue, err := instance.GetProperty("VolumeButtonDisabled")
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
