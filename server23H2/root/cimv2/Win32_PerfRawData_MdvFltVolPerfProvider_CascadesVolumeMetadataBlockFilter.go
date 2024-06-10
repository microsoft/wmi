// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter struct
type Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter struct {
	*Win32_PerfRawData

	//
	CreatesPersec uint64

	//
	FileOpens uint64

	//
	ReadBytesPersec uint64

	//
	ReadsPersec uint64

	//
	WriteBytesPersec uint64

	//
	WritesPersec uint64
}

func NewWin32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilterEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetCreatesPersec sets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyCreatesPersec(value uint64) (err error) {
	return instance.SetProperty("CreatesPersec", (value))
}

// GetCreatesPersec gets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyCreatesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CreatesPersec")
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

// SetFileOpens sets the value of FileOpens for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyFileOpens(value uint64) (err error) {
	return instance.SetProperty("FileOpens", (value))
}

// GetFileOpens gets the value of FileOpens for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyFileOpens() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileOpens")
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

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
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

// SetReadsPersec sets the value of ReadsPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyReadsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsPersec", (value))
}

// GetReadsPersec gets the value of ReadsPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersec")
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

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
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

// SetWritesPersec sets the value of WritesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) SetPropertyWritesPersec(value uint64) (err error) {
	return instance.SetProperty("WritesPersec", (value))
}

// GetWritesPersec gets the value of WritesPersec for the instance
func (instance *Win32_PerfRawData_MdvFltVolPerfProvider_CascadesVolumeMetadataBlockFilter) GetPropertyWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WritesPersec")
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
