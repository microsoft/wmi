// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_Physical3dGraphicsProcessor struct
type Msvm_Physical3dGraphicsProcessor struct {
	*CIM_LogicalDevice

	//
	AdapterIndexID uint64

	//
	AvailableVideoMemory uint64

	//
	CompatibleForVirtualization bool

	//
	DedicatedSystemMemory uint64

	//
	DedicatedVideoMemory uint64

	//
	DirectXVersion string

	//
	DriverDate string

	//
	DriverInstalled string

	//
	DriverModelVersion string

	//
	DriverProvider string

	//
	DriverVersion string

	//
	EnabledForVirtualization bool

	//
	GPUID string

	//
	PixelShaderVersion string

	//
	Rating uint64

	//
	SharedSystemMemory uint64

	//
	TotalVideoMemory uint64
}

func NewMsvm_Physical3dGraphicsProcessorEx1(instance *cim.WmiInstance) (newInstance *Msvm_Physical3dGraphicsProcessor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Physical3dGraphicsProcessor{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_Physical3dGraphicsProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Physical3dGraphicsProcessor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Physical3dGraphicsProcessor{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAdapterIndexID sets the value of AdapterIndexID for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyAdapterIndexID(value uint64) (err error) {
	return instance.SetProperty("AdapterIndexID", (value))
}

// GetAdapterIndexID gets the value of AdapterIndexID for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyAdapterIndexID() (value uint64, err error) {
	retValue, err := instance.GetProperty("AdapterIndexID")
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

// SetAvailableVideoMemory sets the value of AvailableVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyAvailableVideoMemory(value uint64) (err error) {
	return instance.SetProperty("AvailableVideoMemory", (value))
}

// GetAvailableVideoMemory gets the value of AvailableVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyAvailableVideoMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableVideoMemory")
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

// SetCompatibleForVirtualization sets the value of CompatibleForVirtualization for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyCompatibleForVirtualization(value bool) (err error) {
	return instance.SetProperty("CompatibleForVirtualization", (value))
}

// GetCompatibleForVirtualization gets the value of CompatibleForVirtualization for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyCompatibleForVirtualization() (value bool, err error) {
	retValue, err := instance.GetProperty("CompatibleForVirtualization")
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

// SetDedicatedSystemMemory sets the value of DedicatedSystemMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDedicatedSystemMemory(value uint64) (err error) {
	return instance.SetProperty("DedicatedSystemMemory", (value))
}

// GetDedicatedSystemMemory gets the value of DedicatedSystemMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDedicatedSystemMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("DedicatedSystemMemory")
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

// SetDedicatedVideoMemory sets the value of DedicatedVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDedicatedVideoMemory(value uint64) (err error) {
	return instance.SetProperty("DedicatedVideoMemory", (value))
}

// GetDedicatedVideoMemory gets the value of DedicatedVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDedicatedVideoMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("DedicatedVideoMemory")
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

// SetDirectXVersion sets the value of DirectXVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDirectXVersion(value string) (err error) {
	return instance.SetProperty("DirectXVersion", (value))
}

// GetDirectXVersion gets the value of DirectXVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDirectXVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DirectXVersion")
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

// SetDriverDate sets the value of DriverDate for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDriverDate(value string) (err error) {
	return instance.SetProperty("DriverDate", (value))
}

// GetDriverDate gets the value of DriverDate for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDriverDate() (value string, err error) {
	retValue, err := instance.GetProperty("DriverDate")
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

// SetDriverInstalled sets the value of DriverInstalled for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDriverInstalled(value string) (err error) {
	return instance.SetProperty("DriverInstalled", (value))
}

// GetDriverInstalled gets the value of DriverInstalled for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDriverInstalled() (value string, err error) {
	retValue, err := instance.GetProperty("DriverInstalled")
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

// SetDriverModelVersion sets the value of DriverModelVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDriverModelVersion(value string) (err error) {
	return instance.SetProperty("DriverModelVersion", (value))
}

// GetDriverModelVersion gets the value of DriverModelVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDriverModelVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverModelVersion")
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

// SetDriverProvider sets the value of DriverProvider for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDriverProvider(value string) (err error) {
	return instance.SetProperty("DriverProvider", (value))
}

// GetDriverProvider gets the value of DriverProvider for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDriverProvider() (value string, err error) {
	retValue, err := instance.GetProperty("DriverProvider")
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

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyDriverVersion(value string) (err error) {
	return instance.SetProperty("DriverVersion", (value))
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyDriverVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
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

// SetEnabledForVirtualization sets the value of EnabledForVirtualization for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyEnabledForVirtualization(value bool) (err error) {
	return instance.SetProperty("EnabledForVirtualization", (value))
}

// GetEnabledForVirtualization gets the value of EnabledForVirtualization for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyEnabledForVirtualization() (value bool, err error) {
	retValue, err := instance.GetProperty("EnabledForVirtualization")
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

// SetGPUID sets the value of GPUID for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyGPUID(value string) (err error) {
	return instance.SetProperty("GPUID", (value))
}

// GetGPUID gets the value of GPUID for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyGPUID() (value string, err error) {
	retValue, err := instance.GetProperty("GPUID")
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

// SetPixelShaderVersion sets the value of PixelShaderVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyPixelShaderVersion(value string) (err error) {
	return instance.SetProperty("PixelShaderVersion", (value))
}

// GetPixelShaderVersion gets the value of PixelShaderVersion for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyPixelShaderVersion() (value string, err error) {
	retValue, err := instance.GetProperty("PixelShaderVersion")
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

// SetRating sets the value of Rating for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyRating(value uint64) (err error) {
	return instance.SetProperty("Rating", (value))
}

// GetRating gets the value of Rating for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyRating() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rating")
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

// SetSharedSystemMemory sets the value of SharedSystemMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertySharedSystemMemory(value uint64) (err error) {
	return instance.SetProperty("SharedSystemMemory", (value))
}

// GetSharedSystemMemory gets the value of SharedSystemMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertySharedSystemMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("SharedSystemMemory")
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

// SetTotalVideoMemory sets the value of TotalVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) SetPropertyTotalVideoMemory(value uint64) (err error) {
	return instance.SetProperty("TotalVideoMemory", (value))
}

// GetTotalVideoMemory gets the value of TotalVideoMemory for the instance
func (instance *Msvm_Physical3dGraphicsProcessor) GetPropertyTotalVideoMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVideoMemory")
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
