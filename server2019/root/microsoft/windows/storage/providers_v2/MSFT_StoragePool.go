// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_StoragePool struct
type MSFT_StoragePool struct {
	MSFT_StorageObject

	// Indicates the total sum of all the capacity used by this storage pool. If the pool is primordial, this will be the sum of all capacity currently allocated to concrete storage pools. If the pool is concrete, this value should be the sum of all capacity currently allocated to virtual disks and other pool metadata.
	AllocatedSize uint64

	// If TRUE, the storage pool should clear (zero out) physical disks that are removed from the pool.
	ClearOnDeallocate bool

	// Determines the default allocation behavior for virtual disks created in this pool. Enclosure aware virtual disks will intelligently pick the physical disks to use for their redundancy. If TRUE, the storage subsystem will use physical disks from different enclosures to balance the fault tolerance between two (or more) physical enclosures.
	EnclosureAwareDefault bool

	// Determines the default allocation behavior for virtual disks created in this pool. Fault domain aware virtual disks will intelligently pick the physical disks to use for their redundancy to balance the fault tolerance between two (or more) fault domain units of the specified type.
	FaultDomainAwarenessDefault StoragePool_FaultDomainAwarenessDefault

	// A user-friendly string representing the name of the storage pool. Friendly name can be set using the SetFriendlyName method.
	FriendlyName string

	// Denotes the current health status of the storage pool. Health of a storage pool is derived from the health of the backing physical disks, and whether or not the storage pool can maintain the required levels of resiliency.
	/// 0 - 'Healthy': All physical disks are present and in a healthy state.
	///1 - 'Warning': The majority of physical disks are healthy, but one or more may be failing I/O requests.
	///2 - 'Unhealthy': The majority of physical disks are unhealthy or in a failed state, and the pool no longer has data integrity.
	HealthStatus StoragePool_HealthStatus

	// Indicates whether or not the storage pool is used in a clustered environment.
	IsClustered bool

	// This property indicates whether the disks comprising this pool are able to tolerate power loss without data loss, e.g. automatically flush volatile buffers to non-volatile media after external power is disconnected.
	IsPowerProtected bool

	// If this field is set to TRUE, the storage pool is primordial. A primordial pool, also known as the 'available storage' pool is where storage capacity is drawn and returned in the creation and deletion of concrete storage pools. Primordial pools cannot be created or deleted.
	///If this field is set to FALSE, the storage pool is a concrete pool. These pools are subject to all of the management operations defined on the storage pool class. This includes creation, deletion, creation of virtual disks, etc.
	IsPrimordial bool

	// Indicates whether or not the storage pool's configuration is read-only. If TRUE, the storage pool will not allow configuration changes to itself or any of its virtual and physical disks. Note that the data on the virtual disk may still be writable.
	IsReadOnly bool

	// This field indicates the logical sector size of the storage pool, in bytes. This value is derived from the backing physical disks, as well as the preference specified at the time this storage pool was created.
	LogicalSectorSize uint64

	//
	MediaTypeDefault uint16

	// Name is a semi-unique (scoped to the owning storage subsystem), human-readable string used to identify a storage pool.
	Name string

	// Indicates the current operating conditions of the storage pool. Unlike HealthStatus, this field indicates the status of hardware, software, and infrastructure issues related to this storage pool, and can contain multiple values. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail.
	///4 - 'Stressed': indicates that the storage pool is functioning, but needs attention. Examples of 'Stressed' states are overload, overheated, and so on.
	///5 - 'Predictive Failure': indicates that the storage pool is functioning nominally but predicting a failure in the near future.
	///11 - 'In Service': describes a storage pool being configured, maintained, or otherwise administered.
	///12 - 'No Contact': indicates that the storage provider has knowledge of this storage pool, but has never been able to establish communications with it.
	///13 - 'Lost Communication': indicates that the storage pool is known to exist and has been contacted successfully in the past, but is currently unreachable.
	///10 - 'Stopped' and 14 - 'Aborted' are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the storage pool might need to be updated.
	///15 - 'Dormant': indicates that the storage pool is inactive.
	///16 - 'Supporting Entity in Error': indicates that this storage pool might be OK, but that another element, on which it is dependent, is in error.
	///
	OperationalStatus []StoragePool_OperationalStatus

	// A string representation of the vendor defined operational status. This field should only be set if the OperationalStatus array contains 1 - 'Other'.
	OtherOperationalStatusDescription string

	// If Usage is set to 1 - 'Other', this field contains the string representation of the vendor defined usage for the storage pool. This property must be NULL if Usage is not set to 1 - 'Other'.
	OtherUsageDescription string

	// This field indicates the physical sector size of the storage pool, in bytes. This value is derived from the backing physical disks for this storage pool.
	PhysicalSectorSize uint64

	// Indicates the provisioning scheme to use when creating new virtual disks on this storage pool.
	///0 - 'Unknown': May mean that this information is unavailable, or the storage pool uses a proprietary method of allocation.
	///1 - 'Thin': Storage for the virtual disk is allocated on-demand.
	///2 - 'Fixed': Storage for the virtual disk is allocated at the time of virtual disk creation.
	ProvisioningTypeDefault StoragePool_ProvisioningTypeDefault

	// Denotes the reason why the storage pool is read-only.
	///1 - 'None': The pool is not read-only.
	///2 - 'By Policy': The administrator has either requested the pool to be read-only or has enacted a policy on the system that requires the pool to be read-only.
	///3 - 'Majority Disks Unhealthy': The majority of the supporting physical disks are in an unhealthy state that has forced the storage pool into a read-only state.
	ReadOnlyReason StoragePool_ReadOnlyReason

	// This property indicates how the operating system will proceed with repairing of virtual disks for this storage pool.
	///2 - 'Sequential': repair will process one allocation slab at a time. This will result in longer repair times, but small impact on the I/O load.
	///3 - 'Parallel': repair will process as many allocation slabs as it can in parallel. This will result in the shortest repair time, but will have significant impact on I/O load.
	///
	RepairPolicy StoragePool_RepairPolicy

	// Indicates the default resiliency setting used for virtual disk creation. This default can be overridden at the time of virtual disk creation. This property's value should correspond to the resiliency setting's Name field.
	ResiliencySettingNameDefault string

	// If TRUE, the storage subsystem will automatically retire missing physical disks in this storage pool and replace them with hot-spares or other available physical disks (in the storage pool).
	RetireMissingPhysicalDisks StoragePool_RetireMissingPhysicalDisks

	// Indicates the capacity of the storage pool. If the pool is primordial, this is the sum of all the healthy physical disk sizes. If the pool is concrete, this is the sum of all associated physical disks (except hot-spares, and including failed drives).
	Size uint64

	// Denotes the provisioning schemes that this storage pool supports.
	SupportedProvisioningTypes []StoragePool_SupportedProvisioningTypes

	// If TRUE, this storage pool supports data deduplication.
	SupportsDeduplication bool

	// Percentages at which an alert should be generated
	ThinProvisioningAlertThresholds []uint16

	// Denotes the intended usage of the storage pool.
	Usage StoragePool_Usage

	// Denotes the version of this storage pool.
	Version StoragePool_Version

	// Default size of write cache for virtual disk creation
	WriteCacheSizeDefault uint64

	// Maximum size of write cache for virtual disk creation
	WriteCacheSizeMax uint64

	// Minimum size of write cache for virtual disk creation
	WriteCacheSizeMin uint64
}

