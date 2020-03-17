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

// MSFT_SBLTargetDisk struct
type MSFT_SBLTargetDisk struct {
	cim.WmiInstance

	//
	CacheMode uint32

	//
	CurrentUsage uint32

	//
	DesiredUsage uint32

	//
	DeviceNumber uint32

	//
	Identifier string

	//
	IsFlash bool

	//
	IsSblCacheDevice bool

	//
	LastStateChangeTime string

	//
	ReadMediaErrorCount uint64

	//
	ReadTotalErrorCount uint64

	//
	SblAttributes uint32

	//
	State uint32

	//
	WriteMediaErrorCount uint64

	//
	WriteTotalErrorCount uint64
}

// SetCacheMode sets the value of CacheMode for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyCacheMode(value uint32) (err error) {
	return instance.SetProperty("CacheMode", value)
}

// GetCacheMode gets the value of CacheMode for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyCacheMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentUsage sets the value of CurrentUsage for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyCurrentUsage(value uint32) (err error) {
	return instance.SetProperty("CurrentUsage", value)
}

// GetCurrentUsage gets the value of CurrentUsage for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyCurrentUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDesiredUsage sets the value of DesiredUsage for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyDesiredUsage(value uint32) (err error) {
	return instance.SetProperty("DesiredUsage", value)
}

// GetDesiredUsage gets the value of DesiredUsage for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyDesiredUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("DesiredUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("DeviceNumber", value)
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentifier sets the value of Identifier for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyIdentifier(value string) (err error) {
	return instance.SetProperty("Identifier", value)
}

// GetIdentifier gets the value of Identifier for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("Identifier")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsFlash sets the value of IsFlash for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyIsFlash(value bool) (err error) {
	return instance.SetProperty("IsFlash", value)
}

// GetIsFlash gets the value of IsFlash for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyIsFlash() (value bool, err error) {
	retValue, err := instance.GetProperty("IsFlash")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsSblCacheDevice sets the value of IsSblCacheDevice for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyIsSblCacheDevice(value bool) (err error) {
	return instance.SetProperty("IsSblCacheDevice", value)
}

// GetIsSblCacheDevice gets the value of IsSblCacheDevice for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyIsSblCacheDevice() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSblCacheDevice")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastStateChangeTime sets the value of LastStateChangeTime for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyLastStateChangeTime(value string) (err error) {
	return instance.SetProperty("LastStateChangeTime", value)
}

// GetLastStateChangeTime gets the value of LastStateChangeTime for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyLastStateChangeTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastStateChangeTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadMediaErrorCount sets the value of ReadMediaErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyReadMediaErrorCount(value uint64) (err error) {
	return instance.SetProperty("ReadMediaErrorCount", value)
}

// GetReadMediaErrorCount gets the value of ReadMediaErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyReadMediaErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadMediaErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadTotalErrorCount sets the value of ReadTotalErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyReadTotalErrorCount(value uint64) (err error) {
	return instance.SetProperty("ReadTotalErrorCount", value)
}

// GetReadTotalErrorCount gets the value of ReadTotalErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyReadTotalErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadTotalErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSblAttributes sets the value of SblAttributes for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertySblAttributes(value uint32) (err error) {
	return instance.SetProperty("SblAttributes", value)
}

// GetSblAttributes gets the value of SblAttributes for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertySblAttributes() (value uint32, err error) {
	retValue, err := instance.GetProperty("SblAttributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteMediaErrorCount sets the value of WriteMediaErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyWriteMediaErrorCount(value uint64) (err error) {
	return instance.SetProperty("WriteMediaErrorCount", value)
}

// GetWriteMediaErrorCount gets the value of WriteMediaErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyWriteMediaErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteMediaErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteTotalErrorCount sets the value of WriteTotalErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) SetPropertyWriteTotalErrorCount(value uint64) (err error) {
	return instance.SetProperty("WriteTotalErrorCount", value)
}

// GetWriteTotalErrorCount gets the value of WriteTotalErrorCount for the instance
func (instance *MSFT_SBLTargetDisk) GetPropertyWriteTotalErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteTotalErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
