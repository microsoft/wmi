// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_TapeDrive struct
type CIM_TapeDrive struct {
	*CIM_MediaAccessDevice

	// EOTWarningZoneSize indicates the size, in bytes, of the area designated as 'end of tape'. Access in this area generates an 'end of tape' warning.
	EOTWarningZoneSize uint32

	// Maximum partition count for the TapeDrive.
	MaxPartitionCount uint32

	// Time in milliseconds to move from the most physically distant point on the tape to the beginning.
	MaxRewindTime uint64

	// Padding indicates the number of bytes inserted between blocks on a tape Media.
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
	return instance.SetProperty("EOTWarningZoneSize", (value))
}

// GetEOTWarningZoneSize gets the value of EOTWarningZoneSize for the instance
func (instance *CIM_TapeDrive) GetPropertyEOTWarningZoneSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("EOTWarningZoneSize")
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

// SetMaxPartitionCount sets the value of MaxPartitionCount for the instance
func (instance *CIM_TapeDrive) SetPropertyMaxPartitionCount(value uint32) (err error) {
	return instance.SetProperty("MaxPartitionCount", (value))
}

// GetMaxPartitionCount gets the value of MaxPartitionCount for the instance
func (instance *CIM_TapeDrive) GetPropertyMaxPartitionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPartitionCount")
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

// SetMaxRewindTime sets the value of MaxRewindTime for the instance
func (instance *CIM_TapeDrive) SetPropertyMaxRewindTime(value uint64) (err error) {
	return instance.SetProperty("MaxRewindTime", (value))
}

// GetMaxRewindTime gets the value of MaxRewindTime for the instance
func (instance *CIM_TapeDrive) GetPropertyMaxRewindTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxRewindTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPadding sets the value of Padding for the instance
func (instance *CIM_TapeDrive) SetPropertyPadding(value uint32) (err error) {
	return instance.SetProperty("Padding", (value))
}

// GetPadding gets the value of Padding for the instance
func (instance *CIM_TapeDrive) GetPropertyPadding() (value uint32, err error) {
	retValue, err := instance.GetProperty("Padding")
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
