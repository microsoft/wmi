// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_Disk struct
type MSFT_Disk struct {
	*MSFT_StorageObject

	// A string representation of the Adapter's serial number.
	AdapterSerialNumber string

	// The amount of space currently used on the disk.
	AllocatedSize uint64

	// This property indicates that the computer is configured to start off of this disk. On computers with BIOS firmware, this is the first disk that the firmware detects during startup. On computers that use EFI firmware, this is the disk that contains the EFI System Partition (ESP). If there are no disks or multiple disks with an ESP partition, this flag is not set for any disk.
	BootFromDisk bool

	// Denotes the I/O bus type used by this disk.
	BusType Disk_BusType

	// A string representation of the disk's firmware version.
	FirmwareVersion string

	// FriendlyName is a user-friendly, display-oriented string to identify the disk.
	FriendlyName string

	// The GPT guid of the disk. This property is only valid on GPT disks and will be NULL for all other disk types.
	Guid string

	// The health status of the Volume.
	///0 - 'Healthy': The disk is functioning normally.
	///1 - 'Warning': The disk is still functioning, but has detected errors or issues that require administrator intervention.
	///2 - 'Unhealthy': The volume is not functioning, due to errors or failures. The volume needs immediate attention from an administrator.
	HealthStatus Disk_HealthStatus

	// This property indicates that the computer has booted off of this disk.
	IsBoot bool

	// If IsClustered is TRUE, this disk is used in a clustered environment.
	IsClustered bool

	// If IsHighlyAvailable is TRUE, the disk is highly available.
	IsHighlyAvailable bool

	//
	IsOffline bool

	//
	IsReadOnly bool

	// If IsScaleOut is TRUE, the disk is scaled out.
	IsScaleOut bool

	// If IsSystem is TRUE, this disk contains the system partition.
	IsSystem bool

	// This field indicates the largest contiguous block of free space on the disk. This is also the largest size of a partition which can be created on the disk.
	LargestFreeExtent uint64

	// Location contains the PnP location path of the disk. The format of this string depends on the bus type. If the bus type is SCSI, SAS, or PCI RAID, the format is <AdapterPnpLocationPath>#<BusType>(P<PathId>T<TargetId>L<LunId>). If the bus type is IDE, ATA, PATA, or SATA, the format is <AdapterPnpLocationPath>#<BusType>(C<PathId>T<TargetId>L<LunId>). For example, a SCSI location may look like: PCIROOT(0)#PCI(1C00)#PCI(0000)#SCSI(P00T01L01). Note: For Hyper-V and VHD images, this member is NULL because the virtual controller does not return the location path.
	Location string

	// This field indicates the logical sector size of the disk in bytes. For example: a 4K native disk will report 4096, while a 512 emulated disk will report 512.
	LogicalSectorSize uint32

	// A string representation of the disk's hardware manufacturer.
	Manufacturer string

	// A string representation of the disk's model.
	Model string

	// The operating system's number for the disk. Disk 0 is typically the boot device. Disk numbers may not necessarily remain the same across reboots.
	Number uint32

	//
	NumberOfPartitions uint32

	// If IsOffline is TRUE, this property informs the user of the specific reason for the disk being offline.
	///1 - 'Policy': The user requested the disk to be offline.
	///2 - 'Redundant Path': The disk is used for multi-path I/O.
	///3 - 'Snapshot': The disk is a snapshot disk.
	///4 - 'Collision': There was a signature or identifier collision with another disk.
	///5 - 'Resource Exhaustion': There were insufficient resources to bring the disk online.
	///6 - 'Critical Write Failures': There were critical write failures on the disk.
	///7 - 'Data Integrity Scan Required': A data integrity scan is required.
	OfflineReason Disk_OfflineReason

	// An array of values that denote the current operational status of the volume.
	///0 - 'Unknown': The operational status is unknown.
	///1 - 'Other': A vendor-specific OperationalStatus has been specified by setting the OtherOperationalStatusDescription property.
	///2 - 'OK': The disk is responding to commands and is in a normal operating state.
	///3 - 'Degraded': The disk is responding to commands, but is not running in an optimal operating state.
	///4 - 'Stressed': The disk is functioning, but needs attention. For example, the disk might be overloaded or overheated.
	///5 - 'Predictive Failure': The disk is functioning, but a failure is likely to occur in the near future.
	///6 - 'Error': An error has occurred.
	///7 - 'Non-Recoverable Error': A non-recoverable error has occurred.
	///8 - 'Starting': The disk is in the process of starting.
	///9 - 'Stopping': The disk is in the process of stopping.
	///10 - 'Stopped': The disk was stopped or shut down in a clean and orderly fashion.
	///11 - 'In Service': The disk is being configured, maintained, cleaned, or otherwise administered.
	///12 - 'No Contact': The storage provider has knowledge of the disk, but has never been able to establish communication with it.
	///13 - 'Lost Communication': The storage provider has knowledge of the disk and has contacted it successfully in the past, but the disk is currently unreachable.
	///14 - 'Aborted': Similar to Stopped, except that the disk stopped abruptly and may require configuration or maintenance.
	///15 - 'Dormant': The disk is reachable, but it is inactive.
	///16 - 'Supporting Entity in Error': This status value does not necessarily indicate trouble with the disk, but it does indicate that another device or connection that the disk depends on may need attention.
	///17 - 'Completed': The disk has completed an operation. This status value should be combined with OK, Error, or Degraded, depending on the outcome of the operation.
	///0xD010 - 'Online': In Windows-based storage subsystems, this indicates that the object is online.
	///0xD011 - 'Not Ready': In Windows-based storage subsystems, this indicates that the object is not ready.
	///0xD012 - 'No Media': In Windows-based storage subsystems, this indicates that the object has no media present.
	///0xD013 - 'Offline': In Windows-based storage subsystems, this indicates that the object is offline.
	///0xD014 - 'Failed': In Windows-based storage subsystems, this indicates that the object is in a failed state.
	OperationalStatus []Disk_OperationalStatus

	//
	PartitionStyle Disk_PartitionStyle

	// Path can be used to open an operating system handle to the disk device.
	Path string

	// This field indicates the physical sector size of the disk in bytes. For example: both 4K native disks and 512 emulated disks will report 4096.
	PhysicalSectorSize uint32

	// Denotes the provisioning type of the disk device.
	///1 - 'Thin' means that the storage for the disk is allocated on-demand.
	///2 - 'Fixed' means that the storage is allocated up front.
	ProvisioningType Disk_ProvisioningType

	// A string representation of the disk's serial number.
	SerialNumber string

	// The MBR signature of the disk. This property is only valid on MBR disks and will be NULL for all other disk types.
	Signature uint32

	// The total size of the disk, measured in bytes.
	Size uint64

	// UniqueIdFormat informs the user what VPD Page 0x83 descriptor type was used to populate the UniqueId field.
	UniqueIdFormat Disk_UniqueIdFormat
}

