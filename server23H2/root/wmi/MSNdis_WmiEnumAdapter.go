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

// MSNdis_WmiEnumAdapter struct
type MSNdis_WmiEnumAdapter struct {
	*MSNdis

	//
	DeviceName string

	//
	Header MSNdis_ObjectHeader

	//
	IfIndex uint32

	//
	NetLuid uint64
}

func NewMSNdis_WmiEnumAdapterEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiEnumAdapter, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiEnumAdapter{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiEnumAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiEnumAdapter, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiEnumAdapter{
		MSNdis: tmp,
	}
	return
}

// SetDeviceName sets the value of DeviceName for the instance
func (instance *MSNdis_WmiEnumAdapter) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", (value))
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *MSNdis_WmiEnumAdapter) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_WmiEnumAdapter) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiEnumAdapter) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetIfIndex sets the value of IfIndex for the instance
func (instance *MSNdis_WmiEnumAdapter) SetPropertyIfIndex(value uint32) (err error) {
	return instance.SetProperty("IfIndex", (value))
}

// GetIfIndex gets the value of IfIndex for the instance
func (instance *MSNdis_WmiEnumAdapter) GetPropertyIfIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("IfIndex")
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

// SetNetLuid sets the value of NetLuid for the instance
func (instance *MSNdis_WmiEnumAdapter) SetPropertyNetLuid(value uint64) (err error) {
	return instance.SetProperty("NetLuid", (value))
}

// GetNetLuid gets the value of NetLuid for the instance
func (instance *MSNdis_WmiEnumAdapter) GetPropertyNetLuid() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetLuid")
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
