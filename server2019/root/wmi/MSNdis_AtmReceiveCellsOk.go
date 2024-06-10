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

// MSNdis_AtmReceiveCellsOk struct
type MSNdis_AtmReceiveCellsOk struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdisAtmReceiveCellsOk uint64
}

func NewMSNdis_AtmReceiveCellsOkEx1(instance *cim.WmiInstance) (newInstance *MSNdis_AtmReceiveCellsOk, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmReceiveCellsOk{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_AtmReceiveCellsOkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_AtmReceiveCellsOk, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmReceiveCellsOk{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_AtmReceiveCellsOk) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_AtmReceiveCellsOk) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_AtmReceiveCellsOk) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_AtmReceiveCellsOk) GetPropertyInstanceName() (value string, err error) {
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

// SetNdisAtmReceiveCellsOk sets the value of NdisAtmReceiveCellsOk for the instance
func (instance *MSNdis_AtmReceiveCellsOk) SetPropertyNdisAtmReceiveCellsOk(value uint64) (err error) {
	return instance.SetProperty("NdisAtmReceiveCellsOk", (value))
}

// GetNdisAtmReceiveCellsOk gets the value of NdisAtmReceiveCellsOk for the instance
func (instance *MSNdis_AtmReceiveCellsOk) GetPropertyNdisAtmReceiveCellsOk() (value uint64, err error) {
	retValue, err := instance.GetProperty("NdisAtmReceiveCellsOk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
