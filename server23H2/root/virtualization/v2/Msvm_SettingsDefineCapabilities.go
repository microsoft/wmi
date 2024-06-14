// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SettingsDefineCapabilities struct
type Msvm_SettingsDefineCapabilities struct {
	*CIM_SettingsDefineCapabilities

	//
	SupportStatement uint16
}

func NewMsvm_SettingsDefineCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_SettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SettingsDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}

func NewMsvm_SettingsDefineCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SettingsDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}

// SetSupportStatement sets the value of SupportStatement for the instance
func (instance *Msvm_SettingsDefineCapabilities) SetPropertySupportStatement(value uint16) (err error) {
	return instance.SetProperty("SupportStatement", (value))
}

// GetSupportStatement gets the value of SupportStatement for the instance
func (instance *Msvm_SettingsDefineCapabilities) GetPropertySupportStatement() (value uint16, err error) {
	retValue, err := instance.GetProperty("SupportStatement")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
