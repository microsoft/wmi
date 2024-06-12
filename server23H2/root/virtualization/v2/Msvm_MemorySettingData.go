// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_MemorySettingData struct
type Msvm_MemorySettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	DynamicMemoryEnabled bool

	//
	HugePagesEnabled bool

	//
	IsVirtualized bool

	//
	MaxMemoryBlocksPerNumaNode uint64

	//
	MemoryAccessTrackingPolicy uint8

	//
	MemoryEncryptionPolicy uint8

	//
	SgxEnabled bool

	//
	SgxLaunchControlDefault string

	//
	SgxLaunchControlMode uint32

	//
	SgxSize uint64

	//
	SwapFilesInUse bool

	//
	TargetMemoryBuffer uint32
}

func NewMsvm_MemorySettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_MemorySettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MemorySettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_MemorySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MemorySettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MemorySettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDynamicMemoryEnabled sets the value of DynamicMemoryEnabled for the instance
func (instance *Msvm_MemorySettingData) SetPropertyDynamicMemoryEnabled(value bool) (err error) {
	return instance.SetProperty("DynamicMemoryEnabled", (value))
}

// GetDynamicMemoryEnabled gets the value of DynamicMemoryEnabled for the instance
func (instance *Msvm_MemorySettingData) GetPropertyDynamicMemoryEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicMemoryEnabled")
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

// SetHugePagesEnabled sets the value of HugePagesEnabled for the instance
func (instance *Msvm_MemorySettingData) SetPropertyHugePagesEnabled(value bool) (err error) {
	return instance.SetProperty("HugePagesEnabled", (value))
}

// GetHugePagesEnabled gets the value of HugePagesEnabled for the instance
func (instance *Msvm_MemorySettingData) GetPropertyHugePagesEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HugePagesEnabled")
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

// SetIsVirtualized sets the value of IsVirtualized for the instance
func (instance *Msvm_MemorySettingData) SetPropertyIsVirtualized(value bool) (err error) {
	return instance.SetProperty("IsVirtualized", (value))
}

// GetIsVirtualized gets the value of IsVirtualized for the instance
func (instance *Msvm_MemorySettingData) GetPropertyIsVirtualized() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVirtualized")
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

// SetMaxMemoryBlocksPerNumaNode sets the value of MaxMemoryBlocksPerNumaNode for the instance
func (instance *Msvm_MemorySettingData) SetPropertyMaxMemoryBlocksPerNumaNode(value uint64) (err error) {
	return instance.SetProperty("MaxMemoryBlocksPerNumaNode", (value))
}

// GetMaxMemoryBlocksPerNumaNode gets the value of MaxMemoryBlocksPerNumaNode for the instance
func (instance *Msvm_MemorySettingData) GetPropertyMaxMemoryBlocksPerNumaNode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxMemoryBlocksPerNumaNode")
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

// SetMemoryAccessTrackingPolicy sets the value of MemoryAccessTrackingPolicy for the instance
func (instance *Msvm_MemorySettingData) SetPropertyMemoryAccessTrackingPolicy(value uint8) (err error) {
	return instance.SetProperty("MemoryAccessTrackingPolicy", (value))
}

// GetMemoryAccessTrackingPolicy gets the value of MemoryAccessTrackingPolicy for the instance
func (instance *Msvm_MemorySettingData) GetPropertyMemoryAccessTrackingPolicy() (value uint8, err error) {
	retValue, err := instance.GetProperty("MemoryAccessTrackingPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMemoryEncryptionPolicy sets the value of MemoryEncryptionPolicy for the instance
func (instance *Msvm_MemorySettingData) SetPropertyMemoryEncryptionPolicy(value uint8) (err error) {
	return instance.SetProperty("MemoryEncryptionPolicy", (value))
}

// GetMemoryEncryptionPolicy gets the value of MemoryEncryptionPolicy for the instance
func (instance *Msvm_MemorySettingData) GetPropertyMemoryEncryptionPolicy() (value uint8, err error) {
	retValue, err := instance.GetProperty("MemoryEncryptionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSgxEnabled sets the value of SgxEnabled for the instance
func (instance *Msvm_MemorySettingData) SetPropertySgxEnabled(value bool) (err error) {
	return instance.SetProperty("SgxEnabled", (value))
}

// GetSgxEnabled gets the value of SgxEnabled for the instance
func (instance *Msvm_MemorySettingData) GetPropertySgxEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SgxEnabled")
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

// SetSgxLaunchControlDefault sets the value of SgxLaunchControlDefault for the instance
func (instance *Msvm_MemorySettingData) SetPropertySgxLaunchControlDefault(value string) (err error) {
	return instance.SetProperty("SgxLaunchControlDefault", (value))
}

// GetSgxLaunchControlDefault gets the value of SgxLaunchControlDefault for the instance
func (instance *Msvm_MemorySettingData) GetPropertySgxLaunchControlDefault() (value string, err error) {
	retValue, err := instance.GetProperty("SgxLaunchControlDefault")
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

// SetSgxLaunchControlMode sets the value of SgxLaunchControlMode for the instance
func (instance *Msvm_MemorySettingData) SetPropertySgxLaunchControlMode(value uint32) (err error) {
	return instance.SetProperty("SgxLaunchControlMode", (value))
}

// GetSgxLaunchControlMode gets the value of SgxLaunchControlMode for the instance
func (instance *Msvm_MemorySettingData) GetPropertySgxLaunchControlMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("SgxLaunchControlMode")
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

// SetSgxSize sets the value of SgxSize for the instance
func (instance *Msvm_MemorySettingData) SetPropertySgxSize(value uint64) (err error) {
	return instance.SetProperty("SgxSize", (value))
}

// GetSgxSize gets the value of SgxSize for the instance
func (instance *Msvm_MemorySettingData) GetPropertySgxSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SgxSize")
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

// SetSwapFilesInUse sets the value of SwapFilesInUse for the instance
func (instance *Msvm_MemorySettingData) SetPropertySwapFilesInUse(value bool) (err error) {
	return instance.SetProperty("SwapFilesInUse", (value))
}

// GetSwapFilesInUse gets the value of SwapFilesInUse for the instance
func (instance *Msvm_MemorySettingData) GetPropertySwapFilesInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("SwapFilesInUse")
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

// SetTargetMemoryBuffer sets the value of TargetMemoryBuffer for the instance
func (instance *Msvm_MemorySettingData) SetPropertyTargetMemoryBuffer(value uint32) (err error) {
	return instance.SetProperty("TargetMemoryBuffer", (value))
}

// GetTargetMemoryBuffer gets the value of TargetMemoryBuffer for the instance
func (instance *Msvm_MemorySettingData) GetPropertyTargetMemoryBuffer() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetMemoryBuffer")
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
func (instance *Msvm_MemorySettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
