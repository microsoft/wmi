// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SmtpSpecifiedPickupDirectorySettings struct
type SmtpSpecifiedPickupDirectorySettings struct {
	*EmbeddedObject

	//
	PickupDirectoryLocation string
}

func NewSmtpSpecifiedPickupDirectorySettingsEx1(instance *cim.WmiInstance) (newInstance *SmtpSpecifiedPickupDirectorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SmtpSpecifiedPickupDirectorySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSmtpSpecifiedPickupDirectorySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SmtpSpecifiedPickupDirectorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SmtpSpecifiedPickupDirectorySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetPickupDirectoryLocation sets the value of PickupDirectoryLocation for the instance
func (instance *SmtpSpecifiedPickupDirectorySettings) SetPropertyPickupDirectoryLocation(value string) (err error) {
	return instance.SetProperty("PickupDirectoryLocation", (value))
}

// GetPickupDirectoryLocation gets the value of PickupDirectoryLocation for the instance
func (instance *SmtpSpecifiedPickupDirectorySettings) GetPropertyPickupDirectoryLocation() (value string, err error) {
	retValue, err := instance.GetProperty("PickupDirectoryLocation")
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
