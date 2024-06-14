// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_DeviceWakeOnMagicPacketOnly struct
type MSNdis_DeviceWakeOnMagicPacketOnly struct {
	*MSNdis

	//
	Active bool

	//
	EnableWakeOnMagicPacketOnly bool

	//
	InstanceName string
}

func NewMSNdis_DeviceWakeOnMagicPacketOnlyEx1(instance *cim.WmiInstance) (newInstance *MSNdis_DeviceWakeOnMagicPacketOnly, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_DeviceWakeOnMagicPacketOnly{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_DeviceWakeOnMagicPacketOnlyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_DeviceWakeOnMagicPacketOnly, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_DeviceWakeOnMagicPacketOnly{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) GetPropertyActive() (value bool, err error) {
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

// SetEnableWakeOnMagicPacketOnly sets the value of EnableWakeOnMagicPacketOnly for the instance
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) SetPropertyEnableWakeOnMagicPacketOnly(value bool) (err error) {
	return instance.SetProperty("EnableWakeOnMagicPacketOnly", (value))
}

// GetEnableWakeOnMagicPacketOnly gets the value of EnableWakeOnMagicPacketOnly for the instance
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) GetPropertyEnableWakeOnMagicPacketOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableWakeOnMagicPacketOnly")
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
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_DeviceWakeOnMagicPacketOnly) GetPropertyInstanceName() (value string, err error) {
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
