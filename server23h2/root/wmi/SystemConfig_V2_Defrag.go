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

// SystemConfig_V2_Defrag struct
type SystemConfig_V2_Defrag struct {
	*SystemConfig_V2

	//
	AlignmentClusters uint64

	//
	AvgFragmentsPerFile uint32

	//
	AvgFreeSpaceSize uint64

	//
	BytesPerCluster uint32

	//
	ClustersPerSlab uint64

	//
	DirectoryCount uint32

	//
	FragmentedDirectories uint32

	//
	FragmentedDirectoryExtents uint64

	//
	FragmentedExtents uint64

	//
	FragmentedFiles uint32

	//
	FragmentedSpace uint32

	//
	FreeSpaceCount uint64

	//
	HardwareIssue uint32

	//
	InUseMFTRecords uint32

	//
	InUseSlabs uint32

	//
	LargestFreeSpaceSize uint64

	//
	LastRunActualPurgeClusters uint64

	//
	LastRunActualPurgeSlabs uint32

	//
	LastRunClustersTrimmed uint64

	//
	LastRunFullDefragTime uint64

	//
	LastRunInitialBackedSlabs uint32

	//
	LastRunPercentFragmentation uint32

	//
	LastRunPinnedSlabs uint32

	//
	LastRunPotentialPurgeSlabs uint32

	//
	LastRunSpaceInefficientSlabs uint32

	//
	LastRunTime uint64

	//
	LastRunTrimmedSlabs uint32

	//
	LastRunUnknownEvictFailSlabs uint32

	//
	LastRunVolsnapPinnedSlabs uint32

	//
	MFTFragmentCount uint32

	//
	MFTSize uint64

	//
	MovableFiles uint32

	//
	TotalClusters uint64

	//
	TotalMFTRecords uint32

	//
	TotalSlabs uint32

	//
	TotalUsedClusters uint64

	//
	UnmovableFiles uint32

	//
	VolumeId interface{}

	//
	VolumePathNames string
}

func NewSystemConfig_V2_DefragEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_Defrag, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Defrag{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_DefragEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_Defrag, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Defrag{
		SystemConfig_V2: tmp,
	}
	return
}

// SetAlignmentClusters sets the value of AlignmentClusters for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyAlignmentClusters(value uint64) (err error) {
	return instance.SetProperty("AlignmentClusters", (value))
}

// GetAlignmentClusters gets the value of AlignmentClusters for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyAlignmentClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("AlignmentClusters")
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

// SetAvgFragmentsPerFile sets the value of AvgFragmentsPerFile for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyAvgFragmentsPerFile(value uint32) (err error) {
	return instance.SetProperty("AvgFragmentsPerFile", (value))
}

// GetAvgFragmentsPerFile gets the value of AvgFragmentsPerFile for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyAvgFragmentsPerFile() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgFragmentsPerFile")
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

// SetAvgFreeSpaceSize sets the value of AvgFreeSpaceSize for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyAvgFreeSpaceSize(value uint64) (err error) {
	return instance.SetProperty("AvgFreeSpaceSize", (value))
}

// GetAvgFreeSpaceSize gets the value of AvgFreeSpaceSize for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyAvgFreeSpaceSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgFreeSpaceSize")
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

// SetBytesPerCluster sets the value of BytesPerCluster for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyBytesPerCluster(value uint32) (err error) {
	return instance.SetProperty("BytesPerCluster", (value))
}

// GetBytesPerCluster gets the value of BytesPerCluster for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyBytesPerCluster() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesPerCluster")
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

// SetClustersPerSlab sets the value of ClustersPerSlab for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyClustersPerSlab(value uint64) (err error) {
	return instance.SetProperty("ClustersPerSlab", (value))
}

// GetClustersPerSlab gets the value of ClustersPerSlab for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyClustersPerSlab() (value uint64, err error) {
	retValue, err := instance.GetProperty("ClustersPerSlab")
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

// SetDirectoryCount sets the value of DirectoryCount for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyDirectoryCount(value uint32) (err error) {
	return instance.SetProperty("DirectoryCount", (value))
}

// GetDirectoryCount gets the value of DirectoryCount for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyDirectoryCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCount")
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

// SetFragmentedDirectories sets the value of FragmentedDirectories for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFragmentedDirectories(value uint32) (err error) {
	return instance.SetProperty("FragmentedDirectories", (value))
}

// GetFragmentedDirectories gets the value of FragmentedDirectories for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFragmentedDirectories() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentedDirectories")
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

// SetFragmentedDirectoryExtents sets the value of FragmentedDirectoryExtents for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFragmentedDirectoryExtents(value uint64) (err error) {
	return instance.SetProperty("FragmentedDirectoryExtents", (value))
}

// GetFragmentedDirectoryExtents gets the value of FragmentedDirectoryExtents for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFragmentedDirectoryExtents() (value uint64, err error) {
	retValue, err := instance.GetProperty("FragmentedDirectoryExtents")
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

