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

// BcdDeviceUnknownData struct
type BcdDeviceUnknownData struct {
	*BcdDeviceData

	// This is the binary data of the unknown device element.
	Data []uint8
}

func NewBcdDeviceUnknownDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceUnknownData, err error) {
	tmp, err := NewBcdDeviceDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceUnknownData{
		BcdDeviceData: tmp,
	}
	return
}

func NewBcdDeviceUnknownDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceUnknownData, err error) {
	tmp, err := NewBcdDeviceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceUnknownData{
		BcdDeviceData: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *BcdDeviceUnknownData) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *BcdDeviceUnknownData) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
