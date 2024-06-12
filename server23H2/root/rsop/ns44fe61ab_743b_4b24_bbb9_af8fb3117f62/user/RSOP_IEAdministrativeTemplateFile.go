// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS44FE61AB_743B_4B24_BBB9_AF8FB3117F62.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEAdministrativeTemplateFile struct
type RSOP_IEAdministrativeTemplateFile struct {
	*cim.WmiInstance

	//
	GPOID string

	//
	lastWriteTime string

	//
	name string
}

func NewRSOP_IEAdministrativeTemplateFileEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEAdministrativeTemplateFile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAdministrativeTemplateFile{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEAdministrativeTemplateFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEAdministrativeTemplateFile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAdministrativeTemplateFile{
		WmiInstance: tmp,
	}
	return
}

// SetGPOID sets the value of GPOID for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", (value))
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertyGPOID() (value string, err error) {
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
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertylastWriteTime(value string) (err error) {
	return instance.SetProperty("lastWriteTime", (value))
}

// GetlastWriteTime gets the value of lastWriteTime for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertylastWriteTime() (value string, err error) {
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
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", (value))
}

// Getname gets the value of name for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertyname() (value string, err error) {
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
