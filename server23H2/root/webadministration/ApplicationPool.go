// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ApplicationPool struct
type ApplicationPool struct {
	*Object

	//
	AutoStart bool

	//
	CLRConfigFile string

	//
	Cpu ApplicationPoolProcessorSettings

	//
	Enable32BitAppOnWin64 bool

	//
	EnableConfigurationOverride bool

	//
	Failure ApplicationPoolFailureSettings

	//
	ManagedPipelineMode int32

	//
	ManagedRuntimeLoader string

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

func NewApplicationPoolEx1(instance *cim.WmiInstance) (newInstance *ApplicationPool, err error) {
	tmp, err := NewObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ApplicationPool{
		Object: tmp,
	}
	return
}

func NewApplicationPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ApplicationPool, err error) {
	tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ApplicationPool{
		Object: tmp,
	}
	return
}

// SetAutoStart sets the value of AutoStart for the instance
func (instance *ApplicationPool) SetPropertyAutoStart(value bool) (err error) {
	return instance.SetProperty("AutoStart", (value))
}

// GetAutoStart gets the value of AutoStart for the instance
func (instance *ApplicationPool) GetPropertyAutoStart() (value bool, err error) {
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

// SetCLRConfigFile sets the value of CLRConfigFile for the instance
func (instance *ApplicationPool) SetPropertyCLRConfigFile(value string) (err error) {
	return instance.SetProperty("CLRConfigFile", (value))
}

// GetCLRConfigFile gets the value of CLRConfigFile for the instance
func (instance *ApplicationPool) GetPropertyCLRConfigFile() (value string, err error) {
	retValue, err := instance.GetProperty("CLRConfigFile")
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

// SetCpu sets the value of Cpu for the instance
func (instance *ApplicationPool) SetPropertyCpu(value ApplicationPoolProcessorSettings) (err error) {
	return instance.SetProperty("Cpu", (value))
}

// GetCpu gets the value of Cpu for the instance
func (instance *ApplicationPool) GetPropertyCpu() (value ApplicationPoolProcessorSettings, err error) {
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
func (instance *ApplicationPool) SetPropertyEnable32BitAppOnWin64(value bool) (err error) {
	return instance.SetProperty("Enable32BitAppOnWin64", (value))
}

// GetEnable32BitAppOnWin64 gets the value of Enable32BitAppOnWin64 for the instance
func (instance *ApplicationPool) GetPropertyEnable32BitAppOnWin64() (value bool, err error) {
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

// SetEnableConfigurationOverride sets the value of EnableConfigurationOverride for the instance
func (instance *ApplicationPool) SetPropertyEnableConfigurationOverride(value bool) (err error) {
	return instance.SetProperty("EnableConfigurationOverride", (value))
}

// GetEnableConfigurationOverride gets the value of EnableConfigurationOverride for the instance
func (instance *ApplicationPool) GetPropertyEnableConfigurationOverride() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableConfigurationOverride")
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
func (instance *ApplicationPool) SetPropertyFailure(value ApplicationPoolFailureSettings) (err error) {
	return instance.SetProperty("Failure", (value))
}

// GetFailure gets the value of Failure for the instance
func (instance *ApplicationPool) GetPropertyFailure() (value ApplicationPoolFailureSettings, err error) {
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
func (instance *ApplicationPool) SetPropertyManagedPipelineMode(value int32) (err error) {
	return instance.SetProperty("ManagedPipelineMode", (value))
}

// GetManagedPipelineMode gets the value of ManagedPipelineMode for the instance
func (instance *ApplicationPool) GetPropertyManagedPipelineMode() (value int32, err error) {
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

// SetManagedRuntimeLoader sets the value of ManagedRuntimeLoader for the instance
func (instance *ApplicationPool) SetPropertyManagedRuntimeLoader(value string) (err error) {
	return instance.SetProperty("ManagedRuntimeLoader", (value))
}

// GetManagedRuntimeLoader gets the value of ManagedRuntimeLoader for the instance
func (instance *ApplicationPool) GetPropertyManagedRuntimeLoader() (value string, err error) {
	retValue, err := instance.GetProperty("ManagedRuntimeLoader")
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

// SetManagedRuntimeVersion sets the value of ManagedRuntimeVersion for the instance
func (instance *ApplicationPool) SetPropertyManagedRuntimeVersion(value string) (err error) {
	return instance.SetProperty("ManagedRuntimeVersion", (value))
}

// GetManagedRuntimeVersion gets the value of ManagedRuntimeVersion for the instance
func (instance *ApplicationPool) GetPropertyManagedRuntimeVersion() (value string, err error) {
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
func (instance *ApplicationPool) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ApplicationPool) GetPropertyName() (value string, err error) {
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
func (instance *ApplicationPool) SetPropertyPassAnonymousToken(value bool) (err error) {
	return instance.SetProperty("PassAnonymousToken", (value))
}

// GetPassAnonymousToken gets the value of PassAnonymousToken for the instance
func (instance *ApplicationPool) GetPropertyPassAnonymousToken() (value bool, err error) {
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
func (instance *ApplicationPool) SetPropertyProcessModel(value ProcessModelSettings) (err error) {
	return instance.SetProperty("ProcessModel", (value))
}

// GetProcessModel gets the value of ProcessModel for the instance
func (instance *ApplicationPool) GetPropertyProcessModel() (value ProcessModelSettings, err error) {
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
func (instance *ApplicationPool) SetPropertyQueueLength(value uint32) (err error) {
	return instance.SetProperty("QueueLength", (value))
}

// GetQueueLength gets the value of QueueLength for the instance
func (instance *ApplicationPool) GetPropertyQueueLength() (value uint32, err error) {
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
func (instance *ApplicationPool) SetPropertyRecycling(value RecyclingSettings) (err error) {
	return instance.SetProperty("Recycling", (value))
}

// GetRecycling gets the value of Recycling for the instance
func (instance *ApplicationPool) GetPropertyRecycling() (value RecyclingSettings, err error) {
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
func (instance *ApplicationPool) SetPropertyStartMode(value int32) (err error) {
	return instance.SetProperty("StartMode", (value))
}

// GetStartMode gets the value of StartMode for the instance
func (instance *ApplicationPool) GetPropertyStartMode() (value int32, err error) {
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *ApplicationPool) GetState() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetState")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//
func (instance *ApplicationPool) Start() (err error) {
	_, err = instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	return

}

//
func (instance *ApplicationPool) Stop() (err error) {
	_, err = instance.InvokeMethodWithReturn("Stop")
	if err != nil {
		return
	}
	return

}

//
func (instance *ApplicationPool) Recycle() (err error) {
	_, err = instance.InvokeMethodWithReturn("Recycle")
	if err != nil {
		return
	}
	return

}

//

// <param name="ReturnValue" type="string "></param>
func (instance *ApplicationPool) GetApplicationPoolSid() (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetApplicationPoolSid")
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

//

// <param name="AutoStart" type="bool "></param>
// <param name="Name" type="string "></param>
func (instance *ApplicationPool) Create( /* IN */ Name string,
	/* OPTIONAL IN */ AutoStart bool) (err error) {
	_, err = instance.InvokeMethodWithReturn("Create", Name, AutoStart)
	if err != nil {
		return
	}
	return

}

//

// <param name="PropertyName" type="string "></param>
func (instance *ApplicationPool) RevertToParent( /* OPTIONAL IN */ PropertyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RevertToParent", PropertyName)
	if err != nil {
		return
	}
	return

}
