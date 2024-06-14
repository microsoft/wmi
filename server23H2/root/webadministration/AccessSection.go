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

// AccessSection struct
type AccessSection struct {
	*ConfigurationSection

	//
	SslFlags int32
}

func NewAccessSectionEx1(instance *cim.WmiInstance) (newInstance *AccessSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AccessSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewAccessSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AccessSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AccessSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetSslFlags sets the value of SslFlags for the instance
func (instance *AccessSection) SetPropertySslFlags(value int32) (err error) {
	return instance.SetProperty("SslFlags", (value))
}

// GetSslFlags gets the value of SslFlags for the instance
func (instance *AccessSection) GetPropertySslFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("SslFlags")
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