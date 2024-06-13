// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageSubSystem struct
type MSFT_StorageSubSystem struct {
	*MSFT_StorageObject

	// Denotes whether this subsystem supports automatic object clustering.
	AutomaticClusteringEnabled bool

	// This field denotes the cache level that has been discovered. This corresponds to the storage provider's DiscoveryLevel parameter in the Discover method.
	///0 - 'Level 0': The storage provider and storage subsystem objects have been discovered.
	///1 - 'Level 1': Storage pools, resiliency settings, target ports, target portals, and initiator ids belonging to this subsystem have been discovered.
	///2 - 'Level 2': Virtual disks and masking sets belonging to this subsystem have been discovered.
	///3 - 'Level 3': Physical disks belonging to this subsystem have been discovered.
	CurrentCacheLevel StorageSubSystem_CurrentCacheLevel

	// Denotes whether storage tiers are supported by the subsystem.
	DataTieringType StorageSubSystem_DataTieringType

	// A user settable description of the storage subsystem. This field can be used to store extra free-form information, such as notes or details about the subsystem's intended usage.
	Description string

	// Determines the default allocation behavior for storage pools created in this subsystem. If the subsystem does not support storage pool creation, then it determines the default allocation behavior for virtual disks created in this subsystem.
	FaultDomainAwarenessDefault StorageSubSystem_FaultDomainAwarenessDefault

	// This field is a string representation of the subsystem's firmware version.
	FirmwareVersion string

	// A user settable string representing the name of the storage subsystem. The storage provider or subsystem is expected to supply an initial value for this field.
	FriendlyName string

	// Denotes the health of the subsystem.
	///0 - 'Healthy': Indicates that the subsystem is functioning normally.
	///1 - 'Warning': Indicates that the subsystem is still functioning, but has detected errors or issues that may require administrator intervention.
	///2 - 'Unhealthy': Indicates that the subsystem is not functioning due to errors or failures. The subsystem needs immediate attention from an administrator.
	HealthStatus StorageSubSystem_HealthStatus

	// Denotes the iSCSI Target Creation Scheme supported by the subsystem.
	///0 - 'Not Applicable' implies a non-iSCSI subsystem.
	///1 - 'Not Supported' implies the subsystem does not allow creation of a Target.
	///2 - 'Manual' implies the subsystem allows manual creation of the Target.
	///3 - 'Auto' implies the subsystem automatically creates a Target.
	///
	iSCSITargetCreationScheme StorageSubSystem_iSCSITargetCreationScheme

	// This field is a string representation of the company responsible for creating the storage subsystem hardware.
	Manufacturer string

	// If TRUE, the storage provider supports the use of the DeviceNumbers parameter of the CreateMaskingSet and AddVirtualDisk methods.
	MaskingClientSelectableDeviceNumbers bool

	// Indicates the maximum number of masking sets that a particular virtual disk can be added to.
	MaskingMapCountMax uint16

	// If TRUE, the subsystem will only allow one initiator to be added to a masking set.
	MaskingOneInitiatorIdPerView bool

	// If MaskingValidInitiatorIdTypes contains the value 1 - 'Other', this field is used to enumerate the other valid initiator id types for this storage subsystem.
	MaskingOtherValidInitiatorIdTypes []string

	// Indicates the number of target ports that can be used for masking a virtual disk. This applies to both masking sets and the virtual disk Show method.
	MaskingPortsPerView StorageSubSystem_MaskingPortsPerView

	// Indicates which address formats can be inferred by the storage provider and subsystem when working with initiator ids.
	MaskingValidInitiatorIdTypes []StorageSubSystem_MaskingValidInitiatorIdTypes

	// This field is a string representation of the model number of the subsystem array.
	Model string

	// Name is a globally unique, human-readable string used to identify a storage subsystem.
	Name string

	// NameFormat describes the format of the Name identifier.
	NameFormat StorageSubSystem_NameFormat

	// Denotes the total number of physical disk slots in the subsystem or enclosure.
	NumberOfSlots uint32

	// Indicates the current statuses of the subsystem. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail.
	///4 - 'Stressed': indicates that the subsystem is functioning, but needs attention. Examples of 'Stressed' states are overload, overheated, and so on.
	///5 - 'Predictive Failure': indicates that the subsystem is functioning nominally but predicting a failure in the near future.
	///11 - 'In Service': describes a subsystem being configured, maintained, cleaned, or otherwise administered.
	///12 - 'No Contact': indicates that the storage provider has knowledge of this subsystem, but has never been able to establish communications with it.
	///13 - 'Lost Communication': indicates that the subsystem is known to exist and has been contacted successfully in the past, but is currently unreachable.
	///10 - 'Stopped' and 14 - 'Aborted' are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the subsystem might need to be updated.
	///15 - 'Dormant': indicates that the subsystem is inactive.
	///16 - 'Supporting Entity in Error': indicates that this subsystem might be OK, but that another element, on which it is dependent, is in error.
	///
	OperationalStatus []StorageSubSystem_OperationalStatus

	// When the corresponding array entry in SupportedHostType[] is "Other", this entry provides a string describing the manufacturer and OS/Environment. When the corresponding SupportedHostType[] entry is not "Other", this entry allows variations or qualifications of ClientTypes - for example, different versions of Solaris.
	OtherHostTypeDescription []string

	// This field is an array of custom identifier for the subsystem. If this field is set, the OtherIdentifyingInfoDescription field must also be set.
	OtherIdentifyingInfo []string

	// An array of string description of the format used in the custom identifiers defined in the OtherIdentifyingInfo field. There must be a 1:1 mapping between this array and OtherIdentifyingInfo.
	OtherIdentifyingInfoDescription []string

	// A string representation of the vendor defined operational status. This field should only be set if the OperationalStatus array contains 1 - 'Other'.
	OtherOperationalStatusDescription string

	// Denotes the minimum number of physical disks required for creating a storage pool on this subsystem.
	PhysicalDisksPerStoragePoolMin uint16

	// This field is reserved for future releases.
	ReplicasPerSourceCloneMax uint16

	// This field is reserved for future releases.
	ReplicasPerSourceMirrorMax uint16

	// This field is reserved for future releases.
	ReplicasPerSourceSnapshotMax uint16

	// This field is a string representation of the serial number of the subsystem array.
	SerialNumber string

	// The storage transport on this subsystem.
	StorageConnectionType uint16

	// Denotes the file system types supported for Deduplication in this subsystem.
	SupportedDeduplicationFileSystemTypes []StorageSubSystem_SupportedDeduplicationFileSystemTypes

	// Denotes the object types supported for Deduplication in this subsystem.
	SupportedDeduplicationObjectTypes []StorageSubSystem_SupportedDeduplicationObjectTypes

	// This field describes the protocols supported by file servers on this subsystem.
	SupportedFileServerProtocols []StorageSubSystem_SupportedFileServerProtocols

	// File systems supported on this subsystem.
	SupportedFileSystems []StorageSubSystem_SupportedFileSystems

	// An array representing the supported host types.
	SupportedHostType []StorageSubSystem_SupportedHostType

	// If TRUE, the CreateVirtualDisk method on the storage subsystem is supported.
	SupportsAutomaticStoragePoolSelection bool

	// Denotes whether this subsystem supports local cloning. This field must be true if the VirtualDisk::CreateClone method is implemented.
	SupportsCloneLocal bool

	// Denotes whether this subsystem supports remote cloning.
	SupportsCloneRemote bool

	// Denotes whether this subsystem supports continuously available (CA) file servers.
	SupportsContinuouslyAvailableFileServer bool

	// Denotes whether this subsystem supports a file server.
	SupportsFileServer bool

	// Denotes whether this subsystem supports creation of a file server.
	SupportsFileServerCreation bool

	// If TRUE, the storage subsystem supports showing and hiding (masking) a virtual disk to a host initiator through the Show/Hide methods of the virtual disk and by the use of masking sets.
	SupportsMaskingVirtualDiskToHosts bool

	// Denotes whether this subsystem supports local mirror replication.
	SupportsMirrorLocal bool

	// Denotes whether this subsystem supports remote mirror replication.
	SupportsMirrorRemote bool

	// If TRUE, all resiliency settings will be copied from the primordial pool and added to a concrete pool upon its creation. If FALSE, the storage pool should copy the setting specified in the ResiliencySettingNameDefault parameter of CreateStoragePool. If no name was given, the resiliency setting specified by the primordial pool's ResiliencySettingNameDefault property should be used.
	SupportsMultipleResiliencySettingsPerStoragePool bool

	// Denotes whether this subsystem supports local snapshotting. This field must be true if the VirtualDisk::CreateSnapshot method is implemented.
	SupportsSnapshotLocal bool

	// Denotes whether this subsystem supports remote snapshotting.
	SupportsSnapshotRemote bool

	// If TRUE, storage pools on this subsystem support capacity expansion through adding more physical disks.
	SupportsStoragePoolAddPhysicalDisk bool

	// If TRUE, this subsystem supports the ability to create new concrete storage pools from one or more physical disks. If FALSE, either the subsystem uses pre-created storage pools, or it does not support storage pools.
	SupportsStoragePoolCreation bool

	// If TRUE, this subsystem supports the deletion of its storage pools.
	SupportsStoragePoolDeletion bool

	//
	SupportsStoragePoolFriendlyNameModification bool

	// If TRUE, storage pools on this subsystem support the replacement or removal of physical disks by use of the RemovePhysicalDisk method on the storage pool instance.
	SupportsStoragePoolRemovePhysicalDisk bool

	// If TRUE, this subsystem supports the ability to create new storage tiers. If FALSE, either the subsystem uses pre-created storage tiers, or it does not support storage tiers.
	SupportsStorageTierCreation bool

	// If TRUE, this subsystem supports the deletion of storage tiers.
	SupportsStorageTierDeletion bool

	// If TRUE, this subsystem supports the creation of tiered virtual disks.
	SupportsStorageTieredVirtualDiskCreation bool

	// If TRUE, this subsystem supports the modification of the storage tier friendly name.
	SupportsStorageTierFriendlyNameModification bool

	// If TRUE, this subsystem supports the resizing of storage tiers.
	SupportsStorageTierResize bool

	// Indicates if the subsystem allows a virtual disk to be grown in size (using the Resize method of the virtual disk instance).
	SupportsVirtualDiskCapacityExpansion bool

	// Indicates if the subsystem allows a virtual disk to be reduced in size (using the Resize method of the virtual disk instance).
	SupportsVirtualDiskCapacityReduction bool

	// Denotes whether a user can create a virtual disk by using the CreateVirtualDisk method on either the storage subsystem or storage pool objects.
	SupportsVirtualDiskCreation bool

	// Denotes whether a user can delete a virtual disk through the use of the DeleteObject extrinsic method on the virtual disk instance.
	SupportsVirtualDiskDeletion bool

	// Denotes whether a user can modify attributes or other properties on a virtual disk by using the various Set* extrinsic methods. (For example: SetFriendlyname ).
	SupportsVirtualDiskModification bool

	// Indicates if the subsystem supports explicit repairing of a virtual disk through the Repair method of the virtual disk instance.
	SupportsVirtualDiskRepair bool

	// Denotes whether this subsystem supports direct creation of volumes on a storage pool.
	SupportsVolumeCreation bool

	// Tag is an identifier for the subsystem that is independent from any location-based information. Examples of a tag could be the subsystem's serial number or asset tag.
	Tag string

	// Denotes whether virtual disk repair is enabled on this subsystem.
	VirtualDiskRepairEnabled bool

	// Denotes the virtual disk repair queue depth policy in this subsystem.
	VirtualDiskRepairQueueDepth uint32
}

