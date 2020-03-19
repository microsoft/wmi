// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PhysicalMemoryArray struct
type Win32_PhysicalMemoryArray struct {
	*CIM_PhysicalPackage

	//
	Location uint16

	//
	MaxCapacity uint32

	//
	MaxCapacityEx uint64

	//
	MemoryDevices uint16

	//
	MemoryErrorCorrection uint16

	//
	Use uint16
}

func NewWin32_PhysicalMemoryArrayEx1(instance *cim.WmiInstance) (newInstance *Win32_PhysicalMemoryArray, err error) {
	tmp, err := NewCIM_PhysicalPackageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PhysicalMemoryArray{
		CIM_PhysicalPackage: tmp,
	}
	return
}

func NewWin32_PhysicalMemoryArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PhysicalMemoryArray, err error) {
	tmp, err := NewCIM_PhysicalPackageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PhysicalMemoryArray{
		CIM_PhysicalPackage: tmp,
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyLocation(value uint16) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyLocation() (value uint16, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCapacity sets the value of MaxCapacity for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyMaxCapacity(value uint32) (err error) {
	return instance.SetProperty("MaxCapacity", value)
}

// GetMaxCapacity gets the value of MaxCapacity for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyMaxCapacity() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCapacityEx sets the value of MaxCapacityEx for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyMaxCapacityEx(value uint64) (err error) {
	return instance.SetProperty("MaxCapacityEx", value)
}

// GetMaxCapacityEx gets the value of MaxCapacityEx for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyMaxCapacityEx() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxCapacityEx")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryDevices sets the value of MemoryDevices for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyMemoryDevices(value uint16) (err error) {
	return instance.SetProperty("MemoryDevices", value)
}

// GetMemoryDevices gets the value of MemoryDevices for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyMemoryDevices() (value uint16, err error) {
	retValue, err := instance.GetProperty("MemoryDevices")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryErrorCorrection sets the value of MemoryErrorCorrection for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyMemoryErrorCorrection(value uint16) (err error) {
	return instance.SetProperty("MemoryErrorCorrection", value)
}

// GetMemoryErrorCorrection gets the value of MemoryErrorCorrection for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyMemoryErrorCorrection() (value uint16, err error) {
	retValue, err := instance.GetProperty("MemoryErrorCorrection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUse sets the value of Use for the instance
func (instance *Win32_PhysicalMemoryArray) SetPropertyUse(value uint16) (err error) {
	return instance.SetProperty("Use", value)
}

// GetUse gets the value of Use for the instance
func (instance *Win32_PhysicalMemoryArray) GetPropertyUse() (value uint16, err error) {
	retValue, err := instance.GetProperty("Use")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