// SetFragmentedExtents sets the value of FragmentedExtents for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFragmentedExtents(value uint64) (err error) {
	return instance.SetProperty("FragmentedExtents", (value))
}

// GetFragmentedExtents gets the value of FragmentedExtents for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFragmentedExtents() (value uint64, err error) {
	retValue, err := instance.GetProperty("FragmentedExtents")
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

// SetFragmentedFiles sets the value of FragmentedFiles for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFragmentedFiles(value uint32) (err error) {
	return instance.SetProperty("FragmentedFiles", (value))
}

// GetFragmentedFiles gets the value of FragmentedFiles for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFragmentedFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentedFiles")
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

// SetFragmentedSpace sets the value of FragmentedSpace for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFragmentedSpace(value uint32) (err error) {
	return instance.SetProperty("FragmentedSpace", (value))
}

// GetFragmentedSpace gets the value of FragmentedSpace for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFragmentedSpace() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentedSpace")
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

// SetFreeSpaceCount sets the value of FreeSpaceCount for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyFreeSpaceCount(value uint64) (err error) {
	return instance.SetProperty("FreeSpaceCount", (value))
}

// GetFreeSpaceCount gets the value of FreeSpaceCount for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyFreeSpaceCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeSpaceCount")
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

// SetHardwareIssue sets the value of HardwareIssue for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyHardwareIssue(value uint32) (err error) {
	return instance.SetProperty("HardwareIssue", (value))
}

// GetHardwareIssue gets the value of HardwareIssue for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyHardwareIssue() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareIssue")
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

// SetInUseMFTRecords sets the value of InUseMFTRecords for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyInUseMFTRecords(value uint32) (err error) {
	return instance.SetProperty("InUseMFTRecords", (value))
}

// GetInUseMFTRecords gets the value of InUseMFTRecords for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyInUseMFTRecords() (value uint32, err error) {
	retValue, err := instance.GetProperty("InUseMFTRecords")
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

// SetInUseSlabs sets the value of InUseSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyInUseSlabs(value uint32) (err error) {
	return instance.SetProperty("InUseSlabs", (value))
}

// GetInUseSlabs gets the value of InUseSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyInUseSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("InUseSlabs")
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

// SetLargestFreeSpaceSize sets the value of LargestFreeSpaceSize for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLargestFreeSpaceSize(value uint64) (err error) {
	return instance.SetProperty("LargestFreeSpaceSize", (value))
}

// GetLargestFreeSpaceSize gets the value of LargestFreeSpaceSize for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLargestFreeSpaceSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LargestFreeSpaceSize")
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

// SetLastRunActualPurgeClusters sets the value of LastRunActualPurgeClusters for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunActualPurgeClusters(value uint64) (err error) {
	return instance.SetProperty("LastRunActualPurgeClusters", (value))
}

// GetLastRunActualPurgeClusters gets the value of LastRunActualPurgeClusters for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunActualPurgeClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastRunActualPurgeClusters")
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

// SetLastRunActualPurgeSlabs sets the value of LastRunActualPurgeSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunActualPurgeSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunActualPurgeSlabs", (value))
}

// GetLastRunActualPurgeSlabs gets the value of LastRunActualPurgeSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunActualPurgeSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunActualPurgeSlabs")
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

// SetLastRunClustersTrimmed sets the value of LastRunClustersTrimmed for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunClustersTrimmed(value uint64) (err error) {
	return instance.SetProperty("LastRunClustersTrimmed", (value))
}

// GetLastRunClustersTrimmed gets the value of LastRunClustersTrimmed for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunClustersTrimmed() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastRunClustersTrimmed")
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

// SetLastRunFullDefragTime sets the value of LastRunFullDefragTime for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunFullDefragTime(value uint64) (err error) {
	return instance.SetProperty("LastRunFullDefragTime", (value))
}

// GetLastRunFullDefragTime gets the value of LastRunFullDefragTime for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunFullDefragTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastRunFullDefragTime")
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

// SetLastRunInitialBackedSlabs sets the value of LastRunInitialBackedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunInitialBackedSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunInitialBackedSlabs", (value))
}

// GetLastRunInitialBackedSlabs gets the value of LastRunInitialBackedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunInitialBackedSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunInitialBackedSlabs")
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

// SetLastRunPercentFragmentation sets the value of LastRunPercentFragmentation for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunPercentFragmentation(value uint32) (err error) {
	return instance.SetProperty("LastRunPercentFragmentation", (value))
}

// GetLastRunPercentFragmentation gets the value of LastRunPercentFragmentation for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunPercentFragmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunPercentFragmentation")
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

// SetLastRunPinnedSlabs sets the value of LastRunPinnedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunPinnedSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunPinnedSlabs", (value))
}

// GetLastRunPinnedSlabs gets the value of LastRunPinnedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunPinnedSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunPinnedSlabs")
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

// SetLastRunPotentialPurgeSlabs sets the value of LastRunPotentialPurgeSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunPotentialPurgeSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunPotentialPurgeSlabs", (value))
}

