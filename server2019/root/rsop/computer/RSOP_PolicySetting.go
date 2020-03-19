// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolicySetting struct
type RSOP_PolicySetting struct {
	*cim.WmiInstance

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

func NewRSOP_PolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolicySetting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_PolicySetting{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_PolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolicySetting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolicySetting{
		WmiInstance: tmp,
	}
	return
}

// SetcreationTime sets the value of creationTime for the instance
func (instance *RSOP_PolicySetting) SetPropertycreationTime(value string) (err error) {
	return instance.SetProperty("creationTime", value)
}

// GetcreationTime gets the value of creationTime for the instance
func (instance *RSOP_PolicySetting) GetPropertycreationTime() (value string, err error) {
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
func (instance *RSOP_PolicySetting) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", value)
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_PolicySetting) GetPropertyGPOID() (value string, err error) {
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
func (instance *RSOP_PolicySetting) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSOP_PolicySetting) GetPropertyid() (value string, err error) {
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
func (instance *RSOP_PolicySetting) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *RSOP_PolicySetting) GetPropertyname() (value string, err error) {
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
func (instance *RSOP_PolicySetting) SetPropertyprecedence(value uint32) (err error) {
	return instance.SetProperty("precedence", value)
}

// Getprecedence gets the value of precedence for the instance
func (instance *RSOP_PolicySetting) GetPropertyprecedence() (value uint32, err error) {
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
func (instance *RSOP_PolicySetting) SetPropertySOMID(value string) (err error) {
	return instance.SetProperty("SOMID", value)
}

// GetSOMID gets the value of SOMID for the instance
func (instance *RSOP_PolicySetting) GetPropertySOMID() (value string, err error) {
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