// SetAllocatedSize sets the value of AllocatedSize for the instance
func (instance *MSFT_StoragePool) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", value)
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_StoragePool) GetPropertyAllocatedSize() (value uint64, err error) {
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

// SetClearOnDeallocate sets the value of ClearOnDeallocate for the instance
func (instance *MSFT_StoragePool) SetPropertyClearOnDeallocate(value bool) (err error) {
	return instance.SetProperty("ClearOnDeallocate", value)
}

// GetClearOnDeallocate gets the value of ClearOnDeallocate for the instance
func (instance *MSFT_StoragePool) GetPropertyClearOnDeallocate() (value bool, err error) {
	retValue, err := instance.GetProperty("ClearOnDeallocate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnclosureAwareDefault sets the value of EnclosureAwareDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyEnclosureAwareDefault(value bool) (err error) {
	return instance.SetProperty("EnclosureAwareDefault", value)
}

// GetEnclosureAwareDefault gets the value of EnclosureAwareDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyEnclosureAwareDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("EnclosureAwareDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomainAwarenessDefault sets the value of FaultDomainAwarenessDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyFaultDomainAwarenessDefault(value StoragePool_FaultDomainAwarenessDefault) (err error) {
	return instance.SetProperty("FaultDomainAwarenessDefault", value)
}

// GetFaultDomainAwarenessDefault gets the value of FaultDomainAwarenessDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyFaultDomainAwarenessDefault() (value StoragePool_FaultDomainAwarenessDefault, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwarenessDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_FaultDomainAwarenessDefault)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_StoragePool) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_StoragePool) GetPropertyFriendlyName() (value string, err error) {
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
func (instance *MSFT_StoragePool) SetPropertyHealthStatus(value StoragePool_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StoragePool) GetPropertyHealthStatus() (value StoragePool_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsClustered sets the value of IsClustered for the instance
func (instance *MSFT_StoragePool) SetPropertyIsClustered(value bool) (err error) {
	return instance.SetProperty("IsClustered", value)
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *MSFT_StoragePool) GetPropertyIsClustered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPowerProtected sets the value of IsPowerProtected for the instance
func (instance *MSFT_StoragePool) SetPropertyIsPowerProtected(value bool) (err error) {
	return instance.SetProperty("IsPowerProtected", value)
}

// GetIsPowerProtected gets the value of IsPowerProtected for the instance
func (instance *MSFT_StoragePool) GetPropertyIsPowerProtected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPowerProtected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPrimordial sets the value of IsPrimordial for the instance
func (instance *MSFT_StoragePool) SetPropertyIsPrimordial(value bool) (err error) {
	return instance.SetProperty("IsPrimordial", value)
}

// GetIsPrimordial gets the value of IsPrimordial for the instance
func (instance *MSFT_StoragePool) GetPropertyIsPrimordial() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPrimordial")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *MSFT_StoragePool) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", value)
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *MSFT_StoragePool) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
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
func (instance *MSFT_StoragePool) SetPropertyLogicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("LogicalSectorSize", value)
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_StoragePool) GetPropertyLogicalSectorSize() (value uint64, err error) {
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

// SetMediaTypeDefault sets the value of MediaTypeDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyMediaTypeDefault(value uint16) (err error) {
	return instance.SetProperty("MediaTypeDefault", value)
}

// GetMediaTypeDefault gets the value of MediaTypeDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyMediaTypeDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("MediaTypeDefault")
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
func (instance *MSFT_StoragePool) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StoragePool) GetPropertyName() (value string, err error) {
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StoragePool) SetPropertyOperationalStatus(value []StoragePool_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StoragePool) GetPropertyOperationalStatus() (value []StoragePool_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]StoragePool_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherOperationalStatusDescription sets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_StoragePool) SetPropertyOtherOperationalStatusDescription(value string) (err error) {
	return instance.SetProperty("OtherOperationalStatusDescription", value)
}

