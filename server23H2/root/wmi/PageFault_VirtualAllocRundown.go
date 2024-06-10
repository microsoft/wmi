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

// PageFault_VirtualAllocRundown struct
type PageFault_VirtualAllocRundown struct {
	*PageFault_V2

	//
	BaseAddress uint32

	//
	CommitSizeInBytes interface{}

	//
	ProcessId uint32

	//
	RegionSize interface{}
}

func NewPageFault_VirtualAllocRundownEx1(instance *cim.WmiInstance) (newInstance *PageFault_VirtualAllocRundown, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_VirtualAllocRundown{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_VirtualAllocRundownEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_VirtualAllocRundown, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_VirtualAllocRundown{
		PageFault_V2: tmp,
	}
	return
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *PageFault_VirtualAllocRundown) SetPropertyBaseAddress(value uint32) (err error) {
	return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *PageFault_VirtualAllocRundown) GetPropertyBaseAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("BaseAddress")
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

// SetCommitSizeInBytes sets the value of CommitSizeInBytes for the instance
func (instance *PageFault_VirtualAllocRundown) SetPropertyCommitSizeInBytes(value interface{}) (err error) {
	return instance.SetProperty("CommitSizeInBytes", (value))
}

// GetCommitSizeInBytes gets the value of CommitSizeInBytes for the instance
func (instance *PageFault_VirtualAllocRundown) GetPropertyCommitSizeInBytes() (value interface{}, err error) {
	retValue, err := instance.GetProperty("CommitSizeInBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *PageFault_VirtualAllocRundown) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *PageFault_VirtualAllocRundown) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetRegionSize sets the value of RegionSize for the instance
func (instance *PageFault_VirtualAllocRundown) SetPropertyRegionSize(value interface{}) (err error) {
	return instance.SetProperty("RegionSize", (value))
}

// GetRegionSize gets the value of RegionSize for the instance
func (instance *PageFault_VirtualAllocRundown) GetPropertyRegionSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("RegionSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
