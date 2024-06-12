// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_PnP struct
type SystemConfig_PnP struct {
	*SystemConfig

	//
	ClassGuid interface{}

	//
	DeviceDescription string

	//
	DeviceID string

	//
	DevProblem uint32

	//
	DevStatus uint32

	//
	FriendlyName string

	//
	LowerFilters []string

	//
	LowerFiltersCount uint32

	//
	PdoName string

	//
	ServiceName string

	//
	UpperFilters []string

	//
	UpperFiltersCount uint32
}

func NewSystemConfig_PnPEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_PnP, err error) {
	tmp, err := NewSystemConfigEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_PnP{
		SystemConfig: tmp,
	}
	return
}

func NewSystemConfig_PnPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_PnP, err error) {
	tmp, err := NewSystemConfigEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_PnP{
		SystemConfig: tmp,
	}
	return
}

// SetClassGuid sets the value of ClassGuid for the instance
func (instance *SystemConfig_PnP) SetPropertyClassGuid(value interface{}) (err error) {
	return instance.SetProperty("ClassGuid", (value))
}

// GetClassGuid gets the value of ClassGuid for the instance
func (instance *SystemConfig_PnP) GetPropertyClassGuid() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ClassGuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetDeviceDescription sets the value of DeviceDescription for the instance
func (instance *SystemConfig_PnP) SetPropertyDeviceDescription(value string) (err error) {
	return instance.SetProperty("DeviceDescription", (value))
}

// GetDeviceDescription gets the value of DeviceDescription for the instance
func (instance *SystemConfig_PnP) GetPropertyDeviceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceDescription")
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

// SetDeviceID sets the value of DeviceID for the instance
func (instance *SystemConfig_PnP) SetPropertyDeviceID(value string) (err error) {
	return instance.SetProperty("DeviceID", (value))
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *SystemConfig_PnP) GetPropertyDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceID")
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

// SetDevProblem sets the value of DevProblem for the instance
func (instance *SystemConfig_PnP) SetPropertyDevProblem(value uint32) (err error) {
	return instance.SetProperty("DevProblem", (value))
}

// GetDevProblem gets the value of DevProblem for the instance
func (instance *SystemConfig_PnP) GetPropertyDevProblem() (value uint32, err error) {
	retValue, err := instance.GetProperty("DevProblem")
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

// SetDevStatus sets the value of DevStatus for the instance
func (instance *SystemConfig_PnP) SetPropertyDevStatus(value uint32) (err error) {
	return instance.SetProperty("DevStatus", (value))
}

// GetDevStatus gets the value of DevStatus for the instance
func (instance *SystemConfig_PnP) GetPropertyDevStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("DevStatus")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *SystemConfig_PnP) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *SystemConfig_PnP) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetLowerFilters sets the value of LowerFilters for the instance
func (instance *SystemConfig_PnP) SetPropertyLowerFilters(value []string) (err error) {
	return instance.SetProperty("LowerFilters", (value))
}

// GetLowerFilters gets the value of LowerFilters for the instance
func (instance *SystemConfig_PnP) GetPropertyLowerFilters() (value []string, err error) {
	retValue, err := instance.GetProperty("LowerFilters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetLowerFiltersCount sets the value of LowerFiltersCount for the instance
func (instance *SystemConfig_PnP) SetPropertyLowerFiltersCount(value uint32) (err error) {
	return instance.SetProperty("LowerFiltersCount", (value))
}

// GetLowerFiltersCount gets the value of LowerFiltersCount for the instance
func (instance *SystemConfig_PnP) GetPropertyLowerFiltersCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowerFiltersCount")
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

// SetPdoName sets the value of PdoName for the instance
func (instance *SystemConfig_PnP) SetPropertyPdoName(value string) (err error) {
	return instance.SetProperty("PdoName", (value))
}

// GetPdoName gets the value of PdoName for the instance
func (instance *SystemConfig_PnP) GetPropertyPdoName() (value string, err error) {
	retValue, err := instance.GetProperty("PdoName")
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *SystemConfig_PnP) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", (value))
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *SystemConfig_PnP) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
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

// SetUpperFilters sets the value of UpperFilters for the instance
func (instance *SystemConfig_PnP) SetPropertyUpperFilters(value []string) (err error) {
	return instance.SetProperty("UpperFilters", (value))
}

// GetUpperFilters gets the value of UpperFilters for the instance
func (instance *SystemConfig_PnP) GetPropertyUpperFilters() (value []string, err error) {
	retValue, err := instance.GetProperty("UpperFilters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetUpperFiltersCount sets the value of UpperFiltersCount for the instance
func (instance *SystemConfig_PnP) SetPropertyUpperFiltersCount(value uint32) (err error) {
	return instance.SetProperty("UpperFiltersCount", (value))
}

// GetUpperFiltersCount gets the value of UpperFiltersCount for the instance
func (instance *SystemConfig_PnP) GetPropertyUpperFiltersCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpperFiltersCount")
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
