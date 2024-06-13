// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSWmi_MofData struct
type MSWmi_MofData struct {
	*MS_WmiInternal

	//
	Active bool

	//
	BinaryMofData []uint8

	//
	InstanceName string

	//
	Size uint32

	//
	Unused1 uint32

	//
	Unused2 uint32

	//
	Unused4 uint32
}

func NewMSWmi_MofDataEx1(instance *cim.WmiInstance) (newInstance *MSWmi_MofData, err error) {
	tmp, err := NewMS_WmiInternalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSWmi_MofData{
		MS_WmiInternal: tmp,
	}
	return
}

func NewMSWmi_MofDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSWmi_MofData, err error) {
	tmp, err := NewMS_WmiInternalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSWmi_MofData{
		MS_WmiInternal: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSWmi_MofData) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSWmi_MofData) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBinaryMofData sets the value of BinaryMofData for the instance
func (instance *MSWmi_MofData) SetPropertyBinaryMofData(value []uint8) (err error) {
	return instance.SetProperty("BinaryMofData", (value))
}

// GetBinaryMofData gets the value of BinaryMofData for the instance
func (instance *MSWmi_MofData) GetPropertyBinaryMofData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("BinaryMofData")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSWmi_MofData) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSWmi_MofData) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetSize sets the value of Size for the instance
func (instance *MSWmi_MofData) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSWmi_MofData) GetPropertySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetUnused1 sets the value of Unused1 for the instance
func (instance *MSWmi_MofData) SetPropertyUnused1(value uint32) (err error) {
	return instance.SetProperty("Unused1", (value))
}

// GetUnused1 gets the value of Unused1 for the instance
func (instance *MSWmi_MofData) GetPropertyUnused1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Unused1")
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

// SetUnused2 sets the value of Unused2 for the instance
func (instance *MSWmi_MofData) SetPropertyUnused2(value uint32) (err error) {
	return instance.SetProperty("Unused2", (value))
}

// GetUnused2 gets the value of Unused2 for the instance
func (instance *MSWmi_MofData) GetPropertyUnused2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Unused2")
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

// SetUnused4 sets the value of Unused4 for the instance
func (instance *MSWmi_MofData) SetPropertyUnused4(value uint32) (err error) {
	return instance.SetProperty("Unused4", (value))
}

// GetUnused4 gets the value of Unused4 for the instance
func (instance *MSWmi_MofData) GetPropertyUnused4() (value uint32, err error) {
	retValue, err := instance.GetProperty("Unused4")
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
