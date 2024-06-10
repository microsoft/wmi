// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_PMAdminConfig struct
type MSNdis_PMAdminConfig struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	PMAdminConfigParam MSNdis_PMAdminConfigParam
}

func NewMSNdis_PMAdminConfigEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMAdminConfig, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMAdminConfig{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PMAdminConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PMAdminConfig, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMAdminConfig{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_PMAdminConfig) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_PMAdminConfig) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_PMAdminConfig) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_PMAdminConfig) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetPMAdminConfigParam sets the value of PMAdminConfigParam for the instance
func (instance *MSNdis_PMAdminConfig) SetPropertyPMAdminConfigParam(value MSNdis_PMAdminConfigParam) (err error) {
	return instance.SetProperty("PMAdminConfigParam", (value))
}

// GetPMAdminConfigParam gets the value of PMAdminConfigParam for the instance
func (instance *MSNdis_PMAdminConfig) GetPropertyPMAdminConfigParam() (value MSNdis_PMAdminConfigParam, err error) {
	retValue, err := instance.GetProperty("PMAdminConfigParam")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMAdminConfigParam)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigParam is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMAdminConfigParam(valuetmp)

	return
}
