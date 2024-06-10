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

// WNFNameSubRundown struct
type WNFNameSubRundown struct {
	*WNFTrace

	//
	NameSub uint32

	//
	StateName uint64
}

func NewWNFNameSubRundownEx1(instance *cim.WmiInstance) (newInstance *WNFNameSubRundown, err error) {
	tmp, err := NewWNFTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WNFNameSubRundown{
		WNFTrace: tmp,
	}
	return
}

func NewWNFNameSubRundownEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WNFNameSubRundown, err error) {
	tmp, err := NewWNFTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WNFNameSubRundown{
		WNFTrace: tmp,
	}
	return
}

// SetNameSub sets the value of NameSub for the instance
func (instance *WNFNameSubRundown) SetPropertyNameSub(value uint32) (err error) {
	return instance.SetProperty("NameSub", (value))
}

// GetNameSub gets the value of NameSub for the instance
func (instance *WNFNameSubRundown) GetPropertyNameSub() (value uint32, err error) {
	retValue, err := instance.GetProperty("NameSub")
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
func (instance *WNFNameSubRundown) SetPropertyStateName(value uint64) (err error) {
	return instance.SetProperty("StateName", (value))
}

// GetStateName gets the value of StateName for the instance
func (instance *WNFNameSubRundown) GetPropertyStateName() (value uint64, err error) {
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
