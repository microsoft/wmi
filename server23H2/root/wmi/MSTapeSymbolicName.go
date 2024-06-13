// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSTapeSymbolicName struct
type MSTapeSymbolicName struct {
	*MSTapeDriver

	//
	Active bool

	//
	InstanceName string

	//
	TapeSymbolicName string
}

func NewMSTapeSymbolicNameEx1(instance *cim.WmiInstance) (newInstance *MSTapeSymbolicName, err error) {
	tmp, err := NewMSTapeDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSTapeSymbolicName{
		MSTapeDriver: tmp,
	}
	return
}

func NewMSTapeSymbolicNameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSTapeSymbolicName, err error) {
	tmp, err := NewMSTapeDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSTapeSymbolicName{
		MSTapeDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSTapeSymbolicName) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeSymbolicName) GetPropertyActive() (value bool, err error) {
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
func (instance *MSTapeSymbolicName) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeSymbolicName) GetPropertyInstanceName() (value string, err error) {
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

// SetTapeSymbolicName sets the value of TapeSymbolicName for the instance
func (instance *MSTapeSymbolicName) SetPropertyTapeSymbolicName(value string) (err error) {
	return instance.SetProperty("TapeSymbolicName", (value))
}

// GetTapeSymbolicName gets the value of TapeSymbolicName for the instance
func (instance *MSTapeSymbolicName) GetPropertyTapeSymbolicName() (value string, err error) {
	retValue, err := instance.GetProperty("TapeSymbolicName")
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