// GetOtherOperationalStatusDescription gets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_StoragePool) GetPropertyOtherOperationalStatusDescription() (value string, err error) {
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
func (instance *MSFT_StoragePool) SetPropertyOtherUsageDescription(value string) (err error) {
	return instance.SetProperty("OtherUsageDescription", value)
}

// GetOtherUsageDescription gets the value of OtherUsageDescription for the instance
func (instance *MSFT_StoragePool) GetPropertyOtherUsageDescription() (value string, err error) {
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

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *MSFT_StoragePool) SetPropertyPhysicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("PhysicalSectorSize", value)
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *MSFT_StoragePool) GetPropertyPhysicalSectorSize() (value uint64, err error) {
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

// SetProvisioningTypeDefault sets the value of ProvisioningTypeDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyProvisioningTypeDefault(value StoragePool_ProvisioningTypeDefault) (err error) {
	return instance.SetProperty("ProvisioningTypeDefault", value)
}

// GetProvisioningTypeDefault gets the value of ProvisioningTypeDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyProvisioningTypeDefault() (value StoragePool_ProvisioningTypeDefault, err error) {
	retValue, err := instance.GetProperty("ProvisioningTypeDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_ProvisioningTypeDefault)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadOnlyReason sets the value of ReadOnlyReason for the instance
func (instance *MSFT_StoragePool) SetPropertyReadOnlyReason(value StoragePool_ReadOnlyReason) (err error) {
	return instance.SetProperty("ReadOnlyReason", value)
}

// GetReadOnlyReason gets the value of ReadOnlyReason for the instance
func (instance *MSFT_StoragePool) GetPropertyReadOnlyReason() (value StoragePool_ReadOnlyReason, err error) {
	retValue, err := instance.GetProperty("ReadOnlyReason")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_ReadOnlyReason)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRepairPolicy sets the value of RepairPolicy for the instance
func (instance *MSFT_StoragePool) SetPropertyRepairPolicy(value StoragePool_RepairPolicy) (err error) {
	return instance.SetProperty("RepairPolicy", value)
}

// GetRepairPolicy gets the value of RepairPolicy for the instance
func (instance *MSFT_StoragePool) GetPropertyRepairPolicy() (value StoragePool_RepairPolicy, err error) {
	retValue, err := instance.GetProperty("RepairPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_RepairPolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResiliencySettingNameDefault sets the value of ResiliencySettingNameDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyResiliencySettingNameDefault(value string) (err error) {
	return instance.SetProperty("ResiliencySettingNameDefault", value)
}

// GetResiliencySettingNameDefault gets the value of ResiliencySettingNameDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyResiliencySettingNameDefault() (value string, err error) {
	retValue, err := instance.GetProperty("ResiliencySettingNameDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetireMissingPhysicalDisks sets the value of RetireMissingPhysicalDisks for the instance
func (instance *MSFT_StoragePool) SetPropertyRetireMissingPhysicalDisks(value StoragePool_RetireMissingPhysicalDisks) (err error) {
	return instance.SetProperty("RetireMissingPhysicalDisks", value)
}

// GetRetireMissingPhysicalDisks gets the value of RetireMissingPhysicalDisks for the instance
func (instance *MSFT_StoragePool) GetPropertyRetireMissingPhysicalDisks() (value StoragePool_RetireMissingPhysicalDisks, err error) {
	retValue, err := instance.GetProperty("RetireMissingPhysicalDisks")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_RetireMissingPhysicalDisks)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_StoragePool) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_StoragePool) GetPropertySize() (value uint64, err error) {
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

// SetSupportedProvisioningTypes sets the value of SupportedProvisioningTypes for the instance
func (instance *MSFT_StoragePool) SetPropertySupportedProvisioningTypes(value []StoragePool_SupportedProvisioningTypes) (err error) {
	return instance.SetProperty("SupportedProvisioningTypes", value)
}

// GetSupportedProvisioningTypes gets the value of SupportedProvisioningTypes for the instance
func (instance *MSFT_StoragePool) GetPropertySupportedProvisioningTypes() (value []StoragePool_SupportedProvisioningTypes, err error) {
	retValue, err := instance.GetProperty("SupportedProvisioningTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]StoragePool_SupportedProvisioningTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsDeduplication sets the value of SupportsDeduplication for the instance
func (instance *MSFT_StoragePool) SetPropertySupportsDeduplication(value bool) (err error) {
	return instance.SetProperty("SupportsDeduplication", value)
}

// GetSupportsDeduplication gets the value of SupportsDeduplication for the instance
func (instance *MSFT_StoragePool) GetPropertySupportsDeduplication() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsDeduplication")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThinProvisioningAlertThresholds sets the value of ThinProvisioningAlertThresholds for the instance
func (instance *MSFT_StoragePool) SetPropertyThinProvisioningAlertThresholds(value []uint16) (err error) {
	return instance.SetProperty("ThinProvisioningAlertThresholds", value)
}

// GetThinProvisioningAlertThresholds gets the value of ThinProvisioningAlertThresholds for the instance
func (instance *MSFT_StoragePool) GetPropertyThinProvisioningAlertThresholds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ThinProvisioningAlertThresholds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_StoragePool) SetPropertyUsage(value StoragePool_Usage) (err error) {
	return instance.SetProperty("Usage", value)
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_StoragePool) GetPropertyUsage() (value StoragePool_Usage, err error) {
	retValue, err := instance.GetProperty("Usage")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_Usage)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *MSFT_StoragePool) SetPropertyVersion(value StoragePool_Version) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_StoragePool) GetPropertyVersion() (value StoragePool_Version, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(StoragePool_Version)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteCacheSizeDefault sets the value of WriteCacheSizeDefault for the instance
func (instance *MSFT_StoragePool) SetPropertyWriteCacheSizeDefault(value uint64) (err error) {
	return instance.SetProperty("WriteCacheSizeDefault", value)
}

// GetWriteCacheSizeDefault gets the value of WriteCacheSizeDefault for the instance
func (instance *MSFT_StoragePool) GetPropertyWriteCacheSizeDefault() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteCacheSizeDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteCacheSizeMax sets the value of WriteCacheSizeMax for the instance
func (instance *MSFT_StoragePool) SetPropertyWriteCacheSizeMax(value uint64) (err error) {
	return instance.SetProperty("WriteCacheSizeMax", value)
}

// GetWriteCacheSizeMax gets the value of WriteCacheSizeMax for the instance
func (instance *MSFT_StoragePool) GetPropertyWriteCacheSizeMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteCacheSizeMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteCacheSizeMin sets the value of WriteCacheSizeMin for the instance
func (instance *MSFT_StoragePool) SetPropertyWriteCacheSizeMin(value uint64) (err error) {
	return instance.SetProperty("WriteCacheSizeMin", value)
}

// GetWriteCacheSizeMin gets the value of WriteCacheSizeMin for the instance
func (instance *MSFT_StoragePool) GetPropertyWriteCacheSizeMin() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteCacheSizeMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// This method creates a virtual disk using the resources of the storage pool. This method is available only when the SupportsVirtualDiskCreation property on the storage subsystem is set to TRUE. If it is set to FALSE, this method will fail with MI_RESULT_NOT_SUPPORTED. This method is also not supported for primordial pools.
///Creating tiered virtual disks is available only when the SupportsStorageTieredVirtualDiskCreation property on the storage subsystem is set to TRUE. If it is set to FALSE, this method will fail with MI_RESULT_NOT_SUPPORTED.
///CreateVirtualDisk requires only FriendlyName and Size to be specified. Sizes can be specified explicitly through the Size parameter, or you can use the maximum available space from the storage pool by specifying the UseMaximumSize parameter. Both FriendlyName and Size are treated as goals rather than hard requirements. For example, not all SMI-S based arrays support custom friendly names; however, the virtual disk creation will still succeed. If the size specified is not achieved, the actual size used for the virtual disk will be returned in the out parameter structure.
///The usage of this virtual disk can be set using the Usage and OtherUsageDescription parameters. If a value for OtherUsageDescription is given, Usage must be set to 1 - 'Other', otherwise an error will be returned.
///By default, the resiliency setting applied to this virtual disk will be whatever is specified in the storage pool's ResiliencySettingNameDefault property. This can be overridden using the ResiliencySettingName parameter. Note that the name given here must correspond to a resiliency setting associated with this storage pool. Any other value will result in an error.
///Individual settings of the resiliency setting can be overridden using the NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, and Interleave parameters. If these parameters are not used, the defaults from the resiliency setting will be used. These overrides will not persist back to the particular resiliency setting instance; however some storage providers may choose to create a new resiliency setting instance to capture this new configuration. If any of the goals specified in the override parameters are out of range, or are not supported by the storage pool, an error will be returned.
///The provisioning policy for the virtual disk is determined in a similar way to the resiliency setting. If no preference is specified in the ProvisioningType parameter, the policy is determined by the storage pool's ProvisioningTypeDefault property. If the ProvisioningType parameter is specified, the default is ignored and the value specified will be used instead.
///Allocation can be further controlled by the PhysicalDisksToUse parameter. There may be certain scenarios where a storage administrator wants to manually choose which physical disks should back the virtual disk. When this parameter is specified, data for the virtual disk will only be stored on the physical disks in this array and not on any others.

// <param name="AutoNumberOfColumns" type="bool ">If TRUE, this field instructs the storage provider (or subsystem) to automatically pick what it determines to be the best number of columns for the virtual disk. If this field is TRUE, then the NumberOfColumns parameter must be NULL.</param>
// <param name="AutoWriteCacheSize" type="bool ">Indicates whether the provider should pick up the auto write cache size</param>
// <param name="FriendlyName" type="string ">This parameter allows the user to specify the FriendlyName at the time of the virtual disk creation. FriendlyNames are expected to be descriptive, however they are not required to be unique. Note that some storage subsystems do not allow setting a friendly name during virtual disk creation. If a subsystem doesn't support this, virtual disk creation should still succeed, however the disk may have a different name assigned to it.</param>
// <param name="Interleave" type="uint64 ">Specifies the number of bytes that should be used for a strip in the common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk. Thus Interleave * NumberOfColumns will yield the size of one stripe of user data. If this parameter is specified, this value will override the InterleaveDefault which would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="IsEnclosureAware" type="bool ">Determines the allocation behavior for this virtual disk. Enclosure aware virtual disks will intelligently pick the physical disks to use for their redundancy. If TRUE, the virtual disk will attempt to use physical disks from different enclosures to balance the fault tolerance between two (or more) physical enclosures.</param>
// <param name="NumberOfColumns" type="uint16 ">Specifies the number of underlying physical disks across which data should be striped. If specified, this value will override the NumberOfColumnsDefault value that would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="NumberOfDataCopies" type="uint16 ">Specifies the number of complete data copies to maintain for this virtual disk. If specified, this value will override the NumberOfDataCopiesDefault value that would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="OtherUsageDescription" type="string ">Allows a user to set a vendor specific usage for the new virtual disk object. This parameter can only be specified if the Usage parameter is set to 1 - 'Other'.</param>
// <param name="PhysicalDiskRedundancy" type="uint16 ">Specifies how many physical disk failures the virtual disk should be able to withstand before data loss occurs. If specified, this value will override the PhysicalDiskRedundancyDefault value that would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="PhysicalDisksToUse" type="MSFT_PhysicalDisk []">If specified, allocation of this virtual disk's storage is limited to the physical disks in the list. These physical disks must already be added to this storage pool.</param>
// <param name="ProvisioningType" type="StoragePool_ProvisioningType ">Denotes the provisioning type of the virtual disk.
///1 - 'Thin': The storage for the virtual disk is allocated on-demand.
///2 - 'Fixed': The storage for the virtual disk is allocated up front.</param>
// <param name="ResiliencySettingName" type="string ">This parameter specifies the resiliency setting to use as a template for this virtual disk. This property's value should correspond with the particular resiliency setting instance's Name property. Only resiliency settings associated with this storage pool may be used.</param>
// <param name="Size" type="uint64 ">Indicates the size for the virtual disk. Note that some storage subsystems will round the size up or down to a multiple of its allocation unit size. This parameter cannot be used if UseMaximumSize is set to TRUE.</param>
// <param name="StorageTiers" type="MSFT_StorageTier []">Storage tiers on this virtual disk</param>
// <param name="StorageTierSizes" type="uint64 []">Sizes of each tier</param>
// <param name="Usage" type="StoragePool_Usage ">Denotes the intended usage of the virtual disk</param>
// <param name="UseMaximumSize" type="bool ">UseMaximumSize instructs the storage array to create the largest possible virtual disk given the available resources of this storage pool. This parameter cannot be used if the Size parameter is set.</param>
// <param name="WriteCacheSize" type="uint64 ">Size of write cache on the virtual disk</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation. When the operation has completed, an association should exist between the storage job and the created objects.</param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateVirtualDisk( /* IN */ FriendlyName string,
	/* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ ProvisioningType StoragePool_ProvisioningType,
	/* IN */ ResiliencySettingName string,
	/* IN */ Usage StoragePool_Usage,
	/* IN */ OtherUsageDescription string,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ AutoNumberOfColumns bool,
	/* IN */ Interleave uint64,
	/* IN */ IsEnclosureAware bool,
	/* IN */ PhysicalDisksToUse []MSFT_PhysicalDisk,
	/* IN */ StorageTiers []MSFT_StorageTier,
	/* IN */ StorageTierSizes []uint64,
	/* IN */ WriteCacheSize uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVirtualDisk", FriendlyName, Size, UseMaximumSize, ProvisioningType, ResiliencySettingName, Usage, OtherUsageDescription, NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, AutoNumberOfColumns, Interleave, IsEnclosureAware, PhysicalDisksToUse, StorageTiers, StorageTierSizes, WriteCacheSize, AutoWriteCacheSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllocationUnitSize" type="uint64 "></param>
// <param name="AutoNumberOfColumns" type="bool "></param>
// <param name="AutoWriteCacheSize" type="bool "></param>
// <param name="ColumnIsolation" type="uint16 "></param>
// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="NumberOfGroups" type="uint16 "></param>
// <param name="OtherUsageDescription" type="string "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="PhysicalDisksToUse" type="MSFT_PhysicalDisk []"></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ReadCacheSize" type="uint64 "></param>
// <param name="ResiliencySettingName" type="string "></param>
// <param name="Size" type="uint64 "></param>
// <param name="StorageTiers" type="MSFT_StorageTier []"></param>
// <param name="StorageTierSizes" type="uint64 []"></param>
// <param name="Usage" type="uint16 "></param>
// <param name="UseMaximumSize" type="bool "></param>
// <param name="WriteCacheSize" type="uint64 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateVirtualDisk2( /* IN */ FriendlyName string,
	/* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ ProvisioningType uint16,
	/* IN */ AllocationUnitSize uint64,
	/* IN */ MediaType uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ Usage uint16,
	/* IN */ OtherUsageDescription string,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ AutoNumberOfColumns bool,
	/* IN */ Interleave uint64,
	/* IN */ NumberOfGroups uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ PhysicalDisksToUse []MSFT_PhysicalDisk,
	/* IN */ StorageTiers []MSFT_StorageTier,
	/* IN */ StorageTierSizes []uint64,
	/* IN */ WriteCacheSize uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* IN */ ReadCacheSize uint64,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVirtualDisk2", FriendlyName, Size, UseMaximumSize, ProvisioningType, AllocationUnitSize, MediaType, ResiliencySettingName, Usage, OtherUsageDescription, NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, AutoNumberOfColumns, Interleave, NumberOfGroups, FaultDomainAwareness, ColumnIsolation, PhysicalDisksToUse, StorageTiers, StorageTierSizes, WriteCacheSize, AutoWriteCacheSize, ReadCacheSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AddToCluster" type="bool "></param>
// <param name="AllocationUnitSize" type="uint64 "></param>
// <param name="AutoNumberOfColumns" type="bool "></param>
// <param name="AutoWriteCacheSize" type="bool "></param>
// <param name="ColumnIsolation" type="uint16 "></param>
// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="IsManualAttach" type="bool "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="NumberOfGroups" type="uint16 "></param>
// <param name="OtherUsageDescription" type="string "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ReadCacheSize" type="uint64 "></param>
// <param name="ResiliencySettingName" type="string "></param>
// <param name="Size" type="uint64 "></param>
// <param name="StorageFaultDomainsToUse" type="MSFT_StorageFaultDomain []"></param>
// <param name="StorageTiers" type="MSFT_StorageTier []"></param>
// <param name="StorageTierSizes" type="uint64 []"></param>
// <param name="Usage" type="uint16 "></param>
// <param name="UseMaximumSize" type="bool "></param>
// <param name="WriteCacheSize" type="uint64 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateVirtualDisk3( /* IN */ FriendlyName string,
	/* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ ProvisioningType uint16,
	/* IN */ AllocationUnitSize uint64,
	/* IN */ MediaType uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ Usage uint16,
	/* IN */ OtherUsageDescription string,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ AutoNumberOfColumns bool,
	/* IN */ Interleave uint64,
	/* IN */ NumberOfGroups uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ StorageFaultDomainsToUse []MSFT_StorageFaultDomain,
	/* IN */ StorageTiers []MSFT_StorageTier,
	/* IN */ StorageTierSizes []uint64,
	/* IN */ WriteCacheSize uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* IN */ ReadCacheSize uint64,
	/* IN */ IsManualAttach bool,
	/* IN */ AddToCluster bool,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVirtualDisk3", FriendlyName, Size, UseMaximumSize, ProvisioningType, AllocationUnitSize, MediaType, ResiliencySettingName, Usage, OtherUsageDescription, NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, AutoNumberOfColumns, Interleave, NumberOfGroups, FaultDomainAwareness, ColumnIsolation, StorageFaultDomainsToUse, StorageTiers, StorageTierSizes, WriteCacheSize, AutoWriteCacheSize, ReadCacheSize, IsManualAttach, AddToCluster)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method creates a virtual disk and single volume using the resources of the storage pool.

// <param name="AccessPath" type="string ">If set to a valid access path, the system will attempt to use this path as a way to access the local volume. If the access path could not be set, or this parameter was left NULL, a new access path will be automatically assigned.</param>
// <param name="FileServer" type="MSFT_FileServer ">The file server that will own this volume.</param>
// <param name="FileSystem" type="StoragePool_FileSystem ">Specifies the file system to format the created volume. Specifying a CSV file system is only supported on a storage spaces subsystem. For CSV the pool must be clusterable and the volume created will be a cluster shared volume.</param>
// <param name="FriendlyName" type="string ">This parameter allows the user to specify the FriendlyName at the time of the volume creation. FriendlyNames are expected to be descriptive, however they are not required to be unique. The filesystem's label will also be set to this friendly name.</param>
// <param name="NumberOfColumns" type="uint16 ">Specifies the number of underlying physical disks across which data should be striped. If specified, this value will override the NumberOfColumnsDefault value that would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="PhysicalDiskRedundancy" type="uint16 ">Specifies how many physical disk failures the virtual disk should be able to withstand before data loss occurs. If specified, this value will override the PhysicalDiskRedundancyDefault which would have been inherited from the resiliency setting specified by ResiliencySettingName.</param>
// <param name="ProvisioningType" type="StoragePool_ProvisioningType ">Denotes the provisioning type of the volume.
///1 - 'Thin': The storage for the volume is allocated on-demand.
///2 - 'Fixed': The storage for the volume is allocated up front.</param>
// <param name="ResiliencySettingName" type="string ">This parameter specifies the resiliency setting to use as a template for this volume. This property's value should correspond with the particular resiliency setting instance's Name property. Only resiliency settings associated with this storage pool may be used.</param>
// <param name="Size" type="uint64 ">Indicates the size for the virtual disk. Note that some storage subsystems will round the size up or down to a multiple of its allocation unit size. The size of the resulting volume will be the maximum size possible for the resulting virtual disk.</param>
// <param name="StorageTiers" type="MSFT_StorageTier []">Storage tiers on this virtual disk</param>
// <param name="StorageTierSizes" type="uint64 []">Sizes of each tier</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation. When the operation has completed, an association should exist between the storage job and the created objects.</param>
// <param name="CreatedVolume" type="MSFT_Volume "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateVolume( /* IN */ FriendlyName string,
	/* IN */ Size uint64,
	/* IN */ StorageTiers []MSFT_StorageTier,
	/* IN */ StorageTierSizes []uint64,
	/* IN */ ProvisioningType StoragePool_ProvisioningType,
	/* IN */ ResiliencySettingName string,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ FileSystem StoragePool_FileSystem,
	/* IN */ AccessPath string,
	/* IN */ FileServer MSFT_FileServer,
	/* OUT */ CreatedVolume MSFT_Volume,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVolume", FriendlyName, Size, StorageTiers, StorageTierSizes, ProvisioningType, ResiliencySettingName, PhysicalDiskRedundancy, NumberOfColumns, FileSystem, AccessPath, FileServer)
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
// <param name="FileServer" type="MSFT_FileServer "></param>
// <param name="FileSystem" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ReadCacheSize" type="uint64 "></param>
// <param name="ResiliencySettingName" type="string "></param>
// <param name="Size" type="uint64 "></param>
// <param name="StorageTiers" type="MSFT_StorageTier []"></param>
// <param name="StorageTierSizes" type="uint64 []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVolume" type="MSFT_Volume "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateVolume2( /* IN */ FriendlyName string,
	/* IN */ Size uint64,
	/* IN */ StorageTiers []MSFT_StorageTier,
	/* IN */ StorageTierSizes []uint64,
	/* IN */ ProvisioningType uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ FileSystem uint16,
	/* IN */ AccessPath string,
	/* IN */ AllocationUnitSize uint32,
	/* IN */ ReadCacheSize uint64,
	/* IN */ FileServer MSFT_FileServer,
	/* OUT */ CreatedVolume MSFT_Volume,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVolume2", FriendlyName, Size, StorageTiers, StorageTierSizes, ProvisioningType, ResiliencySettingName, PhysicalDiskRedundancy, NumberOfColumns, FileSystem, AccessPath, AllocationUnitSize, ReadCacheSize, FileServer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Creates a storage tier template on the storage pool. This method is available only when the SupportsStorageTierCreation property on the storage subsystem is set to TRUE. If it is set to FALSE, this method will fail with MI_RESULT_NOT_SUPPORTED. This method is also not supported for primordial pools.

// <param name="Description" type="string ">Description of the storage tier</param>
// <param name="FriendlyName" type="string ">Friendly name of the storage tier</param>
// <param name="MediaType" type="StoragePool_MediaType ">Media type of the storage tier</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="CreatedStorageTier" type="MSFT_StorageTier "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateStorageTier( /* IN */ FriendlyName string,
	/* IN */ MediaType StoragePool_MediaType,
	/* IN */ Description string,
	/* OUT */ CreatedStorageTier MSFT_StorageTier,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateStorageTier", FriendlyName, MediaType, Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ColumnIsolation" type="uint16 "></param>
// <param name="Description" type="string "></param>
// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="NumberOfGroups" type="uint16 "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ResiliencySettingName" type="string "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedStorageTier" type="MSFT_StorageTier "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateStorageTier2( /* IN */ FriendlyName string,
	/* IN */ ProvisioningType uint16,
	/* IN */ MediaType uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ Interleave uint64,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ NumberOfGroups uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ Description string,
	/* OUT */ CreatedStorageTier MSFT_StorageTier,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateStorageTier2", FriendlyName, ProvisioningType, MediaType, FaultDomainAwareness, ColumnIsolation, ResiliencySettingName, Interleave, NumberOfDataCopies, NumberOfGroups, NumberOfColumns, PhysicalDiskRedundancy, Description)
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
// <param name="Description" type="string "></param>
// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="MediaType" type="uint16 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="NumberOfGroups" type="uint16 "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="ResiliencySettingName" type="string "></param>
// <param name="StorageFaultDomainsToUse" type="MSFT_StorageFaultDomain []"></param>
// <param name="Usage" type="uint16 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedStorageTier" type="MSFT_StorageTier "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) CreateStorageTier3( /* IN */ FriendlyName string,
	/* IN */ ProvisioningType uint16,
	/* IN */ AllocationUnitSize uint64,
	/* IN */ MediaType uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ StorageFaultDomainsToUse []MSFT_StorageFaultDomain,
	/* IN */ ResiliencySettingName string,
	/* IN */ Usage uint16,
	/* IN */ Interleave uint64,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ NumberOfGroups uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ Description string,
	/* OUT */ CreatedStorageTier MSFT_StorageTier,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateStorageTier3", FriendlyName, ProvisioningType, AllocationUnitSize, MediaType, FaultDomainAwareness, ColumnIsolation, StorageFaultDomainsToUse, ResiliencySettingName, Usage, Interleave, NumberOfDataCopies, NumberOfGroups, NumberOfColumns, PhysicalDiskRedundancy, Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method deletes an empty storage pool. If the storage pool contains any virtual disks, these virtual disks should be removed first.

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method will upgrade the version of the storage pool.

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) Upgrade( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Upgrade")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) Optimize( /* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Optimize", RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method will add one or more physical disks from the primordial storage pool to an existing concrete storage pool.

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>
// <param name="Usage" type="StoragePool_Usage "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) AddPhysicalDisk( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* IN */ Usage StoragePool_Usage,
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
func (instance *MSFT_StoragePool) AddPhysicalDisk2( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
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

// This method removes one or more physical disks from the pool and returns all previously allocated space on the disk to the available capacity in the primordial pool.

// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) RemovePhysicalDisk( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
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
func (instance *MSFT_StoragePool) RemovePhysicalDisk2( /* IN */ PhysicalDisks []MSFT_PhysicalDisk,
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

// This method returns the supported sizes for a virtual disk created on this storage pool. These sizes can either be returned in an array of all supported sizes, through a min, max, and divisor, or both.

// <param name="ResiliencySettingName" type="string ">Specifies the name of the resiliency setting that should be used when determining the supported sizes. Note that the sizes returned may be different depending on the resiliency setting.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedSizes" type="uint64 []">This output parameter will contain an array of all of the supported sizes by the storage pool. This parameter may be NULL if the number of supported sizes is large, but is useful for storage pools that support only a select number of virtual disk sizes.</param>
// <param name="VirtualDiskSizeDivisor" type="uint64 ">This parameter indicates the interval in which the supported sizes increment. For example: If the minimum supported size is 10 GB, and this parameter is 2 GB, then the supported sizes for this pool would be 10 GB, 12 GB, 14 GB, etc. until the maximum supported size is reached.</param>
// <param name="VirtualDiskSizeMax" type="uint64 ">This parameter denotes the maximum supported size that a virtual disk created in this pool can be.</param>
// <param name="VirtualDiskSizeMin" type="uint64 ">This parameter denotes the minimum supported size that a virtual disk created in this pool can be.</param>
func (instance *MSFT_StoragePool) GetSupportedSize( /* IN */ ResiliencySettingName string,
	/* OUT */ SupportedSizes []uint64,
	/* OUT */ VirtualDiskSizeMin uint64,
	/* OUT */ VirtualDiskSizeMax uint64,
	/* OUT */ VirtualDiskSizeDivisor uint64,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedSize", ResiliencySettingName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="ResiliencySettingName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedSizes" type="uint64 []"></param>
// <param name="VirtualDiskSizeDivisor" type="uint64 "></param>
// <param name="VirtualDiskSizeMax" type="uint64 "></param>
// <param name="VirtualDiskSizeMin" type="uint64 "></param>
func (instance *MSFT_StoragePool) GetSupportedSize2( /* IN */ ResiliencySettingName string,
	/* IN */ FaultDomainAwareness uint16,
	/* OUT */ SupportedSizes []uint64,
	/* OUT */ VirtualDiskSizeMin uint64,
	/* OUT */ VirtualDiskSizeMax uint64,
	/* OUT */ VirtualDiskSizeDivisor uint64,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSupportedSize2", ResiliencySettingName, FaultDomainAwareness)
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
func (instance *MSFT_StoragePool) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
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
func (instance *MSFT_StoragePool) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the storage pool to be renamed.

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the storage pool's intended usage to be updated. Not all storage pools may allow this and will return 1 - 'Not Supported' if this operation cannot be performed.

// <param name="OtherUsageDescription" type="string ">If Usage is set to 1 - 'Other', this parameter takes in the string representation of a vendor defined usage for this storage pool. This parameter must not be set if Usage is a value other than 1 - 'Other'.</param>
// <param name="Usage" type="StoragePool_Usage ">Denotes the new intended usage of the storage pool.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) SetUsage( /* IN */ Usage StoragePool_Usage,
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

// This method allows the user to update or set various defaults on the storage pool. Note that not all parameters must be specified, and only those given will be updated.

// <param name="AutoWriteCacheSize" type="bool ">Indicates whether the provider should pick up the auto write cache size</param>
// <param name="EnclosureAwareDefault" type="bool ">This parameter indicates the default allocation policy for virtual disks created in an enclosure aware storage pool. For example, an enclosure aware subsystem could balance each data copy of the virtual disk across multiple physical enclosures such that each enclosure contains a full data copy of the virtual disk.</param>
// <param name="ProvisioningTypeDefault" type="StoragePool_ProvisioningTypeDefault ">Specifies the new default provisioning type of the storage pool.</param>
// <param name="ResiliencySettingNameDefault" type="string ">Specifies the new default resiliency setting that should be used by this storage pool. The resiliency setting specified must already be associated with this storage pool.</param>
// <param name="WriteCacheSizeDefault" type="uint64 ">New default size of write cache for virtual disk creation</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) SetDefaults( /* IN */ ProvisioningTypeDefault StoragePool_ProvisioningTypeDefault,
	/* IN */ ResiliencySettingNameDefault string,
	/* IN */ EnclosureAwareDefault bool,
	/* IN */ WriteCacheSizeDefault uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDefaults", ProvisioningTypeDefault, ResiliencySettingNameDefault, EnclosureAwareDefault, WriteCacheSizeDefault, AutoWriteCacheSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AutoWriteCacheSize" type="bool "></param>
// <param name="FaultDomainAwarenessDefault" type="uint16 "></param>
// <param name="MediaTypeDefault" type="uint16 "></param>
// <param name="ProvisioningTypeDefault" type="uint16 "></param>
// <param name="ResiliencySettingNameDefault" type="string "></param>
// <param name="WriteCacheSizeDefault" type="uint64 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) SetDefaults2( /* IN */ ProvisioningTypeDefault uint16,
	/* IN */ MediaTypeDefault uint16,
	/* IN */ ResiliencySettingNameDefault string,
	/* IN */ FaultDomainAwarenessDefault uint16,
	/* IN */ WriteCacheSizeDefault uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDefaults2", ProvisioningTypeDefault, MediaTypeDefault, ResiliencySettingNameDefault, FaultDomainAwarenessDefault, WriteCacheSizeDefault, AutoWriteCacheSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the user to update or set various attributes on the storage pool. Note that not all parameters must be specified, and only those given will be updated.

// <param name="ClearOnDeallocate" type="bool "></param>
// <param name="IsPowerProtected" type="bool "></param>
// <param name="IsReadOnly" type="bool "></param>
// <param name="RepairPolicy" type="StoragePool_RepairPolicy "></param>
// <param name="RetireMissingPhysicalDisks" type="StoragePool_RetireMissingPhysicalDisks "></param>
// <param name="ThinProvisioningAlertThresholds" type="uint16 []">Percentages at which an alert should be generated</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StoragePool) SetAttributes( /* IN */ IsReadOnly bool,
	/* IN */ ClearOnDeallocate bool,
	/* IN */ IsPowerProtected bool,
	/* IN */ RepairPolicy StoragePool_RepairPolicy,
	/* IN */ RetireMissingPhysicalDisks StoragePool_RetireMissingPhysicalDisks,
	/* IN */ ThinProvisioningAlertThresholds []uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", IsReadOnly, ClearOnDeallocate, IsPowerProtected, RepairPolicy, RetireMissingPhysicalDisks, ThinProvisioningAlertThresholds)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
