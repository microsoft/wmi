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

// HttpCookiesSection struct
type HttpCookiesSection struct {
	*ConfigurationSection

	//
	Domain string

	//
	HttpOnlyCookies bool

	//
	RequireSSL bool
}

func NewHttpCookiesSectionEx1(instance *cim.WmiInstance) (newInstance *HttpCookiesSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpCookiesSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewHttpCookiesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpCookiesSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpCookiesSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *HttpCookiesSection) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *HttpCookiesSection) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
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

// SetHttpOnlyCookies sets the value of HttpOnlyCookies for the instance
func (instance *HttpCookiesSection) SetPropertyHttpOnlyCookies(value bool) (err error) {
	return instance.SetProperty("HttpOnlyCookies", (value))
}

// GetHttpOnlyCookies gets the value of HttpOnlyCookies for the instance
func (instance *HttpCookiesSection) GetPropertyHttpOnlyCookies() (value bool, err error) {
	retValue, err := instance.GetProperty("HttpOnlyCookies")
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

// SetRequireSSL sets the value of RequireSSL for the instance
func (instance *HttpCookiesSection) SetPropertyRequireSSL(value bool) (err error) {
	return instance.SetProperty("RequireSSL", (value))
}

// GetRequireSSL gets the value of RequireSSL for the instance
func (instance *HttpCookiesSection) GetPropertyRequireSSL() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSSL")
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