func NewMSFT_StorageSubSystemEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageSubSystem, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystem{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_StorageSubSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageSubSystem, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystem{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetAutomaticClusteringEnabled sets the value of AutomaticClusteringEnabled for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyAutomaticClusteringEnabled(value bool) (err error) {
	return instance.SetProperty("AutomaticClusteringEnabled", (value))
}

// GetAutomaticClusteringEnabled gets the value of AutomaticClusteringEnabled for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyAutomaticClusteringEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticClusteringEnabled")
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

// SetCurrentCacheLevel sets the value of CurrentCacheLevel for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyCurrentCacheLevel(value StorageSubSystem_CurrentCacheLevel) (err error) {
	return instance.SetProperty("CurrentCacheLevel", (value))
}

// GetCurrentCacheLevel gets the value of CurrentCacheLevel for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyCurrentCacheLevel() (value StorageSubSystem_CurrentCacheLevel, err error) {
	retValue, err := instance.GetProperty("CurrentCacheLevel")
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

	value = StorageSubSystem_CurrentCacheLevel(valuetmp)

	return
}

// SetDataTieringType sets the value of DataTieringType for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyDataTieringType(value StorageSubSystem_DataTieringType) (err error) {
	return instance.SetProperty("DataTieringType", (value))
}

// GetDataTieringType gets the value of DataTieringType for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyDataTieringType() (value StorageSubSystem_DataTieringType, err error) {
	retValue, err := instance.GetProperty("DataTieringType")
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

	value = StorageSubSystem_DataTieringType(valuetmp)

	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetFaultDomainAwarenessDefault sets the value of FaultDomainAwarenessDefault for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyFaultDomainAwarenessDefault(value StorageSubSystem_FaultDomainAwarenessDefault) (err error) {
	return instance.SetProperty("FaultDomainAwarenessDefault", (value))
}

