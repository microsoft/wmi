// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DedupProperties struct
type MSFT_DedupProperties struct {
	*cim.WmiInstance

	//
	DedupMode uint32

	//
	InPolicyFilesCount uint64

	//
	InPolicyFilesSize uint64

	//
	OptimizedFilesCount uint64

	//
	OptimizedFilesSavingsRate uint32

	//
	OptimizedFilesSize uint64

	//
	ReFSDedupCompressionChunkSize uint32

	//
	ReFSDedupCompressionFormat uint32

	//
	ReFSDedupCompressionInProgress bool

	//
	ReFSDedupCompressionLevel uint16

	//
	ReFSDedupLastRunDuration string

	//
	ReFSDedupLastRunStatus uint64

	//
	ReFSDedupLastRunTime string

	//
	ReFSDedupMode uint32

	//
	ReFSDedupNextRunTime string

	//
	ReFSDedupPercentComplete float64

	//
	ReFSDedupProcessedOnLastRun uint64

	//
	ReFSDedupRunning bool

	//
	ReFSDedupSavingsSize uint64

	//
	ReFSDedupSavingsSizeOnLastRun uint64

	//
	ReFSDedupVolSize uint64

	//
	ReFSDedupVolumeClusterSizeBytes uint32

	//
	ReFSDedupVolumeTotalAllocatedClusters uint64

	//
	ReFSDedupVolumeTotalAllocatedCompressibleClusters uint64

	//
	ReFSDedupVolumeTotalClusters uint64

	//
	ReFSDedupVolumeTotalCompressedClusters uint64

	//
	ReFSDedupVolumeTotalCompressionSavings uint64

	//
	ReFSDedupVolumeTotalInUseCompressibleClusters uint64

	//
	SavingsRate uint32

	//
	SavingsSize uint64

	//
	UnoptimizedSize uint64
}

func NewMSFT_DedupPropertiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_DedupProperties, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DedupProperties{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DedupPropertiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DedupProperties, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DedupProperties{
		WmiInstance: tmp,
	}
	return
}

// SetDedupMode sets the value of DedupMode for the instance
func (instance *MSFT_DedupProperties) SetPropertyDedupMode(value uint32) (err error) {
	return instance.SetProperty("DedupMode", (value))
}

// GetDedupMode gets the value of DedupMode for the instance
func (instance *MSFT_DedupProperties) GetPropertyDedupMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DedupMode")
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

// SetInPolicyFilesCount sets the value of InPolicyFilesCount for the instance
func (instance *MSFT_DedupProperties) SetPropertyInPolicyFilesCount(value uint64) (err error) {
	return instance.SetProperty("InPolicyFilesCount", (value))
}

// GetInPolicyFilesCount gets the value of InPolicyFilesCount for the instance
func (instance *MSFT_DedupProperties) GetPropertyInPolicyFilesCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("InPolicyFilesCount")
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

// SetInPolicyFilesSize sets the value of InPolicyFilesSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyInPolicyFilesSize(value uint64) (err error) {
	return instance.SetProperty("InPolicyFilesSize", (value))
}

// GetInPolicyFilesSize gets the value of InPolicyFilesSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyInPolicyFilesSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("InPolicyFilesSize")
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

// SetOptimizedFilesCount sets the value of OptimizedFilesCount for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesCount(value uint64) (err error) {
	return instance.SetProperty("OptimizedFilesCount", (value))
}

// GetOptimizedFilesCount gets the value of OptimizedFilesCount for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesCount")
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

// SetOptimizedFilesSavingsRate sets the value of OptimizedFilesSavingsRate for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesSavingsRate(value uint32) (err error) {
	return instance.SetProperty("OptimizedFilesSavingsRate", (value))
}

// GetOptimizedFilesSavingsRate gets the value of OptimizedFilesSavingsRate for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesSavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesSavingsRate")
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

// SetOptimizedFilesSize sets the value of OptimizedFilesSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesSize(value uint64) (err error) {
	return instance.SetProperty("OptimizedFilesSize", (value))
}

// GetOptimizedFilesSize gets the value of OptimizedFilesSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesSize")
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

// SetReFSDedupCompressionChunkSize sets the value of ReFSDedupCompressionChunkSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupCompressionChunkSize(value uint32) (err error) {
	return instance.SetProperty("ReFSDedupCompressionChunkSize", (value))
}

// GetReFSDedupCompressionChunkSize gets the value of ReFSDedupCompressionChunkSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupCompressionChunkSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReFSDedupCompressionChunkSize")
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

