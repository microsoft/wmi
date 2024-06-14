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

// AssertSettings struct
type AssertSettings struct {
	*EmbeddedObject

	//
	AssertUIEnabled bool

	//
	LogFileName string
}

func NewAssertSettingsEx1(instance *cim.WmiInstance) (newInstance *AssertSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AssertSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewAssertSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AssertSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AssertSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAssertUIEnabled sets the value of AssertUIEnabled for the instance
func (instance *AssertSettings) SetPropertyAssertUIEnabled(value bool) (err error) {
	return instance.SetProperty("AssertUIEnabled", (value))
}

// GetAssertUIEnabled gets the value of AssertUIEnabled for the instance
func (instance *AssertSettings) GetPropertyAssertUIEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AssertUIEnabled")
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

// SetLogFileName sets the value of LogFileName for the instance
func (instance *AssertSettings) SetPropertyLogFileName(value string) (err error) {
	return instance.SetProperty("LogFileName", (value))
}

// GetLogFileName gets the value of LogFileName for the instance
func (instance *AssertSettings) GetPropertyLogFileName() (value string, err error) {
	retValue, err := instance.GetProperty("LogFileName")
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
