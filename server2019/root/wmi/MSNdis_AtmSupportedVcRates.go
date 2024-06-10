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

// MSNdis_AtmSupportedVcRates struct
type MSNdis_AtmSupportedVcRates struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	MaxCellRate uint32

	//
	MinCellRate uint32
}

func NewMSNdis_AtmSupportedVcRatesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_AtmSupportedVcRates, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmSupportedVcRates{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_AtmSupportedVcRatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_AtmSupportedVcRates, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmSupportedVcRates{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_AtmSupportedVcRates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_AtmSupportedVcRates) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_AtmSupportedVcRates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_AtmSupportedVcRates) GetPropertyInstanceName() (value string, err error) {
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

// SetMaxCellRate sets the value of MaxCellRate for the instance
func (instance *MSNdis_AtmSupportedVcRates) SetPropertyMaxCellRate(value uint32) (err error) {
	return instance.SetProperty("MaxCellRate", (value))
}

// GetMaxCellRate gets the value of MaxCellRate for the instance
func (instance *MSNdis_AtmSupportedVcRates) GetPropertyMaxCellRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCellRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMinCellRate sets the value of MinCellRate for the instance
func (instance *MSNdis_AtmSupportedVcRates) SetPropertyMinCellRate(value uint32) (err error) {
	return instance.SetProperty("MinCellRate", (value))
}

// GetMinCellRate gets the value of MinCellRate for the instance
func (instance *MSNdis_AtmSupportedVcRates) GetPropertyMinCellRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinCellRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
