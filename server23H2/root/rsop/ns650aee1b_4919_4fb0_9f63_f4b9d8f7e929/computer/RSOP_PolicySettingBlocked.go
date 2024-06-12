// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS650AEE1B_4919_4FB0_9F63_F4B9D8F7E929.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_PolicySettingBlocked struct
type RSOP_PolicySettingBlocked struct {
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

func NewRSOP_PolicySettingBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolicySettingBlocked, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_PolicySettingBlocked{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_PolicySettingBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolicySettingBlocked, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolicySettingBlocked{
		WmiInstance: tmp,
	}
	return
}

// SetcreationTime sets the value of creationTime for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertycreationTime(value string) (err error) {
	return instance.SetProperty("creationTime", (value))
}

// GetcreationTime gets the value of creationTime for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertycreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("creationTime")
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

// SetGPOID sets the value of GPOID for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyGPOID(value string) (err error) {
	return instance.SetProperty("GPOID", (value))
}

// GetGPOID gets the value of GPOID for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyGPOID() (value string, err error) {
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

// Setid sets the value of id for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", (value))
}

// Getid gets the value of id for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyid() (value string, err error) {
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
func (instance *RSOP_PolicySettingBlocked) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", (value))
}

// Getname gets the value of name for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyname() (value string, err error) {
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

// Setprecedence sets the value of precedence for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertyprecedence(value uint32) (err error) {
	return instance.SetProperty("precedence", (value))
}

// Getprecedence gets the value of precedence for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertyprecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("precedence")
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

// SetSOMID sets the value of SOMID for the instance
func (instance *RSOP_PolicySettingBlocked) SetPropertySOMID(value string) (err error) {
	return instance.SetProperty("SOMID", (value))
}

// GetSOMID gets the value of SOMID for the instance
func (instance *RSOP_PolicySettingBlocked) GetPropertySOMID() (value string, err error) {
	retValue, err := instance.GetProperty("SOMID")
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
