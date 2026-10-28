// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters struct
type Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters struct {
	*Win32_PerfRawData

	//
	AllocationofMetadataClustersonFastTierPersec uint64

	//
	AllocatorCacheCollisions uint32

	//
	AllocatorCacheDeletions uint32

	//
	AllocatorCacheInsertions uint32

	//
	AllocatorCachePurges uint32

	//
	AllocatorCacheSize uint32

	//
	AllocatorRegionsExaminedAverage uint64

	//
	AllocatorRegionsExaminedAverage_Base uint32

	//
	AllocatorRegionsExaminedMax uint32

	//
	Checkpointastreeupdatelatency100ns uint64

	//
	Checkpointastreeupdatelatency100ns_Base uint32

	//
	Checkpointexclusivelockslatency100ns uint64

	//
	Checkpointexclusivelockslatency100ns_Base uint32

	//
	CheckpointsPersec uint64

	//
	DeleteQueueentries uint32

	//
	Dirtymetadatapages uint64

	//
	Dirtytablelistentries uint32

	//
	TotalAllocationofClustersPersec uint64

	//
	Totaltimeneededtoprocesscheckpointsin100ns uint64

	//
	Totaltimeneededtoprocesscheckpointsin100ns_Base uint32

	//
	Treeupdateexclusivelockslatency100ns uint64

	//
	Treeupdateexclusivelockslatency100ns_Base uint32

	//
	Treeupdatelatency100ns uint64

	//
	Treeupdatelatency100ns_Base uint32

	//
	TreeupdatesPersec uint64
}

func NewWin32_PerfRawData_Counters_ReFSDedupMinstorePerfCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_ReFSDedupMinstorePerfCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAllocationofMetadataClustersonFastTierPersec sets the value of AllocationofMetadataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocationofMetadataClustersonFastTierPersec(value uint64) (err error) {
	return instance.SetProperty("AllocationofMetadataClustersonFastTierPersec", (value))
}

// GetAllocationofMetadataClustersonFastTierPersec gets the value of AllocationofMetadataClustersonFastTierPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocationofMetadataClustersonFastTierPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationofMetadataClustersonFastTierPersec")
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

// SetAllocatorCacheCollisions sets the value of AllocatorCacheCollisions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorCacheCollisions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheCollisions", (value))
}

// GetAllocatorCacheCollisions gets the value of AllocatorCacheCollisions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorCacheCollisions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheCollisions")
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

// SetAllocatorCacheDeletions sets the value of AllocatorCacheDeletions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorCacheDeletions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheDeletions", (value))
}

// GetAllocatorCacheDeletions gets the value of AllocatorCacheDeletions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorCacheDeletions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheDeletions")
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

// SetAllocatorCacheInsertions sets the value of AllocatorCacheInsertions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorCacheInsertions(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheInsertions", (value))
}

// GetAllocatorCacheInsertions gets the value of AllocatorCacheInsertions for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorCacheInsertions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheInsertions")
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

// SetAllocatorCachePurges sets the value of AllocatorCachePurges for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorCachePurges(value uint32) (err error) {
	return instance.SetProperty("AllocatorCachePurges", (value))
}

// GetAllocatorCachePurges gets the value of AllocatorCachePurges for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorCachePurges() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCachePurges")
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

// SetAllocatorCacheSize sets the value of AllocatorCacheSize for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorCacheSize(value uint32) (err error) {
	return instance.SetProperty("AllocatorCacheSize", (value))
}

// GetAllocatorCacheSize gets the value of AllocatorCacheSize for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorCacheSize")
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

// SetAllocatorRegionsExaminedAverage sets the value of AllocatorRegionsExaminedAverage for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorRegionsExaminedAverage(value uint64) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedAverage", (value))
}

// GetAllocatorRegionsExaminedAverage gets the value of AllocatorRegionsExaminedAverage for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorRegionsExaminedAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedAverage")
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

// SetAllocatorRegionsExaminedAverage_Base sets the value of AllocatorRegionsExaminedAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorRegionsExaminedAverage_Base(value uint32) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedAverage_Base", (value))
}

// GetAllocatorRegionsExaminedAverage_Base gets the value of AllocatorRegionsExaminedAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorRegionsExaminedAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedAverage_Base")
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

// SetAllocatorRegionsExaminedMax sets the value of AllocatorRegionsExaminedMax for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyAllocatorRegionsExaminedMax(value uint32) (err error) {
	return instance.SetProperty("AllocatorRegionsExaminedMax", (value))
}

// GetAllocatorRegionsExaminedMax gets the value of AllocatorRegionsExaminedMax for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyAllocatorRegionsExaminedMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocatorRegionsExaminedMax")
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

// SetCheckpointastreeupdatelatency100ns sets the value of Checkpointastreeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyCheckpointastreeupdatelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Checkpointastreeupdatelatency100ns", (value))
}

// GetCheckpointastreeupdatelatency100ns gets the value of Checkpointastreeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyCheckpointastreeupdatelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Checkpointastreeupdatelatency100ns")
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

// SetCheckpointastreeupdatelatency100ns_Base sets the value of Checkpointastreeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyCheckpointastreeupdatelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Checkpointastreeupdatelatency100ns_Base", (value))
}

// GetCheckpointastreeupdatelatency100ns_Base gets the value of Checkpointastreeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyCheckpointastreeupdatelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Checkpointastreeupdatelatency100ns_Base")
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

// SetCheckpointexclusivelockslatency100ns sets the value of Checkpointexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyCheckpointexclusivelockslatency100ns(value uint64) (err error) {
	return instance.SetProperty("Checkpointexclusivelockslatency100ns", (value))
}

