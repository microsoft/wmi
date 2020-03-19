// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DiskImage struct
type MSFT_DiskImage struct {
	*cim.WmiInstance

	//
	Attached bool

	//
	BlockSize uint64

	//
	DevicePath string

	//
	FileSize uint64

	//
	ImagePath string

	//
	LogicalSectorSize uint64

	//
	Number uint32

	//
	Size uint64

	//
	StorageType uint32
}

func NewMSFT_DiskImageEx1(instance *cim.WmiInstance) (newInstance *MSFT_DiskImage, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DiskImage{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DiskImageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DiskImage, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DiskImage{
		WmiInstance: tmp,
	}
	return
}

// SetAttached sets the value of Attached for the instance
func (instance *MSFT_DiskImage) SetPropertyAttached(value bool) (err error) {
	return instance.SetProperty("Attached", value)
}

// GetAttached gets the value of Attached for the instance
func (instance *MSFT_DiskImage) GetPropertyAttached() (value bool, err error) {
	retValue, err := instance.GetProperty("Attached")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBlockSize sets the value of BlockSize for the instance
func (instance *MSFT_DiskImage) SetPropertyBlockSize(value uint64) (err error) {
	return instance.SetProperty("BlockSize", value)
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *MSFT_DiskImage) GetPropertyBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("BlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDevicePath sets the value of DevicePath for the instance
func (instance *MSFT_DiskImage) SetPropertyDevicePath(value string) (err error) {
	return instance.SetProperty("DevicePath", value)
}

// GetDevicePath gets the value of DevicePath for the instance
func (instance *MSFT_DiskImage) GetPropertyDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("DevicePath")
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
func (instance *MSFT_DiskImage) SetPropertyFileSize(value uint64) (err error) {
	return instance.SetProperty("FileSize", value)
}

// GetFileSize gets the value of FileSize for the instance
func (instance *MSFT_DiskImage) GetPropertyFileSize() (value uint64, err error) {
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

// SetImagePath sets the value of ImagePath for the instance
func (instance *MSFT_DiskImage) SetPropertyImagePath(value string) (err error) {
	return instance.SetProperty("ImagePath", value)
}

// GetImagePath gets the value of ImagePath for the instance
func (instance *MSFT_DiskImage) GetPropertyImagePath() (value string, err error) {
	retValue, err := instance.GetProperty("ImagePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *MSFT_DiskImage) SetPropertyLogicalSectorSize(value uint64) (err error) {
	return instance.SetProperty("LogicalSectorSize", value)
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_DiskImage) GetPropertyLogicalSectorSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumber sets the value of Number for the instance
func (instance *MSFT_DiskImage) SetPropertyNumber(value uint32) (err error) {
	return instance.SetProperty("Number", value)
}

// GetNumber gets the value of Number for the instance
func (instance *MSFT_DiskImage) GetPropertyNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("Number")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_DiskImage) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_DiskImage) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageType sets the value of StorageType for the instance
func (instance *MSFT_DiskImage) SetPropertyStorageType(value uint32) (err error) {
	return instance.SetProperty("StorageType", value)
}

// GetStorageType gets the value of StorageType for the instance
func (instance *MSFT_DiskImage) GetPropertyStorageType() (value uint32, err error) {
	retValue, err := instance.GetProperty("StorageType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Access" type="uint16 "></param>
// <param name="NoDriveLetter" type="bool "></param>

// <param name="DiskImage" type="MSFT_DiskImage "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DiskImage) Mount( /* IN */ Access uint16,
	/* IN */ NoDriveLetter bool,
	/* OUT */ DiskImage MSFT_DiskImage) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Mount", Access, NoDriveLetter)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DiskImage" type="MSFT_DiskImage "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DiskImage) Dismount( /* OUT */ DiskImage MSFT_DiskImage) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Dismount")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
