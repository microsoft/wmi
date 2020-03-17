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

// RSOP_PolicySettingBlocked struct
type RSOP_PolicySettingBlocked struct {
	cim.WmiInstance

	//
	creationTime string

	//
	GPOID string

	//
	id string

	//
	name string

	//
	precedence uint32

	//
	SOMID string
}

// SetcreationTime sets the value of creationTime for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertycreationTime(value string) (err error) {
	return instance.SetProperty("creationTime", value)
}

// GetcreationTime gets the value of creationTime for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertycreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("creationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGPOID sets the value of GPOID for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", value)
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyGPOID() (value string, err error) {
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

// Setid sets the value of id for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
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
func (instance *RSOP_PolicySettingBlocked) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyname() (value string, err error) {
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

// Setprecedence sets the value of precedence for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyprecedence(value uint32) (err error) {
	return instance.SetProperty("precedence", value)
}

// Getprecedence gets the value of precedence for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyprecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("precedence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSOMID sets the value of SOMID for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertySOMID(value string) (err error) {
	return instance.SetProperty("SOMID", value)
}

// GetSOMID gets the value of SOMID for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertySOMID() (value string, err error) {
	retValue, err := instance.GetProperty("SOMID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