// GetCheckpointexclusivelockslatency100ns gets the value of Checkpointexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyCheckpointexclusivelockslatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Checkpointexclusivelockslatency100ns")
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

// SetCheckpointexclusivelockslatency100ns_Base sets the value of Checkpointexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyCheckpointexclusivelockslatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Checkpointexclusivelockslatency100ns_Base", (value))
}

// GetCheckpointexclusivelockslatency100ns_Base gets the value of Checkpointexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyCheckpointexclusivelockslatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Checkpointexclusivelockslatency100ns_Base")
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

// SetCheckpointsPersec sets the value of CheckpointsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyCheckpointsPersec(value uint64) (err error) {
	return instance.SetProperty("CheckpointsPersec", (value))
}

// GetCheckpointsPersec gets the value of CheckpointsPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyCheckpointsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CheckpointsPersec")
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

// SetDeleteQueueentries sets the value of DeleteQueueentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyDeleteQueueentries(value uint32) (err error) {
	return instance.SetProperty("DeleteQueueentries", (value))
}

// GetDeleteQueueentries gets the value of DeleteQueueentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyDeleteQueueentries() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeleteQueueentries")
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

// SetDirtymetadatapages sets the value of Dirtymetadatapages for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyDirtymetadatapages(value uint64) (err error) {
	return instance.SetProperty("Dirtymetadatapages", (value))
}

// GetDirtymetadatapages gets the value of Dirtymetadatapages for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyDirtymetadatapages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Dirtymetadatapages")
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

// SetDirtytablelistentries sets the value of Dirtytablelistentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyDirtytablelistentries(value uint32) (err error) {
	return instance.SetProperty("Dirtytablelistentries", (value))
}

// GetDirtytablelistentries gets the value of Dirtytablelistentries for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyDirtytablelistentries() (value uint32, err error) {
	retValue, err := instance.GetProperty("Dirtytablelistentries")
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

// SetTotalAllocationofClustersPersec sets the value of TotalAllocationofClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTotalAllocationofClustersPersec(value uint64) (err error) {
	return instance.SetProperty("TotalAllocationofClustersPersec", (value))
}

// GetTotalAllocationofClustersPersec gets the value of TotalAllocationofClustersPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTotalAllocationofClustersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalAllocationofClustersPersec")
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

// SetTotaltimeneededtoprocesscheckpointsin100ns sets the value of Totaltimeneededtoprocesscheckpointsin100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTotaltimeneededtoprocesscheckpointsin100ns(value uint64) (err error) {
	return instance.SetProperty("Totaltimeneededtoprocesscheckpointsin100ns", (value))
}

// GetTotaltimeneededtoprocesscheckpointsin100ns gets the value of Totaltimeneededtoprocesscheckpointsin100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTotaltimeneededtoprocesscheckpointsin100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Totaltimeneededtoprocesscheckpointsin100ns")
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

// SetTotaltimeneededtoprocesscheckpointsin100ns_Base sets the value of Totaltimeneededtoprocesscheckpointsin100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTotaltimeneededtoprocesscheckpointsin100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Totaltimeneededtoprocesscheckpointsin100ns_Base", (value))
}

// GetTotaltimeneededtoprocesscheckpointsin100ns_Base gets the value of Totaltimeneededtoprocesscheckpointsin100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTotaltimeneededtoprocesscheckpointsin100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Totaltimeneededtoprocesscheckpointsin100ns_Base")
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

// SetTreeupdateexclusivelockslatency100ns sets the value of Treeupdateexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTreeupdateexclusivelockslatency100ns(value uint64) (err error) {
	return instance.SetProperty("Treeupdateexclusivelockslatency100ns", (value))
}

// GetTreeupdateexclusivelockslatency100ns gets the value of Treeupdateexclusivelockslatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTreeupdateexclusivelockslatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Treeupdateexclusivelockslatency100ns")
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

// SetTreeupdateexclusivelockslatency100ns_Base sets the value of Treeupdateexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTreeupdateexclusivelockslatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Treeupdateexclusivelockslatency100ns_Base", (value))
}

// GetTreeupdateexclusivelockslatency100ns_Base gets the value of Treeupdateexclusivelockslatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTreeupdateexclusivelockslatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Treeupdateexclusivelockslatency100ns_Base")
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

// SetTreeupdatelatency100ns sets the value of Treeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTreeupdatelatency100ns(value uint64) (err error) {
	return instance.SetProperty("Treeupdatelatency100ns", (value))
}

// GetTreeupdatelatency100ns gets the value of Treeupdatelatency100ns for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTreeupdatelatency100ns() (value uint64, err error) {
	retValue, err := instance.GetProperty("Treeupdatelatency100ns")
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

// SetTreeupdatelatency100ns_Base sets the value of Treeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTreeupdatelatency100ns_Base(value uint32) (err error) {
	return instance.SetProperty("Treeupdatelatency100ns_Base", (value))
}

// GetTreeupdatelatency100ns_Base gets the value of Treeupdatelatency100ns_Base for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTreeupdatelatency100ns_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Treeupdatelatency100ns_Base")
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

// SetTreeupdatesPersec sets the value of TreeupdatesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) SetPropertyTreeupdatesPersec(value uint64) (err error) {
	return instance.SetProperty("TreeupdatesPersec", (value))
}

// GetTreeupdatesPersec gets the value of TreeupdatesPersec for the instance
func (instance *Win32_PerfRawData_Counters_ReFSDedupMinstorePerfCounters) GetPropertyTreeupdatesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TreeupdatesPersec")
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