func NewMSFT_DiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_Disk, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_Disk{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_DiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Disk, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Disk{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetAdapterSerialNumber sets the value of AdapterSerialNumber for the instance
func (instance *MSFT_Disk) SetPropertyAdapterSerialNumber(value string) (err error) {
	return instance.SetProperty("AdapterSerialNumber", (value))
}

// GetAdapterSerialNumber gets the value of AdapterSerialNumber for the instance
func (instance *MSFT_Disk) GetPropertyAdapterSerialNumber() (value string, err error) {
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
func (instance *MSFT_Disk) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", (value))
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_Disk) GetPropertyAllocatedSize() (value uint64, err error) {
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

// SetBootFromDisk sets the value of BootFromDisk for the instance
func (instance *MSFT_Disk) SetPropertyBootFromDisk(value bool) (err error) {
	return instance.SetProperty("BootFromDisk", (value))
}

// GetBootFromDisk gets the value of BootFromDisk for the instance
func (instance *MSFT_Disk) GetPropertyBootFromDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("BootFromDisk")
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

// SetBusType sets the value of BusType for the instance
func (instance *MSFT_Disk) SetPropertyBusType(value Disk_BusType) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *MSFT_Disk) GetPropertyBusType() (value Disk_BusType, err error) {
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

	value = Disk_BusType(valuetmp)

	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_Disk) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_Disk) GetPropertyFirmwareVersion() (value string, err error) {
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_Disk) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_Disk) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_Disk) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_Disk) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_Disk) SetPropertyHealthStatus(value Disk_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_Disk) GetPropertyHealthStatus() (value Disk_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

	value = Disk_HealthStatus(valuetmp)

	return
}

// SetIsBoot sets the value of IsBoot for the instance
func (instance *MSFT_Disk) SetPropertyIsBoot(value bool) (err error) {
	return instance.SetProperty("IsBoot", (value))
}

// GetIsBoot gets the value of IsBoot for the instance
func (instance *MSFT_Disk) GetPropertyIsBoot() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBoot")
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

// SetIsClustered sets the value of IsClustered for the instance
func (instance *MSFT_Disk) SetPropertyIsClustered(value bool) (err error) {
	return instance.SetProperty("IsClustered", (value))
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *MSFT_Disk) GetPropertyIsClustered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
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

