// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEESC struct
type RSOP_IEESC struct {
	*cim.WmiInstance

	//
	EscEnabled bool

	//
	rsopID string

	//
	rsopPrecedence int32
}

func NewRSOP_IEESCEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEESC, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESC{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEESCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEESC, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESC{
		WmiInstance: tmp,
	}
	return
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESC) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", (value))
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESC) GetPropertyEscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EscEnabled")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEESC) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEESC) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEESC) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEESC) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
