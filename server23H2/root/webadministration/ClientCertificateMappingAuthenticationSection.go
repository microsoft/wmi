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

// ClientCertificateMappingAuthenticationSection struct
type ClientCertificateMappingAuthenticationSection struct {
	*ConfigurationSection

	//
	Enabled bool
}

func NewClientCertificateMappingAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *ClientCertificateMappingAuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ClientCertificateMappingAuthenticationSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewClientCertificateMappingAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClientCertificateMappingAuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClientCertificateMappingAuthenticationSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *ClientCertificateMappingAuthenticationSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ClientCertificateMappingAuthenticationSection) GetPropertyEnabled() (value bool, err error) {
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