// GetLastRunPotentialPurgeSlabs gets the value of LastRunPotentialPurgeSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunPotentialPurgeSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunPotentialPurgeSlabs")
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

// SetLastRunSpaceInefficientSlabs sets the value of LastRunSpaceInefficientSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunSpaceInefficientSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunSpaceInefficientSlabs", (value))
}

// GetLastRunSpaceInefficientSlabs gets the value of LastRunSpaceInefficientSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunSpaceInefficientSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunSpaceInefficientSlabs")
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

// SetLastRunTime sets the value of LastRunTime for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunTime(value uint64) (err error) {
	return instance.SetProperty("LastRunTime", (value))
}

// GetLastRunTime gets the value of LastRunTime for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastRunTime")
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

// SetLastRunTrimmedSlabs sets the value of LastRunTrimmedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunTrimmedSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunTrimmedSlabs", (value))
}

// GetLastRunTrimmedSlabs gets the value of LastRunTrimmedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunTrimmedSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunTrimmedSlabs")
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

// SetLastRunUnknownEvictFailSlabs sets the value of LastRunUnknownEvictFailSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunUnknownEvictFailSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunUnknownEvictFailSlabs", (value))
}

// GetLastRunUnknownEvictFailSlabs gets the value of LastRunUnknownEvictFailSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunUnknownEvictFailSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunUnknownEvictFailSlabs")
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

// SetLastRunVolsnapPinnedSlabs sets the value of LastRunVolsnapPinnedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyLastRunVolsnapPinnedSlabs(value uint32) (err error) {
	return instance.SetProperty("LastRunVolsnapPinnedSlabs", (value))
}

// GetLastRunVolsnapPinnedSlabs gets the value of LastRunVolsnapPinnedSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyLastRunVolsnapPinnedSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastRunVolsnapPinnedSlabs")
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

// SetMFTFragmentCount sets the value of MFTFragmentCount for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyMFTFragmentCount(value uint32) (err error) {
	return instance.SetProperty("MFTFragmentCount", (value))
}

// GetMFTFragmentCount gets the value of MFTFragmentCount for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyMFTFragmentCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MFTFragmentCount")
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

// SetMFTSize sets the value of MFTSize for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyMFTSize(value uint64) (err error) {
	return instance.SetProperty("MFTSize", (value))
}

// GetMFTSize gets the value of MFTSize for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyMFTSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MFTSize")
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

// SetMovableFiles sets the value of MovableFiles for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyMovableFiles(value uint32) (err error) {
	return instance.SetProperty("MovableFiles", (value))
}

// GetMovableFiles gets the value of MovableFiles for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyMovableFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("MovableFiles")
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

// SetTotalClusters sets the value of TotalClusters for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyTotalClusters(value uint64) (err error) {
	return instance.SetProperty("TotalClusters", (value))
}

// GetTotalClusters gets the value of TotalClusters for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyTotalClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalClusters")
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

// SetTotalMFTRecords sets the value of TotalMFTRecords for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyTotalMFTRecords(value uint32) (err error) {
	return instance.SetProperty("TotalMFTRecords", (value))
}

// GetTotalMFTRecords gets the value of TotalMFTRecords for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyTotalMFTRecords() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMFTRecords")
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

// SetTotalSlabs sets the value of TotalSlabs for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyTotalSlabs(value uint32) (err error) {
	return instance.SetProperty("TotalSlabs", (value))
}

// GetTotalSlabs gets the value of TotalSlabs for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyTotalSlabs() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalSlabs")
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

// SetTotalUsedClusters sets the value of TotalUsedClusters for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyTotalUsedClusters(value uint64) (err error) {
	return instance.SetProperty("TotalUsedClusters", (value))
}

// GetTotalUsedClusters gets the value of TotalUsedClusters for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyTotalUsedClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalUsedClusters")
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

// SetUnmovableFiles sets the value of UnmovableFiles for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyUnmovableFiles(value uint32) (err error) {
	return instance.SetProperty("UnmovableFiles", (value))
}

// GetUnmovableFiles gets the value of UnmovableFiles for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyUnmovableFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("UnmovableFiles")
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

// SetVolumeId sets the value of VolumeId for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyVolumeId(value interface{}) (err error) {
	return instance.SetProperty("VolumeId", (value))
}

// GetVolumeId gets the value of VolumeId for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyVolumeId() (value interface{}, err error) {
	retValue, err := instance.GetProperty("VolumeId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetVolumePathNames sets the value of VolumePathNames for the instance
func (instance *SystemConfig_V2_Defrag) SetPropertyVolumePathNames(value string) (err error) {
	return instance.SetProperty("VolumePathNames", (value))
}

// GetVolumePathNames gets the value of VolumePathNames for the instance
func (instance *SystemConfig_V2_Defrag) GetPropertyVolumePathNames() (value string, err error) {
	retValue, err := instance.GetProperty("VolumePathNames")
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