// GetFaultDomainAwarenessDefault gets the value of FaultDomainAwarenessDefault for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyFaultDomainAwarenessDefault() (value StorageSubSystem_FaultDomainAwarenessDefault, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwarenessDefault")
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

	value = StorageSubSystem_FaultDomainAwarenessDefault(valuetmp)

	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyFirmwareVersion() (value string, err error) {
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
func (instance *MSFT_StorageSubSystem) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyFriendlyName() (value string, err error) {
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyHealthStatus(value StorageSubSystem_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyHealthStatus() (value StorageSubSystem_HealthStatus, err error) {
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

	value = StorageSubSystem_HealthStatus(valuetmp)

	return
}

// SetiSCSITargetCreationScheme sets the value of iSCSITargetCreationScheme for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyiSCSITargetCreationScheme(value StorageSubSystem_iSCSITargetCreationScheme) (err error) {
	return instance.SetProperty("iSCSITargetCreationScheme", (value))
}

// GetiSCSITargetCreationScheme gets the value of iSCSITargetCreationScheme for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyiSCSITargetCreationScheme() (value StorageSubSystem_iSCSITargetCreationScheme, err error) {
	retValue, err := instance.GetProperty("iSCSITargetCreationScheme")
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

	value = StorageSubSystem_iSCSITargetCreationScheme(valuetmp)

	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyManufacturer() (value string, err error) {
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

// SetMaskingClientSelectableDeviceNumbers sets the value of MaskingClientSelectableDeviceNumbers for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingClientSelectableDeviceNumbers(value bool) (err error) {
	return instance.SetProperty("MaskingClientSelectableDeviceNumbers", (value))
}

// GetMaskingClientSelectableDeviceNumbers gets the value of MaskingClientSelectableDeviceNumbers for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingClientSelectableDeviceNumbers() (value bool, err error) {
	retValue, err := instance.GetProperty("MaskingClientSelectableDeviceNumbers")
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

// SetMaskingMapCountMax sets the value of MaskingMapCountMax for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingMapCountMax(value uint16) (err error) {
	return instance.SetProperty("MaskingMapCountMax", (value))
}

// GetMaskingMapCountMax gets the value of MaskingMapCountMax for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingMapCountMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaskingMapCountMax")
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

// SetMaskingOneInitiatorIdPerView sets the value of MaskingOneInitiatorIdPerView for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingOneInitiatorIdPerView(value bool) (err error) {
	return instance.SetProperty("MaskingOneInitiatorIdPerView", (value))
}

// GetMaskingOneInitiatorIdPerView gets the value of MaskingOneInitiatorIdPerView for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingOneInitiatorIdPerView() (value bool, err error) {
	retValue, err := instance.GetProperty("MaskingOneInitiatorIdPerView")
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

// SetMaskingOtherValidInitiatorIdTypes sets the value of MaskingOtherValidInitiatorIdTypes for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingOtherValidInitiatorIdTypes(value []string) (err error) {
	return instance.SetProperty("MaskingOtherValidInitiatorIdTypes", (value))
}

// GetMaskingOtherValidInitiatorIdTypes gets the value of MaskingOtherValidInitiatorIdTypes for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingOtherValidInitiatorIdTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("MaskingOtherValidInitiatorIdTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMaskingPortsPerView sets the value of MaskingPortsPerView for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingPortsPerView(value StorageSubSystem_MaskingPortsPerView) (err error) {
	return instance.SetProperty("MaskingPortsPerView", (value))
}

// GetMaskingPortsPerView gets the value of MaskingPortsPerView for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingPortsPerView() (value StorageSubSystem_MaskingPortsPerView, err error) {
	retValue, err := instance.GetProperty("MaskingPortsPerView")
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

	value = StorageSubSystem_MaskingPortsPerView(valuetmp)

	return
}

// SetMaskingValidInitiatorIdTypes sets the value of MaskingValidInitiatorIdTypes for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyMaskingValidInitiatorIdTypes(value []StorageSubSystem_MaskingValidInitiatorIdTypes) (err error) {
	return instance.SetProperty("MaskingValidInitiatorIdTypes", (value))
}

// GetMaskingValidInitiatorIdTypes gets the value of MaskingValidInitiatorIdTypes for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyMaskingValidInitiatorIdTypes() (value []StorageSubSystem_MaskingValidInitiatorIdTypes, err error) {
	retValue, err := instance.GetProperty("MaskingValidInitiatorIdTypes")
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
		value = append(value, StorageSubSystem_MaskingValidInitiatorIdTypes(valuetmp))
	}

	return
}

// SetModel sets the value of Model for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyModel() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNameFormat sets the value of NameFormat for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyNameFormat(value StorageSubSystem_NameFormat) (err error) {
	return instance.SetProperty("NameFormat", (value))
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyNameFormat() (value StorageSubSystem_NameFormat, err error) {
	retValue, err := instance.GetProperty("NameFormat")
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

	value = StorageSubSystem_NameFormat(valuetmp)

	return
}

// SetNumberOfSlots sets the value of NumberOfSlots for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyNumberOfSlots(value uint32) (err error) {
	return instance.SetProperty("NumberOfSlots", (value))
}

// GetNumberOfSlots gets the value of NumberOfSlots for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyNumberOfSlots() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSlots")
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyOperationalStatus(value []StorageSubSystem_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyOperationalStatus() (value []StorageSubSystem_OperationalStatus, err error) {
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
		value = append(value, StorageSubSystem_OperationalStatus(valuetmp))
	}

	return
}

// SetOtherHostTypeDescription sets the value of OtherHostTypeDescription for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyOtherHostTypeDescription(value []string) (err error) {
	return instance.SetProperty("OtherHostTypeDescription", (value))
}

// GetOtherHostTypeDescription gets the value of OtherHostTypeDescription for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyOtherHostTypeDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherHostTypeDescription")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", (value))
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherIdentifyingInfoDescription sets the value of OtherIdentifyingInfoDescription for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyOtherIdentifyingInfoDescription(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfoDescription", (value))
}

// GetOtherIdentifyingInfoDescription gets the value of OtherIdentifyingInfoDescription for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyOtherIdentifyingInfoDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfoDescription")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherOperationalStatusDescription sets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyOtherOperationalStatusDescription(value string) (err error) {
	return instance.SetProperty("OtherOperationalStatusDescription", (value))
}

// GetOtherOperationalStatusDescription gets the value of OtherOperationalStatusDescription for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyOtherOperationalStatusDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherOperationalStatusDescription")
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

// SetPhysicalDisksPerStoragePoolMin sets the value of PhysicalDisksPerStoragePoolMin for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyPhysicalDisksPerStoragePoolMin(value uint16) (err error) {
	return instance.SetProperty("PhysicalDisksPerStoragePoolMin", (value))
}

// GetPhysicalDisksPerStoragePoolMin gets the value of PhysicalDisksPerStoragePoolMin for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyPhysicalDisksPerStoragePoolMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDisksPerStoragePoolMin")
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

// SetReplicasPerSourceCloneMax sets the value of ReplicasPerSourceCloneMax for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyReplicasPerSourceCloneMax(value uint16) (err error) {
	return instance.SetProperty("ReplicasPerSourceCloneMax", (value))
}

// GetReplicasPerSourceCloneMax gets the value of ReplicasPerSourceCloneMax for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyReplicasPerSourceCloneMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicasPerSourceCloneMax")
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

// SetReplicasPerSourceMirrorMax sets the value of ReplicasPerSourceMirrorMax for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyReplicasPerSourceMirrorMax(value uint16) (err error) {
	return instance.SetProperty("ReplicasPerSourceMirrorMax", (value))
}

// GetReplicasPerSourceMirrorMax gets the value of ReplicasPerSourceMirrorMax for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyReplicasPerSourceMirrorMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicasPerSourceMirrorMax")
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

// SetReplicasPerSourceSnapshotMax sets the value of ReplicasPerSourceSnapshotMax for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyReplicasPerSourceSnapshotMax(value uint16) (err error) {
	return instance.SetProperty("ReplicasPerSourceSnapshotMax", (value))
}

// GetReplicasPerSourceSnapshotMax gets the value of ReplicasPerSourceSnapshotMax for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyReplicasPerSourceSnapshotMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicasPerSourceSnapshotMax")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySerialNumber() (value string, err error) {
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

// SetStorageConnectionType sets the value of StorageConnectionType for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyStorageConnectionType(value uint16) (err error) {
	return instance.SetProperty("StorageConnectionType", (value))
}

