// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData struct
type Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData struct {
	*Win32_PerfRawData

	//
	CreatesPersec uint64

	//
	FileOpens uint64

	//
	OpenedFiles uint64

	//
	OpenedHandles uint64

	//
	OpenedStreams uint64

	//
	ReadBytesPersec uint64

	//
	ReadsPersec uint64

	//
	SplitIOPerSec uint64

	//
	WriteBytesPersec uint64

	//
	WritesPersec uint64
}

func NewWin32_PerfRawData_DvFilterPerfProvider_CascadesVolumeDataEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_DvFilterPerfProvider_CascadesVolumeDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetCreatesPersec sets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyCreatesPersec(value uint64) (err error) {
	return instance.SetProperty("CreatesPersec", (value))
}

// GetCreatesPersec gets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyCreatesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyFileOpens(value uint64) (err error) {
	return instance.SetProperty("FileOpens", (value))
}

// GetFileOpens gets the value of FileOpens for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyFileOpens() (value uint64, err error) {
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

// SetOpenedFiles sets the value of OpenedFiles for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyOpenedFiles(value uint64) (err error) {
	return instance.SetProperty("OpenedFiles", (value))
}

// GetOpenedFiles gets the value of OpenedFiles for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyOpenedFiles() (value uint64, err error) {
	retValue, err := instance.GetProperty("OpenedFiles")
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

// SetOpenedHandles sets the value of OpenedHandles for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyOpenedHandles(value uint64) (err error) {
	return instance.SetProperty("OpenedHandles", (value))
}

// GetOpenedHandles gets the value of OpenedHandles for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyOpenedHandles() (value uint64, err error) {
	retValue, err := instance.GetProperty("OpenedHandles")
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

// SetOpenedStreams sets the value of OpenedStreams for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyOpenedStreams(value uint64) (err error) {
	return instance.SetProperty("OpenedStreams", (value))
}

// GetOpenedStreams gets the value of OpenedStreams for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyOpenedStreams() (value uint64, err error) {
	retValue, err := instance.GetProperty("OpenedStreams")
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
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyReadBytesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyReadsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsPersec", (value))
}

// GetReadsPersec gets the value of ReadsPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyReadsPersec() (value uint64, err error) {
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

// SetSplitIOPerSec sets the value of SplitIOPerSec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertySplitIOPerSec(value uint64) (err error) {
	return instance.SetProperty("SplitIOPerSec", (value))
}

// GetSplitIOPerSec gets the value of SplitIOPerSec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertySplitIOPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SplitIOPerSec")
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
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyWriteBytesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) SetPropertyWritesPersec(value uint64) (err error) {
	return instance.SetProperty("WritesPersec", (value))
}

// GetWritesPersec gets the value of WritesPersec for the instance
func (instance *Win32_PerfRawData_DvFilterPerfProvider_CascadesVolumeData) GetPropertyWritesPersec() (value uint64, err error) {
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
