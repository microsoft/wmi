// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEAdministrativeTemplateFile struct
type RSOP_IEAdministrativeTemplateFile struct {
	cim.WmiInstance

	//
	GPOID string

	//
	lastWriteTime string

	//
	name string
}

// SetGPOID sets the value of GPOID for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", value)
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertyGPOID() (value string, err error) {
	retValue, err := instance.GetProperty("GPOID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetlastWriteTime sets the value of lastWriteTime for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertylastWriteTime(value string) (err error) {
	return instance.SetProperty("lastWriteTime", value)
}

// GetlastWriteTime gets the value of lastWriteTime for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertylastWriteTime() (value string, err error) {
	retValue, err := instance.GetProperty("lastWriteTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setname sets the value of name for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *RSOP_IEAdministrativeTemplateFile) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
