// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SMBIOSMemory struct
type Win32_SMBIOSMemory struct {
	*CIM_StorageExtent

	//
	AdditionalErrorData []uint8

	//
	CorrectableError bool

	//
	EndingAddress uint64

	//
	ErrorAccess uint16

	//
	ErrorAddress uint64

	//
	ErrorData []uint8

	//
	ErrorDataOrder uint16

	//
	ErrorInfo uint16

	//
	ErrorResolution uint64

	//
	ErrorTime string

	//
	ErrorTransferSize uint32

	//
	OtherErrorDescription string

	//
	StartingAddress uint64

	//
	SystemLevelAddress bool
}

func NewWin32_SMBIOSMemoryEx1(instance *cim.WmiInstance) (newInstance *Win32_SMBIOSMemory, err error) {
	tmp, err := NewCIM_StorageExtentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_SMBIOSMemory{
		CIM_StorageExtent: tmp,
	}
	return
}

func NewWin32_SMBIOSMemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SMBIOSMemory, err error) {
	tmp, err := NewCIM_StorageExtentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SMBIOSMemory{
		CIM_StorageExtent: tmp,
	}
	return
}

// SetAdditionalErrorData sets the value of AdditionalErrorData for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyAdditionalErrorData(value []uint8) (err error) {
	return instance.SetProperty("AdditionalErrorData", value)
}

// GetAdditionalErrorData gets the value of AdditionalErrorData for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyAdditionalErrorData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AdditionalErrorData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCorrectableError sets the value of CorrectableError for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyCorrectableError(value bool) (err error) {
	return instance.SetProperty("CorrectableError", value)
}

// GetCorrectableError gets the value of CorrectableError for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyCorrectableError() (value bool, err error) {
	retValue, err := instance.GetProperty("CorrectableError")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEndingAddress sets the value of EndingAddress for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyEndingAddress(value uint64) (err error) {
	return instance.SetProperty("EndingAddress", value)
}

// GetEndingAddress gets the value of EndingAddress for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyEndingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("EndingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorAccess sets the value of ErrorAccess for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorAccess(value uint16) (err error) {
	return instance.SetProperty("ErrorAccess", value)
}

// GetErrorAccess gets the value of ErrorAccess for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorAccess() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorAddress sets the value of ErrorAddress for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorAddress(value uint64) (err error) {
	return instance.SetProperty("ErrorAddress", value)
}

// GetErrorAddress gets the value of ErrorAddress for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorData sets the value of ErrorData for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorData(value []uint8) (err error) {
	return instance.SetProperty("ErrorData", value)
}

// GetErrorData gets the value of ErrorData for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ErrorData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDataOrder sets the value of ErrorDataOrder for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorDataOrder(value uint16) (err error) {
	return instance.SetProperty("ErrorDataOrder", value)
}

// GetErrorDataOrder gets the value of ErrorDataOrder for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorDataOrder() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorDataOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorInfo sets the value of ErrorInfo for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorInfo(value uint16) (err error) {
	return instance.SetProperty("ErrorInfo", value)
}

// GetErrorInfo gets the value of ErrorInfo for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorResolution sets the value of ErrorResolution for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorResolution(value uint64) (err error) {
	return instance.SetProperty("ErrorResolution", value)
}

// GetErrorResolution gets the value of ErrorResolution for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorResolution() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorTime sets the value of ErrorTime for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorTime(value string) (err error) {
	return instance.SetProperty("ErrorTime", value)
}

// GetErrorTime gets the value of ErrorTime for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorTime() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorTransferSize sets the value of ErrorTransferSize for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyErrorTransferSize(value uint32) (err error) {
	return instance.SetProperty("ErrorTransferSize", value)
}

// GetErrorTransferSize gets the value of ErrorTransferSize for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyErrorTransferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorTransferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherErrorDescription sets the value of OtherErrorDescription for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyOtherErrorDescription(value string) (err error) {
	return instance.SetProperty("OtherErrorDescription", value)
}

// GetOtherErrorDescription gets the value of OtherErrorDescription for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyOtherErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartingAddress sets the value of StartingAddress for the instance
func (instance *Win32_SMBIOSMemory) SetPropertyStartingAddress(value uint64) (err error) {
	return instance.SetProperty("StartingAddress", value)
}

// GetStartingAddress gets the value of StartingAddress for the instance
func (instance *Win32_SMBIOSMemory) GetPropertyStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemLevelAddress sets the value of SystemLevelAddress for the instance
func (instance *Win32_SMBIOSMemory) SetPropertySystemLevelAddress(value bool) (err error) {
	return instance.SetProperty("SystemLevelAddress", value)
}

// GetSystemLevelAddress gets the value of SystemLevelAddress for the instance
func (instance *Win32_SMBIOSMemory) GetPropertySystemLevelAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("SystemLevelAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
