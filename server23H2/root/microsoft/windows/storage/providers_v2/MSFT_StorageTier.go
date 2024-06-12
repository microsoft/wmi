// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageTier struct
type MSFT_StorageTier struct {
	*MSFT_StorageObject

	//
	AllocatedSize uint64

	//
	AllocationUnitSize uint64

	//
	ColumnIsolation uint16

	// A user settable description of the storage tier
	Description string

	//
	FaultDomainAwareness uint16

	//
	FootprintOnPool uint64

	// Friendly name of the storage tier, defined by the user
	FriendlyName string

	//
	Interleave uint64

	// Media type of this storage tier
	MediaType StorageTier_MediaType

	//
	NumberOfColumns uint16

	//
	NumberOfDataCopies uint16

	//
	NumberOfGroups uint16

	//
	ParityLayout uint16

	//
	PhysicalDiskRedundancy uint16

	//
	ProvisioningType uint16

	// Specifies the name of the resiliency setting that should be used for storage tier creation.
	ResiliencySettingName string

	// Size of the tier on the virtual disk. This property is available only when the storage tier is part of a virtual disk. The property is unspecified for pool-level storage tiers.
	Size uint64

	//
	TierClass uint16

	//
	Usage uint16
}

func NewMSFT_StorageTierEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageTier, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageTier{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_StorageTierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageTier, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageTier{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetAllocatedSize sets the value of AllocatedSize for the instance
func (instance *MSFT_StorageTier) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", (value))
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_StorageTier) GetPropertyAllocatedSize() (value uint64, err error) {
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

// SetAllocationUnitSize sets the value of AllocationUnitSize for the instance
func (instance *MSFT_StorageTier) SetPropertyAllocationUnitSize(value uint64) (err error) {
	return instance.SetProperty("AllocationUnitSize", (value))
}

// GetAllocationUnitSize gets the value of AllocationUnitSize for the instance
func (instance *MSFT_StorageTier) GetPropertyAllocationUnitSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationUnitSize")
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

// SetColumnIsolation sets the value of ColumnIsolation for the instance
func (instance *MSFT_StorageTier) SetPropertyColumnIsolation(value uint16) (err error) {
	return instance.SetProperty("ColumnIsolation", (value))
}

// GetColumnIsolation gets the value of ColumnIsolation for the instance
func (instance *MSFT_StorageTier) GetPropertyColumnIsolation() (value uint16, err error) {
	retValue, err := instance.GetProperty("ColumnIsolation")
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageTier) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageTier) GetPropertyDescription() (value string, err error) {
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

// SetFaultDomainAwareness sets the value of FaultDomainAwareness for the instance
func (instance *MSFT_StorageTier) SetPropertyFaultDomainAwareness(value uint16) (err error) {
	return instance.SetProperty("FaultDomainAwareness", (value))
}

// GetFaultDomainAwareness gets the value of FaultDomainAwareness for the instance
func (instance *MSFT_StorageTier) GetPropertyFaultDomainAwareness() (value uint16, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwareness")
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

// SetFootprintOnPool sets the value of FootprintOnPool for the instance
func (instance *MSFT_StorageTier) SetPropertyFootprintOnPool(value uint64) (err error) {
	return instance.SetProperty("FootprintOnPool", (value))
}

// GetFootprintOnPool gets the value of FootprintOnPool for the instance
func (instance *MSFT_StorageTier) GetPropertyFootprintOnPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("FootprintOnPool")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_StorageTier) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_StorageTier) GetPropertyFriendlyName() (value string, err error) {
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

// SetInterleave sets the value of Interleave for the instance
func (instance *MSFT_StorageTier) SetPropertyInterleave(value uint64) (err error) {
	return instance.SetProperty("Interleave", (value))
}

// GetInterleave gets the value of Interleave for the instance
func (instance *MSFT_StorageTier) GetPropertyInterleave() (value uint64, err error) {
	retValue, err := instance.GetProperty("Interleave")
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
func (instance *MSFT_StorageTier) SetPropertyMediaType(value StorageTier_MediaType) (err error) {
	return instance.SetProperty("MediaType", (value))
}

// GetMediaType gets the value of MediaType for the instance
func (instance *MSFT_StorageTier) GetPropertyMediaType() (value StorageTier_MediaType, err error) {
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

	value = StorageTier_MediaType(valuetmp)

	return
}

// SetNumberOfColumns sets the value of NumberOfColumns for the instance
func (instance *MSFT_StorageTier) SetPropertyNumberOfColumns(value uint16) (err error) {
	return instance.SetProperty("NumberOfColumns", (value))
}

// GetNumberOfColumns gets the value of NumberOfColumns for the instance
func (instance *MSFT_StorageTier) GetPropertyNumberOfColumns() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColumns")
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

// SetNumberOfDataCopies sets the value of NumberOfDataCopies for the instance
func (instance *MSFT_StorageTier) SetPropertyNumberOfDataCopies(value uint16) (err error) {
	return instance.SetProperty("NumberOfDataCopies", (value))
}

// GetNumberOfDataCopies gets the value of NumberOfDataCopies for the instance
func (instance *MSFT_StorageTier) GetPropertyNumberOfDataCopies() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfDataCopies")
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

// SetNumberOfGroups sets the value of NumberOfGroups for the instance
func (instance *MSFT_StorageTier) SetPropertyNumberOfGroups(value uint16) (err error) {
	return instance.SetProperty("NumberOfGroups", (value))
}

// GetNumberOfGroups gets the value of NumberOfGroups for the instance
func (instance *MSFT_StorageTier) GetPropertyNumberOfGroups() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfGroups")
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

// SetParityLayout sets the value of ParityLayout for the instance
func (instance *MSFT_StorageTier) SetPropertyParityLayout(value uint16) (err error) {
	return instance.SetProperty("ParityLayout", (value))
}

// GetParityLayout gets the value of ParityLayout for the instance
func (instance *MSFT_StorageTier) GetPropertyParityLayout() (value uint16, err error) {
	retValue, err := instance.GetProperty("ParityLayout")
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

// SetPhysicalDiskRedundancy sets the value of PhysicalDiskRedundancy for the instance
func (instance *MSFT_StorageTier) SetPropertyPhysicalDiskRedundancy(value uint16) (err error) {
	return instance.SetProperty("PhysicalDiskRedundancy", (value))
}

// GetPhysicalDiskRedundancy gets the value of PhysicalDiskRedundancy for the instance
func (instance *MSFT_StorageTier) GetPropertyPhysicalDiskRedundancy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskRedundancy")
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

// SetProvisioningType sets the value of ProvisioningType for the instance
func (instance *MSFT_StorageTier) SetPropertyProvisioningType(value uint16) (err error) {
	return instance.SetProperty("ProvisioningType", (value))
}

// GetProvisioningType gets the value of ProvisioningType for the instance
func (instance *MSFT_StorageTier) GetPropertyProvisioningType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProvisioningType")
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

// SetResiliencySettingName sets the value of ResiliencySettingName for the instance
func (instance *MSFT_StorageTier) SetPropertyResiliencySettingName(value string) (err error) {
	return instance.SetProperty("ResiliencySettingName", (value))
}

// GetResiliencySettingName gets the value of ResiliencySettingName for the instance
func (instance *MSFT_StorageTier) GetPropertyResiliencySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("ResiliencySettingName")
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

// SetSize sets the value of Size for the instance
func (instance *MSFT_StorageTier) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_StorageTier) GetPropertySize() (value uint64, err error) {
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

// SetTierClass sets the value of TierClass for the instance
func (instance *MSFT_StorageTier) SetPropertyTierClass(value uint16) (err error) {
	return instance.SetProperty("TierClass", (value))
}

// GetTierClass gets the value of TierClass for the instance
func (instance *MSFT_StorageTier) GetPropertyTierClass() (value uint16, err error) {
	retValue, err := instance.GetProperty("TierClass")
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

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_StorageTier) SetPropertyUsage(value uint16) (err error) {
	return instance.SetProperty("Usage", (value))
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_StorageTier) GetPropertyUsage() (value uint16, err error) {
	retValue, err := instance.GetProperty("Usage")
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

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="PhysicalExtents" type="MSFT_PhysicalExtent []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) GetPhysicalExtent( /* OUT */ PhysicalExtents []MSFT_PhysicalExtent,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalExtent")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method deletes the storage tier.This method is available only when the SupportsStorageTierDeletion propertyon the storage subsystem is set to TRUE. If it is set to FALSE, this methodwill fail with MI_RESULT_NOT_SUPPORTED.

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Resizes the storage tier on the virtual disk. This method is not available for pool-level storage tiers.

// <param name="Size" type="uint64 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Size" type="uint64 "></param>
func (instance *MSFT_StorageTier) Resize( /* IN/OUT */ Size uint64,
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

// This method allows the storage tier to be renamed.

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the user to update or set various attributes on the storage tier. Note that not all parameters must be specified, and only those given will be updated.

// <param name="MediaType" type="StorageTier_MediaType ">Media type of this storage tier</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) SetAttributes( /* IN */ MediaType StorageTier_MediaType,
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

// <param name="ExtendedStatus" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) SetAttributes2( /* IN */ ProvisioningType uint16,
	/* IN */ MediaType uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ NumberOfGroups uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ Interleave uint64,
	/* OUT */ ExtendedStatus string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes2", ProvisioningType, MediaType, FaultDomainAwareness, ColumnIsolation, ResiliencySettingName, PhysicalDiskRedundancy, NumberOfDataCopies, NumberOfGroups, NumberOfColumns, Interleave)
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
// <param name="Usage" type="uint16 "></param>

// <param name="ExtendedStatus" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) SetAttributes3( /* IN */ ProvisioningType uint16,
	/* IN */ AllocationUnitSize uint64,
	/* IN */ MediaType uint16,
	/* IN */ FaultDomainAwareness uint16,
	/* IN */ ColumnIsolation uint16,
	/* IN */ ResiliencySettingName string,
	/* IN */ Usage uint16,
	/* IN */ PhysicalDiskRedundancy uint16,
	/* IN */ NumberOfDataCopies uint16,
	/* IN */ NumberOfGroups uint16,
	/* IN */ NumberOfColumns uint16,
	/* IN */ Interleave uint64,
	/* OUT */ ExtendedStatus string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes3", ProvisioningType, AllocationUnitSize, MediaType, FaultDomainAwareness, ColumnIsolation, ResiliencySettingName, Usage, PhysicalDiskRedundancy, NumberOfDataCopies, NumberOfGroups, NumberOfColumns, Interleave)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the storage tier's description to be changed.

// <param name="Description" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) SetDescription( /* IN */ Description string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDescription", Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method returns the supported sizes for a new storage tier. These sizes can either be returned in an array of all supported sizes, through a min, max, and divisor, or both.

// <param name="ResiliencySettingName" type="string ">Specifies the name of the resiliency setting that should be used when determining the supported sizes. Note that the sizes returned may be different depending on the resiliency setting.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedSizes" type="uint64 []">This output parameter will contain an array of all of the supported sizes for the storage tier. This parameter may be NULL if the number of supported sizes is large, but is useful for storage tiers that only support a select number of tier sizes.</param>
// <param name="TierSizeDivisor" type="uint64 ">This parameter indicates the interval in which the supported sizes increment. For example: If the minimum supported size is 10 GB, and this parameter is 2 GB, then the supported sizes for this pool would be 10 GB, 12 GB, 14 GB, etc. until the maximum supported size is reached.</param>
// <param name="TierSizeMax" type="uint64 ">This parameter denotes the maximum supported size that a tier created in this pool can be.</param>
// <param name="TierSizeMin" type="uint64 ">This parameter denotes the minimum supported size that a tier created in this pool can be.</param>
func (instance *MSFT_StorageTier) GetSupportedSize( /* IN */ ResiliencySettingName string,
	/* OUT */ SupportedSizes []uint64,
	/* OUT */ TierSizeMin uint64,
	/* OUT */ TierSizeMax uint64,
	/* OUT */ TierSizeDivisor uint64,
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

// <param name="StorageFaultDomains" type="MSFT_StorageFaultDomain []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageTier) AddStorageFaultDomain( /* IN */ StorageFaultDomains []MSFT_StorageFaultDomain,
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
func (instance *MSFT_StorageTier) RemoveStorageFaultDomain( /* IN */ StorageFaultDomains []MSFT_StorageFaultDomain,
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