// GetStorageConnectionType gets the value of StorageConnectionType for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyStorageConnectionType() (value uint16, err error) {
	retValue, err := instance.GetProperty("StorageConnectionType")
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

// SetSupportedDeduplicationFileSystemTypes sets the value of SupportedDeduplicationFileSystemTypes for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportedDeduplicationFileSystemTypes(value []StorageSubSystem_SupportedDeduplicationFileSystemTypes) (err error) {
	return instance.SetProperty("SupportedDeduplicationFileSystemTypes", (value))
}

// GetSupportedDeduplicationFileSystemTypes gets the value of SupportedDeduplicationFileSystemTypes for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportedDeduplicationFileSystemTypes() (value []StorageSubSystem_SupportedDeduplicationFileSystemTypes, err error) {
	retValue, err := instance.GetProperty("SupportedDeduplicationFileSystemTypes")
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
		value = append(value, StorageSubSystem_SupportedDeduplicationFileSystemTypes(valuetmp))
	}

	return
}

// SetSupportedDeduplicationObjectTypes sets the value of SupportedDeduplicationObjectTypes for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportedDeduplicationObjectTypes(value []StorageSubSystem_SupportedDeduplicationObjectTypes) (err error) {
	return instance.SetProperty("SupportedDeduplicationObjectTypes", (value))
}

// GetSupportedDeduplicationObjectTypes gets the value of SupportedDeduplicationObjectTypes for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportedDeduplicationObjectTypes() (value []StorageSubSystem_SupportedDeduplicationObjectTypes, err error) {
	retValue, err := instance.GetProperty("SupportedDeduplicationObjectTypes")
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
		value = append(value, StorageSubSystem_SupportedDeduplicationObjectTypes(valuetmp))
	}

	return
}

// SetSupportedFileServerProtocols sets the value of SupportedFileServerProtocols for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportedFileServerProtocols(value []StorageSubSystem_SupportedFileServerProtocols) (err error) {
	return instance.SetProperty("SupportedFileServerProtocols", (value))
}

// GetSupportedFileServerProtocols gets the value of SupportedFileServerProtocols for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportedFileServerProtocols() (value []StorageSubSystem_SupportedFileServerProtocols, err error) {
	retValue, err := instance.GetProperty("SupportedFileServerProtocols")
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
		value = append(value, StorageSubSystem_SupportedFileServerProtocols(valuetmp))
	}

	return
}

// SetSupportedFileSystems sets the value of SupportedFileSystems for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportedFileSystems(value []StorageSubSystem_SupportedFileSystems) (err error) {
	return instance.SetProperty("SupportedFileSystems", (value))
}

// GetSupportedFileSystems gets the value of SupportedFileSystems for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportedFileSystems() (value []StorageSubSystem_SupportedFileSystems, err error) {
	retValue, err := instance.GetProperty("SupportedFileSystems")
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
		value = append(value, StorageSubSystem_SupportedFileSystems(valuetmp))
	}

	return
}

// SetSupportedHostType sets the value of SupportedHostType for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportedHostType(value []StorageSubSystem_SupportedHostType) (err error) {
	return instance.SetProperty("SupportedHostType", (value))
}

// GetSupportedHostType gets the value of SupportedHostType for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportedHostType() (value []StorageSubSystem_SupportedHostType, err error) {
	retValue, err := instance.GetProperty("SupportedHostType")
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
		value = append(value, StorageSubSystem_SupportedHostType(valuetmp))
	}

	return
}

// SetSupportsAutomaticStoragePoolSelection sets the value of SupportsAutomaticStoragePoolSelection for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsAutomaticStoragePoolSelection(value bool) (err error) {
	return instance.SetProperty("SupportsAutomaticStoragePoolSelection", (value))
}

// GetSupportsAutomaticStoragePoolSelection gets the value of SupportsAutomaticStoragePoolSelection for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsAutomaticStoragePoolSelection() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsAutomaticStoragePoolSelection")
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

// SetSupportsCloneLocal sets the value of SupportsCloneLocal for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsCloneLocal(value bool) (err error) {
	return instance.SetProperty("SupportsCloneLocal", (value))
}

// GetSupportsCloneLocal gets the value of SupportsCloneLocal for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsCloneLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCloneLocal")
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

// SetSupportsCloneRemote sets the value of SupportsCloneRemote for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsCloneRemote(value bool) (err error) {
	return instance.SetProperty("SupportsCloneRemote", (value))
}

// GetSupportsCloneRemote gets the value of SupportsCloneRemote for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsCloneRemote() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCloneRemote")
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

// SetSupportsContinuouslyAvailableFileServer sets the value of SupportsContinuouslyAvailableFileServer for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsContinuouslyAvailableFileServer(value bool) (err error) {
	return instance.SetProperty("SupportsContinuouslyAvailableFileServer", (value))
}

// GetSupportsContinuouslyAvailableFileServer gets the value of SupportsContinuouslyAvailableFileServer for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsContinuouslyAvailableFileServer() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsContinuouslyAvailableFileServer")
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

// SetSupportsFileServer sets the value of SupportsFileServer for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsFileServer(value bool) (err error) {
	return instance.SetProperty("SupportsFileServer", (value))
}

// GetSupportsFileServer gets the value of SupportsFileServer for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsFileServer() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsFileServer")
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

// SetSupportsFileServerCreation sets the value of SupportsFileServerCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsFileServerCreation(value bool) (err error) {
	return instance.SetProperty("SupportsFileServerCreation", (value))
}

// GetSupportsFileServerCreation gets the value of SupportsFileServerCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsFileServerCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsFileServerCreation")
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

// SetSupportsMaskingVirtualDiskToHosts sets the value of SupportsMaskingVirtualDiskToHosts for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsMaskingVirtualDiskToHosts(value bool) (err error) {
	return instance.SetProperty("SupportsMaskingVirtualDiskToHosts", (value))
}

// GetSupportsMaskingVirtualDiskToHosts gets the value of SupportsMaskingVirtualDiskToHosts for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsMaskingVirtualDiskToHosts() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMaskingVirtualDiskToHosts")
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

// SetSupportsMirrorLocal sets the value of SupportsMirrorLocal for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsMirrorLocal(value bool) (err error) {
	return instance.SetProperty("SupportsMirrorLocal", (value))
}

// GetSupportsMirrorLocal gets the value of SupportsMirrorLocal for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsMirrorLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMirrorLocal")
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

// SetSupportsMirrorRemote sets the value of SupportsMirrorRemote for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsMirrorRemote(value bool) (err error) {
	return instance.SetProperty("SupportsMirrorRemote", (value))
}

// GetSupportsMirrorRemote gets the value of SupportsMirrorRemote for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsMirrorRemote() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMirrorRemote")
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

// SetSupportsMultipleResiliencySettingsPerStoragePool sets the value of SupportsMultipleResiliencySettingsPerStoragePool for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsMultipleResiliencySettingsPerStoragePool(value bool) (err error) {
	return instance.SetProperty("SupportsMultipleResiliencySettingsPerStoragePool", (value))
}

// GetSupportsMultipleResiliencySettingsPerStoragePool gets the value of SupportsMultipleResiliencySettingsPerStoragePool for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsMultipleResiliencySettingsPerStoragePool() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMultipleResiliencySettingsPerStoragePool")
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

