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

// SystemConfig_V1_LogDisk struct
type SystemConfig_V1_LogDisk struct {
	*SystemConfig_V1

	//
	BytesPerSector uint32

	//
	DiskNumber uint32

	//
	DriveLetterString []byte

	//
	DriveType uint32

	//
	FileSystem []byte

	//
	NumberOfFreeClusters int64

	//
	Pad1 uint32

	//
	Pad2 uint32

	//
	PartitionNumber uint32

	//
	PartitionSize uint64

	//
	SectorsPerCluster uint32

	//
	Size uint32

	//
	StartOffset uint64

	//
	TotalNumberOfClusters int64

	//
	VolumeExt uint32
}

func NewSystemConfig_V1_LogDiskEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V1_LogDisk, err error) {
	tmp, err := NewSystemConfig_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V1_LogDisk{
		SystemConfig_V1: tmp,
	}
	return
}

func NewSystemConfig_V1_LogDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V1_LogDisk, err error) {
	tmp, err := NewSystemConfig_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V1_LogDisk{
		SystemConfig_V1: tmp,
	}
	return
}

// SetBytesPerSector sets the value of BytesPerSector for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyBytesPerSector(value uint32) (err error) {
	return instance.SetProperty("BytesPerSector", (value))
}

// GetBytesPerSector gets the value of BytesPerSector for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyBytesPerSector() (value uint32, err error) {
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

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
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

// SetDriveLetterString sets the value of DriveLetterString for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyDriveLetterString(value []byte) (err error) {
	return instance.SetProperty("DriveLetterString", (value))
}

// GetDriveLetterString gets the value of DriveLetterString for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyDriveLetterString() (value []byte, err error) {
	retValue, err := instance.GetProperty("DriveLetterString")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetDriveType sets the value of DriveType for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyDriveType(value uint32) (err error) {
	return instance.SetProperty("DriveType", (value))
}

// GetDriveType gets the value of DriveType for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyDriveType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveType")
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

// SetFileSystem sets the value of FileSystem for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyFileSystem(value []byte) (err error) {
	return instance.SetProperty("FileSystem", (value))
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyFileSystem() (value []byte, err error) {
	retValue, err := instance.GetProperty("FileSystem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetNumberOfFreeClusters sets the value of NumberOfFreeClusters for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyNumberOfFreeClusters(value int64) (err error) {
	return instance.SetProperty("NumberOfFreeClusters", (value))
}

// GetNumberOfFreeClusters gets the value of NumberOfFreeClusters for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyNumberOfFreeClusters() (value int64, err error) {
	retValue, err := instance.GetProperty("NumberOfFreeClusters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetPad1 sets the value of Pad1 for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyPad1(value uint32) (err error) {
	return instance.SetProperty("Pad1", (value))
}

// GetPad1 gets the value of Pad1 for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyPad1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Pad1")
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

// SetPad2 sets the value of Pad2 for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyPad2(value uint32) (err error) {
	return instance.SetProperty("Pad2", (value))
}

// GetPad2 gets the value of Pad2 for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyPad2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Pad2")
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

// SetPartitionNumber sets the value of PartitionNumber for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyPartitionNumber(value uint32) (err error) {
	return instance.SetProperty("PartitionNumber", (value))
}

// GetPartitionNumber gets the value of PartitionNumber for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyPartitionNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PartitionNumber")
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

// SetPartitionSize sets the value of PartitionSize for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyPartitionSize(value uint64) (err error) {
	return instance.SetProperty("PartitionSize", (value))
}

// GetPartitionSize gets the value of PartitionSize for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyPartitionSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PartitionSize")
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

// SetSectorsPerCluster sets the value of SectorsPerCluster for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertySectorsPerCluster(value uint32) (err error) {
	return instance.SetProperty("SectorsPerCluster", (value))
}

// GetSectorsPerCluster gets the value of SectorsPerCluster for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertySectorsPerCluster() (value uint32, err error) {
	retValue, err := instance.GetProperty("SectorsPerCluster")
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
func (instance *SystemConfig_V1_LogDisk) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertySize() (value uint32, err error) {
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

// SetStartOffset sets the value of StartOffset for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyStartOffset(value uint64) (err error) {
	return instance.SetProperty("StartOffset", (value))
}

// GetStartOffset gets the value of StartOffset for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyStartOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartOffset")
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

// SetTotalNumberOfClusters sets the value of TotalNumberOfClusters for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyTotalNumberOfClusters(value int64) (err error) {
	return instance.SetProperty("TotalNumberOfClusters", (value))
}

// GetTotalNumberOfClusters gets the value of TotalNumberOfClusters for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyTotalNumberOfClusters() (value int64, err error) {
	retValue, err := instance.GetProperty("TotalNumberOfClusters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetVolumeExt sets the value of VolumeExt for the instance
func (instance *SystemConfig_V1_LogDisk) SetPropertyVolumeExt(value uint32) (err error) {
	return instance.SetProperty("VolumeExt", (value))
}

// GetVolumeExt gets the value of VolumeExt for the instance
func (instance *SystemConfig_V1_LogDisk) GetPropertyVolumeExt() (value uint32, err error) {
	retValue, err := instance.GetProperty("VolumeExt")
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
