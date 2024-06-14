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

// TransactionsMachineSettingsSection struct
type TransactionsMachineSettingsSection struct {
	*ConfigurationSection

	//
	MaxTimeout string
}

func NewTransactionsMachineSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *TransactionsMachineSettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransactionsMachineSettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewTransactionsMachineSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransactionsMachineSettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransactionsMachineSettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetMaxTimeout sets the value of MaxTimeout for the instance
func (instance *TransactionsMachineSettingsSection) SetPropertyMaxTimeout(value string) (err error) {
	return instance.SetProperty("MaxTimeout", (value))
}

// GetMaxTimeout gets the value of MaxTimeout for the instance
func (instance *TransactionsMachineSettingsSection) GetPropertyMaxTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("MaxTimeout")
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