// SetSupportsSnapshotLocal sets the value of SupportsSnapshotLocal for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsSnapshotLocal(value bool) (err error) {
	return instance.SetProperty("SupportsSnapshotLocal", (value))
}

// GetSupportsSnapshotLocal gets the value of SupportsSnapshotLocal for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsSnapshotLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsSnapshotLocal")
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

// SetSupportsSnapshotRemote sets the value of SupportsSnapshotRemote for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsSnapshotRemote(value bool) (err error) {
	return instance.SetProperty("SupportsSnapshotRemote", (value))
}

// GetSupportsSnapshotRemote gets the value of SupportsSnapshotRemote for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsSnapshotRemote() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsSnapshotRemote")
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

// SetSupportsStoragePoolAddPhysicalDisk sets the value of SupportsStoragePoolAddPhysicalDisk for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStoragePoolAddPhysicalDisk(value bool) (err error) {
	return instance.SetProperty("SupportsStoragePoolAddPhysicalDisk", (value))
}

// GetSupportsStoragePoolAddPhysicalDisk gets the value of SupportsStoragePoolAddPhysicalDisk for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStoragePoolAddPhysicalDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStoragePoolAddPhysicalDisk")
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

// SetSupportsStoragePoolCreation sets the value of SupportsStoragePoolCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStoragePoolCreation(value bool) (err error) {
	return instance.SetProperty("SupportsStoragePoolCreation", (value))
}

// GetSupportsStoragePoolCreation gets the value of SupportsStoragePoolCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStoragePoolCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStoragePoolCreation")
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

// SetSupportsStoragePoolDeletion sets the value of SupportsStoragePoolDeletion for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStoragePoolDeletion(value bool) (err error) {
	return instance.SetProperty("SupportsStoragePoolDeletion", (value))
}

// GetSupportsStoragePoolDeletion gets the value of SupportsStoragePoolDeletion for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStoragePoolDeletion() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStoragePoolDeletion")
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

// SetSupportsStoragePoolFriendlyNameModification sets the value of SupportsStoragePoolFriendlyNameModification for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStoragePoolFriendlyNameModification(value bool) (err error) {
	return instance.SetProperty("SupportsStoragePoolFriendlyNameModification", (value))
}

// GetSupportsStoragePoolFriendlyNameModification gets the value of SupportsStoragePoolFriendlyNameModification for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStoragePoolFriendlyNameModification() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStoragePoolFriendlyNameModification")
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

// SetSupportsStoragePoolRemovePhysicalDisk sets the value of SupportsStoragePoolRemovePhysicalDisk for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStoragePoolRemovePhysicalDisk(value bool) (err error) {
	return instance.SetProperty("SupportsStoragePoolRemovePhysicalDisk", (value))
}

// GetSupportsStoragePoolRemovePhysicalDisk gets the value of SupportsStoragePoolRemovePhysicalDisk for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStoragePoolRemovePhysicalDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStoragePoolRemovePhysicalDisk")
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

// SetSupportsStorageTierCreation sets the value of SupportsStorageTierCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStorageTierCreation(value bool) (err error) {
	return instance.SetProperty("SupportsStorageTierCreation", (value))
}

// GetSupportsStorageTierCreation gets the value of SupportsStorageTierCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStorageTierCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStorageTierCreation")
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

// SetSupportsStorageTierDeletion sets the value of SupportsStorageTierDeletion for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStorageTierDeletion(value bool) (err error) {
	return instance.SetProperty("SupportsStorageTierDeletion", (value))
}

// GetSupportsStorageTierDeletion gets the value of SupportsStorageTierDeletion for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStorageTierDeletion() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStorageTierDeletion")
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

// SetSupportsStorageTieredVirtualDiskCreation sets the value of SupportsStorageTieredVirtualDiskCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStorageTieredVirtualDiskCreation(value bool) (err error) {
	return instance.SetProperty("SupportsStorageTieredVirtualDiskCreation", (value))
}

// GetSupportsStorageTieredVirtualDiskCreation gets the value of SupportsStorageTieredVirtualDiskCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStorageTieredVirtualDiskCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStorageTieredVirtualDiskCreation")
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

// SetSupportsStorageTierFriendlyNameModification sets the value of SupportsStorageTierFriendlyNameModification for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStorageTierFriendlyNameModification(value bool) (err error) {
	return instance.SetProperty("SupportsStorageTierFriendlyNameModification", (value))
}

// GetSupportsStorageTierFriendlyNameModification gets the value of SupportsStorageTierFriendlyNameModification for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStorageTierFriendlyNameModification() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStorageTierFriendlyNameModification")
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

// SetSupportsStorageTierResize sets the value of SupportsStorageTierResize for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsStorageTierResize(value bool) (err error) {
	return instance.SetProperty("SupportsStorageTierResize", (value))
}

// GetSupportsStorageTierResize gets the value of SupportsStorageTierResize for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsStorageTierResize() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsStorageTierResize")
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

// SetSupportsVirtualDiskCapacityExpansion sets the value of SupportsVirtualDiskCapacityExpansion for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskCapacityExpansion(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskCapacityExpansion", (value))
}

// GetSupportsVirtualDiskCapacityExpansion gets the value of SupportsVirtualDiskCapacityExpansion for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskCapacityExpansion() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskCapacityExpansion")
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

// SetSupportsVirtualDiskCapacityReduction sets the value of SupportsVirtualDiskCapacityReduction for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskCapacityReduction(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskCapacityReduction", (value))
}

// GetSupportsVirtualDiskCapacityReduction gets the value of SupportsVirtualDiskCapacityReduction for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskCapacityReduction() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskCapacityReduction")
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

// SetSupportsVirtualDiskCreation sets the value of SupportsVirtualDiskCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskCreation(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskCreation", (value))
}

// GetSupportsVirtualDiskCreation gets the value of SupportsVirtualDiskCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskCreation")
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

// SetSupportsVirtualDiskDeletion sets the value of SupportsVirtualDiskDeletion for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskDeletion(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskDeletion", (value))
}

// GetSupportsVirtualDiskDeletion gets the value of SupportsVirtualDiskDeletion for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskDeletion() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskDeletion")
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

// SetSupportsVirtualDiskModification sets the value of SupportsVirtualDiskModification for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskModification(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskModification", (value))
}

// GetSupportsVirtualDiskModification gets the value of SupportsVirtualDiskModification for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskModification() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskModification")
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

// SetSupportsVirtualDiskRepair sets the value of SupportsVirtualDiskRepair for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVirtualDiskRepair(value bool) (err error) {
	return instance.SetProperty("SupportsVirtualDiskRepair", (value))
}

// GetSupportsVirtualDiskRepair gets the value of SupportsVirtualDiskRepair for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVirtualDiskRepair() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVirtualDiskRepair")
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

// SetSupportsVolumeCreation sets the value of SupportsVolumeCreation for the instance
func (instance *MSFT_StorageSubSystem) SetPropertySupportsVolumeCreation(value bool) (err error) {
	return instance.SetProperty("SupportsVolumeCreation", (value))
}

// GetSupportsVolumeCreation gets the value of SupportsVolumeCreation for the instance
func (instance *MSFT_StorageSubSystem) GetPropertySupportsVolumeCreation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsVolumeCreation")
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

// SetTag sets the value of Tag for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyTag(value string) (err error) {
	return instance.SetProperty("Tag", (value))
}

// GetTag gets the value of Tag for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyTag() (value string, err error) {
	retValue, err := instance.GetProperty("Tag")
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

