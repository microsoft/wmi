// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_NTEventlogFile struct
type Win32_NTEventlogFile struct {
	CIM_DataFile

	//
	LogfileName string

	//
	MaxFileSize uint32

	//
	NumberOfRecords uint32

	//
	OverwriteOutDated uint32

	//
	OverWritePolicy string

	//
	Sources []string
}

// SetLogfileName sets the value of LogfileName for the instance
func (instance *Win32_NTEventlogFile) SetPropertyLogfileName(value string) (err error) {
	return instance.SetProperty("LogfileName", value)
}

// GetLogfileName gets the value of LogfileName for the instance
func (instance *Win32_NTEventlogFile) GetPropertyLogfileName() (value string, err error) {
	retValue, err := instance.GetProperty("LogfileName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxFileSize sets the value of MaxFileSize for the instance
func (instance *Win32_NTEventlogFile) SetPropertyMaxFileSize(value uint32) (err error) {
	return instance.SetProperty("MaxFileSize", value)
}

// GetMaxFileSize gets the value of MaxFileSize for the instance
func (instance *Win32_NTEventlogFile) GetPropertyMaxFileSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfRecords sets the value of NumberOfRecords for the instance
func (instance *Win32_NTEventlogFile) SetPropertyNumberOfRecords(value uint32) (err error) {
	return instance.SetProperty("NumberOfRecords", value)
}

// GetNumberOfRecords gets the value of NumberOfRecords for the instance
func (instance *Win32_NTEventlogFile) GetPropertyNumberOfRecords() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfRecords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOverwriteOutDated sets the value of OverwriteOutDated for the instance
func (instance *Win32_NTEventlogFile) SetPropertyOverwriteOutDated(value uint32) (err error) {
	return instance.SetProperty("OverwriteOutDated", value)
}

// GetOverwriteOutDated gets the value of OverwriteOutDated for the instance
func (instance *Win32_NTEventlogFile) GetPropertyOverwriteOutDated() (value uint32, err error) {
	retValue, err := instance.GetProperty("OverwriteOutDated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOverWritePolicy sets the value of OverWritePolicy for the instance
func (instance *Win32_NTEventlogFile) SetPropertyOverWritePolicy(value string) (err error) {
	return instance.SetProperty("OverWritePolicy", value)
}

// GetOverWritePolicy gets the value of OverWritePolicy for the instance
func (instance *Win32_NTEventlogFile) GetPropertyOverWritePolicy() (value string, err error) {
	retValue, err := instance.GetProperty("OverWritePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSources sets the value of Sources for the instance
func (instance *Win32_NTEventlogFile) SetPropertySources(value []string) (err error) {
	return instance.SetProperty("Sources", value)
}

// GetSources gets the value of Sources for the instance
func (instance *Win32_NTEventlogFile) GetPropertySources() (value []string, err error) {
	retValue, err := instance.GetProperty("Sources")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ArchiveFileName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_NTEventlogFile) ClearEventlog( /* IN */ ArchiveFileName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearEventlog", ArchiveFileName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ArchiveFileName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_NTEventlogFile) BackupEventlog( /* IN */ ArchiveFileName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("BackupEventlog", ArchiveFileName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
