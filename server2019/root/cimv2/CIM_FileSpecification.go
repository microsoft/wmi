// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_FileSpecification struct
type CIM_FileSpecification struct {
	CIM_Check

	//
	CheckSum uint32

	//
	CRC1 uint32

	//
	CRC2 uint32

	//
	CreateTimeStamp string

	//
	FileSize uint64

	//
	MD5Checksum string
}

// SetCheckSum sets the value of CheckSum for the instance
func (instance *CIM_FileSpecification) SetPropertyCheckSum(value uint32) (err error) {
	return instance.SetProperty("CheckSum", value)
}

// GetCheckSum gets the value of CheckSum for the instance
func (instance *CIM_FileSpecification) GetPropertyCheckSum() (value uint32, err error) {
	retValue, err := instance.GetProperty("CheckSum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCRC1 sets the value of CRC1 for the instance
func (instance *CIM_FileSpecification) SetPropertyCRC1(value uint32) (err error) {
	return instance.SetProperty("CRC1", value)
}

// GetCRC1 gets the value of CRC1 for the instance
func (instance *CIM_FileSpecification) GetPropertyCRC1() (value uint32, err error) {
	retValue, err := instance.GetProperty("CRC1")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCRC2 sets the value of CRC2 for the instance
func (instance *CIM_FileSpecification) SetPropertyCRC2(value uint32) (err error) {
	return instance.SetProperty("CRC2", value)
}

// GetCRC2 gets the value of CRC2 for the instance
func (instance *CIM_FileSpecification) GetPropertyCRC2() (value uint32, err error) {
	retValue, err := instance.GetProperty("CRC2")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreateTimeStamp sets the value of CreateTimeStamp for the instance
func (instance *CIM_FileSpecification) SetPropertyCreateTimeStamp(value string) (err error) {
	return instance.SetProperty("CreateTimeStamp", value)
}

// GetCreateTimeStamp gets the value of CreateTimeStamp for the instance
func (instance *CIM_FileSpecification) GetPropertyCreateTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("CreateTimeStamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSize sets the value of FileSize for the instance
func (instance *CIM_FileSpecification) SetPropertyFileSize(value uint64) (err error) {
	return instance.SetProperty("FileSize", value)
}

// GetFileSize gets the value of FileSize for the instance
func (instance *CIM_FileSpecification) GetPropertyFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMD5Checksum sets the value of MD5Checksum for the instance
func (instance *CIM_FileSpecification) SetPropertyMD5Checksum(value string) (err error) {
	return instance.SetProperty("MD5Checksum", value)
}

// GetMD5Checksum gets the value of MD5Checksum for the instance
func (instance *CIM_FileSpecification) GetPropertyMD5Checksum() (value string, err error) {
	retValue, err := instance.GetProperty("MD5Checksum")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
