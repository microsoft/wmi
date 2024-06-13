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

// SystemConfig_V0_PhyDisk struct
type SystemConfig_V0_PhyDisk struct {
	*SystemConfig_V0

	//
	BootDriveLetter []byte

	//
	BytesPerSector uint32

	//
	Cylinders uint64

	//
	DiskNumber uint32

	//
	Manufacturer []byte

	//
	Pad uint8

	//
	PartitionCount uint32

	//
	SCSILun uint32

	//
	SCSIPath uint32

	//
	SCSIPort uint32

	//
	SCSITarget uint32

	//
	SectorsPerTrack uint32

	//
	Spare []byte

	//
	TracksPerCylinder uint32

	//
	WriteCacheEnabled uint8
}

func NewSystemConfig_V0_PhyDiskEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V0_PhyDisk, err error) {
	tmp, err := NewSystemConfig_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_PhyDisk{
		SystemConfig_V0: tmp,
	}
	return
}

func NewSystemConfig_V0_PhyDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V0_PhyDisk, err error) {
	tmp, err := NewSystemConfig_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_PhyDisk{
		SystemConfig_V0: tmp,
	}
	return
}

// SetBootDriveLetter sets the value of BootDriveLetter for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyBootDriveLetter(value []byte) (err error) {
	return instance.SetProperty("BootDriveLetter", (value))
}

// GetBootDriveLetter gets the value of BootDriveLetter for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyBootDriveLetter() (value []byte, err error) {
	retValue, err := instance.GetProperty("BootDriveLetter")
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

// SetBytesPerSector sets the value of BytesPerSector for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyBytesPerSector(value uint32) (err error) {
	return instance.SetProperty("BytesPerSector", (value))
}

// GetBytesPerSector gets the value of BytesPerSector for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyBytesPerSector() (value uint32, err error) {
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

// SetCylinders sets the value of Cylinders for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyCylinders(value uint64) (err error) {
	return instance.SetProperty("Cylinders", (value))
}

// GetCylinders gets the value of Cylinders for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyCylinders() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cylinders")
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

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyDiskNumber() (value uint32, err error) {
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyManufacturer(value []byte) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyManufacturer() (value []byte, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetPad sets the value of Pad for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyPad(value uint8) (err error) {
	return instance.SetProperty("Pad", (value))
}

// GetPad gets the value of Pad for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyPad() (value uint8, err error) {
	retValue, err := instance.GetProperty("Pad")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetPartitionCount sets the value of PartitionCount for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyPartitionCount(value uint32) (err error) {
	return instance.SetProperty("PartitionCount", (value))
}

// GetPartitionCount gets the value of PartitionCount for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyPartitionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PartitionCount")
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

// SetSCSILun sets the value of SCSILun for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySCSILun(value uint32) (err error) {
	return instance.SetProperty("SCSILun", (value))
}

// GetSCSILun gets the value of SCSILun for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySCSILun() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSILun")
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

// SetSCSIPath sets the value of SCSIPath for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySCSIPath(value uint32) (err error) {
	return instance.SetProperty("SCSIPath", (value))
}

// GetSCSIPath gets the value of SCSIPath for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySCSIPath() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSIPath")
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

// SetSCSIPort sets the value of SCSIPort for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySCSIPort(value uint32) (err error) {
	return instance.SetProperty("SCSIPort", (value))
}

// GetSCSIPort gets the value of SCSIPort for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySCSIPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSIPort")
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

// SetSCSITarget sets the value of SCSITarget for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySCSITarget(value uint32) (err error) {
	return instance.SetProperty("SCSITarget", (value))
}

// GetSCSITarget gets the value of SCSITarget for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySCSITarget() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSITarget")
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

// SetSectorsPerTrack sets the value of SectorsPerTrack for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySectorsPerTrack(value uint32) (err error) {
	return instance.SetProperty("SectorsPerTrack", (value))
}

// GetSectorsPerTrack gets the value of SectorsPerTrack for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySectorsPerTrack() (value uint32, err error) {
	retValue, err := instance.GetProperty("SectorsPerTrack")
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

// SetSpare sets the value of Spare for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertySpare(value []byte) (err error) {
	return instance.SetProperty("Spare", (value))
}

// GetSpare gets the value of Spare for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertySpare() (value []byte, err error) {
	retValue, err := instance.GetProperty("Spare")
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

// SetTracksPerCylinder sets the value of TracksPerCylinder for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyTracksPerCylinder(value uint32) (err error) {
	return instance.SetProperty("TracksPerCylinder", (value))
}

// GetTracksPerCylinder gets the value of TracksPerCylinder for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyTracksPerCylinder() (value uint32, err error) {
	retValue, err := instance.GetProperty("TracksPerCylinder")
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

// SetWriteCacheEnabled sets the value of WriteCacheEnabled for the instance
func (instance *SystemConfig_V0_PhyDisk) SetPropertyWriteCacheEnabled(value uint8) (err error) {
	return instance.SetProperty("WriteCacheEnabled", (value))
}

// GetWriteCacheEnabled gets the value of WriteCacheEnabled for the instance
func (instance *SystemConfig_V0_PhyDisk) GetPropertyWriteCacheEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("WriteCacheEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
