// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PhysicalExtent struct
type MSFT_PhysicalExtent struct {
	cim.WmiInstance

	//
	ColumnNumber uint16

	//
	CopyNumber uint16

	//
	Flags uint64

	//
	OperationalDetails []string

	//
	OperationalStatus []uint16

	//
	PhysicalDiskOffset uint64

	//
	PhysicalDiskUniqueId string

	//
	ReplacementCopyNumber uint16

	//
	Size uint64

	//
	StorageTierUniqueId string

	//
	VirtualDiskOffset uint64

	//
	VirtualDiskUniqueId string
}

// SetColumnNumber sets the value of ColumnNumber for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyColumnNumber(value uint16) (err error) {
	return instance.SetProperty("ColumnNumber", value)
}

// GetColumnNumber gets the value of ColumnNumber for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyColumnNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("ColumnNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyNumber sets the value of CopyNumber for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyCopyNumber(value uint16) (err error) {
	return instance.SetProperty("CopyNumber", value)
}

// GetCopyNumber gets the value of CopyNumber for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyCopyNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("CopyNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyFlags(value uint64) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalDetails sets the value of OperationalDetails for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyOperationalDetails(value []string) (err error) {
	return instance.SetProperty("OperationalDetails", value)
}

// GetOperationalDetails gets the value of OperationalDetails for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyOperationalDetails() (value []string, err error) {
	retValue, err := instance.GetProperty("OperationalDetails")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskOffset sets the value of PhysicalDiskOffset for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyPhysicalDiskOffset(value uint64) (err error) {
	return instance.SetProperty("PhysicalDiskOffset", value)
}

// GetPhysicalDiskOffset gets the value of PhysicalDiskOffset for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyPhysicalDiskOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskOffset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskUniqueId sets the value of PhysicalDiskUniqueId for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyPhysicalDiskUniqueId(value string) (err error) {
	return instance.SetProperty("PhysicalDiskUniqueId", value)
}

// GetPhysicalDiskUniqueId gets the value of PhysicalDiskUniqueId for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyPhysicalDiskUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplacementCopyNumber sets the value of ReplacementCopyNumber for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyReplacementCopyNumber(value uint16) (err error) {
	return instance.SetProperty("ReplacementCopyNumber", value)
}

// GetReplacementCopyNumber gets the value of ReplacementCopyNumber for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyReplacementCopyNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplacementCopyNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_PhysicalExtent) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_PhysicalExtent) GetPropertySize() (value uint64, err error) {
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

// SetStorageTierUniqueId sets the value of StorageTierUniqueId for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyStorageTierUniqueId(value string) (err error) {
	return instance.SetProperty("StorageTierUniqueId", value)
}

// GetStorageTierUniqueId gets the value of StorageTierUniqueId for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyStorageTierUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("StorageTierUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDiskOffset sets the value of VirtualDiskOffset for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyVirtualDiskOffset(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskOffset", value)
}

// GetVirtualDiskOffset gets the value of VirtualDiskOffset for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyVirtualDiskOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskOffset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDiskUniqueId sets the value of VirtualDiskUniqueId for the instance
func (instance *MSFT_PhysicalExtent) SetPropertyVirtualDiskUniqueId(value string) (err error) {
	return instance.SetProperty("VirtualDiskUniqueId", value)
}

// GetVirtualDiskUniqueId gets the value of VirtualDiskUniqueId for the instance
func (instance *MSFT_PhysicalExtent) GetPropertyVirtualDiskUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualDiskUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
