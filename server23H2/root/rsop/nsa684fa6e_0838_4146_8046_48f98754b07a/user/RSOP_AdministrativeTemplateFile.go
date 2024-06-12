// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA684FA6E_0838_4146_8046_48F98754B07A.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_AdministrativeTemplateFile struct
type RSOP_AdministrativeTemplateFile struct {
	*cim.WmiInstance

	//
	GPOID string

	//
	lastWriteTime string

	//
	name string
}

func NewRSOP_AdministrativeTemplateFileEx1(instance *cim.WmiInstance) (newInstance *RSOP_AdministrativeTemplateFile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_AdministrativeTemplateFile{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_AdministrativeTemplateFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_AdministrativeTemplateFile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_AdministrativeTemplateFile{
		WmiInstance: tmp,
	}
	return
}

// SetGPOID sets the value of GPOID for the instance
func (instance *RSOP_AdministrativeTemplateFile) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", (value))
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_AdministrativeTemplateFile) GetPropertyGPOID() (value string, err error) {
	retValue, err := instance.GetProperty("GPOID")
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

// SetlastWriteTime sets the value of lastWriteTime for the instance
func (instance *RSOP_AdministrativeTemplateFile) SetPropertylastWriteTime(value string) (err error) {
	return instance.SetProperty("lastWriteTime", (value))
}

// GetlastWriteTime gets the value of lastWriteTime for the instance
func (instance *RSOP_AdministrativeTemplateFile) GetPropertylastWriteTime() (value string, err error) {
	retValue, err := instance.GetProperty("lastWriteTime")
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
func (instance *RSOP_AdministrativeTemplateFile) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", (value))
}

// Getname gets the value of name for the instance
func (instance *RSOP_AdministrativeTemplateFile) GetPropertyname() (value string, err error) {
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
