// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PhysicalDisk struct
type MSFT_PhysicalDisk struct {
	*MSFT_StorageFaultDomain

	// A string representation of the Adapter's serial number.
	AdapterSerialNumber string

	// This field indicates the sum of used space on this physical disk. This should include usage from all storage pools and other data stored on the disk.
	AllocatedSize uint64

	//
	BusType PhysicalDisk_BusType

	// Indicates the reason why this physical disk cannot be added to a concrete pool
	CannotPoolReason []PhysicalDisk_CannotPoolReason

	// Indicates whether this physical disk can be added to a concrete pool or not
	CanPool bool

	// DeviceId is an address or other identifier that uniquely names the physical disk.
	DeviceId string

	// Indicates the enclosure number in which the disk physically resides
	EnclosureNumber uint16

	// This field is a string representation of the physical disk's firmware version.
	FirmwareVersion string

	// FruId is an identifier of the replacement unit housing the physical disk.
	FruId string

	// Indicates whether the physical disk's identification LEDs are active or not. This is typically used in maintenance operations.
	IsIndicationEnabled bool

	// Indicates whether this physical disk is partially consumed by a system or service whose use is outside of normal storage pool operations.
	IsPartial bool

	// This field indicates the logical sector size of the physical disk in bytes. For example: a 4K native disk should report 4096, while a 512 emulated disk should report 512.
	LogicalSectorSize uint64

	// Media type of this physical disk
	MediaType PhysicalDisk_MediaType

	// If CannotPoolReason contains 1 - 'Other', this field contains the string representing the vendor defined reason why this physical disk cannot be added to a concrete pool. This property must be NULL if CannotPoolReason does not contain 1 - 'Other'.
	OtherCannotPoolReasonDescription string

	// This field is a string representation of the physical disk's part number or SKU.
	PartNumber string

	// This field indicates the physical sector size of the physical disk in bytes. For example: for 4K native and 512 emulated disks, the value should be 4096.
	PhysicalSectorSize uint64

	// Indicates the total physical storage size of the disk in bytes
	Size uint64

	// Indicates the enclosure slot number in which the disk physically resides
	SlotNumber uint16

	// This field is a string representation of the physical disk's software version.
	SoftwareVersion string

	// This field indicates the rotational speed of spindle-based physical disks. For solid state devices (SSDs) or other non-rotational media, this field should set to 0. For rotating media which has an unknown speed, this field should be set to -1 (UINT32_MAX).
	SpindleSpeed uint32

	//
	StoragePoolUniqueId string

	// This field describes the supported usages of this physical disk.
	SupportedUsages []PhysicalDisk_SupportedUsages

	// UniqueIdFormat indicates the type of identifier used in the UniqueId field. The identifier used in UniqueId must be the highest available identifier using the following order of preference: 8 (highest), 3, 2, 1, 0 (lowest). For example: if the physical disk device exposes identifiers of type 0, 1, and 3, UniqueId must be the identifier of type 3, and UniqueIdFormat should be set to 3.
	UniqueIdFormat PhysicalDisk_UniqueIdFormat

	// This field describes the intended usage of this physical disk within a concrete pool. Storage pools are required to follow the assigned policy for a physical disk.
	///1 - 'Auto-Select': This physical disk should only be used for data storage.
	///2 - 'Manual-Select': This physical disk should only be used if manually selected by an administrator at the time of virtual disk creation. A manual-select disk is selected using the PhysicalDisksToUse parameter to CreateVirtualDisk.
	///3 - 'Hot Spare': This physical disk should be used as a hot spare.
	///4 - 'Retired': This physical disk should be retired from use. At a minimum, no new allocations should go to this disk. If the virtual disks that reside on this disk are repaired, the data should be moved to another active physical disk.
	Usage PhysicalDisk_Usage

	// This field indicates the size in bytes of the user data footprint from virtual disks on this physical disk.
	VirtualDiskFootprint uint64
}

func NewMSFT_PhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_PhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_PhysicalDisk{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}

func NewMSFT_PhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PhysicalDisk{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}

// SetAdapterSerialNumber sets the value of AdapterSerialNumber for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyAdapterSerialNumber(value string) (err error) {
	return instance.SetProperty("AdapterSerialNumber", (value))
}

// GetAdapterSerialNumber gets the value of AdapterSerialNumber for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyAdapterSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("AdapterSerialNumber")
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

