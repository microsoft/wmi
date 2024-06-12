// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_WmiHDSplitCurrentConfig struct
type MSNdis_WmiHDSplitCurrentConfig struct {
	*MSNdis

	//
	BackfillSize uint32

	//
	CurrentCapabilities uint32

	//
	HardwareCapabilities uint32

	//
	HDSplitCombineFlags uint32

	//
	HDSplitFlags uint32

	//
	Header MSNdis_ObjectHeader

	//
	MaxHeaderSize uint32
}

func NewMSNdis_WmiHDSplitCurrentConfigEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiHDSplitCurrentConfig, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiHDSplitCurrentConfig{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiHDSplitCurrentConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiHDSplitCurrentConfig, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiHDSplitCurrentConfig{
		MSNdis: tmp,
	}
	return
}

// SetBackfillSize sets the value of BackfillSize for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyBackfillSize(value uint32) (err error) {
	return instance.SetProperty("BackfillSize", (value))
}

// GetBackfillSize gets the value of BackfillSize for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyBackfillSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("BackfillSize")
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

// SetCurrentCapabilities sets the value of CurrentCapabilities for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyCurrentCapabilities(value uint32) (err error) {
	return instance.SetProperty("CurrentCapabilities", (value))
}

// GetCurrentCapabilities gets the value of CurrentCapabilities for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyCurrentCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentCapabilities")
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

// SetHardwareCapabilities sets the value of HardwareCapabilities for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyHardwareCapabilities(value uint32) (err error) {
	return instance.SetProperty("HardwareCapabilities", (value))
}

// GetHardwareCapabilities gets the value of HardwareCapabilities for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyHardwareCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareCapabilities")
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

// SetHDSplitCombineFlags sets the value of HDSplitCombineFlags for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyHDSplitCombineFlags(value uint32) (err error) {
	return instance.SetProperty("HDSplitCombineFlags", (value))
}

// GetHDSplitCombineFlags gets the value of HDSplitCombineFlags for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyHDSplitCombineFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("HDSplitCombineFlags")
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

// SetHDSplitFlags sets the value of HDSplitFlags for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyHDSplitFlags(value uint32) (err error) {
	return instance.SetProperty("HDSplitFlags", (value))
}

// GetHDSplitFlags gets the value of HDSplitFlags for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyHDSplitFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("HDSplitFlags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetMaxHeaderSize sets the value of MaxHeaderSize for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) SetPropertyMaxHeaderSize(value uint32) (err error) {
	return instance.SetProperty("MaxHeaderSize", (value))
}

// GetMaxHeaderSize gets the value of MaxHeaderSize for the instance
func (instance *MSNdis_WmiHDSplitCurrentConfig) GetPropertyMaxHeaderSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxHeaderSize")
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
