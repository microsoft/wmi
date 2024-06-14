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

// FastCgiApplicationElement struct
type FastCgiApplicationElement struct {
	*CollectionElement

	//
	ActivityTimeout uint32

	//
	Arguments string

	//
	EnvironmentVariables FastCgiEnvironmentSettings

	//
	FlushNamedPipe bool

	//
	FullPath string

	//
	IdleTimeout uint32

	//
	InstanceMaxRequests uint32

	//
	MaxInstances uint32

	//
	MonitorChangesTo string

	//
	Protocol int32

	//
	QueueLength uint32

	//
	RapidFailsPerMinute uint32

	//
	RequestTimeout uint32

	//
	SignalBeforeTerminateSeconds uint32

	//
	StderrMode int32
}

func NewFastCgiApplicationElementEx1(instance *cim.WmiInstance) (newInstance *FastCgiApplicationElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FastCgiApplicationElement{
		CollectionElement: tmp,
	}
	return
}

func NewFastCgiApplicationElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FastCgiApplicationElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FastCgiApplicationElement{
		CollectionElement: tmp,
	}
	return
}

// SetActivityTimeout sets the value of ActivityTimeout for the instance
func (instance *FastCgiApplicationElement) SetPropertyActivityTimeout(value uint32) (err error) {
	return instance.SetProperty("ActivityTimeout", (value))
}

// GetActivityTimeout gets the value of ActivityTimeout for the instance
func (instance *FastCgiApplicationElement) GetPropertyActivityTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActivityTimeout")
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

// SetArguments sets the value of Arguments for the instance
func (instance *FastCgiApplicationElement) SetPropertyArguments(value string) (err error) {
	return instance.SetProperty("Arguments", (value))
}

// GetArguments gets the value of Arguments for the instance
func (instance *FastCgiApplicationElement) GetPropertyArguments() (value string, err error) {
	retValue, err := instance.GetProperty("Arguments")
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

// SetEnvironmentVariables sets the value of EnvironmentVariables for the instance
func (instance *FastCgiApplicationElement) SetPropertyEnvironmentVariables(value FastCgiEnvironmentSettings) (err error) {
	return instance.SetProperty("EnvironmentVariables", (value))
}

// GetEnvironmentVariables gets the value of EnvironmentVariables for the instance
func (instance *FastCgiApplicationElement) GetPropertyEnvironmentVariables() (value FastCgiEnvironmentSettings, err error) {
	retValue, err := instance.GetProperty("EnvironmentVariables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FastCgiEnvironmentSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FastCgiEnvironmentSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FastCgiEnvironmentSettings(valuetmp)

	return
}

// SetFlushNamedPipe sets the value of FlushNamedPipe for the instance
func (instance *FastCgiApplicationElement) SetPropertyFlushNamedPipe(value bool) (err error) {
	return instance.SetProperty("FlushNamedPipe", (value))
}

// GetFlushNamedPipe gets the value of FlushNamedPipe for the instance
func (instance *FastCgiApplicationElement) GetPropertyFlushNamedPipe() (value bool, err error) {
	retValue, err := instance.GetProperty("FlushNamedPipe")
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

// SetFullPath sets the value of FullPath for the instance
func (instance *FastCgiApplicationElement) SetPropertyFullPath(value string) (err error) {
	return instance.SetProperty("FullPath", (value))
}

// GetFullPath gets the value of FullPath for the instance
func (instance *FastCgiApplicationElement) GetPropertyFullPath() (value string, err error) {
	retValue, err := instance.GetProperty("FullPath")
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

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *FastCgiApplicationElement) SetPropertyIdleTimeout(value uint32) (err error) {
	return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *FastCgiApplicationElement) GetPropertyIdleTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleTimeout")
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

// SetInstanceMaxRequests sets the value of InstanceMaxRequests for the instance
func (instance *FastCgiApplicationElement) SetPropertyInstanceMaxRequests(value uint32) (err error) {
	return instance.SetProperty("InstanceMaxRequests", (value))
}

// GetInstanceMaxRequests gets the value of InstanceMaxRequests for the instance
func (instance *FastCgiApplicationElement) GetPropertyInstanceMaxRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("InstanceMaxRequests")
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

// SetMaxInstances sets the value of MaxInstances for the instance
func (instance *FastCgiApplicationElement) SetPropertyMaxInstances(value uint32) (err error) {
	return instance.SetProperty("MaxInstances", (value))
}

// GetMaxInstances gets the value of MaxInstances for the instance
func (instance *FastCgiApplicationElement) GetPropertyMaxInstances() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInstances")
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

// SetMonitorChangesTo sets the value of MonitorChangesTo for the instance
func (instance *FastCgiApplicationElement) SetPropertyMonitorChangesTo(value string) (err error) {
	return instance.SetProperty("MonitorChangesTo", (value))
}

// GetMonitorChangesTo gets the value of MonitorChangesTo for the instance
func (instance *FastCgiApplicationElement) GetPropertyMonitorChangesTo() (value string, err error) {
	retValue, err := instance.GetProperty("MonitorChangesTo")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *FastCgiApplicationElement) SetPropertyProtocol(value int32) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *FastCgiApplicationElement) GetPropertyProtocol() (value int32, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetQueueLength sets the value of QueueLength for the instance
func (instance *FastCgiApplicationElement) SetPropertyQueueLength(value uint32) (err error) {
	return instance.SetProperty("QueueLength", (value))
}

// GetQueueLength gets the value of QueueLength for the instance
func (instance *FastCgiApplicationElement) GetPropertyQueueLength() (value uint32, err error) {
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

// SetRapidFailsPerMinute sets the value of RapidFailsPerMinute for the instance
func (instance *FastCgiApplicationElement) SetPropertyRapidFailsPerMinute(value uint32) (err error) {
	return instance.SetProperty("RapidFailsPerMinute", (value))
}

// GetRapidFailsPerMinute gets the value of RapidFailsPerMinute for the instance
func (instance *FastCgiApplicationElement) GetPropertyRapidFailsPerMinute() (value uint32, err error) {
	retValue, err := instance.GetProperty("RapidFailsPerMinute")
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

// SetRequestTimeout sets the value of RequestTimeout for the instance
func (instance *FastCgiApplicationElement) SetPropertyRequestTimeout(value uint32) (err error) {
	return instance.SetProperty("RequestTimeout", (value))
}

// GetRequestTimeout gets the value of RequestTimeout for the instance
func (instance *FastCgiApplicationElement) GetPropertyRequestTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestTimeout")
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

// SetSignalBeforeTerminateSeconds sets the value of SignalBeforeTerminateSeconds for the instance
func (instance *FastCgiApplicationElement) SetPropertySignalBeforeTerminateSeconds(value uint32) (err error) {
	return instance.SetProperty("SignalBeforeTerminateSeconds", (value))
}

// GetSignalBeforeTerminateSeconds gets the value of SignalBeforeTerminateSeconds for the instance
func (instance *FastCgiApplicationElement) GetPropertySignalBeforeTerminateSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignalBeforeTerminateSeconds")
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

// SetStderrMode sets the value of StderrMode for the instance
func (instance *FastCgiApplicationElement) SetPropertyStderrMode(value int32) (err error) {
	return instance.SetProperty("StderrMode", (value))
}

// GetStderrMode gets the value of StderrMode for the instance
func (instance *FastCgiApplicationElement) GetPropertyStderrMode() (value int32, err error) {
	retValue, err := instance.GetProperty("StderrMode")
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
