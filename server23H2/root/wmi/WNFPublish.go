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

// WNFPublish struct
type WNFPublish struct {
	*WNFTrace

	//
	DataLength uint32

	//
	StateName uint64
}

func NewWNFPublishEx1(instance *cim.WmiInstance) (newInstance *WNFPublish, err error) {
	tmp, err := NewWNFTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WNFPublish{
		WNFTrace: tmp,
	}
	return
}

func NewWNFPublishEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WNFPublish, err error) {
	tmp, err := NewWNFTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WNFPublish{
		WNFTrace: tmp,
	}
	return
}

// SetDataLength sets the value of DataLength for the instance
func (instance *WNFPublish) SetPropertyDataLength(value uint32) (err error) {
	return instance.SetProperty("DataLength", (value))
}

// GetDataLength gets the value of DataLength for the instance
func (instance *WNFPublish) GetPropertyDataLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataLength")
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

// SetStateName sets the value of StateName for the instance
func (instance *WNFPublish) SetPropertyStateName(value uint64) (err error) {
	return instance.SetProperty("StateName", (value))
}

// GetStateName gets the value of StateName for the instance
func (instance *WNFPublish) GetPropertyStateName() (value uint64, err error) {
	retValue, err := instance.GetProperty("StateName")
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
