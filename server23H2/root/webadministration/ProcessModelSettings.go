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

// ProcessModelSettings struct
type ProcessModelSettings struct {
	*EmbeddedObject

	//
	IdentityType int32

	//
	IdleTimeout string

	//
	IdleTimeoutAction int32

	//
	LoadUserProfile bool

	//
	LogEventOnProcessModel int32

	//
	LogonType int32

	//
	ManualGroupMembership bool

	//
	MaxProcesses uint32

	//
	Password string

	//
	PingingEnabled bool

	//
	PingInterval string

	//
	PingResponseTime string

	//
	SetProfileEnvironment bool

	//
	ShutdownTimeLimit string

	//
	StartupTimeLimit string

	//
	UserName string
}

func NewProcessModelSettingsEx1(instance *cim.WmiInstance) (newInstance *ProcessModelSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessModelSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewProcessModelSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessModelSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessModelSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetIdentityType sets the value of IdentityType for the instance
func (instance *ProcessModelSettings) SetPropertyIdentityType(value int32) (err error) {
	return instance.SetProperty("IdentityType", (value))
}

// GetIdentityType gets the value of IdentityType for the instance
func (instance *ProcessModelSettings) GetPropertyIdentityType() (value int32, err error) {
	retValue, err := instance.GetProperty("IdentityType")
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

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *ProcessModelSettings) SetPropertyIdleTimeout(value string) (err error) {
	return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *ProcessModelSettings) GetPropertyIdleTimeout() (value string, err error) {
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

// SetIdleTimeoutAction sets the value of IdleTimeoutAction for the instance
func (instance *ProcessModelSettings) SetPropertyIdleTimeoutAction(value int32) (err error) {
	return instance.SetProperty("IdleTimeoutAction", (value))
}

// GetIdleTimeoutAction gets the value of IdleTimeoutAction for the instance
func (instance *ProcessModelSettings) GetPropertyIdleTimeoutAction() (value int32, err error) {
	retValue, err := instance.GetProperty("IdleTimeoutAction")
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

// SetLoadUserProfile sets the value of LoadUserProfile for the instance
func (instance *ProcessModelSettings) SetPropertyLoadUserProfile(value bool) (err error) {
	return instance.SetProperty("LoadUserProfile", (value))
}

// GetLoadUserProfile gets the value of LoadUserProfile for the instance
func (instance *ProcessModelSettings) GetPropertyLoadUserProfile() (value bool, err error) {
	retValue, err := instance.GetProperty("LoadUserProfile")
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

// SetLogEventOnProcessModel sets the value of LogEventOnProcessModel for the instance
func (instance *ProcessModelSettings) SetPropertyLogEventOnProcessModel(value int32) (err error) {
	return instance.SetProperty("LogEventOnProcessModel", (value))
}

// GetLogEventOnProcessModel gets the value of LogEventOnProcessModel for the instance
func (instance *ProcessModelSettings) GetPropertyLogEventOnProcessModel() (value int32, err error) {
	retValue, err := instance.GetProperty("LogEventOnProcessModel")
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

// SetLogonType sets the value of LogonType for the instance
func (instance *ProcessModelSettings) SetPropertyLogonType(value int32) (err error) {
	return instance.SetProperty("LogonType", (value))
}

// GetLogonType gets the value of LogonType for the instance
func (instance *ProcessModelSettings) GetPropertyLogonType() (value int32, err error) {
	retValue, err := instance.GetProperty("LogonType")
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

// SetManualGroupMembership sets the value of ManualGroupMembership for the instance
func (instance *ProcessModelSettings) SetPropertyManualGroupMembership(value bool) (err error) {
	return instance.SetProperty("ManualGroupMembership", (value))
}

// GetManualGroupMembership gets the value of ManualGroupMembership for the instance
func (instance *ProcessModelSettings) GetPropertyManualGroupMembership() (value bool, err error) {
	retValue, err := instance.GetProperty("ManualGroupMembership")
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

// SetMaxProcesses sets the value of MaxProcesses for the instance
func (instance *ProcessModelSettings) SetPropertyMaxProcesses(value uint32) (err error) {
	return instance.SetProperty("MaxProcesses", (value))
}

// GetMaxProcesses gets the value of MaxProcesses for the instance
func (instance *ProcessModelSettings) GetPropertyMaxProcesses() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxProcesses")
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

// SetPassword sets the value of Password for the instance
func (instance *ProcessModelSettings) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *ProcessModelSettings) GetPropertyPassword() (value string, err error) {
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

// SetPingingEnabled sets the value of PingingEnabled for the instance
func (instance *ProcessModelSettings) SetPropertyPingingEnabled(value bool) (err error) {
	return instance.SetProperty("PingingEnabled", (value))
}

// GetPingingEnabled gets the value of PingingEnabled for the instance
func (instance *ProcessModelSettings) GetPropertyPingingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PingingEnabled")
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

// SetPingInterval sets the value of PingInterval for the instance
func (instance *ProcessModelSettings) SetPropertyPingInterval(value string) (err error) {
	return instance.SetProperty("PingInterval", (value))
}

// GetPingInterval gets the value of PingInterval for the instance
func (instance *ProcessModelSettings) GetPropertyPingInterval() (value string, err error) {
	retValue, err := instance.GetProperty("PingInterval")
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

// SetPingResponseTime sets the value of PingResponseTime for the instance
func (instance *ProcessModelSettings) SetPropertyPingResponseTime(value string) (err error) {
	return instance.SetProperty("PingResponseTime", (value))
}

// GetPingResponseTime gets the value of PingResponseTime for the instance
func (instance *ProcessModelSettings) GetPropertyPingResponseTime() (value string, err error) {
	retValue, err := instance.GetProperty("PingResponseTime")
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

// SetSetProfileEnvironment sets the value of SetProfileEnvironment for the instance
func (instance *ProcessModelSettings) SetPropertySetProfileEnvironment(value bool) (err error) {
	return instance.SetProperty("SetProfileEnvironment", (value))
}

// GetSetProfileEnvironment gets the value of SetProfileEnvironment for the instance
func (instance *ProcessModelSettings) GetPropertySetProfileEnvironment() (value bool, err error) {
	retValue, err := instance.GetProperty("SetProfileEnvironment")
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

// SetShutdownTimeLimit sets the value of ShutdownTimeLimit for the instance
func (instance *ProcessModelSettings) SetPropertyShutdownTimeLimit(value string) (err error) {
	return instance.SetProperty("ShutdownTimeLimit", (value))
}

// GetShutdownTimeLimit gets the value of ShutdownTimeLimit for the instance
func (instance *ProcessModelSettings) GetPropertyShutdownTimeLimit() (value string, err error) {
	retValue, err := instance.GetProperty("ShutdownTimeLimit")
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

// SetStartupTimeLimit sets the value of StartupTimeLimit for the instance
func (instance *ProcessModelSettings) SetPropertyStartupTimeLimit(value string) (err error) {
	return instance.SetProperty("StartupTimeLimit", (value))
}

// GetStartupTimeLimit gets the value of StartupTimeLimit for the instance
func (instance *ProcessModelSettings) GetPropertyStartupTimeLimit() (value string, err error) {
	retValue, err := instance.GetProperty("StartupTimeLimit")
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
func (instance *ProcessModelSettings) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *ProcessModelSettings) GetPropertyUserName() (value string, err error) {
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
