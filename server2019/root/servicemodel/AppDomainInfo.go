// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppDomainInfo struct
type AppDomainInfo struct {
	*cim.WmiInstance

	// The Id of the appdomain.
	AppDomainId int32

	// Indicates if the appdomain is the default appdomain.
	IsDefault bool

	// A value that specifies whether malformed messages are logged.
	LogMalformedMessages bool

	// A value that specifies whether messages are traced at the service level (before encryption and transport-related transforms).
	LogMessagesAtServiceLevel bool

	// A value that specifies whether messages are traced at the transport level.
	LogMessagesAtTransportLevel bool

	// The collection trace listeners that listen to the System.ServiceModel.MessageLogging trace source.
	MessageLoggingTraceListeners []TraceListener

	// The name of the appdomain.
	Name string

	// The scope of active performance counters in the appdomain.
	PerformanceCounters string

	// The process Id.
	ProcessId int32

	// The path to the configuration of the service.
	ServiceConfigPath string

	// The collection trace listeners that listen to the System.ServiceModel trace source.
	ServiceModelTraceListeners []TraceListener

	// The trace level of the System.ServiceModel trace source.
	TraceLevel string
}

func NewAppDomainInfoEx1(instance *cim.WmiInstance) (newInstance *AppDomainInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppDomainInfo{
		WmiInstance: tmp,
	}
	return
}

func NewAppDomainInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppDomainInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppDomainInfo{
		WmiInstance: tmp,
	}
	return
}

// SetAppDomainId sets the value of AppDomainId for the instance
func (instance *AppDomainInfo) SetPropertyAppDomainId(value int32) (err error) {
	return instance.SetProperty("AppDomainId", (value))
}

// GetAppDomainId gets the value of AppDomainId for the instance
func (instance *AppDomainInfo) GetPropertyAppDomainId() (value int32, err error) {
	retValue, err := instance.GetProperty("AppDomainId")
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

// SetIsDefault sets the value of IsDefault for the instance
func (instance *AppDomainInfo) SetPropertyIsDefault(value bool) (err error) {
	return instance.SetProperty("IsDefault", (value))
}

// GetIsDefault gets the value of IsDefault for the instance
func (instance *AppDomainInfo) GetPropertyIsDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDefault")
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

// SetLogMalformedMessages sets the value of LogMalformedMessages for the instance
func (instance *AppDomainInfo) SetPropertyLogMalformedMessages(value bool) (err error) {
	return instance.SetProperty("LogMalformedMessages", (value))
}

// GetLogMalformedMessages gets the value of LogMalformedMessages for the instance
func (instance *AppDomainInfo) GetPropertyLogMalformedMessages() (value bool, err error) {
	retValue, err := instance.GetProperty("LogMalformedMessages")
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

// SetLogMessagesAtServiceLevel sets the value of LogMessagesAtServiceLevel for the instance
func (instance *AppDomainInfo) SetPropertyLogMessagesAtServiceLevel(value bool) (err error) {
	return instance.SetProperty("LogMessagesAtServiceLevel", (value))
}

// GetLogMessagesAtServiceLevel gets the value of LogMessagesAtServiceLevel for the instance
func (instance *AppDomainInfo) GetPropertyLogMessagesAtServiceLevel() (value bool, err error) {
	retValue, err := instance.GetProperty("LogMessagesAtServiceLevel")
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

// SetLogMessagesAtTransportLevel sets the value of LogMessagesAtTransportLevel for the instance
func (instance *AppDomainInfo) SetPropertyLogMessagesAtTransportLevel(value bool) (err error) {
	return instance.SetProperty("LogMessagesAtTransportLevel", (value))
}

// GetLogMessagesAtTransportLevel gets the value of LogMessagesAtTransportLevel for the instance
func (instance *AppDomainInfo) GetPropertyLogMessagesAtTransportLevel() (value bool, err error) {
	retValue, err := instance.GetProperty("LogMessagesAtTransportLevel")
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

// SetMessageLoggingTraceListeners sets the value of MessageLoggingTraceListeners for the instance
func (instance *AppDomainInfo) SetPropertyMessageLoggingTraceListeners(value []TraceListener) (err error) {
	return instance.SetProperty("MessageLoggingTraceListeners", (value))
}

// GetMessageLoggingTraceListeners gets the value of MessageLoggingTraceListeners for the instance
func (instance *AppDomainInfo) GetPropertyMessageLoggingTraceListeners() (value []TraceListener, err error) {
	retValue, err := instance.GetProperty("MessageLoggingTraceListeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceListener)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceListener is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceListener(valuetmp))
	}

	return
}

// SetName sets the value of Name for the instance
func (instance *AppDomainInfo) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *AppDomainInfo) GetPropertyName() (value string, err error) {
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

// SetPerformanceCounters sets the value of PerformanceCounters for the instance
func (instance *AppDomainInfo) SetPropertyPerformanceCounters(value string) (err error) {
	return instance.SetProperty("PerformanceCounters", (value))
}

// GetPerformanceCounters gets the value of PerformanceCounters for the instance
func (instance *AppDomainInfo) GetPropertyPerformanceCounters() (value string, err error) {
	retValue, err := instance.GetProperty("PerformanceCounters")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *AppDomainInfo) SetPropertyProcessId(value int32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *AppDomainInfo) GetPropertyProcessId() (value int32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetServiceConfigPath sets the value of ServiceConfigPath for the instance
func (instance *AppDomainInfo) SetPropertyServiceConfigPath(value string) (err error) {
	return instance.SetProperty("ServiceConfigPath", (value))
}

// GetServiceConfigPath gets the value of ServiceConfigPath for the instance
func (instance *AppDomainInfo) GetPropertyServiceConfigPath() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceConfigPath")
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

// SetServiceModelTraceListeners sets the value of ServiceModelTraceListeners for the instance
func (instance *AppDomainInfo) SetPropertyServiceModelTraceListeners(value []TraceListener) (err error) {
	return instance.SetProperty("ServiceModelTraceListeners", (value))
}

// GetServiceModelTraceListeners gets the value of ServiceModelTraceListeners for the instance
func (instance *AppDomainInfo) GetPropertyServiceModelTraceListeners() (value []TraceListener, err error) {
	retValue, err := instance.GetProperty("ServiceModelTraceListeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceListener)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceListener is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceListener(valuetmp))
	}

	return
}

// SetTraceLevel sets the value of TraceLevel for the instance
func (instance *AppDomainInfo) SetPropertyTraceLevel(value string) (err error) {
	return instance.SetProperty("TraceLevel", (value))
}

// GetTraceLevel gets the value of TraceLevel for the instance
func (instance *AppDomainInfo) GetPropertyTraceLevel() (value string, err error) {
	retValue, err := instance.GetProperty("TraceLevel")
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
