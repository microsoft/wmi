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

// CIM_PhysicalMedia struct
type CIM_PhysicalMedia struct {
	*CIM_PhysicalComponent

	//
	Capacity uint64

	//
	CleanerMedia bool

	//
	MediaDescription string

	//
	MediaType uint16

	//
	WriteProtectOn bool
}

func NewCIM_PhysicalMediaEx1(instance *cim.WmiInstance) (newInstance *CIM_PhysicalMedia, err error) {
	tmp, err := NewCIM_PhysicalComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalMedia{
		CIM_PhysicalComponent: tmp,
	}
	return
}

func NewCIM_PhysicalMediaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PhysicalMedia, err error) {
	tmp, err := NewCIM_PhysicalComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalMedia{
		CIM_PhysicalComponent: tmp,
	}
	return
}

// SetCapacity sets the value of Capacity for the instance
func (instance *CIM_PhysicalMedia) SetPropertyCapacity(value uint64) (err error) {
	return instance.SetProperty("Capacity", value)
}

// GetCapacity gets the value of Capacity for the instance
func (instance *CIM_PhysicalMedia) GetPropertyCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("Capacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCleanerMedia sets the value of CleanerMedia for the instance
func (instance *CIM_PhysicalMedia) SetPropertyCleanerMedia(value bool) (err error) {
	return instance.SetProperty("CleanerMedia", value)
}

// GetCleanerMedia gets the value of CleanerMedia for the instance
func (instance *CIM_PhysicalMedia) GetPropertyCleanerMedia() (value bool, err error) {
	retValue, err := instance.GetProperty("CleanerMedia")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaDescription sets the value of MediaDescription for the instance
func (instance *CIM_PhysicalMedia) SetPropertyMediaDescription(value string) (err error) {
	return instance.SetProperty("MediaDescription", value)
}

// GetMediaDescription gets the value of MediaDescription for the instance
func (instance *CIM_PhysicalMedia) GetPropertyMediaDescription() (value string, err error) {
	retValue, err := instance.GetProperty("MediaDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *CIM_PhysicalMedia) SetPropertyMediaType(value uint16) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *CIM_PhysicalMedia) GetPropertyMediaType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteProtectOn sets the value of WriteProtectOn for the instance
func (instance *CIM_PhysicalMedia) SetPropertyWriteProtectOn(value bool) (err error) {
	return instance.SetProperty("WriteProtectOn", value)
}

// GetWriteProtectOn gets the value of WriteProtectOn for the instance
func (instance *CIM_PhysicalMedia) GetPropertyWriteProtectOn() (value bool, err error) {
	retValue, err := instance.GetProperty("WriteProtectOn")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
