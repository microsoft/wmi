// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_TapeDrive struct
type CIM_TapeDrive struct {
	CIM_MediaAccessDevice

	//
	EOTWarningZoneSize uint32

	//
	MaxPartitionCount uint32

	//
	Padding uint32
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
