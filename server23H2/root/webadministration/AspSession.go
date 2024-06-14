// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AspSession struct
type AspSession struct {
	*EmbeddedObject

	//
	AllowSessionState bool

	//
	KeepSessionIdSecure bool

	//
	Max uint32

	//
	Timeout string
}

func NewAspSessionEx1(instance *cim.WmiInstance) (newInstance *AspSession, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AspSession{
		EmbeddedObject: tmp,
	}
	return
}

func NewAspSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AspSession, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AspSession{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowSessionState sets the value of AllowSessionState for the instance
func (instance *AspSession) SetPropertyAllowSessionState(value bool) (err error) {
	return instance.SetProperty("AllowSessionState", (value))
}

// GetAllowSessionState gets the value of AllowSessionState for the instance
func (instance *AspSession) GetPropertyAllowSessionState() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowSessionState")
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

// SetKeepSessionIdSecure sets the value of KeepSessionIdSecure for the instance
func (instance *AspSession) SetPropertyKeepSessionIdSecure(value bool) (err error) {
	return instance.SetProperty("KeepSessionIdSecure", (value))
}

// GetKeepSessionIdSecure gets the value of KeepSessionIdSecure for the instance
func (instance *AspSession) GetPropertyKeepSessionIdSecure() (value bool, err error) {
	retValue, err := instance.GetProperty("KeepSessionIdSecure")
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

// SetMax sets the value of Max for the instance
func (instance *AspSession) SetPropertyMax(value uint32) (err error) {
	return instance.SetProperty("Max", (value))
}

// GetMax gets the value of Max for the instance
func (instance *AspSession) GetPropertyMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("Max")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *AspSession) SetPropertyTimeout(value string) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *AspSession) GetPropertyTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("Timeout")
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