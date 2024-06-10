// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ApplicationPoolElementDefaults struct
type ApplicationPoolElementDefaults struct {
	*EmbeddedObject

	//
	AutoStart bool

	//
	Cpu ApplicationPoolProcessorSettings

	//
	Enable32BitAppOnWin64 bool

	//
	Failure ApplicationPoolFailureSettings

	//
	ManagedPipelineMode int32

	//
	ManagedRuntimeVersion string

	//
	Name string

	//
	PassAnonymousToken bool

	//
	ProcessModel ProcessModelSettings

	//
	QueueLength uint32

	//
	Recycling RecyclingSettings

	//
	StartMode int32
}

func NewApplicationPoolElementDefaultsEx1(instance *cim.WmiInstance) (newInstance *ApplicationPoolElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ApplicationPoolElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

func NewApplicationPoolElementDefaultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ApplicationPoolElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ApplicationPoolElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

// SetAutoStart sets the value of AutoStart for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyAutoStart(value bool) (err error) {
	return instance.SetProperty("AutoStart", (value))
}

// GetAutoStart gets the value of AutoStart for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyAutoStart() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoStart")
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

// SetCpu sets the value of Cpu for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyCpu(value ApplicationPoolProcessorSettings) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyCpu() (value ApplicationPoolProcessorSettings, err error) {
	retValue, err := instance.GetProperty("Cpu")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ApplicationPoolProcessorSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ApplicationPoolProcessorSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ApplicationPoolProcessorSettings(valuetmp)

	return
}

// SetEnable32BitAppOnWin64 sets the value of Enable32BitAppOnWin64 for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyEnable32BitAppOnWin64(value bool) (err error) {
	return instance.SetProperty("Enable32BitAppOnWin64", (value))
}

// GetEnable32BitAppOnWin64 gets the value of Enable32BitAppOnWin64 for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyEnable32BitAppOnWin64() (value bool, err error) {
	retValue, err := instance.GetProperty("Enable32BitAppOnWin64")
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

// SetFailure sets the value of Failure for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyFailure(value ApplicationPoolFailureSettings) (err error) {
	return instance.SetProperty("Failure", (value))
}

// GetFailure gets the value of Failure for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyFailure() (value ApplicationPoolFailureSettings, err error) {
	retValue, err := instance.GetProperty("Failure")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ApplicationPoolFailureSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ApplicationPoolFailureSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ApplicationPoolFailureSettings(valuetmp)

	return
}

// SetManagedPipelineMode sets the value of ManagedPipelineMode for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyManagedPipelineMode(value int32) (err error) {
	return instance.SetProperty("ManagedPipelineMode", (value))
}

// GetManagedPipelineMode gets the value of ManagedPipelineMode for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyManagedPipelineMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ManagedPipelineMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetManagedRuntimeVersion sets the value of ManagedRuntimeVersion for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyManagedRuntimeVersion(value string) (err error) {
	return instance.SetProperty("ManagedRuntimeVersion", (value))
}

// GetManagedRuntimeVersion gets the value of ManagedRuntimeVersion for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyManagedRuntimeVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ManagedRuntimeVersion")
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

// SetName sets the value of Name for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetPassAnonymousToken sets the value of PassAnonymousToken for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyPassAnonymousToken(value bool) (err error) {
	return instance.SetProperty("PassAnonymousToken", (value))
}

// GetPassAnonymousToken gets the value of PassAnonymousToken for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyPassAnonymousToken() (value bool, err error) {
	retValue, err := instance.GetProperty("PassAnonymousToken")
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

// SetProcessModel sets the value of ProcessModel for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyProcessModel(value ProcessModelSettings) (err error) {
	return instance.SetProperty("ProcessModel", (value))
}

// GetProcessModel gets the value of ProcessModel for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyProcessModel() (value ProcessModelSettings, err error) {
	retValue, err := instance.GetProperty("ProcessModel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProcessModelSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProcessModelSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProcessModelSettings(valuetmp)

	return
}

// SetQueueLength sets the value of QueueLength for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyQueueLength(value uint32) (err error) {
	return instance.SetProperty("QueueLength", (value))
}

// GetQueueLength gets the value of QueueLength for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyQueueLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueLength")
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

// SetRecycling sets the value of Recycling for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyRecycling(value RecyclingSettings) (err error) {
	return instance.SetProperty("Recycling", (value))
}

// GetRecycling gets the value of Recycling for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyRecycling() (value RecyclingSettings, err error) {
	retValue, err := instance.GetProperty("Recycling")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RecyclingSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RecyclingSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RecyclingSettings(valuetmp)

	return
}

// SetStartMode sets the value of StartMode for the instance
func (instance *ApplicationPoolElementDefaults) SetPropertyStartMode(value int32) (err error) {
	return instance.SetProperty("StartMode", (value))
}

// GetStartMode gets the value of StartMode for the instance
func (instance *ApplicationPoolElementDefaults) GetPropertyStartMode() (value int32, err error) {
	retValue, err := instance.GetProperty("StartMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
