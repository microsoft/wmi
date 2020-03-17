// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_VirtualDisk struct
type MSFT_VirtualDisk struct {
	MSFT_StorageObject

	// Indicates whether the virtual disk is available for read and/or write access
	Access VirtualDisk_Access

	// The currently allocated size of the virtual disk. If the virtual disk's ProvisioningType is 2 - 'Fixed', this value should equal Size. If the ProvisioningType is 1 - 'Thin', this value is the amount of space actually allocated (i.e. some value less than Size).
	AllocatedSize uint64

	// Specifies the allocation unit size in bytes for this virtual disk.
	AllocationUnitSize uint64

	//
	ColumnIsolation uint16

	// Denotes the reason why this virtual disk is detached. This field will only be set when the virtual disk's OperationalStatus includes 0xD002 - 'Detached'. Note that this field is specific to Storage Spaces.
	DetachedReason VirtualDisk_DetachedReason

	// Determines the current allocation behavior for this virtual disk. Fault domain aware virtual disks will intelligently pick the physical disks to use for their redundancy to balance the fault tolerance between two (or more) fault domain units of the specified type.
	FaultDomainAwareness VirtualDisk_FaultDomainAwareness

	// This field indicates the total storage pool capacity being consumed by this virtual disk. For example: in the case of a 2-way mirrored virtual disk of size 1 GB, the footprint on the pool will be approximately 2 GB.
	FootprintOnPool uint64

	// A user-settable, display-oriented string representing the name of the virtual disk.
	FriendlyName string

	// Denotes the current health status of the virtual disk. Health of a virtual disk is derived from the health of the backing physical disks, and whether or not the virtual disk can maintain the required levels of resiliency.
	/// 0 - 'Healthy': All physical disks are present and in a healthy state.
	///1 - 'Warning': The majority of physical disks are healthy, but one or more may be failing I/O requests.
	///2 - 'Unhealthy': The majority of physical disks are unhealthy or in a failed state, and the virtual disk no longer has data integrity.
	HealthStatus VirtualDisk_HealthStatus

	// This field indicates the number of bytes that will form a strip in common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk. Thus, Interleave * NumberOfColumns will yield the size of one stripe of user data.
	Interleave uint64

	//
	IsDeduplicationEnabled bool

	// Determines the current allocation behavior for this virtual disk. Enclosure aware virtual disks will intelligently pick the physical disks to use for their redundancy. If TRUE, the virtual disk will attempt to use physical disks from different enclosures to balance the fault tolerance between two (or more) physical enclosures.
	IsEnclosureAware bool

	// If TRUE, this virtual disk will only be attached to the system if an explicit call is made to the Attach method. Note that this property is specific to Storage Spaces.
	IsManualAttach bool

	// Indicates whether this virtual disk is a snapshot of another virtual disk
	IsSnapshot bool

	// Indicates whether or not there are tiers associated with this virtual disk.
	IsTiered bool

	//
	LogicalSectorSize uint64

	//
	MediaType uint16

	// Name is a semi-unique (scoped to the owning storage subsystem), human-readable string used to identify the virtual disk.
	Name string

	//
	NameFormat VirtualDisk_NameFormat

	//
	NumberOfAvailableCopies uint16

	// This field indicates the number of underlying physical disks across which data for this virtual disk is striped.
	NumberOfColumns uint16

	// This field indicates the number of complete data copies that are being maintained. For example, RAID 5 maintains 1 copy of data, whereas RAID 1 maintains at least 2 copies.
	NumberOfDataCopies uint16

	//
	NumberOfGroups uint16

	// Indicates the current operating conditions of the virtual disk. Unlike HealthStatus, this field indicates the status of hardware, software, and infrastructure issues related to this virtual disk, and can contain multiple values. Various operational statuses are defined.
	///11 - 'In Service': describes a virtual disk being configured, maintained, or otherwise administered.
	///0xD002 - 'Detached': This value is reserved for Windows. This value indicates a virtual disk that is visible to the host system but does not have a disk device object.
	///0xD003 - 'Incomplete': describes a virtual disk which does not have enough redundancy remaining to successfully repair or regenerate its data.
	OperationalStatus []VirtualDisk_OperationalStatus

	// If OperationalStatus contains 1 - 'Other', this field contains the string representing the vendor defined operational status. This property must be NULL if OperationalStatus does not contain 1 - 'Other'.
	OtherOperationalStatusDescription string

	// If the virtual disk's Usage field is set to 1 - 'Other', this field must contain a description of the vendor or user defined usage. If Usage is not set to 1 - 'Other', this field must not be set.
	OtherUsageDescription string

	// This field indicates what type of parity layout is being used for parity resiliency settings. This field should be NULL if the virtual disk does not use a parity resiliency setting.
	ParityLayout VirtualDisk_ParityLayout

	// This field indicates how many backing physical disks can fail without compromising data redundancy. For example: RAID 0 cannot tolerate any failures, RAID 5 can tolerate a single drive failure, and RAID 6 can tolerate 2 failures.
	PhysicalDiskRedundancy uint16

	//
	PhysicalSectorSize uint64

	// Denotes the provisioning scheme of the virtual disk.
	///1 - 'Thin' indicates that the virtual disk's capacity is allocated on demand.
	///2 - 'Fixed' indicates that the virtual disk's capacity is fully allocated upon creation.
	ProvisioningType VirtualDisk_ProvisioningType

	// Size of the read cache for the virtual disk
	ReadCacheSize uint64

	//
	RequestNoSinglePointOfFailure bool

	// The name of the resiliency setting used to create this virtual disk.
	ResiliencySettingName string

	// The logical size of the virtual disk measured in bytes
	Size uint64

	// UniqueIdFormat indicates the type of identifier used in the UniqueId field. The identifier used in UniqueId must be the highest available identifier using the following order of preference: 8 (highest), 3, 2, 1, 0 (lowest). For example: if the virtual disk device exposes identifiers of type 0, 1, and 3, UniqueId must be the identifier of type 3, and UniqueIdFormat should be set to 3.
	UniqueIdFormat VirtualDisk_UniqueIdFormat

	// Certain values for UniqueIdFormat may include various sub-formats. This field is a free-form string used to describe the specific format used in UniqueId.
	UniqueIdFormatDescription string

	// This field indicates the intended usage for this virtual disk.
	Usage VirtualDisk_Usage

	// Size of the write cache for the virtual disk
	WriteCacheSize uint64
}