// SetReFSDedupCompressionFormat sets the value of ReFSDedupCompressionFormat for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupCompressionFormat(value uint32) (err error) {
	return instance.SetProperty("ReFSDedupCompressionFormat", (value))
}

// GetReFSDedupCompressionFormat gets the value of ReFSDedupCompressionFormat for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupCompressionFormat() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReFSDedupCompressionFormat")
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

// SetReFSDedupCompressionInProgress sets the value of ReFSDedupCompressionInProgress for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupCompressionInProgress(value bool) (err error) {
	return instance.SetProperty("ReFSDedupCompressionInProgress", (value))
}

// GetReFSDedupCompressionInProgress gets the value of ReFSDedupCompressionInProgress for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupCompressionInProgress() (value bool, err error) {
	retValue, err := instance.GetProperty("ReFSDedupCompressionInProgress")
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

// SetReFSDedupCompressionLevel sets the value of ReFSDedupCompressionLevel for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupCompressionLevel(value uint16) (err error) {
	return instance.SetProperty("ReFSDedupCompressionLevel", (value))
}

// GetReFSDedupCompressionLevel gets the value of ReFSDedupCompressionLevel for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupCompressionLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReFSDedupCompressionLevel")
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

// SetReFSDedupLastRunDuration sets the value of ReFSDedupLastRunDuration for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupLastRunDuration(value string) (err error) {
	return instance.SetProperty("ReFSDedupLastRunDuration", (value))
}

// GetReFSDedupLastRunDuration gets the value of ReFSDedupLastRunDuration for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupLastRunDuration() (value string, err error) {
	retValue, err := instance.GetProperty("ReFSDedupLastRunDuration")
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

// SetReFSDedupLastRunStatus sets the value of ReFSDedupLastRunStatus for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupLastRunStatus(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupLastRunStatus", (value))
}

// GetReFSDedupLastRunStatus gets the value of ReFSDedupLastRunStatus for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupLastRunStatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupLastRunStatus")
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

// SetReFSDedupLastRunTime sets the value of ReFSDedupLastRunTime for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupLastRunTime(value string) (err error) {
	return instance.SetProperty("ReFSDedupLastRunTime", (value))
}

// GetReFSDedupLastRunTime gets the value of ReFSDedupLastRunTime for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupLastRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("ReFSDedupLastRunTime")
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

// SetReFSDedupMode sets the value of ReFSDedupMode for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupMode(value uint32) (err error) {
	return instance.SetProperty("ReFSDedupMode", (value))
}

// GetReFSDedupMode gets the value of ReFSDedupMode for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReFSDedupMode")
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

// SetReFSDedupNextRunTime sets the value of ReFSDedupNextRunTime for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupNextRunTime(value string) (err error) {
	return instance.SetProperty("ReFSDedupNextRunTime", (value))
}

// GetReFSDedupNextRunTime gets the value of ReFSDedupNextRunTime for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupNextRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("ReFSDedupNextRunTime")
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

// SetReFSDedupPercentComplete sets the value of ReFSDedupPercentComplete for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupPercentComplete(value float64) (err error) {
	return instance.SetProperty("ReFSDedupPercentComplete", (value))
}

// GetReFSDedupPercentComplete gets the value of ReFSDedupPercentComplete for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupPercentComplete() (value float64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupPercentComplete")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetReFSDedupProcessedOnLastRun sets the value of ReFSDedupProcessedOnLastRun for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupProcessedOnLastRun(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupProcessedOnLastRun", (value))
}

// GetReFSDedupProcessedOnLastRun gets the value of ReFSDedupProcessedOnLastRun for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupProcessedOnLastRun() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupProcessedOnLastRun")
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

// SetReFSDedupRunning sets the value of ReFSDedupRunning for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupRunning(value bool) (err error) {
	return instance.SetProperty("ReFSDedupRunning", (value))
}

// GetReFSDedupRunning gets the value of ReFSDedupRunning for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupRunning() (value bool, err error) {
	retValue, err := instance.GetProperty("ReFSDedupRunning")
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

// SetReFSDedupSavingsSize sets the value of ReFSDedupSavingsSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupSavingsSize(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupSavingsSize", (value))
}

// GetReFSDedupSavingsSize gets the value of ReFSDedupSavingsSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupSavingsSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupSavingsSize")
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

// SetReFSDedupSavingsSizeOnLastRun sets the value of ReFSDedupSavingsSizeOnLastRun for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupSavingsSizeOnLastRun(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupSavingsSizeOnLastRun", (value))
}

