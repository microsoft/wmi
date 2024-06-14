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

// LoaderBaseEvent struct
type LoaderBaseEvent struct {
	*Image_V2

	//
	BaseAddress uint64
}

func NewLoaderBaseEventEx1(instance *cim.WmiInstance) (newInstance *LoaderBaseEvent, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderBaseEvent{
		Image_V2: tmp,
	}
	return
}

func NewLoaderBaseEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderBaseEvent, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderBaseEvent{
		Image_V2: tmp,
	}
	return
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *LoaderBaseEvent) SetPropertyBaseAddress(value uint64) (err error) {
	return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *LoaderBaseEvent) GetPropertyBaseAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("BaseAddress")
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
