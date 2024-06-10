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

// MDM_Policy_Config01_DisplayEnhancement02 struct
type MDM_Policy_Config01_DisplayEnhancement02 struct {
	*cim.WmiInstance

	//
	AutobrightnessLuxToNitsCurve string

	//
	DefaultAdaptiveColorAdaptationStrength int32

	//
	DefaultBatterySaverBrightnessMultiplier int32

	//
	DefaultBrightnessSliderPercentage int32

	//
	DefaultDimBrightnessMultiplier int32

	//
	InstanceID string

	//
	IsAdaptiveColorOnByDefault int32

	//
	IsAutobrightnessOnByDefault int32

	//
	ParentID string

	//
	ShouldStopTransitionDuringHandsOnDisplay int32
}

func NewMDM_Policy_Config01_DisplayEnhancement02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_DisplayEnhancement02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DisplayEnhancement02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_DisplayEnhancement02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_DisplayEnhancement02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DisplayEnhancement02{
		WmiInstance: tmp,
	}
	return
}

// SetAutobrightnessLuxToNitsCurve sets the value of AutobrightnessLuxToNitsCurve for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyAutobrightnessLuxToNitsCurve(value string) (err error) {
	return instance.SetProperty("AutobrightnessLuxToNitsCurve", (value))
}

// GetAutobrightnessLuxToNitsCurve gets the value of AutobrightnessLuxToNitsCurve for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyAutobrightnessLuxToNitsCurve() (value string, err error) {
	retValue, err := instance.GetProperty("AutobrightnessLuxToNitsCurve")
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

// SetDefaultAdaptiveColorAdaptationStrength sets the value of DefaultAdaptiveColorAdaptationStrength for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyDefaultAdaptiveColorAdaptationStrength(value int32) (err error) {
	return instance.SetProperty("DefaultAdaptiveColorAdaptationStrength", (value))
}

// GetDefaultAdaptiveColorAdaptationStrength gets the value of DefaultAdaptiveColorAdaptationStrength for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyDefaultAdaptiveColorAdaptationStrength() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultAdaptiveColorAdaptationStrength")
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

// SetDefaultBatterySaverBrightnessMultiplier sets the value of DefaultBatterySaverBrightnessMultiplier for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyDefaultBatterySaverBrightnessMultiplier(value int32) (err error) {
	return instance.SetProperty("DefaultBatterySaverBrightnessMultiplier", (value))
}

// GetDefaultBatterySaverBrightnessMultiplier gets the value of DefaultBatterySaverBrightnessMultiplier for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyDefaultBatterySaverBrightnessMultiplier() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultBatterySaverBrightnessMultiplier")
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

// SetDefaultBrightnessSliderPercentage sets the value of DefaultBrightnessSliderPercentage for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyDefaultBrightnessSliderPercentage(value int32) (err error) {
	return instance.SetProperty("DefaultBrightnessSliderPercentage", (value))
}

// GetDefaultBrightnessSliderPercentage gets the value of DefaultBrightnessSliderPercentage for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyDefaultBrightnessSliderPercentage() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultBrightnessSliderPercentage")
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

// SetDefaultDimBrightnessMultiplier sets the value of DefaultDimBrightnessMultiplier for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyDefaultDimBrightnessMultiplier(value int32) (err error) {
	return instance.SetProperty("DefaultDimBrightnessMultiplier", (value))
}

// GetDefaultDimBrightnessMultiplier gets the value of DefaultDimBrightnessMultiplier for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyDefaultDimBrightnessMultiplier() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultDimBrightnessMultiplier")
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
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyInstanceID() (value string, err error) {
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

// SetIsAdaptiveColorOnByDefault sets the value of IsAdaptiveColorOnByDefault for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyIsAdaptiveColorOnByDefault(value int32) (err error) {
	return instance.SetProperty("IsAdaptiveColorOnByDefault", (value))
}

// GetIsAdaptiveColorOnByDefault gets the value of IsAdaptiveColorOnByDefault for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyIsAdaptiveColorOnByDefault() (value int32, err error) {
	retValue, err := instance.GetProperty("IsAdaptiveColorOnByDefault")
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

// SetIsAutobrightnessOnByDefault sets the value of IsAutobrightnessOnByDefault for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyIsAutobrightnessOnByDefault(value int32) (err error) {
	return instance.SetProperty("IsAutobrightnessOnByDefault", (value))
}

// GetIsAutobrightnessOnByDefault gets the value of IsAutobrightnessOnByDefault for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyIsAutobrightnessOnByDefault() (value int32, err error) {
	retValue, err := instance.GetProperty("IsAutobrightnessOnByDefault")
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
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyParentID() (value string, err error) {
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

// SetShouldStopTransitionDuringHandsOnDisplay sets the value of ShouldStopTransitionDuringHandsOnDisplay for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) SetPropertyShouldStopTransitionDuringHandsOnDisplay(value int32) (err error) {
	return instance.SetProperty("ShouldStopTransitionDuringHandsOnDisplay", (value))
}

// GetShouldStopTransitionDuringHandsOnDisplay gets the value of ShouldStopTransitionDuringHandsOnDisplay for the instance
func (instance *MDM_Policy_Config01_DisplayEnhancement02) GetPropertyShouldStopTransitionDuringHandsOnDisplay() (value int32, err error) {
	retValue, err := instance.GetProperty("ShouldStopTransitionDuringHandsOnDisplay")
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
