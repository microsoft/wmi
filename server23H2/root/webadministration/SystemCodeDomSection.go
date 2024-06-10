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

// SystemCodeDomSection struct
type SystemCodeDomSection struct {
	*ConfigurationSectionWithCollection

	//
	Compilers CompilerSettings
}

func NewSystemCodeDomSectionEx1(instance *cim.WmiInstance) (newInstance *SystemCodeDomSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemCodeDomSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSystemCodeDomSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemCodeDomSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemCodeDomSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetCompilers sets the value of Compilers for the instance
func (instance *SystemCodeDomSection) SetPropertyCompilers(value CompilerSettings) (err error) {
	return instance.SetProperty("Compilers", (value))
}

// GetCompilers gets the value of Compilers for the instance
func (instance *SystemCodeDomSection) GetPropertyCompilers() (value CompilerSettings, err error) {
	retValue, err := instance.GetProperty("Compilers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CompilerSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CompilerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CompilerSettings(valuetmp)

	return
}
