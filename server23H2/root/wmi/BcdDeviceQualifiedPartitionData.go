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

// BcdDeviceQualifiedPartitionData struct
type BcdDeviceQualifiedPartitionData struct {
	*BcdDeviceData

	// This is the disk signature.
	DiskSignature string

	// This is the partition identifier.
	PartitionIdentifier string

	// This is the type of partition, MBR or GPT..
	PartitionStyle BcdDeviceQualifiedPartitionData_PartitionStyle
}

func NewBcdDeviceQualifiedPartitionDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceQualifiedPartitionData, err error) {
	tmp, err := NewBcdDeviceDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceQualifiedPartitionData{
		BcdDeviceData: tmp,
	}
	return
}

func NewBcdDeviceQualifiedPartitionDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceQualifiedPartitionData, err error) {
	tmp, err := NewBcdDeviceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceQualifiedPartitionData{
		BcdDeviceData: tmp,
	}
	return
}

// SetDiskSignature sets the value of DiskSignature for the instance
func (instance *BcdDeviceQualifiedPartitionData) SetPropertyDiskSignature(value string) (err error) {
	return instance.SetProperty("DiskSignature", (value))
}

// GetDiskSignature gets the value of DiskSignature for the instance
func (instance *BcdDeviceQualifiedPartitionData) GetPropertyDiskSignature() (value string, err error) {
	retValue, err := instance.GetProperty("DiskSignature")
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

// SetPartitionIdentifier sets the value of PartitionIdentifier for the instance
func (instance *BcdDeviceQualifiedPartitionData) SetPropertyPartitionIdentifier(value string) (err error) {
	return instance.SetProperty("PartitionIdentifier", (value))
}

// GetPartitionIdentifier gets the value of PartitionIdentifier for the instance
func (instance *BcdDeviceQualifiedPartitionData) GetPropertyPartitionIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("PartitionIdentifier")
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

// SetPartitionStyle sets the value of PartitionStyle for the instance
func (instance *BcdDeviceQualifiedPartitionData) SetPropertyPartitionStyle(value BcdDeviceQualifiedPartitionData_PartitionStyle) (err error) {
	return instance.SetProperty("PartitionStyle", (value))
}

// GetPartitionStyle gets the value of PartitionStyle for the instance
func (instance *BcdDeviceQualifiedPartitionData) GetPropertyPartitionStyle() (value BcdDeviceQualifiedPartitionData_PartitionStyle, err error) {
	retValue, err := instance.GetProperty("PartitionStyle")
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

	value = BcdDeviceQualifiedPartitionData_PartitionStyle(valuetmp)

	return
}
