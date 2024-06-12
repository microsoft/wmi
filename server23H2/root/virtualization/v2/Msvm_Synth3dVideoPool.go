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

// Msvm_Synth3dVideoPool struct
type Msvm_Synth3dVideoPool struct {
	*CIM_ResourcePool

	//
	DirectXVersion string

	//
	Is3dVideoSupported bool

	//
	IsGPUCapable bool

	//
	IsSLATCapable bool

	//
	RequiredMinimumDirectXVersion string
}

func NewMsvm_Synth3dVideoPoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_Synth3dVideoPool, err error) {
	tmp, err := NewCIM_ResourcePoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synth3dVideoPool{
		CIM_ResourcePool: tmp,
	}
	return
}

func NewMsvm_Synth3dVideoPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Synth3dVideoPool, err error) {
	tmp, err := NewCIM_ResourcePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synth3dVideoPool{
		CIM_ResourcePool: tmp,
	}
	return
}

// SetDirectXVersion sets the value of DirectXVersion for the instance
func (instance *Msvm_Synth3dVideoPool) SetPropertyDirectXVersion(value string) (err error) {
	return instance.SetProperty("DirectXVersion", (value))
}

// GetDirectXVersion gets the value of DirectXVersion for the instance
func (instance *Msvm_Synth3dVideoPool) GetPropertyDirectXVersion() (value string, err error) {
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

// SetIs3dVideoSupported sets the value of Is3dVideoSupported for the instance
func (instance *Msvm_Synth3dVideoPool) SetPropertyIs3dVideoSupported(value bool) (err error) {
	return instance.SetProperty("Is3dVideoSupported", (value))
}

// GetIs3dVideoSupported gets the value of Is3dVideoSupported for the instance
func (instance *Msvm_Synth3dVideoPool) GetPropertyIs3dVideoSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("Is3dVideoSupported")
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

// SetIsGPUCapable sets the value of IsGPUCapable for the instance
func (instance *Msvm_Synth3dVideoPool) SetPropertyIsGPUCapable(value bool) (err error) {
	return instance.SetProperty("IsGPUCapable", (value))
}

// GetIsGPUCapable gets the value of IsGPUCapable for the instance
func (instance *Msvm_Synth3dVideoPool) GetPropertyIsGPUCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsGPUCapable")
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

// SetIsSLATCapable sets the value of IsSLATCapable for the instance
func (instance *Msvm_Synth3dVideoPool) SetPropertyIsSLATCapable(value bool) (err error) {
	return instance.SetProperty("IsSLATCapable", (value))
}

// GetIsSLATCapable gets the value of IsSLATCapable for the instance
func (instance *Msvm_Synth3dVideoPool) GetPropertyIsSLATCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSLATCapable")
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

// SetRequiredMinimumDirectXVersion sets the value of RequiredMinimumDirectXVersion for the instance
func (instance *Msvm_Synth3dVideoPool) SetPropertyRequiredMinimumDirectXVersion(value string) (err error) {
	return instance.SetProperty("RequiredMinimumDirectXVersion", (value))
}

// GetRequiredMinimumDirectXVersion gets the value of RequiredMinimumDirectXVersion for the instance
func (instance *Msvm_Synth3dVideoPool) GetPropertyRequiredMinimumDirectXVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RequiredMinimumDirectXVersion")
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

//

// <param name="monitorResolution" type="uint32 "></param>
// <param name="numberOfMonitors" type="uint32 "></param>

// <param name="requiredVideoMemory" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Synth3dVideoPool) CalculateVideoMemoryRequirements( /* IN */ monitorResolution uint32,
	/* IN */ numberOfMonitors uint32,
	/* OUT */ requiredVideoMemory uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CalculateVideoMemoryRequirements", monitorResolution, numberOfMonitors)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_Synth3dVideoPool) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_Synth3dVideoPool) GetRelatedResourcePoolSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolSettingData")
}

func (instance *Msvm_Synth3dVideoPool) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}

func (instance *Msvm_Synth3dVideoPool) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
