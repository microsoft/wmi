// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator struct
type Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator struct {
	Win32_PerfRawData

	//
	CreateFile uint64

	//
	CreateFilePersec uint64

	//
	IOReadBytes uint64

	//
	IOReadBytesPersec uint64

	//
	IOReads uint64

	//
	IOReadsPersec uint64

	//
	IOWriteBytesPersec uint64

	//
	IOWrites uint64

	//
	IOWritesBytes uint64

	//
	IOWritesPersec uint64

	//
	MetadataIO uint64

	//
	MetadataIOPersec uint64
}

// SetCreateFile sets the value of CreateFile for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyCreateFile(value uint64) (err error) {
	return instance.SetProperty("CreateFile", value)
}

// GetCreateFile gets the value of CreateFile for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyCreateFile() (value uint64, err error) {
	retValue, err := instance.GetProperty("CreateFile")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreateFilePersec sets the value of CreateFilePersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyCreateFilePersec(value uint64) (err error) {
	return instance.SetProperty("CreateFilePersec", value)
}

// GetCreateFilePersec gets the value of CreateFilePersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyCreateFilePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CreateFilePersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytes sets the value of IOReadBytes for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOReadBytes(value uint64) (err error) {
	return instance.SetProperty("IOReadBytes", value)
}

// GetIOReadBytes gets the value of IOReadBytes for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytesPersec sets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesPersec", value)
}

// GetIOReadBytesPersec gets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReads sets the value of IOReads for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOReads(value uint64) (err error) {
	return instance.SetProperty("IOReads", value)
}

// GetIOReads gets the value of IOReads for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadsPersec sets the value of IOReadsPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOReadsPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadsPersec", value)
}

// GetIOReadsPersec gets the value of IOReadsPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWriteBytesPersec sets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytesPersec", value)
}

// GetIOWriteBytesPersec gets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWrites sets the value of IOWrites for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOWrites(value uint64) (err error) {
	return instance.SetProperty("IOWrites", value)
}

// GetIOWrites gets the value of IOWrites for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOWrites() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWrites")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWritesBytes sets the value of IOWritesBytes for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOWritesBytes(value uint64) (err error) {
	return instance.SetProperty("IOWritesBytes", value)
}

// GetIOWritesBytes gets the value of IOWritesBytes for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOWritesBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWritesBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWritesPersec sets the value of IOWritesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyIOWritesPersec(value uint64) (err error) {
	return instance.SetProperty("IOWritesPersec", value)
}

// GetIOWritesPersec gets the value of IOWritesPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyIOWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetadataIO sets the value of MetadataIO for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyMetadataIO(value uint64) (err error) {
	return instance.SetProperty("MetadataIO", value)
}

// GetMetadataIO gets the value of MetadataIO for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyMetadataIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("MetadataIO")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetadataIOPersec sets the value of MetadataIOPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) SetPropertyMetadataIOPersec(value uint64) (err error) {
	return instance.SetProperty("MetadataIOPersec", value)
}

// GetMetadataIOPersec gets the value of MetadataIOPersec for the instance
func (instance *Win32_PerfRawData_Csv20FilterPerfProvider_ClusterCSVCoordinator) GetPropertyMetadataIOPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MetadataIOPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
