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

// FtpAuthorizationSection struct
type FtpAuthorizationSection struct {
	*ConfigurationSectionWithCollection

	//
	Authorization []FtpAuthorizationRule
}

func NewFtpAuthorizationSectionEx1(instance *cim.WmiInstance) (newInstance *FtpAuthorizationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpAuthorizationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFtpAuthorizationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpAuthorizationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpAuthorizationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAuthorization sets the value of Authorization for the instance
func (instance *FtpAuthorizationSection) SetPropertyAuthorization(value []FtpAuthorizationRule) (err error) {
	return instance.SetProperty("Authorization", (value))
}

// GetAuthorization gets the value of Authorization for the instance
func (instance *FtpAuthorizationSection) GetPropertyAuthorization() (value []FtpAuthorizationRule, err error) {
	retValue, err := instance.GetProperty("Authorization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpAuthorizationRule)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpAuthorizationRule is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpAuthorizationRule(valuetmp))
	}

	return
}