// SetVirtualDiskRepairEnabled sets the value of VirtualDiskRepairEnabled for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyVirtualDiskRepairEnabled(value bool) (err error) {
	return instance.SetProperty("VirtualDiskRepairEnabled", (value))
}

// GetVirtualDiskRepairEnabled gets the value of VirtualDiskRepairEnabled for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyVirtualDiskRepairEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VirtualDiskRepairEnabled")
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

// SetVirtualDiskRepairQueueDepth sets the value of VirtualDiskRepairQueueDepth for the instance
func (instance *MSFT_StorageSubSystem) SetPropertyVirtualDiskRepairQueueDepth(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskRepairQueueDepth", (value))
}

// GetVirtualDiskRepairQueueDepth gets the value of VirtualDiskRepairQueueDepth for the instance
func (instance *MSFT_StorageSubSystem) GetPropertyVirtualDiskRepairQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskRepairQueueDepth")
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

// This method creates a storage pool from available physical disks contained within a common primordial pool. A physical disk is available for storage pool creation if its CanPool property is set to TRUE. Storage pool creation is only available when the SupportsStoragePoolCreation field of the storage subsystem is TRUE.

// <param name="AutoWriteCacheSize" type="bool ">Indicates if provider should pick up the auto write cache size or not</param>
// <param name="EnclosureAwareDefault" type="bool ">This parameter indicates the default allocation policy for virtual disks created in an enclosure aware storage pool. For example, an enclosure aware subsystem could balance each data copy of the virtual disk across multiple physical enclosures such that each enclosure contains a full data copy of the virtual disk.</param>
// <param name="FriendlyName" type="string ">This parameter allows the user to specify the FriendlyName at the time of the storage pool creation. FriendlyNames are expected to be descriptive, however they are not required to be unique. Note that some storage subsystems do not allow setting a friendly name during pool creation. If a subsystem doesn't support this, storage pool creation should still succeed, however the pool may have a different name assigned to it.</param>
// <param name="LogicalSectorSizeDefault" type="uint64 ">This parameter indicates the default logical sector size for the storage pool. This is useful when a storage pool may contain a mix of 512 emulated and either 4K native or 512 native physical disks.</param>
// <param name="OtherUsageDescription" type="string ">Allows a user to set a custom usage type for the new storage pool object. This parameter can only be specified if the Usage parameter is set to 1 - 'Other'. </param>
// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []">This parameter is used to specify an array of physical disk objects that will be used as the backing data storage for the created storage pool. The physical disks must come from a primordial pool on the subsystem on which you are creating this pool. Only the disks from a single primordial pool may be used.</param>
// <param name="ProvisioningTypeDefault" type="StorageSubSystem_ProvisioningTypeDefault ">This parameter indicates the provisioning type to be used by default when creating a new virtual disk on this storage pool. If no default is specified, the default is inherited from the primordial pool.</param>
// <param name="ResiliencySettingNameDefault" type="string ">This parameter indicates the resiliency setting to be used by default when creating a new virtual disk on this storage pool. If the subsystem's SupportsMultipleResiliencySettingsPerStoragePool property is set to FALSE, this parameter also acts as a hint to the Storage Management Provider on which resiliency setting should be inherited by this storage pool. If no value is given, it is up to the Storage Management Provider to pick the most appropriate resiliency setting.</param>
// <param name="Usage" type="StorageSubSystem_Usage ">Denotes the intended usage of the storage pool.</param>
// <param name="WriteCacheSizeDefault" type="uint64 ">Default size of write cache for virtual disk creation</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation. When the operation has completed, an association should exist between the storage job and the created objects.</param>
// <param name="CreatedStoragePool" type="MSFT_StoragePool "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateStoragePool( /* IN */ FriendlyName string,
	/* IN */ Usage StorageSubSystem_Usage,
	/* IN */ OtherUsageDescription string,
	/* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* IN */ ResiliencySettingNameDefault string,
	/* IN */ ProvisioningTypeDefault StorageSubSystem_ProvisioningTypeDefault,
	/* IN */ LogicalSectorSizeDefault uint64,
	/* IN */ EnclosureAwareDefault bool,
	/* IN */ WriteCacheSizeDefault uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* OUT */ CreatedStoragePool MSFT_StoragePool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateStoragePool", FriendlyName, Usage, OtherUsageDescription, PhysicalDisks, ResiliencySettingNameDefault, ProvisioningTypeDefault, LogicalSectorSizeDefault, EnclosureAwareDefault, WriteCacheSizeDefault, AutoWriteCacheSize)
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
// <param name="FriendlyName" type="string "></param>
// <param name="LogicalSectorSizeDefault" type="uint64 "></param>
// <param name="MediaTypeDefault" type="uint16 "></param>
// <param name="MetadataLength" type="uint64 "></param>
// <param name="MinimumAllocationSize" type="uint64 "></param>
// <param name="OtherUsageDescription" type="string "></param>
// <param name="PhysicalDisks" type="MSFT_PhysicalDisk []"></param>
// <param name="ProvisioningTypeDefault" type="uint16 "></param>
// <param name="ResiliencySettingNameDefault" type="string "></param>
// <param name="Usage" type="uint16 "></param>
// <param name="Version" type="uint16 "></param>
// <param name="WriteCacheSizeDefault" type="uint64 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedStoragePool" type="MSFT_StoragePool "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateStoragePool2( /* IN */ FriendlyName string,
	/* IN */ Usage uint16,
	/* IN */ OtherUsageDescription string,
	/* IN */ PhysicalDisks []MSFT_PhysicalDisk,
	/* IN */ ResiliencySettingNameDefault string,
	/* IN */ ProvisioningTypeDefault uint16,
	/* IN */ MediaTypeDefault uint16,
	/* IN */ LogicalSectorSizeDefault uint64,
	/* IN */ MetadataLength uint64,
	/* IN */ MinimumAllocationSize uint64,
	/* IN */ FaultDomainAwarenessDefault uint16,
	/* IN */ WriteCacheSizeDefault uint64,
	/* IN */ AutoWriteCacheSize bool,
	/* IN */ Version uint16,
	/* OUT */ CreatedStoragePool MSFT_StoragePool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateStoragePool2", FriendlyName, Usage, OtherUsageDescription, PhysicalDisks, ResiliencySettingNameDefault, ProvisioningTypeDefault, MediaTypeDefault, LogicalSectorSizeDefault, MetadataLength, MinimumAllocationSize, FaultDomainAwarenessDefault, WriteCacheSizeDefault, AutoWriteCacheSize, Version)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows for the creation of virtual disks on a storage subsystem. This method is typically used when either a) the subsystem's storage pools do not allow virtual disk creation directly, or b) the subsystem does not support storage pools. Storage Management Providers may also choose to implement this method to 'intelligently' pick a storage pool for the user. If this method is supported, the subsystem's SupportsAutomaticStoragePoolSelection property should be set to TRUE.