// SetAllocatedSize sets the value of AllocatedSize for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", (value))
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyAllocatedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatedSize")
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

// SetBusType sets the value of BusType for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyBusType(value PhysicalDisk_BusType) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyBusType() (value PhysicalDisk_BusType, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalDisk_BusType(valuetmp)

	return
}

// SetCannotPoolReason sets the value of CannotPoolReason for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyCannotPoolReason(value []PhysicalDisk_CannotPoolReason) (err error) {
	return instance.SetProperty("CannotPoolReason", (value))
}

// GetCannotPoolReason gets the value of CannotPoolReason for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyCannotPoolReason() (value []PhysicalDisk_CannotPoolReason, err error) {
	retValue, err := instance.GetProperty("CannotPoolReason")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, PhysicalDisk_CannotPoolReason(valuetmp))
	}

	return
}

// SetCanPool sets the value of CanPool for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyCanPool(value bool) (err error) {
	return instance.SetProperty("CanPool", (value))
}

// GetCanPool gets the value of CanPool for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyCanPool() (value bool, err error) {
	retValue, err := instance.GetProperty("CanPool")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyDeviceId(value string) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceId")
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

// SetEnclosureNumber sets the value of EnclosureNumber for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyEnclosureNumber(value uint16) (err error) {
	return instance.SetProperty("EnclosureNumber", (value))
}

// GetEnclosureNumber gets the value of EnclosureNumber for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyEnclosureNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnclosureNumber")
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

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
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

// SetFruId sets the value of FruId for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyFruId(value string) (err error) {
	return instance.SetProperty("FruId", (value))
}

// GetFruId gets the value of FruId for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyFruId() (value string, err error) {
	retValue, err := instance.GetProperty("FruId")
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

// SetIsIndicationEnabled sets the value of IsIndicationEnabled for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyIsIndicationEnabled(value bool) (err error) {
	return instance.SetProperty("IsIndicationEnabled", (value))
}

// GetIsIndicationEnabled gets the value of IsIndicationEnabled for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyIsIndicationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIndicationEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIsPartial sets the value of IsPartial for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyIsPartial(value bool) (err error) {
	return instance.SetProperty("IsPartial", (value))
}

// GetIsPartial gets the value of IsPartial for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyIsPartial() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPartial")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyLogicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("LogicalSectorSize", (value))
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyLogicalSectorSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
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

// SetMediaType sets the value of MediaType for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyMediaType(value PhysicalDisk_MediaType) (err error) {
	return instance.SetProperty("MediaType", (value))
}

// GetMediaType gets the value of MediaType for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyMediaType() (value PhysicalDisk_MediaType, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalDisk_MediaType(valuetmp)

	return
}

// SetOtherCannotPoolReasonDescription sets the value of OtherCannotPoolReasonDescription for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyOtherCannotPoolReasonDescription(value string) (err error) {
	return instance.SetProperty("OtherCannotPoolReasonDescription", (value))
}

// GetOtherCannotPoolReasonDescription gets the value of OtherCannotPoolReasonDescription for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyOtherCannotPoolReasonDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCannotPoolReasonDescription")
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

// SetPartNumber sets the value of PartNumber for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyPartNumber(value string) (err error) {
	return instance.SetProperty("PartNumber", (value))
}

// GetPartNumber gets the value of PartNumber for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyPartNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PartNumber")
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

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyPhysicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("PhysicalSectorSize", (value))
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyPhysicalSectorSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
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

// SetSize sets the value of Size for the instance
func (instance *MSFT_PhysicalDisk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_PhysicalDisk) GetPropertySize() (value uint64, err error) {
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

// SetSlotNumber sets the value of SlotNumber for the instance
func (instance *MSFT_PhysicalDisk) SetPropertySlotNumber(value uint16) (err error) {
	return instance.SetProperty("SlotNumber", (value))
}

// GetSlotNumber gets the value of SlotNumber for the instance
func (instance *MSFT_PhysicalDisk) GetPropertySlotNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("SlotNumber")
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

// SetSoftwareVersion sets the value of SoftwareVersion for the instance
func (instance *MSFT_PhysicalDisk) SetPropertySoftwareVersion(value string) (err error) {
	return instance.SetProperty("SoftwareVersion", (value))
}

// GetSoftwareVersion gets the value of SoftwareVersion for the instance
func (instance *MSFT_PhysicalDisk) GetPropertySoftwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SoftwareVersion")
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

// SetSpindleSpeed sets the value of SpindleSpeed for the instance
func (instance *MSFT_PhysicalDisk) SetPropertySpindleSpeed(value uint32) (err error) {
	return instance.SetProperty("SpindleSpeed", (value))
}

// GetSpindleSpeed gets the value of SpindleSpeed for the instance
func (instance *MSFT_PhysicalDisk) GetPropertySpindleSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpindleSpeed")
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

// SetStoragePoolUniqueId sets the value of StoragePoolUniqueId for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyStoragePoolUniqueId(value string) (err error) {
	return instance.SetProperty("StoragePoolUniqueId", (value))
}