// SetAccess sets the value of Access for the instance
func (instance *MSFT_VirtualDisk) SetPropertyAccess(value VirtualDisk_Access) (err error) {
	return instance.SetProperty("Access", value)
}

// GetAccess gets the value of Access for the instance
func (instance *MSFT_VirtualDisk) GetPropertyAccess() (value VirtualDisk_Access, err error) {
	retValue, err := instance.GetProperty("Access")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_Access)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllocatedSize sets the value of AllocatedSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", value)
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyAllocatedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatedSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllocationUnitSize sets the value of AllocationUnitSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyAllocationUnitSize(value uint64) (err error) {
	return instance.SetProperty("AllocationUnitSize", value)
}

// GetAllocationUnitSize gets the value of AllocationUnitSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyAllocationUnitSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationUnitSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColumnIsolation sets the value of ColumnIsolation for the instance
func (instance *MSFT_VirtualDisk) SetPropertyColumnIsolation(value uint16) (err error) {
	return instance.SetProperty("ColumnIsolation", value)
}

// GetColumnIsolation gets the value of ColumnIsolation for the instance
func (instance *MSFT_VirtualDisk) GetPropertyColumnIsolation() (value uint16, err error) {
	retValue, err := instance.GetProperty("ColumnIsolation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDetachedReason sets the value of DetachedReason for the instance
func (instance *MSFT_VirtualDisk) SetPropertyDetachedReason(value VirtualDisk_DetachedReason) (err error) {
	return instance.SetProperty("DetachedReason", value)
}

// GetDetachedReason gets the value of DetachedReason for the instance
func (instance *MSFT_VirtualDisk) GetPropertyDetachedReason() (value VirtualDisk_DetachedReason, err error) {
	retValue, err := instance.GetProperty("DetachedReason")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_DetachedReason)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomainAwareness sets the value of FaultDomainAwareness for the instance
func (instance *MSFT_VirtualDisk) SetPropertyFaultDomainAwareness(value VirtualDisk_FaultDomainAwareness) (err error) {
	return instance.SetProperty("FaultDomainAwareness", value)
}

// GetFaultDomainAwareness gets the value of FaultDomainAwareness for the instance
func (instance *MSFT_VirtualDisk) GetPropertyFaultDomainAwareness() (value VirtualDisk_FaultDomainAwareness, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwareness")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_FaultDomainAwareness)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFootprintOnPool sets the value of FootprintOnPool for the instance
func (instance *MSFT_VirtualDisk) SetPropertyFootprintOnPool(value uint64) (err error) {
	return instance.SetProperty("FootprintOnPool", value)
}

// GetFootprintOnPool gets the value of FootprintOnPool for the instance
func (instance *MSFT_VirtualDisk) GetPropertyFootprintOnPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("FootprintOnPool")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_VirtualDisk) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_VirtualDisk) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_VirtualDisk) SetPropertyHealthStatus(value VirtualDisk_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_VirtualDisk) GetPropertyHealthStatus() (value VirtualDisk_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterleave sets the value of Interleave for the instance
func (instance *MSFT_VirtualDisk) SetPropertyInterleave(value uint64) (err error) {
	return instance.SetProperty("Interleave", value)
}

// GetInterleave gets the value of Interleave for the instance
func (instance *MSFT_VirtualDisk) GetPropertyInterleave() (value uint64, err error) {
	retValue, err := instance.GetProperty("Interleave")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDeduplicationEnabled sets the value of IsDeduplicationEnabled for the instance
func (instance *MSFT_VirtualDisk) SetPropertyIsDeduplicationEnabled(value bool) (err error) {
	return instance.SetProperty("IsDeduplicationEnabled", value)
}

// GetIsDeduplicationEnabled gets the value of IsDeduplicationEnabled for the instance
func (instance *MSFT_VirtualDisk) GetPropertyIsDeduplicationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDeduplicationEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEnclosureAware sets the value of IsEnclosureAware for the instance
func (instance *MSFT_VirtualDisk) SetPropertyIsEnclosureAware(value bool) (err error) {
	return instance.SetProperty("IsEnclosureAware", value)
}

// GetIsEnclosureAware gets the value of IsEnclosureAware for the instance
func (instance *MSFT_VirtualDisk) GetPropertyIsEnclosureAware() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnclosureAware")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsManualAttach sets the value of IsManualAttach for the instance
func (instance *MSFT_VirtualDisk) SetPropertyIsManualAttach(value bool) (err error) {
	return instance.SetProperty("IsManualAttach", value)
}

// GetIsManualAttach gets the value of IsManualAttach for the instance
func (instance *MSFT_VirtualDisk) GetPropertyIsManualAttach() (value bool, err error) {
	retValue, err := instance.GetProperty("IsManualAttach")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsSnapshot sets the value of IsSnapshot for the instance
func (instance *MSFT_VirtualDisk) SetPropertyIsSnapshot(value bool) (err error) {
	return instance.SetProperty("IsSnapshot", value)
}

// GetIsSnapshot gets the value of IsSnapshot for the instance
func (instance *MSFT_VirtualDisk) GetPropertyIsSnapshot() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSnapshot")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsTiered sets the value of IsTiered for the instance
func (instance *MSFT_VirtualDisk) SetPropertyIsTiered(value bool) (err error) {
	return instance.SetProperty("IsTiered", value)
}

// GetIsTiered gets the value of IsTiered for the instance
func (instance *MSFT_VirtualDisk) GetPropertyIsTiered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTiered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyLogicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("LogicalSectorSize", value)
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyLogicalSectorSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *MSFT_VirtualDisk) SetPropertyMediaType(value uint16) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *MSFT_VirtualDisk) GetPropertyMediaType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_VirtualDisk) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_VirtualDisk) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *MSFT_VirtualDisk) SetPropertyNameFormat(value VirtualDisk_NameFormat) (err error) {
	return instance.SetProperty("NameFormat", value)
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *MSFT_VirtualDisk) GetPropertyNameFormat() (value VirtualDisk_NameFormat, err error) {
	retValue, err := instance.GetProperty("NameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_NameFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfAvailableCopies sets the value of NumberOfAvailableCopies for the instance
func (instance *MSFT_VirtualDisk) SetPropertyNumberOfAvailableCopies(value uint16) (err error) {
	return instance.SetProperty("NumberOfAvailableCopies", value)
}

// GetNumberOfAvailableCopies gets the value of NumberOfAvailableCopies for the instance
func (instance *MSFT_VirtualDisk) GetPropertyNumberOfAvailableCopies() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfAvailableCopies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfColumns sets the value of NumberOfColumns for the instance
func (instance *MSFT_VirtualDisk) SetPropertyNumberOfColumns(value uint16) (err error) {
	return instance.SetProperty("NumberOfColumns", value)
}

// GetNumberOfColumns gets the value of NumberOfColumns for the instance
func (instance *MSFT_VirtualDisk) GetPropertyNumberOfColumns() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColumns")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfDataCopies sets the value of NumberOfDataCopies for the instance
func (instance *MSFT_VirtualDisk) SetPropertyNumberOfDataCopies(value uint16) (err error) {
	return instance.SetProperty("NumberOfDataCopies", value)
}

// GetNumberOfDataCopies gets the value of NumberOfDataCopies for the instance
func (instance *MSFT_VirtualDisk) GetPropertyNumberOfDataCopies() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfDataCopies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfGroups sets the value of NumberOfGroups for the instance
func (instance *MSFT_VirtualDisk) SetPropertyNumberOfGroups(value uint16) (err error) {
	return instance.SetProperty("NumberOfGroups", value)
}

// GetNumberOfGroups gets the value of NumberOfGroups for the instance
func (instance *MSFT_VirtualDisk) GetPropertyNumberOfGroups() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfGroups")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_VirtualDisk) SetPropertyOperationalStatus(value []VirtualDisk_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_VirtualDisk) GetPropertyOperationalStatus() (value []VirtualDisk_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]VirtualDisk_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherOperationalStatusDescription sets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_VirtualDisk) SetPropertyOtherOperationalStatusDescription(value string) (err error) {
	return instance.SetProperty("OtherOperationalStatusDescription", value)
}

// GetOtherOperationalStatusDescription gets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_VirtualDisk) GetPropertyOtherOperationalStatusDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherOperationalStatusDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherUsageDescription sets the value of OtherUsageDescription for the instance
func (instance *MSFT_VirtualDisk) SetPropertyOtherUsageDescription(value string) (err error) {
	return instance.SetProperty("OtherUsageDescription", value)
}

// GetOtherUsageDescription gets the value of OtherUsageDescription for the instance
func (instance *MSFT_VirtualDisk) GetPropertyOtherUsageDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherUsageDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParityLayout sets the value of ParityLayout for the instance
func (instance *MSFT_VirtualDisk) SetPropertyParityLayout(value VirtualDisk_ParityLayout) (err error) {
	return instance.SetProperty("ParityLayout", value)
}

// GetParityLayout gets the value of ParityLayout for the instance
func (instance *MSFT_VirtualDisk) GetPropertyParityLayout() (value VirtualDisk_ParityLayout, err error) {
	retValue, err := instance.GetProperty("ParityLayout")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_ParityLayout)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskRedundancy sets the value of PhysicalDiskRedundancy for the instance
func (instance *MSFT_VirtualDisk) SetPropertyPhysicalDiskRedundancy(value uint16) (err error) {
	return instance.SetProperty("PhysicalDiskRedundancy", value)
}

// GetPhysicalDiskRedundancy gets the value of PhysicalDiskRedundancy for the instance
func (instance *MSFT_VirtualDisk) GetPropertyPhysicalDiskRedundancy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskRedundancy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyPhysicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("PhysicalSectorSize", value)
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyPhysicalSectorSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProvisioningType sets the value of ProvisioningType for the instance
func (instance *MSFT_VirtualDisk) SetPropertyProvisioningType(value VirtualDisk_ProvisioningType) (err error) {
	return instance.SetProperty("ProvisioningType", value)
}

// GetProvisioningType gets the value of ProvisioningType for the instance
func (instance *MSFT_VirtualDisk) GetPropertyProvisioningType() (value VirtualDisk_ProvisioningType, err error) {
	retValue, err := instance.GetProperty("ProvisioningType")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_ProvisioningType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadCacheSize sets the value of ReadCacheSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyReadCacheSize(value uint64) (err error) {
	return instance.SetProperty("ReadCacheSize", value)
}

// GetReadCacheSize gets the value of ReadCacheSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyReadCacheSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadCacheSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestNoSinglePointOfFailure sets the value of RequestNoSinglePointOfFailure for the instance
func (instance *MSFT_VirtualDisk) SetPropertyRequestNoSinglePointOfFailure(value bool) (err error) {
	return instance.SetProperty("RequestNoSinglePointOfFailure", value)
}

// GetRequestNoSinglePointOfFailure gets the value of RequestNoSinglePointOfFailure for the instance
func (instance *MSFT_VirtualDisk) GetPropertyRequestNoSinglePointOfFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("RequestNoSinglePointOfFailure")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResiliencySettingName sets the value of ResiliencySettingName for the instance
func (instance *MSFT_VirtualDisk) SetPropertyResiliencySettingName(value string) (err error) {
	return instance.SetProperty("ResiliencySettingName", value)
}

// GetResiliencySettingName gets the value of ResiliencySettingName for the instance
func (instance *MSFT_VirtualDisk) GetPropertyResiliencySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("ResiliencySettingName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_VirtualDisk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_VirtualDisk) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSFT_VirtualDisk) SetPropertyUniqueIdFormat(value VirtualDisk_UniqueIdFormat) (err error) {
	return instance.SetProperty("UniqueIdFormat", value)
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSFT_VirtualDisk) GetPropertyUniqueIdFormat() (value VirtualDisk_UniqueIdFormat, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_UniqueIdFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueIdFormatDescription sets the value of UniqueIdFormatDescription for the instance
func (instance *MSFT_VirtualDisk) SetPropertyUniqueIdFormatDescription(value string) (err error) {
	return instance.SetProperty("UniqueIdFormatDescription", value)
}

// GetUniqueIdFormatDescription gets the value of UniqueIdFormatDescription for the instance
func (instance *MSFT_VirtualDisk) GetPropertyUniqueIdFormatDescription() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormatDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_VirtualDisk) SetPropertyUsage(value VirtualDisk_Usage) (err error) {
	return instance.SetProperty("Usage", value)
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_VirtualDisk) GetPropertyUsage() (value VirtualDisk_Usage, err error) {
	retValue, err := instance.GetProperty("Usage")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualDisk_Usage)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteCacheSize sets the value of WriteCacheSize for the instance
func (instance *MSFT_VirtualDisk) SetPropertyWriteCacheSize(value uint64) (err error) {
	return instance.SetProperty("WriteCacheSize", value)
}

// GetWriteCacheSize gets the value of WriteCacheSize for the instance
func (instance *MSFT_VirtualDisk) GetPropertyWriteCacheSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteCacheSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="PhysicalExtents" type="MSFT_PhysicalExtent []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) GetPhysicalExtent( /* OUT */ PhysicalExtents []MSFT_PhysicalExtent,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalExtent")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method deletes the virtual disk. After this method is called, the space used by the virtual disk will be reclaimed and the user will be unable to reverse the delete operation.

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method shows a virtual disk to an initiator. This operation is also known as 'exposing' or 'unmasking' a virtual disk.

// <param name="HostType" type="VirtualDisk_HostType ">This field indicates the operating system type running on the host of the initiator port.</param>
// <param name="InitiatorAddress" type="string ">The address of the initiator to which the virtual disk should be shown</param>
// <param name="TargetPortAddresses" type="string []">An array of target port addresses from which the virtual disk should be shown</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) Show( /* IN */ TargetPortAddresses []string,
	/* IN */ InitiatorAddress string,
	/* IN */ HostType VirtualDisk_HostType,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Show", TargetPortAddresses, InitiatorAddress, HostType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method hides a virtual disk from an initiator. This operation is also known as 'unexposing' or 'masking' a virtual disk.

// <param name="InitiatorAddress" type="string ">The address of the initiator to which the virtual disk should be hidden</param>
// <param name="TargetPortAddresses" type="string []">An array of target port addresses from which the virtual disk should be hidden. Note: this array may contain a subset of the addresses originally given in Show.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) Hide( /* IN */ TargetPortAddresses []string,
	/* IN */ InitiatorAddress string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Hide", TargetPortAddresses, InitiatorAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method creates a point in time snapshot of the virtual disk.

// <param name="FriendlyName" type="string ">The desired name of the snapshot virtual disk</param>
// <param name="TargetStoragePoolName" type="string ">This field indicates which storage pool should be used to hold the created snapshot. If this field is not set, this method will default to using the same storage pool that contains the source virtual disk.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) CreateSnapshot( /* IN */ FriendlyName string,
	/* IN */ TargetStoragePoolName string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateSnapshot", FriendlyName, TargetStoragePoolName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method creates a clone of the virtual disk, resulting in another virtual disk with identical data to the source.

// <param name="FriendlyName" type="string ">The desired name of the virtual disk clone</param>
// <param name="TargetStoragePoolName" type="string ">This field indicates which storage pool should be used to hold the created clone. If this field is not set, this method will default to using the same storage pool that contains the source virtual disk.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) CreateClone( /* IN */ FriendlyName string,
	/* IN */ TargetStoragePoolName string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateClone", FriendlyName, TargetStoragePoolName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a virtual disk to be resized. The size specified must be in the range of valid values given by the GetSupportedSize method on the storage pool object.

// <param name="Size" type="uint64 ">As input, this parameter contains the requested size for the virtual disk to become. As output, this parameter contains the size that was actually achieved after the resize operation.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Size" type="uint64 ">As input, this parameter contains the requested size for the virtual disk to become. As output, this parameter contains the size that was actually achieved after the resize operation.</param>
func (instance *MSFT_VirtualDisk) Resize( /* IN/OUT */ Size uint64,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Resize")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method initiates a repair of the virtual disk - restoring data and redundancy to different (or new) physical disks within the storage pool.

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) Repair( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Repair")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method returns the security descriptor that controls access to this specific object instance.

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the access control list of the object.</param>
func (instance *MSFT_VirtualDisk) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user with sufficient privileges to set the security descriptor that control access to this specific object instance. If the call is not made in the context of a user specified in the security descriptor's access control list, this method will fail with 40001 - 'Access Denied'. If an empty security descriptor is passed to this function, the behavior is left to the specific implementation so long as there is some user context (typically domain administrators) that can access and administer the object.

// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the desired access control list for this object.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the virtual disk to be renamed.

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the virtual disk's intended usage to be updated. Not all virtual disks may allow this and will return 1 - 'Not Supported' if this operation cannot be performed.

// <param name="OtherUsageDescription" type="string ">If Usage is set to 1 - 'Other', this parameter takes in the string representation of a vendor defined usage for this virtual disk. This parameter must not be set if Usage is a value other than 1 - 'Other'.</param>
// <param name="Usage" type="VirtualDisk_Usage "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetUsage( /* IN */ Usage VirtualDisk_Usage,
	/* IN */ OtherUsageDescription string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetUsage", Usage, OtherUsageDescription)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the user to update or set various attributes on the virtual disk. Note that not all parameters must be specified, and only those given will be updated.

// <param name="Access" type="VirtualDisk_Access "></param>
// <param name="IsManualAttach" type="bool "></param>
// <param name="StorageNodeName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetAttributes( /* IN */ IsManualAttach bool,
	/* IN */ StorageNodeName string,
	/* IN */ Access VirtualDisk_Access,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", IsManualAttach, StorageNodeName, Access)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllocationUnitSize" type="uint64 "></param>
// <param name="ColumnIsolation" type="uint16 "></param>
// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="NumberOfGroups" type="uint16 "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ResiliencySettingName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetProperties( /* IN */ ProvisioningType uint16,
	/* IN */ AllocationUnitSize uint64,
	/* IN */ MediaType uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ NumberOfGroups uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ Interleave uint64,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetProperties", ProvisioningType, AllocationUnitSize, MediaType, FaultDomainAwareness, ColumnIsolation, ResiliencySettingName, PhysicalDiskRedundancy, NumberOfDataCopies, NumberOfGroups, NumberOfColumns, Interleave)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Attaches a Storage Spaces based virtual disk to the system. This operation is similar to Show and Hide, however there is no need for target and initiator configuration since everything is done locally. Depending on the system's NewDiskPolicy (formerly SAN policy), a Storage Space may need to be Attached before it can be used.

// <param name="StorageNodeName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) Attach( /* IN */ StorageNodeName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Attach", StorageNodeName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Detaches a Storage Spaces based virtual disk from the system. This operation is similar to Hide, however there is no need for target and initiator configuration since everything is done locally. Detaching a Storage Space will result in it's corresponding disk object to be suprise removed from the system. Note that detaching can happen in response to certain failure and warning conditions (such as failing redundancy, or thin provisioning capacity limits being reached).

// <param name="StorageNodeName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) Detach( /* IN */ StorageNodeName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Detach", StorageNodeName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method will add one or more physical disks for manual allocation.

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>
// <param name="Usage" type="VirtualDisk_Usage "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) AddPhysicalDisk( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* IN */ Usage VirtualDisk_Usage,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddPhysicalDisk", PhysicalDisks, Usage)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>
// <param name="Usage" type="uint16 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) AddPhysicalDisk2( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* IN */ Usage uint16,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddPhysicalDisk2", PhysicalDisks, Usage)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method will remove one or more physical disks from manual allocation.

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) RemovePhysicalDisk( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemovePhysicalDisk", PhysicalDisks)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) RemovePhysicalDisk2( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemovePhysicalDisk2", PhysicalDisks)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="StorageFaultDomains" type="MSFT_StorageFaultDomain []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) AddStorageFaultDomain( /* IN */ StorageFaultDomains []MSFT_StorageFaultDomain,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddStorageFaultDomain", StorageFaultDomains)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="StorageFaultDomains" type="MSFT_StorageFaultDomain []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) RemoveStorageFaultDomain( /* IN */ StorageFaultDomains []MSFT_StorageFaultDomain,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveStorageFaultDomain", StorageFaultDomains)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FriendlyName" type="string "></param>
// <param name="RecoveryPointObjective" type="uint16 "></param>
// <param name="ReplicationSettings" type="MSFT_ReplicationSettings "></param>
// <param name="SyncType" type="uint16 "></param>
// <param name="TargetStoragePoolObjectId" type="string "></param>
// <param name="TargetStorageSubsystem" type="MSFT_ReplicaPeer "></param>
// <param name="TargetVirtualDiskObjectId" type="string "></param>

// <param name="CreatedReplicaPeer" type="MSFT_ReplicaPeer "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) CreateReplica( /* IN */ FriendlyName string,
	/* IN */ TargetStorageSubsystem MSFT_ReplicaPeer,
	/* IN */ TargetVirtualDiskObjectId string,
	/* IN */ TargetStoragePoolObjectId string,
	/* IN */ RecoveryPointObjective uint16,
	/* IN */ ReplicationSettings MSFT_ReplicationSettings,
	/* IN */ SyncType uint16,
	/* OUT */ CreatedReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplica", FriendlyName, TargetStorageSubsystem, TargetVirtualDiskObjectId, TargetStoragePoolObjectId, RecoveryPointObjective, ReplicationSettings, SyncType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Operation" type="uint16 "></param>
// <param name="VirtualDiskReplicaPeer" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VirtualDisk) SetReplicationRelationship( /* IN */ Operation uint16,
	/* IN */ VirtualDiskReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetReplicationRelationship", Operation, VirtualDiskReplicaPeer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
