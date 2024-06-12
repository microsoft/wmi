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

// BcdDevicePartitionData struct
type BcdDevicePartitionData struct {
	*BcdDeviceData

	// This is the device path.
	Path string
}

func NewBcdDevicePartitionDataEx1(instance *cim.WmiInstance) (newInstance *BcdDevicePartitionData, err error) {
	tmp, err := NewBcdDeviceDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDevicePartitionData{
		BcdDeviceData: tmp,
	}
	return
}

func NewBcdDevicePartitionDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDevicePartitionData, err error) {
	tmp, err := NewBcdDeviceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDevicePartitionData{
		BcdDeviceData: tmp,
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *BcdDevicePartitionData) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *BcdDevicePartitionData) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
