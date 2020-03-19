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

// CIM_TapeDrive struct
type CIM_TapeDrive struct {
	*CIM_MediaAccessDevice

	//
	EOTWarningZoneSize uint32

	//
	MaxPartitionCount uint32

	//
	Padding uint32
}

func NewCIM_TapeDriveEx1(instance *cim.WmiInstance) (newInstance *CIM_TapeDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_TapeDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}

func NewCIM_TapeDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_TapeDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_TapeDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}

// SetEOTWarningZoneSize sets the value of EOTWarningZoneSize for the instance
func (instance *CIM_TapeDrive) SetPropertyEOTWarningZoneSize(value uint32) (err error) {
	return instance.SetProperty("EOTWarningZoneSize", value)
}

// GetEOTWarningZoneSize gets the value of EOTWarningZoneSize for the instance
func (instance *CIM_TapeDrive) GetPropertyEOTWarningZoneSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("EOTWarningZoneSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPartitionCount sets the value of MaxPartitionCount for the instance
func (instance *CIM_TapeDrive) SetPropertyMaxPartitionCount(value uint32) (err error) {
	return instance.SetProperty("MaxPartitionCount", value)
}

// GetMaxPartitionCount gets the value of MaxPartitionCount for the instance
func (instance *CIM_TapeDrive) GetPropertyMaxPartitionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPartitionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPadding sets the value of Padding for the instance
func (instance *CIM_TapeDrive) SetPropertyPadding(value uint32) (err error) {
	return instance.SetProperty("Padding", value)
}

// GetPadding gets the value of Padding for the instance
func (instance *CIM_TapeDrive) GetPropertyPadding() (value uint32, err error) {
	retValue, err := instance.GetProperty("Padding")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
