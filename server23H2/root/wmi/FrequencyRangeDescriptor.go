// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FrequencyRangeDescriptor struct
type FrequencyRangeDescriptor struct {
	*cim.WmiInstance

	//
	ActiveHeight uint32

	//
	ActiveWidth uint32

	//
	ConstraintType uint32

	//
	MaxHSyncDenominator uint32

	//
	MaxHSyncNumerator uint32

	//
	MaxPixelRate uint32

	//
	MaxVSyncDenominator uint32

	//
	MaxVSyncNumerator uint32

	//
	MinHSyncDenominator uint32

	//
	MinHSyncNumerator uint32

	//
	MinVSyncDenominator uint32

	//
	MinVSyncNumerator uint32

	//
	Origin uint8
}

func NewFrequencyRangeDescriptorEx1(instance *cim.WmiInstance) (newInstance *FrequencyRangeDescriptor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &FrequencyRangeDescriptor{
		WmiInstance: tmp,
	}
	return
}

func NewFrequencyRangeDescriptorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FrequencyRangeDescriptor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FrequencyRangeDescriptor{
		WmiInstance: tmp,
	}
	return
}

// SetActiveHeight sets the value of ActiveHeight for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyActiveHeight(value uint32) (err error) {
	return instance.SetProperty("ActiveHeight", (value))
}

// GetActiveHeight gets the value of ActiveHeight for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyActiveHeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveHeight")
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

// SetActiveWidth sets the value of ActiveWidth for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyActiveWidth(value uint32) (err error) {
	return instance.SetProperty("ActiveWidth", (value))
}

// GetActiveWidth gets the value of ActiveWidth for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyActiveWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveWidth")
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

// SetConstraintType sets the value of ConstraintType for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyConstraintType(value uint32) (err error) {
	return instance.SetProperty("ConstraintType", (value))
}

// GetConstraintType gets the value of ConstraintType for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyConstraintType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConstraintType")
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

// SetMaxHSyncDenominator sets the value of MaxHSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMaxHSyncDenominator(value uint32) (err error) {
	return instance.SetProperty("MaxHSyncDenominator", (value))
}

// GetMaxHSyncDenominator gets the value of MaxHSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMaxHSyncDenominator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxHSyncDenominator")
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

// SetMaxHSyncNumerator sets the value of MaxHSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMaxHSyncNumerator(value uint32) (err error) {
	return instance.SetProperty("MaxHSyncNumerator", (value))
}

// GetMaxHSyncNumerator gets the value of MaxHSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMaxHSyncNumerator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxHSyncNumerator")
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

// SetMaxPixelRate sets the value of MaxPixelRate for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMaxPixelRate(value uint32) (err error) {
	return instance.SetProperty("MaxPixelRate", (value))
}

// GetMaxPixelRate gets the value of MaxPixelRate for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMaxPixelRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPixelRate")
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

// SetMaxVSyncDenominator sets the value of MaxVSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMaxVSyncDenominator(value uint32) (err error) {
	return instance.SetProperty("MaxVSyncDenominator", (value))
}

// GetMaxVSyncDenominator gets the value of MaxVSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMaxVSyncDenominator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxVSyncDenominator")
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

// SetMaxVSyncNumerator sets the value of MaxVSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMaxVSyncNumerator(value uint32) (err error) {
	return instance.SetProperty("MaxVSyncNumerator", (value))
}

// GetMaxVSyncNumerator gets the value of MaxVSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMaxVSyncNumerator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxVSyncNumerator")
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

// SetMinHSyncDenominator sets the value of MinHSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMinHSyncDenominator(value uint32) (err error) {
	return instance.SetProperty("MinHSyncDenominator", (value))
}

// GetMinHSyncDenominator gets the value of MinHSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMinHSyncDenominator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinHSyncDenominator")
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

// SetMinHSyncNumerator sets the value of MinHSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMinHSyncNumerator(value uint32) (err error) {
	return instance.SetProperty("MinHSyncNumerator", (value))
}

// GetMinHSyncNumerator gets the value of MinHSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMinHSyncNumerator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinHSyncNumerator")
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

// SetMinVSyncDenominator sets the value of MinVSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMinVSyncDenominator(value uint32) (err error) {
	return instance.SetProperty("MinVSyncDenominator", (value))
}

// GetMinVSyncDenominator gets the value of MinVSyncDenominator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMinVSyncDenominator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinVSyncDenominator")
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

// SetMinVSyncNumerator sets the value of MinVSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyMinVSyncNumerator(value uint32) (err error) {
	return instance.SetProperty("MinVSyncNumerator", (value))
}

// GetMinVSyncNumerator gets the value of MinVSyncNumerator for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyMinVSyncNumerator() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinVSyncNumerator")
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

// SetOrigin sets the value of Origin for the instance
func (instance *FrequencyRangeDescriptor) SetPropertyOrigin(value uint8) (err error) {
	return instance.SetProperty("Origin", (value))
}

// GetOrigin gets the value of Origin for the instance
func (instance *FrequencyRangeDescriptor) GetPropertyOrigin() (value uint8, err error) {
	retValue, err := instance.GetProperty("Origin")
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
