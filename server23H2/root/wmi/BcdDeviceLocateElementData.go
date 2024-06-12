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

// BcdDeviceLocateElementData struct
type BcdDeviceLocateElementData struct {
	*BcdDeviceLocateData

	// This provides the locate device element.
	Element uint32
}

func NewBcdDeviceLocateElementDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceLocateElementData, err error) {
	tmp, err := NewBcdDeviceLocateDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateElementData{
		BcdDeviceLocateData: tmp,
	}
	return
}

func NewBcdDeviceLocateElementDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceLocateElementData, err error) {
	tmp, err := NewBcdDeviceLocateDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateElementData{
		BcdDeviceLocateData: tmp,
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *BcdDeviceLocateElementData) SetPropertyElement(value uint32) (err error) {
	return instance.SetProperty("Element", (value))
}

// GetElement gets the value of Element for the instance
func (instance *BcdDeviceLocateElementData) GetPropertyElement() (value uint32, err error) {
	retValue, err := instance.GetProperty("Element")
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
