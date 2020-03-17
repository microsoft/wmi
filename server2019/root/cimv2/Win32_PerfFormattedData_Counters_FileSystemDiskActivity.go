// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_FileSystemDiskActivity struct
type Win32_PerfFormattedData_Counters_FileSystemDiskActivity struct {
	Win32_PerfFormattedData

	//
	FileSystemBytesRead uint64

	//
	FileSystemBytesWritten uint64
}

// SetFileSystemBytesRead sets the value of FileSystemBytesRead for the instance
func (instance *Win32_PerfFormattedData_Counters_FileSystemDiskActivity) SetPropertyFileSystemBytesRead(value uint64) (err error) {
	return instance.SetProperty("FileSystemBytesRead", value)
}

// GetFileSystemBytesRead gets the value of FileSystemBytesRead for the instance
func (instance *Win32_PerfFormattedData_Counters_FileSystemDiskActivity) GetPropertyFileSystemBytesRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSystemBytesRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSystemBytesWritten sets the value of FileSystemBytesWritten for the instance
func (instance *Win32_PerfFormattedData_Counters_FileSystemDiskActivity) SetPropertyFileSystemBytesWritten(value uint64) (err error) {
	return instance.SetProperty("FileSystemBytesWritten", value)
}

// GetFileSystemBytesWritten gets the value of FileSystemBytesWritten for the instance
func (instance *Win32_PerfFormattedData_Counters_FileSystemDiskActivity) GetPropertyFileSystemBytesWritten() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSystemBytesWritten")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
