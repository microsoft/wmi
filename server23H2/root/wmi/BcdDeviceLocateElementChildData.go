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

// BcdDeviceLocateElementChildData struct
type BcdDeviceLocateElementChildData struct {
	*BcdDeviceLocateData

	// This provides the locate device element.
	Element uint32

	// This is the parent device of this locate device.
	Parent BcdDeviceData
}

func NewBcdDeviceLocateElementChildDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceLocateElementChildData, err error) {
	tmp, err := NewBcdDeviceLocateDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateElementChildData{
		BcdDeviceLocateData: tmp,
	}
	return
}

func NewBcdDeviceLocateElementChildDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceLocateElementChildData, err error) {
	tmp, err := NewBcdDeviceLocateDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateElementChildData{
		BcdDeviceLocateData: tmp,
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *BcdDeviceLocateElementChildData) SetPropertyElement(value uint32) (err error) {
	return instance.SetProperty("Element", (value))
}

// GetElement gets the value of Element for the instance
func (instance *BcdDeviceLocateElementChildData) GetPropertyElement() (value uint32, err error) {
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

// SetParent sets the value of Parent for the instance
func (instance *BcdDeviceLocateElementChildData) SetPropertyParent(value BcdDeviceData) (err error) {
	return instance.SetProperty("Parent", (value))
}

// GetParent gets the value of Parent for the instance
func (instance *BcdDeviceLocateElementChildData) GetPropertyParent() (value BcdDeviceData, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BcdDeviceData)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BcdDeviceData is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BcdDeviceData(valuetmp)

	return
}