// <param name="FriendlyName" type="string ">This parameter allows the user to specify the desired FriendlyName at the time of the virtual disk creation. FriendlyNames are expected to be descriptive, however they are not required to be unique. Note that some storage subsystems do not allow setting a friendly name during virtual disk creation. If a subsystem doesn't support this, virtual disk creation should still succeed, however the disk may have a different name assigned to it.</param>
// <param name="Interleave" type="uint64 ">Specifies the number of bytes used to form a strip in common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk. Thus Interleave * NumberOfColumns will yield the total size of one stripe.</param>
// <param name="IsEnclosureAware" type="bool ">Determines the allocation behavior for this virtual disk. Enclosure aware virtual disks will intelligently pick the physical disks to use for their redundancy. If TRUE, the virtual disk will attempt to use physical disks from different enclosures to balance the fault tolerance between two (or more) physical enclosures.</param>
// <param name="NumberOfColumns" type="uint16 ">Specifies the number of underlying physical disks across which data should be striped.</param>
// <param name="NumberOfDataCopies" type="uint16 ">Specifies the number of complete data copies to maintain for this virtual disk.</param>
// <param name="OtherUsageDescription" type="string ">Allows a user to set a custom usage type for the new virtual disk object. This parameter can only be specified if the Usage parameter is set to 1 - 'Other'. </param>
// <param name="ParityLayout" type="StorageSubSystem_ParityLayout ">This field specifies whether a parity-based resiliency setting is using a rotated or non-rotated parity layout. If the resiliency setting is not parity based, this field must be set to NULL</param>
// <param name="PhysicalDiskRedundancy" type="uint16 ">Specifies how many physical disk failures the virtual disk should be able to withstand before data loss occurs.</param>
// <param name="ProvisioningType" type="StorageSubSystem_ProvisioningType ">Denotes the provisioning type of the virtual disk. A value of 1 - 'Thin' means that the storage for the disk is allocated on-demand. A value of 2 - 'Fixed' means that the storage is allocated up front.</param>
// <param name="RequestNoSinglePointOfFailure" type="bool "></param>
// <param name="Size" type="uint64 ">Indicates the desired size for the virtual disk. Note that some storage subsystems will round the size up or down to a multiple of its allocation unit size. If this parameter is specified, UseMaximumSize must be NULL or FALSE.</param>
// <param name="Usage" type="StorageSubSystem_Usage ">Denotes the intended usage of the virtual disk</param>
// <param name="UseMaximumSize" type="bool ">Create a virtual disk using the largest supported size. This parameter cannot be used with the Size parameter.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation. When the operation has completed, an association should exist between the storage job and the created objects.</param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Size" type="uint64 ">Indicates the desired size for the virtual disk. Note that some storage subsystems will round the size up or down to a multiple of its allocation unit size. If this parameter is specified, UseMaximumSize must be NULL or FALSE.</param>
func (instance *MSFT_StorageSubSystem) CreateVirtualDisk( /* IN */ FriendlyName string,
	/* IN */ Usage StorageSubSystem_Usage,
	/* IN */ OtherUsageDescription string,
	/* IN/OUT */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ Interleave uint64,
	/* IN */ ParityLayout StorageSubSystem_ParityLayout,
	/* IN */ RequestNoSinglePointOfFailure bool,
	/* IN */ IsEnclosureAware bool,
	/* IN */ ProvisioningType StorageSubSystem_ProvisioningType,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVirtualDisk", FriendlyName, Usage, OtherUsageDescription, UseMaximumSize, NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, Interleave, ParityLayout, RequestNoSinglePointOfFailure, IsEnclosureAware, ProvisioningType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FaultDomainAwareness" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="Interleave" type="uint64 "></param>
// <param name="NumberOfColumns" type="uint16 "></param>
// <param name="NumberOfDataCopies" type="uint16 "></param>
// <param name="OtherUsageDescription" type="string "></param>
// <param name="ParityLayout" type="uint16 "></param>
// <param name="PhysicalDiskRedundancy" type="uint16 "></param>
// <param name="ProvisioningType" type="uint16 "></param>
// <param name="RequestNoSinglePointOfFailure" type="bool "></param>
// <param name="Size" type="uint64 "></param>
// <param name="Usage" type="uint16 "></param>
// <param name="UseMaximumSize" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVirtualDisk" type="MSFT_VirtualDisk "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateVirtualDisk2( /* IN */ FriendlyName string,
	/* IN */ Usage uint16,
	/* IN */ OtherUsageDescription string,
	/* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ Interleave uint64,
	/* IN */ ParityLayout uint16,
	/* IN */ RequestNoSinglePointOfFailure bool,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ProvisioningType uint16,
	/* OUT */ CreatedVirtualDisk MSFT_VirtualDisk,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVirtualDisk2", FriendlyName, Usage, OtherUsageDescription, Size, UseMaximumSize, NumberOfDataCopies, PhysicalDiskRedundancy, NumberOfColumns, Interleave, ParityLayout, RequestNoSinglePointOfFailure, FaultDomainAwareness, ProvisioningType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Creates logical grouping of virtual disks, target ports, and initiators for the purpose of showing virtual disks to host systems.

// <param name="DeviceAccesses" type="StorageSubSystem_DeviceAccesses []">This parameter specifies the level of access the initiator should have to each virtual disk specified by VirtualDiskNames. This parameter has a 1:1 mapping with the VirtualDiskNames parameter (the arrays must be the same length and have the same order).</param>
// <param name="DeviceNumbers" type="string []">Specifies the order in which the virtual disks should be exposed to the initiator. This capability is only available if the storage subsystem's MaskingClientSelectableDeviceNumbers property is set to TRUE. If specified, this parameter must have a 1:1 mapping with the VirtualDiskNames parameter.</param>
// <param name="FriendlyName" type="string ">This parameter allows the user to specify the desired FriendlyName for the masking set at the time of its creation. FriendlyNames are expected to be descriptive, however they are not requried to be unique.</param>
// <param name="HostType" type="StorageSubSystem_HostType ">Designates the host operating system or other host environment factors that may influence the behavior the storage subsystem should take when showing a virtual disk to an initiator.</param>
// <param name="InitiatorAddresses" type="string []">This parameter specifies the initiators for which the virtual disks should be shown. If the subsystem's MaskingOneInitiatorIdPerView property is TRUE, only one initiator can be specified for this masking set. The list of valid initiator address formats can be determined through the subsystem's MaskingValidInitiatorIdTypes property.</param>
// <param name="TargetPortAddresses" type="string []">This parameter specifies the target ports which should be used when showing the virtual disks to the initiators. The number of target ports that can be specified depends on the subsystem's MaskingPortsPerView property. If MaskingPortsPerView is set to 4 - 'All target ports share the same view', this parameter is essentially ignored as all target ports on the system will be associated with this masking set.</param>
// <param name="VirtualDiskNames" type="string []">This parameter specifies the list of virtual disks to show to the initiators in the masking set. The identifier used by this parameter is the virtual disk Name property. This parameter has a 1:1 mapping with the DeviceAccesses parameter (the arrays must be the same length and have the same order).</param>

// <param name="CreatedMaskingSet" type="MSFT_MaskingSet "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation. When the operation has completed, an association should exist between the storage job and the created objects.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateMaskingSet( /* IN */ FriendlyName string,
	/* IN */ VirtualDiskNames []string,
	/* IN */ DeviceAccesses []StorageSubSystem_DeviceAccesses,
	/* IN */ DeviceNumbers []string,
	/* IN */ TargetPortAddresses []string,
	/* IN */ InitiatorAddresses []string,
	/* IN */ HostType StorageSubSystem_HostType,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ CreatedMaskingSet MSFT_MaskingSet,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateMaskingSet", FriendlyName, VirtualDiskNames, DeviceAccesses, DeviceNumbers, TargetPortAddresses, InitiatorAddresses, HostType)
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
func (instance *MSFT_StorageSubSystem) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
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
func (instance *MSFT_StorageSubSystem) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user to set the description field of the storage subsystem.

// <param name="Description" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) SetDescription( /* IN */ Description string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDescription", Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user to set the SupportsAutomaticObjectClustering field of the storage subsystem.

// <param name="AutomaticClusteringEnabled" type="bool "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) SetAttributes( /* IN */ AutomaticClusteringEnabled bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", AutomaticClusteringEnabled)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AutomaticClusteringEnabled" type="bool "></param>
// <param name="FaultDomainAwarenessDefault" type="uint16 "></param>
// <param name="VirtualDiskRepairEnabled" type="bool "></param>
// <param name="VirtualDiskRepairQueueDepth" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) SetAttributes2( /* IN */ AutomaticClusteringEnabled bool,
	/* IN */ VirtualDiskRepairEnabled bool,
	/* IN */ VirtualDiskRepairQueueDepth uint32,
	/* IN */ FaultDomainAwarenessDefault uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes2", AutomaticClusteringEnabled, VirtualDiskRepairEnabled, VirtualDiskRepairQueueDepth, FaultDomainAwarenessDefault)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FriendlyName" type="string "></param>
// <param name="RecoveryPointObjective" type="uint32 "></param>
// <param name="SourceGroup" type="MSFT_ReplicationGroup "></param>
// <param name="SourceGroupSettings" type="MSFT_ReplicationSettings "></param>
// <param name="SourceReplicationGroupDescription" type="string "></param>
// <param name="SourceReplicationGroupFriendlyName" type="string "></param>
// <param name="SourceStorageElements" type="MSFT_StorageObject []"></param>
// <param name="SyncType" type="uint16 "></param>
// <param name="TargetGroup" type="MSFT_ReplicationGroup "></param>
// <param name="TargetGroupSettings" type="MSFT_ReplicationSettings "></param>
// <param name="TargetReplicationGroupDescription" type="string "></param>
// <param name="TargetReplicationGroupFriendlyName" type="string "></param>
// <param name="TargetStorageElements" type="MSFT_StorageObject []"></param>
// <param name="TargetStoragePool" type="MSFT_StoragePool "></param>
// <param name="TargetStoragePools" type="MSFT_StoragePool []"></param>
// <param name="TargetStorageSubsystem" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedReplicaPeer" type="MSFT_ReplicaPeer "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SourceGroup" type="MSFT_ReplicationGroup "></param>
func (instance *MSFT_StorageSubSystem) CreateReplicationRelationship( /* IN */ FriendlyName string,
	/* IN */ SyncType uint16,
	/* IN */ TargetStorageSubsystem MSFT_ReplicaPeer,
	/* IN */ SourceReplicationGroupFriendlyName string,
	/* IN */ SourceReplicationGroupDescription string,
	/* IN */ SourceStorageElements []MSFT_StorageObject,
	/* IN */ SourceGroupSettings MSFT_ReplicationSettings,
	/* IN */ TargetReplicationGroupFriendlyName string,
	/* IN */ TargetReplicationGroupDescription string,
	/* IN */ TargetStorageElements []MSFT_StorageObject,
	/* IN */ TargetStoragePool MSFT_StoragePool,
	/* IN */ TargetStoragePools []MSFT_StoragePool,
	/* IN */ TargetGroupSettings MSFT_ReplicationSettings,
	/* IN */ RecoveryPointObjective uint32,
	/* IN/OUT */ SourceGroup MSFT_ReplicationGroup,
	/* IN */ TargetGroup MSFT_ReplicationGroup,
	/* OUT */ CreatedReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplicationRelationship", FriendlyName, SyncType, TargetStorageSubsystem, SourceReplicationGroupFriendlyName, SourceReplicationGroupDescription, SourceStorageElements, SourceGroupSettings, TargetReplicationGroupFriendlyName, TargetReplicationGroupDescription, TargetStorageElements, TargetStoragePool, TargetStoragePools, TargetGroupSettings, RecoveryPointObjective, TargetGroup)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SourceReplicationGroup" type="MSFT_ReplicationGroup "></param>
// <param name="TargetGroupReplicaPeer" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) DeleteReplicationRelationship( /* IN */ SourceReplicationGroup MSFT_ReplicationGroup,
	/* IN */ TargetGroupReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteReplicationRelationship", SourceReplicationGroup, TargetGroupReplicaPeer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Description" type="string "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="ReplicationSettings" type="MSFT_ReplicationSettings "></param>
// <param name="StorageElements" type="MSFT_StorageObject []"></param>

// <param name="CreatedReplicationGroup" type="MSFT_ReplicationGroup "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateReplicationGroup( /* IN */ FriendlyName string,
	/* IN */ Description string,
	/* IN */ StorageElements []MSFT_StorageObject,
	/* IN */ ReplicationSettings MSFT_ReplicationSettings,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ CreatedReplicationGroup MSFT_ReplicationGroup,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplicationGroup", FriendlyName, Description, StorageElements, ReplicationSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FileSharingProtocols" type="uint16 []"></param>
// <param name="FriendlyName" type="string "></param>
// <param name="HostNames" type="string []"></param>

// <param name="CreatedFileServer" type="MSFT_FileServer "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) CreateFileServer( /* IN */ FriendlyName string,
	/* IN */ FileSharingProtocols []uint16,
	/* IN */ HostNames []string,
	/* OUT */ CreatedFileServer MSFT_FileServer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateFileServer", FriendlyName, FileSharingProtocols, HostNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ActivityId" type="string "></param>
// <param name="CopyExistingInfoOnly" type="bool "></param>
// <param name="DestinationPath" type="string "></param>
// <param name="ExcludeDiagnosticLog" type="bool "></param>
// <param name="ExcludeOperationalLog" type="bool "></param>
// <param name="IncludeLiveDump" type="bool "></param>
// <param name="TimeSpan" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) GetDiagnosticInfo( /* IN */ DestinationPath string,
	/* IN */ TimeSpan uint32,
	/* IN */ ActivityId string,
	/* IN */ ExcludeOperationalLog bool,
	/* IN */ ExcludeDiagnosticLog bool,
	/* IN */ IncludeLiveDump bool,
	/* IN */ CopyExistingInfoOnly bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDiagnosticInfo", DestinationPath, TimeSpan, ActivityId, ExcludeOperationalLog, ExcludeDiagnosticLog, IncludeLiveDump, CopyExistingInfoOnly)
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
func (instance *MSFT_StorageSubSystem) ClearDiagnosticInfo( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ClearDiagnosticInfo")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Level" type="uint16 "></param>
// <param name="MaxLogSize" type="uint64 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) StartDiagnosticLog( /* IN */ Level uint16,
	/* IN */ MaxLogSize uint64,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("StartDiagnosticLog", Level, MaxLogSize)
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
func (instance *MSFT_StorageSubSystem) StopDiagnosticLog( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("StopDiagnosticLog")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DiagnoseResults" type="MSFT_StorageDiagnoseResult []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) Diagnose( /* OUT */ DiagnoseResults []MSFT_StorageDiagnoseResult,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Diagnose")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ActionResults" type="MSFT_HealthAction []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSubSystem) GetActions( /* OUT */ ActionResults []MSFT_HealthAction,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetActions")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