// GetReFSDedupSavingsSizeOnLastRun gets the value of ReFSDedupSavingsSizeOnLastRun for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupSavingsSizeOnLastRun() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupSavingsSizeOnLastRun")
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

// SetReFSDedupVolSize sets the value of ReFSDedupVolSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolSize(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolSize", (value))
}

// GetReFSDedupVolSize gets the value of ReFSDedupVolSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolSize")
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

// SetReFSDedupVolumeClusterSizeBytes sets the value of ReFSDedupVolumeClusterSizeBytes for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeClusterSizeBytes(value uint32) (err error) {
	return instance.SetProperty("ReFSDedupVolumeClusterSizeBytes", (value))
}

// GetReFSDedupVolumeClusterSizeBytes gets the value of ReFSDedupVolumeClusterSizeBytes for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeClusterSizeBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeClusterSizeBytes")
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

// SetReFSDedupVolumeTotalAllocatedClusters sets the value of ReFSDedupVolumeTotalAllocatedClusters for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalAllocatedClusters(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalAllocatedClusters", (value))
}

// GetReFSDedupVolumeTotalAllocatedClusters gets the value of ReFSDedupVolumeTotalAllocatedClusters for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalAllocatedClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalAllocatedClusters")
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

// SetReFSDedupVolumeTotalAllocatedCompressibleClusters sets the value of ReFSDedupVolumeTotalAllocatedCompressibleClusters for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalAllocatedCompressibleClusters(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalAllocatedCompressibleClusters", (value))
}

// GetReFSDedupVolumeTotalAllocatedCompressibleClusters gets the value of ReFSDedupVolumeTotalAllocatedCompressibleClusters for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalAllocatedCompressibleClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalAllocatedCompressibleClusters")
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

// SetReFSDedupVolumeTotalClusters sets the value of ReFSDedupVolumeTotalClusters for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalClusters(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalClusters", (value))
}

// GetReFSDedupVolumeTotalClusters gets the value of ReFSDedupVolumeTotalClusters for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalClusters")
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

// SetReFSDedupVolumeTotalCompressedClusters sets the value of ReFSDedupVolumeTotalCompressedClusters for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalCompressedClusters(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalCompressedClusters", (value))
}

// GetReFSDedupVolumeTotalCompressedClusters gets the value of ReFSDedupVolumeTotalCompressedClusters for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalCompressedClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalCompressedClusters")
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

// SetReFSDedupVolumeTotalCompressionSavings sets the value of ReFSDedupVolumeTotalCompressionSavings for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalCompressionSavings(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalCompressionSavings", (value))
}

// GetReFSDedupVolumeTotalCompressionSavings gets the value of ReFSDedupVolumeTotalCompressionSavings for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalCompressionSavings() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalCompressionSavings")
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

// SetReFSDedupVolumeTotalInUseCompressibleClusters sets the value of ReFSDedupVolumeTotalInUseCompressibleClusters for the instance
func (instance *MSFT_DedupProperties) SetPropertyReFSDedupVolumeTotalInUseCompressibleClusters(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupVolumeTotalInUseCompressibleClusters", (value))
}

// GetReFSDedupVolumeTotalInUseCompressibleClusters gets the value of ReFSDedupVolumeTotalInUseCompressibleClusters for the instance
func (instance *MSFT_DedupProperties) GetPropertyReFSDedupVolumeTotalInUseCompressibleClusters() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupVolumeTotalInUseCompressibleClusters")
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

// SetSavingsRate sets the value of SavingsRate for the instance
func (instance *MSFT_DedupProperties) SetPropertySavingsRate(value uint32) (err error) {
	return instance.SetProperty("SavingsRate", (value))
}

// GetSavingsRate gets the value of SavingsRate for the instance
func (instance *MSFT_DedupProperties) GetPropertySavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("SavingsRate")
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

// SetSavingsSize sets the value of SavingsSize for the instance
func (instance *MSFT_DedupProperties) SetPropertySavingsSize(value uint64) (err error) {
	return instance.SetProperty("SavingsSize", (value))
}

// GetSavingsSize gets the value of SavingsSize for the instance
func (instance *MSFT_DedupProperties) GetPropertySavingsSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SavingsSize")
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

// SetUnoptimizedSize sets the value of UnoptimizedSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyUnoptimizedSize(value uint64) (err error) {
	return instance.SetProperty("UnoptimizedSize", (value))
}

// GetUnoptimizedSize gets the value of UnoptimizedSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyUnoptimizedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnoptimizedSize")
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
