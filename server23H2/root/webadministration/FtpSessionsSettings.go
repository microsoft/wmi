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

// FtpSessionsSettings struct
type FtpSessionsSettings struct {
	*EmbeddedObject

	//
	Sessions []FtpSessionElement
}

func NewFtpSessionsSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpSessionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpSessionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpSessionsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpSessionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpSessionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetSessions sets the value of Sessions for the instance
func (instance *FtpSessionsSettings) SetPropertySessions(value []FtpSessionElement) (err error) {
	return instance.SetProperty("Sessions", (value))
}

// GetSessions gets the value of Sessions for the instance
func (instance *FtpSessionsSettings) GetPropertySessions() (value []FtpSessionElement, err error) {
	retValue, err := instance.GetProperty("Sessions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpSessionElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpSessionElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpSessionElement(valuetmp))
	}

	return
}
