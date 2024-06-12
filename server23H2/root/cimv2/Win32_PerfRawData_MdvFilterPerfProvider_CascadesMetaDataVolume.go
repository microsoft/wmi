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

// Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume struct
type Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume struct {
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
}

func NewWin32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolumeEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetCreatesPersec sets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) SetPropertyCreatesPersec(value uint64) (err error) {
	return instance.SetProperty("CreatesPersec", (value))
}

// GetCreatesPersec gets the value of CreatesPersec for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) GetPropertyCreatesPersec() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) SetPropertyFileOpens(value uint64) (err error) {
	return instance.SetProperty("FileOpens", (value))
}

// GetFileOpens gets the value of FileOpens for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) GetPropertyFileOpens() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) SetPropertyOpenedFiles(value uint64) (err error) {
	return instance.SetProperty("OpenedFiles", (value))
}

// GetOpenedFiles gets the value of OpenedFiles for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) GetPropertyOpenedFiles() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) SetPropertyOpenedHandles(value uint64) (err error) {
	return instance.SetProperty("OpenedHandles", (value))
}

// GetOpenedHandles gets the value of OpenedHandles for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) GetPropertyOpenedHandles() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) SetPropertyOpenedStreams(value uint64) (err error) {
	return instance.SetProperty("OpenedStreams", (value))
}

// GetOpenedStreams gets the value of OpenedStreams for the instance
func (instance *Win32_PerfRawData_MdvFilterPerfProvider_CascadesMetaDataVolume) GetPropertyOpenedStreams() (value uint64, err error) {
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