// SetIsHighlyAvailable sets the value of IsHighlyAvailable for the instance
func (instance *MSFT_Disk) SetPropertyIsHighlyAvailable(value bool) (err error) {
	return instance.SetProperty("IsHighlyAvailable", (value))
}

// GetIsHighlyAvailable gets the value of IsHighlyAvailable for the instance
func (instance *MSFT_Disk) GetPropertyIsHighlyAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHighlyAvailable")
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

// SetIsOffline sets the value of IsOffline for the instance
func (instance *MSFT_Disk) SetPropertyIsOffline(value bool) (err error) {
	return instance.SetProperty("IsOffline", (value))
}

// GetIsOffline gets the value of IsOffline for the instance
func (instance *MSFT_Disk) GetPropertyIsOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOffline")
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

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *MSFT_Disk) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", (value))
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *MSFT_Disk) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
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

// SetIsScaleOut sets the value of IsScaleOut for the instance
func (instance *MSFT_Disk) SetPropertyIsScaleOut(value bool) (err error) {
	return instance.SetProperty("IsScaleOut", (value))
}

// GetIsScaleOut gets the value of IsScaleOut for the instance
func (instance *MSFT_Disk) GetPropertyIsScaleOut() (value bool, err error) {
	retValue, err := instance.GetProperty("IsScaleOut")
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

// SetIsSystem sets the value of IsSystem for the instance
func (instance *MSFT_Disk) SetPropertyIsSystem(value bool) (err error) {
	return instance.SetProperty("IsSystem", (value))
}

// GetIsSystem gets the value of IsSystem for the instance
func (instance *MSFT_Disk) GetPropertyIsSystem() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSystem")
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

// SetLargestFreeExtent sets the value of LargestFreeExtent for the instance
func (instance *MSFT_Disk) SetPropertyLargestFreeExtent(value uint64) (err error) {
	return instance.SetProperty("LargestFreeExtent", (value))
}

// GetLargestFreeExtent gets the value of LargestFreeExtent for the instance
func (instance *MSFT_Disk) GetPropertyLargestFreeExtent() (value uint64, err error) {
	retValue, err := instance.GetProperty("LargestFreeExtent")
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

// SetLocation sets the value of Location for the instance
func (instance *MSFT_Disk) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *MSFT_Disk) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
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

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *MSFT_Disk) SetPropertyLogicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("LogicalSectorSize", (value))
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_Disk) GetPropertyLogicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
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
func (instance *MSFT_Disk) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_Disk) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetModel sets the value of Model for the instance
func (instance *MSFT_Disk) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_Disk) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
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

// SetNumber sets the value of Number for the instance
func (instance *MSFT_Disk) SetPropertyNumber(value uint32) (err error) {
	return instance.SetProperty("Number", (value))
}

// GetNumber gets the value of Number for the instance
func (instance *MSFT_Disk) GetPropertyNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("Number")
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

// SetNumberOfPartitions sets the value of NumberOfPartitions for the instance
func (instance *MSFT_Disk) SetPropertyNumberOfPartitions(value uint32) (err error) {
	return instance.SetProperty("NumberOfPartitions", (value))
}

// GetNumberOfPartitions gets the value of NumberOfPartitions for the instance
func (instance *MSFT_Disk) GetPropertyNumberOfPartitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPartitions")
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

// SetOfflineReason sets the value of OfflineReason for the instance
func (instance *MSFT_Disk) SetPropertyOfflineReason(value Disk_OfflineReason) (err error) {
	return instance.SetProperty("OfflineReason", (value))
}

// GetOfflineReason gets the value of OfflineReason for the instance
func (instance *MSFT_Disk) GetPropertyOfflineReason() (value Disk_OfflineReason, err error) {
	retValue, err := instance.GetProperty("OfflineReason")
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

	value = Disk_OfflineReason(valuetmp)

	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_Disk) SetPropertyOperationalStatus(value []Disk_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_Disk) GetPropertyOperationalStatus() (value []Disk_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
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
		value = append(value, Disk_OperationalStatus(valuetmp))
	}

	return
}

// SetPartitionStyle sets the value of PartitionStyle for the instance
func (instance *MSFT_Disk) SetPropertyPartitionStyle(value Disk_PartitionStyle) (err error) {
	return instance.SetProperty("PartitionStyle", (value))
}

// GetPartitionStyle gets the value of PartitionStyle for the instance
func (instance *MSFT_Disk) GetPropertyPartitionStyle() (value Disk_PartitionStyle, err error) {
	retValue, err := instance.GetProperty("PartitionStyle")
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

	value = Disk_PartitionStyle(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *MSFT_Disk) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_Disk) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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
