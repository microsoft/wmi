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

// SystemConfig_V2_OpticalMedia struct
type SystemConfig_V2_OpticalMedia struct {
	*SystemConfig_V2

	//
	BusType uint16

	//
	BytesPerSector uint32

	//
	DeviceName string

	//
	DeviceType uint16

	//
	DiscStatus uint16

	//
	DiskNumber uint16

	//
	DriveLetter string

	//
	FileSystemName string

	//
	LastSessionStatus uint16

	//
	ManufacturerName string

	//
	MediaType uint16

	//
	NextWritableAddress uint64

	//
	NumberOfFreeBlocks uint64

	//
	NumberOfSessions uint32

	//
	NumberOfTracks uint32

	//
	Size uint64

	//
	StartingOffset uint64

	//
	TotalNumberOfBlocks uint64
}

func NewSystemConfig_V2_OpticalMediaEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_OpticalMedia, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_OpticalMedia{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_OpticalMediaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_OpticalMedia, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_OpticalMedia{
		SystemConfig_V2: tmp,
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyBusType(value uint16) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyBusType() (value uint16, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetBytesPerSector sets the value of BytesPerSector for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyBytesPerSector(value uint32) (err error) {
	return instance.SetProperty("BytesPerSector", (value))
}

// GetBytesPerSector gets the value of BytesPerSector for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyBytesPerSector() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesPerSector")
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

// SetDeviceName sets the value of DeviceName for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", (value))
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyDeviceType(value uint16) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyDeviceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("DeviceType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDiscStatus sets the value of DiscStatus for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyDiscStatus(value uint16) (err error) {
	return instance.SetProperty("DiscStatus", (value))
}

// GetDiscStatus gets the value of DiscStatus for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyDiscStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("DiscStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyDiskNumber(value uint16) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyDiskNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDriveLetter sets the value of DriveLetter for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyDriveLetter(value string) (err error) {
	return instance.SetProperty("DriveLetter", (value))
}

// GetDriveLetter gets the value of DriveLetter for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyDriveLetter() (value string, err error) {
	retValue, err := instance.GetProperty("DriveLetter")
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

// SetFileSystemName sets the value of FileSystemName for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyFileSystemName(value string) (err error) {
	return instance.SetProperty("FileSystemName", (value))
}

// GetFileSystemName gets the value of FileSystemName for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyFileSystemName() (value string, err error) {
	retValue, err := instance.GetProperty("FileSystemName")
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

// SetLastSessionStatus sets the value of LastSessionStatus for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyLastSessionStatus(value uint16) (err error) {
	return instance.SetProperty("LastSessionStatus", (value))
}

// GetLastSessionStatus gets the value of LastSessionStatus for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyLastSessionStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastSessionStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetManufacturerName sets the value of ManufacturerName for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyManufacturerName(value string) (err error) {
	return instance.SetProperty("ManufacturerName", (value))
}

// GetManufacturerName gets the value of ManufacturerName for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyManufacturerName() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerName")
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

// SetMediaType sets the value of MediaType for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyMediaType(value uint16) (err error) {
	return instance.SetProperty("MediaType", (value))
}

// GetMediaType gets the value of MediaType for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyMediaType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetNextWritableAddress sets the value of NextWritableAddress for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyNextWritableAddress(value uint64) (err error) {
	return instance.SetProperty("NextWritableAddress", (value))
}

// GetNextWritableAddress gets the value of NextWritableAddress for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyNextWritableAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("NextWritableAddress")
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

// SetNumberOfFreeBlocks sets the value of NumberOfFreeBlocks for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyNumberOfFreeBlocks(value uint64) (err error) {
	return instance.SetProperty("NumberOfFreeBlocks", (value))
}

// GetNumberOfFreeBlocks gets the value of NumberOfFreeBlocks for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyNumberOfFreeBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberOfFreeBlocks")
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

// SetNumberOfSessions sets the value of NumberOfSessions for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyNumberOfSessions(value uint32) (err error) {
	return instance.SetProperty("NumberOfSessions", (value))
}

// GetNumberOfSessions gets the value of NumberOfSessions for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyNumberOfSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSessions")
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

// SetNumberOfTracks sets the value of NumberOfTracks for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyNumberOfTracks(value uint32) (err error) {
	return instance.SetProperty("NumberOfTracks", (value))
}

// GetNumberOfTracks gets the value of NumberOfTracks for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyNumberOfTracks() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfTracks")
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

// SetSize sets the value of Size for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetStartingOffset sets the value of StartingOffset for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyStartingOffset(value uint64) (err error) {
	return instance.SetProperty("StartingOffset", (value))
}

// GetStartingOffset gets the value of StartingOffset for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyStartingOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartingOffset")
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

// SetTotalNumberOfBlocks sets the value of TotalNumberOfBlocks for the instance
func (instance *SystemConfig_V2_OpticalMedia) SetPropertyTotalNumberOfBlocks(value uint64) (err error) {
	return instance.SetProperty("TotalNumberOfBlocks", (value))
}

// GetTotalNumberOfBlocks gets the value of TotalNumberOfBlocks for the instance
func (instance *SystemConfig_V2_OpticalMedia) GetPropertyTotalNumberOfBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalNumberOfBlocks")
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
