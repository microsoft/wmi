// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_DeviceOnSession struct
type MSiSCSIInitiator_DeviceOnSession struct {
	*cim.WmiInstance

	// PNP Device interface guid
	DeviceInterfaceGuid string

	// PNP Device interface name
	DeviceInterfaceName string

	// The number of this device
	DeviceNumber uint32

	// The FILE_DEVICE_XXX type for this device.
	DeviceType uint32

	// Name of initiator
	InitiatorName string

	// Legacy Device interface name
	LegacyName string

	// If the device is partitionable, the partition number of the device. Otherwise -1
	PartitionNumber uint32

	// OS SCSI Logical Unit Number
	ScsiLun uint8

	// OS SCSI path identifier or bus number
	ScsiPathId uint8

	// OS SCSI port number
	ScsiPortNumber uint8

	// OS SCSI Target Id
	ScsiTargetId uint8

	// Name of target
	TargetName string
}

func NewMSiSCSIInitiator_DeviceOnSessionEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_DeviceOnSession, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_DeviceOnSession{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_DeviceOnSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_DeviceOnSession, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_DeviceOnSession{
		WmiInstance: tmp,
	}
	return
}

// SetDeviceInterfaceGuid sets the value of DeviceInterfaceGuid for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyDeviceInterfaceGuid(value string) (err error) {
	return instance.SetProperty("DeviceInterfaceGuid", (value))
}

// GetDeviceInterfaceGuid gets the value of DeviceInterfaceGuid for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyDeviceInterfaceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInterfaceGuid")
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

// SetDeviceInterfaceName sets the value of DeviceInterfaceName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyDeviceInterfaceName(value string) (err error) {
	return instance.SetProperty("DeviceInterfaceName", (value))
}

// GetDeviceInterfaceName gets the value of DeviceInterfaceName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyDeviceInterfaceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInterfaceName")
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

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("DeviceNumber", (value))
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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

// SetLegacyName sets the value of LegacyName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyLegacyName(value string) (err error) {
	return instance.SetProperty("LegacyName", (value))
}

// GetLegacyName gets the value of LegacyName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyLegacyName() (value string, err error) {
	retValue, err := instance.GetProperty("LegacyName")
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

// SetPartitionNumber sets the value of PartitionNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyPartitionNumber(value uint32) (err error) {
	return instance.SetProperty("PartitionNumber", (value))
}

// GetPartitionNumber gets the value of PartitionNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyPartitionNumber() (value uint32, err error) {
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

// SetScsiLun sets the value of ScsiLun for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyScsiLun(value uint8) (err error) {
	return instance.SetProperty("ScsiLun", (value))
}

// GetScsiLun gets the value of ScsiLun for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyScsiLun() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScsiLun")
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

// SetScsiPathId sets the value of ScsiPathId for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyScsiPathId(value uint8) (err error) {
	return instance.SetProperty("ScsiPathId", (value))
}

// GetScsiPathId gets the value of ScsiPathId for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyScsiPathId() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScsiPathId")
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

// SetScsiPortNumber sets the value of ScsiPortNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyScsiPortNumber(value uint8) (err error) {
	return instance.SetProperty("ScsiPortNumber", (value))
}

// GetScsiPortNumber gets the value of ScsiPortNumber for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyScsiPortNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScsiPortNumber")
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

// SetScsiTargetId sets the value of ScsiTargetId for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyScsiTargetId(value uint8) (err error) {
	return instance.SetProperty("ScsiTargetId", (value))
}

// GetScsiTargetId gets the value of ScsiTargetId for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyScsiTargetId() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScsiTargetId")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_DeviceOnSession) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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
