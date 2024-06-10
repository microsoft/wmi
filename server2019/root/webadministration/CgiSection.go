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

// CgiSection struct
type CgiSection struct {
	*ConfigurationSection

	//
	CreateCGIWithNewConsole bool

	//
	CreateProcessAsUser bool

	//
	Timeout string
}

func NewCgiSectionEx1(instance *cim.WmiInstance) (newInstance *CgiSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CgiSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewCgiSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CgiSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CgiSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCreateCGIWithNewConsole sets the value of CreateCGIWithNewConsole for the instance
func (instance *CgiSection) SetPropertyCreateCGIWithNewConsole(value bool) (err error) {
	return instance.SetProperty("CreateCGIWithNewConsole", (value))
}

// GetCreateCGIWithNewConsole gets the value of CreateCGIWithNewConsole for the instance
func (instance *CgiSection) GetPropertyCreateCGIWithNewConsole() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateCGIWithNewConsole")
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

// SetCreateProcessAsUser sets the value of CreateProcessAsUser for the instance
func (instance *CgiSection) SetPropertyCreateProcessAsUser(value bool) (err error) {
	return instance.SetProperty("CreateProcessAsUser", (value))
}

// GetCreateProcessAsUser gets the value of CreateProcessAsUser for the instance
func (instance *CgiSection) GetPropertyCreateProcessAsUser() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateProcessAsUser")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *CgiSection) SetPropertyTimeout(value string) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *CgiSection) GetPropertyTimeout() (value string, err error) {
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
