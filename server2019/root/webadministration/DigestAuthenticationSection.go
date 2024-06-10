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

// DigestAuthenticationSection struct
type DigestAuthenticationSection struct {
	*ConfigurationSection

	//
	Enabled bool

	//
	Realm string
}

func NewDigestAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *DigestAuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DigestAuthenticationSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewDigestAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DigestAuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DigestAuthenticationSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *DigestAuthenticationSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *DigestAuthenticationSection) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetRealm sets the value of Realm for the instance
func (instance *DigestAuthenticationSection) SetPropertyRealm(value string) (err error) {
	return instance.SetProperty("Realm", (value))
}

// GetRealm gets the value of Realm for the instance
func (instance *DigestAuthenticationSection) GetPropertyRealm() (value string, err error) {
	retValue, err := instance.GetProperty("Realm")
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
