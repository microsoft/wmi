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

// MSSerial_PortName struct
type MSSerial_PortName struct {
	*MSSerial

	//
	Active bool

	//
	InstanceName string

	//
	PortName string
}

func NewMSSerial_PortNameEx1(instance *cim.WmiInstance) (newInstance *MSSerial_PortName, err error) {
	tmp, err := NewMSSerialEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSerial_PortName{
		MSSerial: tmp,
	}
	return
}

func NewMSSerial_PortNameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSerial_PortName, err error) {
	tmp, err := NewMSSerialEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSerial_PortName{
		MSSerial: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSerial_PortName) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSerial_PortName) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSerial_PortName) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSerial_PortName) GetPropertyInstanceName() (value string, err error) {
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

// SetPortName sets the value of PortName for the instance
func (instance *MSSerial_PortName) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", (value))
}

// GetPortName gets the value of PortName for the instance
func (instance *MSSerial_PortName) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
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
