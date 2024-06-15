// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker struct
type Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker struct {
	*Win32_PerfFormattedData

	//
	MaxReadSplitSize uint32

	//
	MaxWriteSplitSize uint32

	//
	MinReadSplitSize uint32

	//
	MinWriteSplitSize uint32

	//
	ReadRequestsPersec uint64

	//
	ReadSubrequestsPersec uint64

	//
	WriteRequestsPersec uint64

	//
	WriteSubrequestsPersec uint64
}

func NewWin32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLockerEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLockerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetMaxReadSplitSize sets the value of MaxReadSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyMaxReadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MaxReadSplitSize", (value))
}

// GetMaxReadSplitSize gets the value of MaxReadSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyMaxReadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReadSplitSize")
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

// SetMaxWriteSplitSize sets the value of MaxWriteSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyMaxWriteSplitSize(value uint32) (err error) {
	return instance.SetProperty("MaxWriteSplitSize", (value))
}

// GetMaxWriteSplitSize gets the value of MaxWriteSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyMaxWriteSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxWriteSplitSize")
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

// SetMinReadSplitSize sets the value of MinReadSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyMinReadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MinReadSplitSize", (value))
}

// GetMinReadSplitSize gets the value of MinReadSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyMinReadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinReadSplitSize")
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

// SetMinWriteSplitSize sets the value of MinWriteSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyMinWriteSplitSize(value uint32) (err error) {
	return instance.SetProperty("MinWriteSplitSize", (value))
}

// GetMinWriteSplitSize gets the value of MinWriteSplitSize for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyMinWriteSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinWriteSplitSize")
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

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyReadRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadRequestsPersec", (value))
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyReadRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadRequestsPersec")
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

// SetReadSubrequestsPersec sets the value of ReadSubrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyReadSubrequestsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadSubrequestsPersec", (value))
}

// GetReadSubrequestsPersec gets the value of ReadSubrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyReadSubrequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadSubrequestsPersec")
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

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyWriteRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("WriteRequestsPersec", (value))
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyWriteRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteRequestsPersec")
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

// SetWriteSubrequestsPersec sets the value of WriteSubrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) SetPropertyWriteSubrequestsPersec(value uint64) (err error) {
	return instance.SetProperty("WriteSubrequestsPersec", (value))
}

// GetWriteSubrequestsPersec gets the value of WriteSubrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsBitLockerDriverCountersProvider_BitLocker) GetPropertyWriteSubrequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteSubrequestsPersec")
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
