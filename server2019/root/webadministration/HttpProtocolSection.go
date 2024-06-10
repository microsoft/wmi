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

// HttpProtocolSection struct
type HttpProtocolSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowKeepAlive bool

	//
	CustomHeaders CustomHeaderSettings

	//
	RedirectHeaders RedirectHeaderSettings
}

func NewHttpProtocolSectionEx1(instance *cim.WmiInstance) (newInstance *HttpProtocolSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpProtocolSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHttpProtocolSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpProtocolSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpProtocolSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowKeepAlive sets the value of AllowKeepAlive for the instance
func (instance *HttpProtocolSection) SetPropertyAllowKeepAlive(value bool) (err error) {
	return instance.SetProperty("AllowKeepAlive", (value))
}

// GetAllowKeepAlive gets the value of AllowKeepAlive for the instance
func (instance *HttpProtocolSection) GetPropertyAllowKeepAlive() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowKeepAlive")
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

// SetCustomHeaders sets the value of CustomHeaders for the instance
func (instance *HttpProtocolSection) SetPropertyCustomHeaders(value CustomHeaderSettings) (err error) {
	return instance.SetProperty("CustomHeaders", (value))
}

// GetCustomHeaders gets the value of CustomHeaders for the instance
func (instance *HttpProtocolSection) GetPropertyCustomHeaders() (value CustomHeaderSettings, err error) {
	retValue, err := instance.GetProperty("CustomHeaders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CustomHeaderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CustomHeaderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CustomHeaderSettings(valuetmp)

	return
}

// SetRedirectHeaders sets the value of RedirectHeaders for the instance
func (instance *HttpProtocolSection) SetPropertyRedirectHeaders(value RedirectHeaderSettings) (err error) {
	return instance.SetProperty("RedirectHeaders", (value))
}

// GetRedirectHeaders gets the value of RedirectHeaders for the instance
func (instance *HttpProtocolSection) GetPropertyRedirectHeaders() (value RedirectHeaderSettings, err error) {
	retValue, err := instance.GetProperty("RedirectHeaders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RedirectHeaderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RedirectHeaderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RedirectHeaderSettings(valuetmp)

	return
}
