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

// MSChangerProblemDeviceError struct
type MSChangerProblemDeviceError struct {
	*MSChangerDriver

	//
	Active bool

	//
	ChangerProblemType uint32

	//
	InstanceName string
}

func NewMSChangerProblemDeviceErrorEx1(instance *cim.WmiInstance) (newInstance *MSChangerProblemDeviceError, err error) {
	tmp, err := NewMSChangerDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSChangerProblemDeviceError{
		MSChangerDriver: tmp,
	}
	return
}

func NewMSChangerProblemDeviceErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSChangerProblemDeviceError, err error) {
	tmp, err := NewMSChangerDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSChangerProblemDeviceError{
		MSChangerDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSChangerProblemDeviceError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSChangerProblemDeviceError) GetPropertyActive() (value bool, err error) {
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

// SetChangerProblemType sets the value of ChangerProblemType for the instance
func (instance *MSChangerProblemDeviceError) SetPropertyChangerProblemType(value uint32) (err error) {
	return instance.SetProperty("ChangerProblemType", (value))
}

// GetChangerProblemType gets the value of ChangerProblemType for the instance
func (instance *MSChangerProblemDeviceError) GetPropertyChangerProblemType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChangerProblemType")
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
func (instance *MSChangerProblemDeviceError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSChangerProblemDeviceError) GetPropertyInstanceName() (value string, err error) {
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
