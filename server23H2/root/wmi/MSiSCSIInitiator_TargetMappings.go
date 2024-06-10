// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_TargetMappings struct
type MSiSCSIInitiator_TargetMappings struct {
	*cim.WmiInstance

	//
	InitiatorName string

	//
	LUNList []MSiSCSIInitiator_LUNList

	//
	OSBusNumber uint32

	//
	OSDeviceName string

	//
	OSTargetNumber uint32

	//
	TargetName string
}

func NewMSiSCSIInitiator_TargetMappingsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_TargetMappings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetMappings{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_TargetMappingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_TargetMappings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetMappings{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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

// SetLUNList sets the value of LUNList for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyLUNList(value []MSiSCSIInitiator_LUNList) (err error) {
	return instance.SetProperty("LUNList", (value))
}

// GetLUNList gets the value of LUNList for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyLUNList() (value []MSiSCSIInitiator_LUNList, err error) {
	retValue, err := instance.GetProperty("LUNList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSiSCSIInitiator_LUNList)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_LUNList is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSiSCSIInitiator_LUNList(valuetmp))
	}

	return
}

// SetOSBusNumber sets the value of OSBusNumber for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyOSBusNumber(value uint32) (err error) {
	return instance.SetProperty("OSBusNumber", (value))
}

// GetOSBusNumber gets the value of OSBusNumber for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyOSBusNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSBusNumber")
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

// SetOSDeviceName sets the value of OSDeviceName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyOSDeviceName(value string) (err error) {
	return instance.SetProperty("OSDeviceName", (value))
}

// GetOSDeviceName gets the value of OSDeviceName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyOSDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("OSDeviceName")
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

// SetOSTargetNumber sets the value of OSTargetNumber for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyOSTargetNumber(value uint32) (err error) {
	return instance.SetProperty("OSTargetNumber", (value))
}

// GetOSTargetNumber gets the value of OSTargetNumber for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyOSTargetNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSTargetNumber")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_TargetMappings) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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
