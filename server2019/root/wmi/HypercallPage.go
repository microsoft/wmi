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

// HypercallPage struct
type HypercallPage struct {
	*Image_V2

	//
	HypercallPageVa uint32
}

func NewHypercallPageEx1(instance *cim.WmiInstance) (newInstance *HypercallPage, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HypercallPage{
		Image_V2: tmp,
	}
	return
}

func NewHypercallPageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HypercallPage, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HypercallPage{
		Image_V2: tmp,
	}
	return
}

// SetHypercallPageVa sets the value of HypercallPageVa for the instance
func (instance *HypercallPage) SetPropertyHypercallPageVa(value uint32) (err error) {
	return instance.SetProperty("HypercallPageVa", (value))
}

// GetHypercallPageVa gets the value of HypercallPageVa for the instance
func (instance *HypercallPage) GetPropertyHypercallPageVa() (value uint32, err error) {
	retValue, err := instance.GetProperty("HypercallPageVa")
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
