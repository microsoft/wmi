// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_PMCapabilities struct
type MSNdis_PMCapabilities struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	PMCapabilitiesParam MSNdis_PMCapabilitiesParam
}

func NewMSNdis_PMCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PMCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PMCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_PMCapabilities) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_PMCapabilities) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_PMCapabilities) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_PMCapabilities) GetPropertyInstanceName() (value string, err error) {
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

// SetPMCapabilitiesParam sets the value of PMCapabilitiesParam for the instance
func (instance *MSNdis_PMCapabilities) SetPropertyPMCapabilitiesParam(value MSNdis_PMCapabilitiesParam) (err error) {
	return instance.SetProperty("PMCapabilitiesParam", (value))
}

// GetPMCapabilitiesParam gets the value of PMCapabilitiesParam for the instance
func (instance *MSNdis_PMCapabilities) GetPropertyPMCapabilitiesParam() (value MSNdis_PMCapabilitiesParam, err error) {
	retValue, err := instance.GetProperty("PMCapabilitiesParam")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilitiesParam)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilitiesParam is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilitiesParam(valuetmp)

	return
}
