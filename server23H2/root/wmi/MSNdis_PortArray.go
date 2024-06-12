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

// MSNdis_PortArray struct
type MSNdis_PortArray struct {
	*MSNdis

	//
	ElementSize uint32

	//
	Header MSNdis_ObjectHeader

	//
	NumberOfPorts uint32

	//
	OffsetFirstPort uint32

	//
	Port []MSNdis_PortChar
}

func NewMSNdis_PortArrayEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PortArray, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PortArray{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PortArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PortArray, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PortArray{
		MSNdis: tmp,
	}
	return
}

// SetElementSize sets the value of ElementSize for the instance
func (instance *MSNdis_PortArray) SetPropertyElementSize(value uint32) (err error) {
	return instance.SetProperty("ElementSize", (value))
}

// GetElementSize gets the value of ElementSize for the instance
func (instance *MSNdis_PortArray) GetPropertyElementSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ElementSize")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PortArray) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PortArray) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetNumberOfPorts sets the value of NumberOfPorts for the instance
func (instance *MSNdis_PortArray) SetPropertyNumberOfPorts(value uint32) (err error) {
	return instance.SetProperty("NumberOfPorts", (value))
}

// GetNumberOfPorts gets the value of NumberOfPorts for the instance
func (instance *MSNdis_PortArray) GetPropertyNumberOfPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPorts")
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

// SetOffsetFirstPort sets the value of OffsetFirstPort for the instance
func (instance *MSNdis_PortArray) SetPropertyOffsetFirstPort(value uint32) (err error) {
	return instance.SetProperty("OffsetFirstPort", (value))
}

// GetOffsetFirstPort gets the value of OffsetFirstPort for the instance
func (instance *MSNdis_PortArray) GetPropertyOffsetFirstPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("OffsetFirstPort")
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

// SetPort sets the value of Port for the instance
func (instance *MSNdis_PortArray) SetPropertyPort(value []MSNdis_PortChar) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *MSNdis_PortArray) GetPropertyPort() (value []MSNdis_PortChar, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSNdis_PortChar)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSNdis_PortChar is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSNdis_PortChar(valuetmp))
	}

	return
}
