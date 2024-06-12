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

// MSTapeDriveParam struct
type MSTapeDriveParam struct {
	*MSTapeDriver

	//
	Active bool

	//
	CompressionCapable bool

	//
	CompressionEnabled bool

	//
	DefaultBlockSize uint32

	//
	HardwareErrorCorrection bool

	//
	InstanceName string

	//
	MaximumBlockSize uint32

	//
	MaximumPartitionCount uint32

	//
	MinimumBlockSize uint32

	//
	ReportSetmarks bool
}

func NewMSTapeDriveParamEx1(instance *cim.WmiInstance) (newInstance *MSTapeDriveParam, err error) {
	tmp, err := NewMSTapeDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSTapeDriveParam{
		MSTapeDriver: tmp,
	}
	return
}

func NewMSTapeDriveParamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSTapeDriveParam, err error) {
	tmp, err := NewMSTapeDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSTapeDriveParam{
		MSTapeDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSTapeDriveParam) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeDriveParam) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCompressionCapable sets the value of CompressionCapable for the instance
func (instance *MSTapeDriveParam) SetPropertyCompressionCapable(value bool) (err error) {
	return instance.SetProperty("CompressionCapable", (value))
}

// GetCompressionCapable gets the value of CompressionCapable for the instance
func (instance *MSTapeDriveParam) GetPropertyCompressionCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("CompressionCapable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCompressionEnabled sets the value of CompressionEnabled for the instance
func (instance *MSTapeDriveParam) SetPropertyCompressionEnabled(value bool) (err error) {
	return instance.SetProperty("CompressionEnabled", (value))
}

// GetCompressionEnabled gets the value of CompressionEnabled for the instance
func (instance *MSTapeDriveParam) GetPropertyCompressionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CompressionEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDefaultBlockSize sets the value of DefaultBlockSize for the instance
func (instance *MSTapeDriveParam) SetPropertyDefaultBlockSize(value uint32) (err error) {
	return instance.SetProperty("DefaultBlockSize", (value))
}

// GetDefaultBlockSize gets the value of DefaultBlockSize for the instance
func (instance *MSTapeDriveParam) GetPropertyDefaultBlockSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultBlockSize")
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

// SetHardwareErrorCorrection sets the value of HardwareErrorCorrection for the instance
func (instance *MSTapeDriveParam) SetPropertyHardwareErrorCorrection(value bool) (err error) {
	return instance.SetProperty("HardwareErrorCorrection", (value))
}

// GetHardwareErrorCorrection gets the value of HardwareErrorCorrection for the instance
func (instance *MSTapeDriveParam) GetPropertyHardwareErrorCorrection() (value bool, err error) {
	retValue, err := instance.GetProperty("HardwareErrorCorrection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSTapeDriveParam) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeDriveParam) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMaximumBlockSize sets the value of MaximumBlockSize for the instance
func (instance *MSTapeDriveParam) SetPropertyMaximumBlockSize(value uint32) (err error) {
	return instance.SetProperty("MaximumBlockSize", (value))
}

// GetMaximumBlockSize gets the value of MaximumBlockSize for the instance
func (instance *MSTapeDriveParam) GetPropertyMaximumBlockSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumBlockSize")
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

// SetMaximumPartitionCount sets the value of MaximumPartitionCount for the instance
func (instance *MSTapeDriveParam) SetPropertyMaximumPartitionCount(value uint32) (err error) {
	return instance.SetProperty("MaximumPartitionCount", (value))
}

// GetMaximumPartitionCount gets the value of MaximumPartitionCount for the instance
func (instance *MSTapeDriveParam) GetPropertyMaximumPartitionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumPartitionCount")
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

// SetMinimumBlockSize sets the value of MinimumBlockSize for the instance
func (instance *MSTapeDriveParam) SetPropertyMinimumBlockSize(value uint32) (err error) {
	return instance.SetProperty("MinimumBlockSize", (value))
}

// GetMinimumBlockSize gets the value of MinimumBlockSize for the instance
func (instance *MSTapeDriveParam) GetPropertyMinimumBlockSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumBlockSize")
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

// SetReportSetmarks sets the value of ReportSetmarks for the instance
func (instance *MSTapeDriveParam) SetPropertyReportSetmarks(value bool) (err error) {
	return instance.SetProperty("ReportSetmarks", (value))
}

// GetReportSetmarks gets the value of ReportSetmarks for the instance
func (instance *MSTapeDriveParam) GetPropertyReportSetmarks() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportSetmarks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
