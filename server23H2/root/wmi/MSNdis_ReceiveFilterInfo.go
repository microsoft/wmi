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

// MSNdis_ReceiveFilterInfo struct
type MSNdis_ReceiveFilterInfo struct {
	*MSNdis

	//
	FilterId uint32

	//
	FilterType uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader
}

func NewMSNdis_ReceiveFilterInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterInfo, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterInfo{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterInfo, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterInfo{
		MSNdis: tmp,
	}
	return
}

// SetFilterId sets the value of FilterId for the instance
func (instance *MSNdis_ReceiveFilterInfo) SetPropertyFilterId(value uint32) (err error) {
	return instance.SetProperty("FilterId", (value))
}

// GetFilterId gets the value of FilterId for the instance
func (instance *MSNdis_ReceiveFilterInfo) GetPropertyFilterId() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterId")
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

// SetFilterType sets the value of FilterType for the instance
func (instance *MSNdis_ReceiveFilterInfo) SetPropertyFilterType(value uint32) (err error) {
	return instance.SetProperty("FilterType", (value))
}

// GetFilterType gets the value of FilterType for the instance
func (instance *MSNdis_ReceiveFilterInfo) GetPropertyFilterType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterType")
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
func (instance *MSNdis_ReceiveFilterInfo) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterInfo) GetPropertyFlags() (value uint32, err error) {
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
func (instance *MSNdis_ReceiveFilterInfo) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterInfo) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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