func (instance *MSFT_Disk) SetPropertyPhysicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("PhysicalSectorSize", (value))
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *MSFT_Disk) GetPropertyPhysicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
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

// SetProvisioningType sets the value of ProvisioningType for the instance
func (instance *MSFT_Disk) SetPropertyProvisioningType(value Disk_ProvisioningType) (err error) {
	return instance.SetProperty("ProvisioningType", (value))
}

// GetProvisioningType gets the value of ProvisioningType for the instance
func (instance *MSFT_Disk) GetPropertyProvisioningType() (value Disk_ProvisioningType, err error) {
	retValue, err := instance.GetProperty("ProvisioningType")
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

	value = Disk_ProvisioningType(valuetmp)

	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_Disk) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_Disk) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
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

// SetSignature sets the value of Signature for the instance
func (instance *MSFT_Disk) SetPropertySignature(value uint32) (err error) {
	return instance.SetProperty("Signature", (value))
}

// GetSignature gets the value of Signature for the instance
func (instance *MSFT_Disk) GetPropertySignature() (value uint32, err error) {
	retValue, err := instance.GetProperty("Signature")
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
func (instance *MSFT_Disk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_Disk) GetPropertySize() (value uint64, err error) {
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

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSFT_Disk) SetPropertyUniqueIdFormat(value Disk_UniqueIdFormat) (err error) {
	return instance.SetProperty("UniqueIdFormat", (value))
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSFT_Disk) GetPropertyUniqueIdFormat() (value Disk_UniqueIdFormat, err error) {
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

	value = Disk_UniqueIdFormat(valuetmp)

	return
}

//

// <param name="Alignment" type="uint32 "></param>
// <param name="AssignDriveLetter" type="bool "></param>
// <param name="DriveLetter" type="byte "></param>
// <param name="GptType" type="string "></param>
// <param name="IsActive" type="bool "></param>
// <param name="IsHidden" type="bool "></param>
// <param name="MbrType" type="uint16 "></param>
// <param name="Offset" type="uint64 "></param>
// <param name="Size" type="uint64 "></param>
// <param name="UseMaximumSize" type="bool "></param>

// <param name="CreatedPartition" type="MSFT_Partition "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) CreatePartition( /* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ Offset uint64,
	/* IN */ Alignment uint32,
	/* IN */ DriveLetter byte,
	/* IN */ AssignDriveLetter bool,
	/* IN */ MbrType uint16,
	/* IN */ GptType string,
	/* IN */ IsHidden bool,
	/* IN */ IsActive bool,
	/* OUT */ CreatedPartition MSFT_Partition,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreatePartition", Size, UseMaximumSize, Offset, Alignment, DriveLetter, AssignDriveLetter, MbrType, GptType, IsHidden, IsActive)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PartitionStyle" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Initialize( /* IN */ PartitionStyle uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Initialize", PartitionStyle)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RemoveData" type="bool "></param>
// <param name="RemoveOEM" type="bool "></param>
// <param name="Sanitize" type="bool "></param>
// <param name="ZeroOutEntireDisk" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Clear( /* IN */ RemoveData bool,
	/* IN */ RemoveOEM bool,
	/* IN */ ZeroOutEntireDisk bool,
	/* IN */ Sanitize bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Clear", RemoveData, RemoveOEM, ZeroOutEntireDisk, Sanitize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PartitionStyle" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) ConvertStyle( /* IN */ PartitionStyle uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ConvertStyle", PartitionStyle)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Offline( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Offline")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Online( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Online")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Guid" type="string "></param>
// <param name="IsReadOnly" type="bool "></param>
// <param name="Signature" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) SetAttributes( /* IN */ IsReadOnly bool,
	/* IN */ Signature uint32,
	/* IN */ Guid string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", IsReadOnly, Signature, Guid)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Refresh( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Refresh")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccessPath" type="string "></param>
// <param name="AllocationUnitSize" type="uint32 "></param>
// <param name="FileSystem" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVolume" type="MSFT_Volume "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) CreateVolume( /* IN */ FriendlyName string,
	/* IN */ FileSystem uint16,
	/* IN */ AccessPath string,
	/* IN */ AllocationUnitSize uint32,
	/* OUT */ CreatedVolume MSFT_Volume,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVolume", FriendlyName, FileSystem, AccessPath, AllocationUnitSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ScaleOut" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) EnableHighAvailability( /* IN */ ScaleOut bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableHighAvailability", ScaleOut)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) DisableHighAvailability( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableHighAvailability")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
