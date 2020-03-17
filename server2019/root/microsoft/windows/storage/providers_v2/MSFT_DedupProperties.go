// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DedupProperties struct
type MSFT_DedupProperties struct {
	cim.WmiInstance

	// The number of files that currently qualify for optimization.
	InPolicyFilesCount uint64

	// The aggregate size of all files that currently qualify for optimization.
	InPolicyFilesSize uint64

	// The number of optimized files on the volume.
	OptimizedFilesCount uint64

	// The ratio of deduplication savings to the logical size of all optimized files on the volume, expressed as a percentage.
	OptimizedFilesSavingsRate uint32

	// The total logical size of all optimized files on the volume, in bytes.
	OptimizedFilesSize uint64

	// The ratio of deduplication savings to the logical size of all of the files on the volume, expressed as a percentage.
	SavingsRate uint32

	// The difference between the logical size of the optimized files and the logical size of the store (the deduplicated user data plus deduplication metadata).
	SavingsSize uint64

	// The total logical size of all files on the volume, in bytes. This is an estimate of the volume used space if deduplication feature was disabled.
	UnoptimizedSize uint64
}

// SetInPolicyFilesCount sets the value of InPolicyFilesCount for the instance
func (instance *MSFT_DedupProperties) SetPropertyInPolicyFilesCount(value uint64) (err error) {
	return instance.SetProperty("InPolicyFilesCount", value)
}

// GetInPolicyFilesCount gets the value of InPolicyFilesCount for the instance
func (instance *MSFT_DedupProperties) GetPropertyInPolicyFilesCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("InPolicyFilesCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInPolicyFilesSize sets the value of InPolicyFilesSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyInPolicyFilesSize(value uint64) (err error) {
	return instance.SetProperty("InPolicyFilesSize", value)
}

// GetInPolicyFilesSize gets the value of InPolicyFilesSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyInPolicyFilesSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("InPolicyFilesSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimizedFilesCount sets the value of OptimizedFilesCount for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesCount(value uint64) (err error) {
	return instance.SetProperty("OptimizedFilesCount", value)
}

// GetOptimizedFilesCount gets the value of OptimizedFilesCount for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimizedFilesSavingsRate sets the value of OptimizedFilesSavingsRate for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesSavingsRate(value uint32) (err error) {
	return instance.SetProperty("OptimizedFilesSavingsRate", value)
}

// GetOptimizedFilesSavingsRate gets the value of OptimizedFilesSavingsRate for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesSavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesSavingsRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimizedFilesSize sets the value of OptimizedFilesSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyOptimizedFilesSize(value uint64) (err error) {
	return instance.SetProperty("OptimizedFilesSize", value)
}

// GetOptimizedFilesSize gets the value of OptimizedFilesSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyOptimizedFilesSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimizedFilesSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSavingsRate sets the value of SavingsRate for the instance
func (instance *MSFT_DedupProperties) SetPropertySavingsRate(value uint32) (err error) {
	return instance.SetProperty("SavingsRate", value)
}

// GetSavingsRate gets the value of SavingsRate for the instance
func (instance *MSFT_DedupProperties) GetPropertySavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("SavingsRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSavingsSize sets the value of SavingsSize for the instance
func (instance *MSFT_DedupProperties) SetPropertySavingsSize(value uint64) (err error) {
	return instance.SetProperty("SavingsSize", value)
}

// GetSavingsSize gets the value of SavingsSize for the instance
func (instance *MSFT_DedupProperties) GetPropertySavingsSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SavingsSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnoptimizedSize sets the value of UnoptimizedSize for the instance
func (instance *MSFT_DedupProperties) SetPropertyUnoptimizedSize(value uint64) (err error) {
	return instance.SetProperty("UnoptimizedSize", value)
}

// GetUnoptimizedSize gets the value of UnoptimizedSize for the instance
func (instance *MSFT_DedupProperties) GetPropertyUnoptimizedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnoptimizedSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
