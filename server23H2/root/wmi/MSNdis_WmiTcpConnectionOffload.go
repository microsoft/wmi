// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_WmiTcpConnectionOffload struct
type MSNdis_WmiTcpConnectionOffload struct {
	*MSNdis

	//
	Encapsulation uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	SupportIp4 uint32

	//
	SupportIp6 uint32

	//
	SupportIp6ExtensionHeaders uint32

	//
	SupportSack uint32

	//
	TcpConnectionOffloadCapacity uint32
}

func NewMSNdis_WmiTcpConnectionOffloadEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpConnectionOffload, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpConnectionOffload{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiTcpConnectionOffloadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiTcpConnectionOffload, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiTcpConnectionOffload{
		MSNdis: tmp,
	}
	return
}

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertyEncapsulation(value uint32) (err error) {
	return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertyEncapsulation() (value uint32, err error) {
	retValue, err := instance.GetProperty("Encapsulation")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetSupportIp4 sets the value of SupportIp4 for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertySupportIp4(value uint32) (err error) {
	return instance.SetProperty("SupportIp4", (value))
}

// GetSupportIp4 gets the value of SupportIp4 for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertySupportIp4() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportIp4")
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

// SetSupportIp6 sets the value of SupportIp6 for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertySupportIp6(value uint32) (err error) {
	return instance.SetProperty("SupportIp6", (value))
}

// GetSupportIp6 gets the value of SupportIp6 for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertySupportIp6() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportIp6")
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

// SetSupportIp6ExtensionHeaders sets the value of SupportIp6ExtensionHeaders for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertySupportIp6ExtensionHeaders(value uint32) (err error) {
	return instance.SetProperty("SupportIp6ExtensionHeaders", (value))
}

// GetSupportIp6ExtensionHeaders gets the value of SupportIp6ExtensionHeaders for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertySupportIp6ExtensionHeaders() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportIp6ExtensionHeaders")
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

// SetSupportSack sets the value of SupportSack for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertySupportSack(value uint32) (err error) {
	return instance.SetProperty("SupportSack", (value))
}

// GetSupportSack gets the value of SupportSack for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertySupportSack() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportSack")
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

// SetTcpConnectionOffloadCapacity sets the value of TcpConnectionOffloadCapacity for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) SetPropertyTcpConnectionOffloadCapacity(value uint32) (err error) {
	return instance.SetProperty("TcpConnectionOffloadCapacity", (value))
}

// GetTcpConnectionOffloadCapacity gets the value of TcpConnectionOffloadCapacity for the instance
func (instance *MSNdis_WmiTcpConnectionOffload) GetPropertyTcpConnectionOffloadCapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpConnectionOffloadCapacity")
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