// GetStoragePoolUniqueId gets the value of StoragePoolUniqueId for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyStoragePoolUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("StoragePoolUniqueId")
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

// SetSupportedUsages sets the value of SupportedUsages for the instance
func (instance *MSFT_PhysicalDisk) SetPropertySupportedUsages(value []PhysicalDisk_SupportedUsages) (err error) {
	return instance.SetProperty("SupportedUsages", (value))
}

// GetSupportedUsages gets the value of SupportedUsages for the instance
func (instance *MSFT_PhysicalDisk) GetPropertySupportedUsages() (value []PhysicalDisk_SupportedUsages, err error) {
	retValue, err := instance.GetProperty("SupportedUsages")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, PhysicalDisk_SupportedUsages(valuetmp))
	}

	return
}

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyUniqueIdFormat(value PhysicalDisk_UniqueIdFormat) (err error) {
	return instance.SetProperty("UniqueIdFormat", (value))
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyUniqueIdFormat() (value PhysicalDisk_UniqueIdFormat, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalDisk_UniqueIdFormat(valuetmp)

	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyUsage(value PhysicalDisk_Usage) (err error) {
	return instance.SetProperty("Usage", (value))
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyUsage() (value PhysicalDisk_Usage, err error) {
	retValue, err := instance.GetProperty("Usage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalDisk_Usage(valuetmp)

	return
}

// SetVirtualDiskFootprint sets the value of VirtualDiskFootprint for the instance
func (instance *MSFT_PhysicalDisk) SetPropertyVirtualDiskFootprint(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskFootprint", (value))
}

// GetVirtualDiskFootprint gets the value of VirtualDiskFootprint for the instance
func (instance *MSFT_PhysicalDisk) GetPropertyVirtualDiskFootprint() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskFootprint")
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

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="PhysicalExtents" type="MSFT_PhysicalExtent []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) GetPhysicalExtent( /* OUT */ PhysicalExtents []MSFT_PhysicalExtent,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalExtent")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user to perform certain maintenance tasks on the physical disk.

// <param name="EnableIndication" type="bool ">If set to TRUE, this instructs the physical disk to enable its indication LED. The indication LED should remain enabled until a second call to Maintenance is made with this parameter specified as FALSE.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) Maintenance( /* IN */ EnableIndication bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Maintenance", EnableIndication)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EnableIndication" type="bool "></param>
// <param name="EnableMaintenanceMode" type="bool "></param>
// <param name="IgnoreDetachedVirtualDisks" type="bool "></param>
// <param name="Timeout" type="uint32 "></param>
// <param name="ValidateMaintenanceMode" type="bool "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) Maintenance2( /* IN */ EnableIndication bool,
	/* IN */ ValidateMaintenanceMode bool,
	/* IN */ EnableMaintenanceMode bool,
	/* IN */ Timeout uint32,
	/* IN */ IgnoreDetachedVirtualDisks bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Maintenance2", EnableIndication, ValidateMaintenanceMode, EnableMaintenanceMode, Timeout, IgnoreDetachedVirtualDisks)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method resets the health and operational status of the physical disk. Exact behavior of this method is dependent on whether this physical disk belongs to a concrete pool.
///If it is a member of a concrete pool, the health and operational statuses should be reset to 1 - 'Healthy', and 1 - 'OK', respectively. If any additional errors are detected after Reset, the health and operational statuses should reflect these new errors.
///If the physical disk is not a member of a concrete pool, then this method should not only reset the health and operational statuses, but it should return the disk into a state where it is usable as storage for a concrete pool. For example: If a physical disk had become missing and then has reappeared (after it has been replaced) this physical disk is expected to be in the primordial pool only with an operational status indicating its data is either split or unrecognized. Calling Reset should clear the physical disk of any data, remove any remaining ties to its former concrete pool, and return the disk to a healthy, usable state.

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) Reset( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the physical disk to be renamed.

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the physical disk's description to be changed.

// <param name="Description" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) SetDescription( /* IN */ Description string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDescription", Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the physical disk's usage to be updated.

// <param name="Usage" type="PhysicalDisk_Usage ">This field describes the intended usage of this physical disk within a concrete pool. Storage pools are required to follow the assigned policy for a physical disk.
///1 - 'Auto-Select': This physical disk should only be used for data storage.
///2 - 'Manual-Select': This physical disk should only be used if manually selected by an administrator at the time of virtual disk creation. A manual-select disk is selected using the PhysicalDisksToUse parameter to CreateVirtualDisk.
///3 - 'Hot Spare': This physical disk should be used as a hot spare.
///4 - 'Retired': This physical disk should be retired from use. At a minimum, no new allocations should go to this disk. If the virtual disks that reside on this disk are repaired, the data should be moved to another active physical disk.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) SetUsage( /* IN */ Usage PhysicalDisk_Usage,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetUsage", Usage)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the physical disk's attributes to be updated.

// <param name="MediaType" type="PhysicalDisk_MediaType ">Media type of this physical disk</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) SetAttributes( /* IN */ MediaType PhysicalDisk_MediaType,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", MediaType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsHidden" type="bool "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="StorageEnclosureId" type="string "></param>
// <param name="StorageScaleUnitId" type="string "></param>

// <param name="ExtendedStatus" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) SetAttributes2( /* IN */ MediaType uint16,
	/* IN */ StorageEnclosureId string,
	/* IN */ StorageScaleUnitId string,
	/* IN */ IsHidden bool,
	/* OUT */ ExtendedStatus string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes2", MediaType, StorageEnclosureId, StorageScaleUnitId, IsHidden)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="IsDeviceCacheEnabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) IsDeviceCacheEnabled( /* OUT */ IsDeviceCacheEnabled bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsDeviceCacheEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="IsPowerProtected" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) IsPowerProtected( /* OUT */ IsPowerProtected bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsPowerProtected")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ActiveSlotNumber" type="uint16 "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="FirmwareVersionInSlot" type="string []"></param>
// <param name="IsSlotWritable" type="bool []"></param>
// <param name="NumberOfSlots" type="uint16 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SlotNumber" type="uint16 []"></param>
// <param name="SupportsUpdate" type="bool "></param>
func (instance *MSFT_PhysicalDisk) GetFirmwareInformation( /* OUT */ SupportsUpdate bool,
	/* OUT */ NumberOfSlots uint16,
	/* OUT */ ActiveSlotNumber uint16,
	/* OUT */ SlotNumber []uint16,
	/* OUT */ IsSlotWritable []bool,
	/* OUT */ FirmwareVersionInSlot []string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetFirmwareInformation")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ImagePath" type="string "></param>
// <param name="SlotNumber" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) UpdateFirmware( /* IN */ ImagePath string,
	/* IN */ SlotNumber uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UpdateFirmware", ImagePath, SlotNumber)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Format" type="uint16 "></param>
// <param name="StoragePoolFriendlyName" type="string "></param>
// <param name="StoragePoolMetadataLength" type="uint64 "></param>
// <param name="StoragePoolMinimumAllocationSize" type="uint64 "></param>
// <param name="StoragePoolVersion" type="uint16 "></param>
// <param name="VirtualDiskAllocationUnitSize" type="uint64 "></param>
// <param name="VirtualDiskFriendlyName" type="string "></param>
// <param name="VirtualDiskProvisioningType" type="uint16 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedStorageObject" type="MSFT_StorageObject "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PhysicalDisk) Convert( /* IN */ Format uint16,
	/* IN */ StoragePoolFriendlyName string,
	/* IN */ StoragePoolVersion uint16,
	/* IN */ StoragePoolMetadataLength uint64,
	/* IN */ StoragePoolMinimumAllocationSize uint64,
	/* IN */ VirtualDiskFriendlyName string,
	/* IN */ VirtualDiskProvisioningType uint16,
	/* IN */ VirtualDiskAllocationUnitSize uint64,
	/* OUT */ CreatedStorageObject MSFT_StorageObject,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Convert", Format, StoragePoolFriendlyName, StoragePoolVersion, StoragePoolMetadataLength, StoragePoolMinimumAllocationSize, VirtualDiskFriendlyName, VirtualDiskProvisioningType, VirtualDiskAllocationUnitSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
