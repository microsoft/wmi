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

// ProcessModelSection struct
type ProcessModelSection struct {
	*ConfigurationSection

	//
	AutoConfig bool

	//
	ClientConnectedCheck string

	//
	ComAuthenticationLevel int32

	//
	ComImpersonationLevel int32

	//
	CpuMask int32

	//
	Enable bool

	//
	IdleTimeout string

	//
	LogLevel int32

	//
	MaxAppDomains int32

	//
	MaxIOThreads int32

	//
	MaxWorkerThreads int32

	//
	MemoryLimit int32

	//
	MinIOThreads int32

	//
	MinWorkerThreads int32

	//
	Password string

	//
	PingFrequency string

	//
	PingTimeout string

	//
	RequestLimit int32

	//
	RequestQueueLimit int32

	//
	ResponseDeadlockInterval string

	//
	ResponseRestartDeadlockInterval string

	//
	RestartQueueLimit int32

	//
	ServerErrorMessageFile string

	//
	ShutdownTimeout string

	//
	Timeout string

	//
	UserName string

	//
	WebGarden bool
}

func NewProcessModelSectionEx1(instance *cim.WmiInstance) (newInstance *ProcessModelSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessModelSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewProcessModelSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessModelSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessModelSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetAutoConfig sets the value of AutoConfig for the instance
func (instance *ProcessModelSection) SetPropertyAutoConfig(value bool) (err error) {
	return instance.SetProperty("AutoConfig", (value))
}

// GetAutoConfig gets the value of AutoConfig for the instance
func (instance *ProcessModelSection) GetPropertyAutoConfig() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoConfig")
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

// SetClientConnectedCheck sets the value of ClientConnectedCheck for the instance
func (instance *ProcessModelSection) SetPropertyClientConnectedCheck(value string) (err error) {
	return instance.SetProperty("ClientConnectedCheck", (value))
}

// GetClientConnectedCheck gets the value of ClientConnectedCheck for the instance
func (instance *ProcessModelSection) GetPropertyClientConnectedCheck() (value string, err error) {
	retValue, err := instance.GetProperty("ClientConnectedCheck")
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

// SetComAuthenticationLevel sets the value of ComAuthenticationLevel for the instance
func (instance *ProcessModelSection) SetPropertyComAuthenticationLevel(value int32) (err error) {
	return instance.SetProperty("ComAuthenticationLevel", (value))
}

// GetComAuthenticationLevel gets the value of ComAuthenticationLevel for the instance
func (instance *ProcessModelSection) GetPropertyComAuthenticationLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("ComAuthenticationLevel")
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

// SetComImpersonationLevel sets the value of ComImpersonationLevel for the instance
func (instance *ProcessModelSection) SetPropertyComImpersonationLevel(value int32) (err error) {
	return instance.SetProperty("ComImpersonationLevel", (value))
}

// GetComImpersonationLevel gets the value of ComImpersonationLevel for the instance
func (instance *ProcessModelSection) GetPropertyComImpersonationLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("ComImpersonationLevel")
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

// SetCpuMask sets the value of CpuMask for the instance
func (instance *ProcessModelSection) SetPropertyCpuMask(value int32) (err error) {
	return instance.SetProperty("CpuMask", (value))
}

// GetCpuMask gets the value of CpuMask for the instance
func (instance *ProcessModelSection) GetPropertyCpuMask() (value int32, err error) {
	retValue, err := instance.GetProperty("CpuMask")
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

// SetEnable sets the value of Enable for the instance
func (instance *ProcessModelSection) SetPropertyEnable(value bool) (err error) {
	return instance.SetProperty("Enable", (value))
}

// GetEnable gets the value of Enable for the instance
func (instance *ProcessModelSection) GetPropertyEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("Enable")
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

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *ProcessModelSection) SetPropertyIdleTimeout(value string) (err error) {
	return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *ProcessModelSection) GetPropertyIdleTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("IdleTimeout")
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

// SetLogLevel sets the value of LogLevel for the instance
func (instance *ProcessModelSection) SetPropertyLogLevel(value int32) (err error) {
	return instance.SetProperty("LogLevel", (value))
}

// GetLogLevel gets the value of LogLevel for the instance
func (instance *ProcessModelSection) GetPropertyLogLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("LogLevel")
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

// SetMaxAppDomains sets the value of MaxAppDomains for the instance
func (instance *ProcessModelSection) SetPropertyMaxAppDomains(value int32) (err error) {
	return instance.SetProperty("MaxAppDomains", (value))
}

// GetMaxAppDomains gets the value of MaxAppDomains for the instance
func (instance *ProcessModelSection) GetPropertyMaxAppDomains() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxAppDomains")
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

// SetMaxIOThreads sets the value of MaxIOThreads for the instance
func (instance *ProcessModelSection) SetPropertyMaxIOThreads(value int32) (err error) {
	return instance.SetProperty("MaxIOThreads", (value))
}

// GetMaxIOThreads gets the value of MaxIOThreads for the instance
func (instance *ProcessModelSection) GetPropertyMaxIOThreads() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxIOThreads")
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

// SetMaxWorkerThreads sets the value of MaxWorkerThreads for the instance
func (instance *ProcessModelSection) SetPropertyMaxWorkerThreads(value int32) (err error) {
	return instance.SetProperty("MaxWorkerThreads", (value))
}

// GetMaxWorkerThreads gets the value of MaxWorkerThreads for the instance
func (instance *ProcessModelSection) GetPropertyMaxWorkerThreads() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxWorkerThreads")
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

// SetMemoryLimit sets the value of MemoryLimit for the instance
func (instance *ProcessModelSection) SetPropertyMemoryLimit(value int32) (err error) {
	return instance.SetProperty("MemoryLimit", (value))
}

// GetMemoryLimit gets the value of MemoryLimit for the instance
func (instance *ProcessModelSection) GetPropertyMemoryLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("MemoryLimit")
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

// SetMinIOThreads sets the value of MinIOThreads for the instance
func (instance *ProcessModelSection) SetPropertyMinIOThreads(value int32) (err error) {
	return instance.SetProperty("MinIOThreads", (value))
}

// GetMinIOThreads gets the value of MinIOThreads for the instance
func (instance *ProcessModelSection) GetPropertyMinIOThreads() (value int32, err error) {
	retValue, err := instance.GetProperty("MinIOThreads")
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

// SetMinWorkerThreads sets the value of MinWorkerThreads for the instance
func (instance *ProcessModelSection) SetPropertyMinWorkerThreads(value int32) (err error) {
	return instance.SetProperty("MinWorkerThreads", (value))
}

// GetMinWorkerThreads gets the value of MinWorkerThreads for the instance
func (instance *ProcessModelSection) GetPropertyMinWorkerThreads() (value int32, err error) {
	retValue, err := instance.GetProperty("MinWorkerThreads")
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

// SetPassword sets the value of Password for the instance
func (instance *ProcessModelSection) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *ProcessModelSection) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetPingFrequency sets the value of PingFrequency for the instance
func (instance *ProcessModelSection) SetPropertyPingFrequency(value string) (err error) {
	return instance.SetProperty("PingFrequency", (value))
}

// GetPingFrequency gets the value of PingFrequency for the instance
func (instance *ProcessModelSection) GetPropertyPingFrequency() (value string, err error) {
	retValue, err := instance.GetProperty("PingFrequency")
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

// SetPingTimeout sets the value of PingTimeout for the instance
func (instance *ProcessModelSection) SetPropertyPingTimeout(value string) (err error) {
	return instance.SetProperty("PingTimeout", (value))
}

// GetPingTimeout gets the value of PingTimeout for the instance
func (instance *ProcessModelSection) GetPropertyPingTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("PingTimeout")
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

// SetRequestLimit sets the value of RequestLimit for the instance
func (instance *ProcessModelSection) SetPropertyRequestLimit(value int32) (err error) {
	return instance.SetProperty("RequestLimit", (value))
}

// GetRequestLimit gets the value of RequestLimit for the instance
func (instance *ProcessModelSection) GetPropertyRequestLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("RequestLimit")
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

// SetRequestQueueLimit sets the value of RequestQueueLimit for the instance
func (instance *ProcessModelSection) SetPropertyRequestQueueLimit(value int32) (err error) {
	return instance.SetProperty("RequestQueueLimit", (value))
}

// GetRequestQueueLimit gets the value of RequestQueueLimit for the instance
func (instance *ProcessModelSection) GetPropertyRequestQueueLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("RequestQueueLimit")
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

// SetResponseDeadlockInterval sets the value of ResponseDeadlockInterval for the instance
func (instance *ProcessModelSection) SetPropertyResponseDeadlockInterval(value string) (err error) {
	return instance.SetProperty("ResponseDeadlockInterval", (value))
}

// GetResponseDeadlockInterval gets the value of ResponseDeadlockInterval for the instance
func (instance *ProcessModelSection) GetPropertyResponseDeadlockInterval() (value string, err error) {
	retValue, err := instance.GetProperty("ResponseDeadlockInterval")
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

// SetResponseRestartDeadlockInterval sets the value of ResponseRestartDeadlockInterval for the instance
func (instance *ProcessModelSection) SetPropertyResponseRestartDeadlockInterval(value string) (err error) {
	return instance.SetProperty("ResponseRestartDeadlockInterval", (value))
}

// GetResponseRestartDeadlockInterval gets the value of ResponseRestartDeadlockInterval for the instance
func (instance *ProcessModelSection) GetPropertyResponseRestartDeadlockInterval() (value string, err error) {
	retValue, err := instance.GetProperty("ResponseRestartDeadlockInterval")
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

// SetRestartQueueLimit sets the value of RestartQueueLimit for the instance
func (instance *ProcessModelSection) SetPropertyRestartQueueLimit(value int32) (err error) {
	return instance.SetProperty("RestartQueueLimit", (value))
}

// GetRestartQueueLimit gets the value of RestartQueueLimit for the instance
func (instance *ProcessModelSection) GetPropertyRestartQueueLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("RestartQueueLimit")
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

// SetServerErrorMessageFile sets the value of ServerErrorMessageFile for the instance
func (instance *ProcessModelSection) SetPropertyServerErrorMessageFile(value string) (err error) {
	return instance.SetProperty("ServerErrorMessageFile", (value))
}

// GetServerErrorMessageFile gets the value of ServerErrorMessageFile for the instance
func (instance *ProcessModelSection) GetPropertyServerErrorMessageFile() (value string, err error) {
	retValue, err := instance.GetProperty("ServerErrorMessageFile")
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

// SetShutdownTimeout sets the value of ShutdownTimeout for the instance
func (instance *ProcessModelSection) SetPropertyShutdownTimeout(value string) (err error) {
	return instance.SetProperty("ShutdownTimeout", (value))
}

// GetShutdownTimeout gets the value of ShutdownTimeout for the instance
func (instance *ProcessModelSection) GetPropertyShutdownTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ShutdownTimeout")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *ProcessModelSection) SetPropertyTimeout(value string) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *ProcessModelSection) GetPropertyTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("Timeout")
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

// SetUserName sets the value of UserName for the instance
func (instance *ProcessModelSection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *ProcessModelSection) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

// SetWebGarden sets the value of WebGarden for the instance
func (instance *ProcessModelSection) SetPropertyWebGarden(value bool) (err error) {
	return instance.SetProperty("WebGarden", (value))
}

// GetWebGarden gets the value of WebGarden for the instance
func (instance *ProcessModelSection) GetPropertyWebGarden() (value bool, err error) {
	retValue, err := instance.GetProperty("WebGarden")
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
