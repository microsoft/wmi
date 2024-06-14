// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSAcpi_ThermalZoneTemperature struct
type MSAcpi_ThermalZoneTemperature struct {
	*MSAcpi

	//
	Active bool

	//
	ActiveTripPoint []uint32

	//
	ActiveTripPointCount uint32

	//
	CriticalTripPoint uint32

	//
	CurrentTemperature uint32

	//
	InstanceName string

	//
	PassiveTripPoint uint32

	//
	Reserved uint32

	//
	SamplingPeriod uint32

	//
	ThermalConstant1 uint32

	//
	ThermalConstant2 uint32

	//
	ThermalStamp uint32
}

func NewMSAcpi_ThermalZoneTemperatureEx1(instance *cim.WmiInstance) (newInstance *MSAcpi_ThermalZoneTemperature, err error) {
	tmp, err := NewMSAcpiEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSAcpi_ThermalZoneTemperature{
		MSAcpi: tmp,
	}
	return
}

func NewMSAcpi_ThermalZoneTemperatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSAcpi_ThermalZoneTemperature, err error) {
	tmp, err := NewMSAcpiEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSAcpi_ThermalZoneTemperature{
		MSAcpi: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetActiveTripPoint sets the value of ActiveTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyActiveTripPoint(value []uint32) (err error) {
	return instance.SetProperty("ActiveTripPoint", (value))
}

// GetActiveTripPoint gets the value of ActiveTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyActiveTripPoint() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ActiveTripPoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetActiveTripPointCount sets the value of ActiveTripPointCount for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyActiveTripPointCount(value uint32) (err error) {
	return instance.SetProperty("ActiveTripPointCount", (value))
}

// GetActiveTripPointCount gets the value of ActiveTripPointCount for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyActiveTripPointCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveTripPointCount")
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

// SetCriticalTripPoint sets the value of CriticalTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyCriticalTripPoint(value uint32) (err error) {
	return instance.SetProperty("CriticalTripPoint", (value))
}

// GetCriticalTripPoint gets the value of CriticalTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyCriticalTripPoint() (value uint32, err error) {
	retValue, err := instance.GetProperty("CriticalTripPoint")
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

// SetCurrentTemperature sets the value of CurrentTemperature for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyCurrentTemperature(value uint32) (err error) {
	return instance.SetProperty("CurrentTemperature", (value))
}

// GetCurrentTemperature gets the value of CurrentTemperature for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyCurrentTemperature() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentTemperature")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetPassiveTripPoint sets the value of PassiveTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyPassiveTripPoint(value uint32) (err error) {
	return instance.SetProperty("PassiveTripPoint", (value))
}

// GetPassiveTripPoint gets the value of PassiveTripPoint for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyPassiveTripPoint() (value uint32, err error) {
	retValue, err := instance.GetProperty("PassiveTripPoint")
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

// SetReserved sets the value of Reserved for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetSamplingPeriod sets the value of SamplingPeriod for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertySamplingPeriod(value uint32) (err error) {
	return instance.SetProperty("SamplingPeriod", (value))
}

// GetSamplingPeriod gets the value of SamplingPeriod for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertySamplingPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("SamplingPeriod")
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

// SetThermalConstant1 sets the value of ThermalConstant1 for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyThermalConstant1(value uint32) (err error) {
	return instance.SetProperty("ThermalConstant1", (value))
}

// GetThermalConstant1 gets the value of ThermalConstant1 for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyThermalConstant1() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThermalConstant1")
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

// SetThermalConstant2 sets the value of ThermalConstant2 for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyThermalConstant2(value uint32) (err error) {
	return instance.SetProperty("ThermalConstant2", (value))
}

// GetThermalConstant2 gets the value of ThermalConstant2 for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyThermalConstant2() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThermalConstant2")
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

// SetThermalStamp sets the value of ThermalStamp for the instance
func (instance *MSAcpi_ThermalZoneTemperature) SetPropertyThermalStamp(value uint32) (err error) {
	return instance.SetProperty("ThermalStamp", (value))
}

// GetThermalStamp gets the value of ThermalStamp for the instance
func (instance *MSAcpi_ThermalZoneTemperature) GetPropertyThermalStamp() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThermalStamp")
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
