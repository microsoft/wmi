// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_GPO struct
type RSOP_GPO struct {
	*cim.WmiInstance

	//
	accessDenied bool

	//
	enabled bool

	//
	extensionIds []string

	//
	fileSystemPath string

	//
	filterAllowed bool

	//
	filterId string

	//
	guidName string

	//
	id string

	//
	name string

	//
	securityDescriptor []uint8

	//
	version uint32
}

func NewRSOP_GPOEx1(instance *cim.WmiInstance) (newInstance *RSOP_GPO, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_GPO{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_GPOEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_GPO, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_GPO{
		WmiInstance: tmp,
	}
	return
}

// SetaccessDenied sets the value of accessDenied for the instance
func (instance *RSOP_GPO) SetPropertyaccessDenied(value bool) (err error) {
	return instance.SetProperty("accessDenied", (value))
}

// GetaccessDenied gets the value of accessDenied for the instance
func (instance *RSOP_GPO) GetPropertyaccessDenied() (value bool, err error) {
	retValue, err := instance.GetProperty("accessDenied")
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

// Setenabled sets the value of enabled for the instance
func (instance *RSOP_GPO) SetPropertyenabled(value bool) (err error) {
	return instance.SetProperty("enabled", (value))
}

// Getenabled gets the value of enabled for the instance
func (instance *RSOP_GPO) GetPropertyenabled() (value bool, err error) {
	retValue, err := instance.GetProperty("enabled")
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

// SetextensionIds sets the value of extensionIds for the instance
func (instance *RSOP_GPO) SetPropertyextensionIds(value []string) (err error) {
	return instance.SetProperty("extensionIds", (value))
}

// GetextensionIds gets the value of extensionIds for the instance
func (instance *RSOP_GPO) GetPropertyextensionIds() (value []string, err error) {
	retValue, err := instance.GetProperty("extensionIds")
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

// SetfileSystemPath sets the value of fileSystemPath for the instance
func (instance *RSOP_GPO) SetPropertyfileSystemPath(value string) (err error) {
	return instance.SetProperty("fileSystemPath", (value))
}

// GetfileSystemPath gets the value of fileSystemPath for the instance
func (instance *RSOP_GPO) GetPropertyfileSystemPath() (value string, err error) {
	retValue, err := instance.GetProperty("fileSystemPath")
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

// SetfilterAllowed sets the value of filterAllowed for the instance
func (instance *RSOP_GPO) SetPropertyfilterAllowed(value bool) (err error) {
	return instance.SetProperty("filterAllowed", (value))
}

// GetfilterAllowed gets the value of filterAllowed for the instance
func (instance *RSOP_GPO) GetPropertyfilterAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("filterAllowed")
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

// SetfilterId sets the value of filterId for the instance
func (instance *RSOP_GPO) SetPropertyfilterId(value string) (err error) {
	return instance.SetProperty("filterId", (value))
}

// GetfilterId gets the value of filterId for the instance
func (instance *RSOP_GPO) GetPropertyfilterId() (value string, err error) {
	retValue, err := instance.GetProperty("filterId")
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

// SetguidName sets the value of guidName for the instance
func (instance *RSOP_GPO) SetPropertyguidName(value string) (err error) {
	return instance.SetProperty("guidName", (value))
}

// GetguidName gets the value of guidName for the instance
func (instance *RSOP_GPO) GetPropertyguidName() (value string, err error) {
	retValue, err := instance.GetProperty("guidName")
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

// Setid sets the value of id for the instance
func (instance *RSOP_GPO) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", (value))
}

// Getid gets the value of id for the instance
func (instance *RSOP_GPO) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
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

// Setname sets the value of name for the instance
func (instance *RSOP_GPO) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", (value))
}

// Getname gets the value of name for the instance
func (instance *RSOP_GPO) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
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

// SetsecurityDescriptor sets the value of securityDescriptor for the instance
func (instance *RSOP_GPO) SetPropertysecurityDescriptor(value []uint8) (err error) {
	return instance.SetProperty("securityDescriptor", (value))
}

// GetsecurityDescriptor gets the value of securityDescriptor for the instance
func (instance *RSOP_GPO) GetPropertysecurityDescriptor() (value []uint8, err error) {
	retValue, err := instance.GetProperty("securityDescriptor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// Setversion sets the value of version for the instance
func (instance *RSOP_GPO) SetPropertyversion(value uint32) (err error) {
	return instance.SetProperty("version", (value))
}

// Getversion gets the value of version for the instance
func (instance *RSOP_GPO) GetPropertyversion() (value uint32, err error) {
	retValue, err := instance.GetProperty("version")
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
