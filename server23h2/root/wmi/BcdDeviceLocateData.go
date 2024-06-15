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

// BcdDeviceLocateData struct
type BcdDeviceLocateData struct {
	*BcdDeviceData

	// This provides the locate device type.
	Type BcdDeviceLocateData_Type
}

func NewBcdDeviceLocateDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceLocateData, err error) {
	tmp, err := NewBcdDeviceDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateData{
		BcdDeviceData: tmp,
	}
	return
}

func NewBcdDeviceLocateDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceLocateData, err error) {
	tmp, err := NewBcdDeviceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceLocateData{
		BcdDeviceData: tmp,
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *BcdDeviceLocateData) SetPropertyType(value BcdDeviceLocateData_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *BcdDeviceLocateData) GetPropertyType() (value BcdDeviceLocateData_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BcdDeviceLocateData_Type(valuetmp)

	return
}
