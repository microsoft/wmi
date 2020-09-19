// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.SDNDiagnostics.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DiagnosticInfo struct
type DiagnosticInfo struct {
	*cim.WmiInstance

	//
	DeviceType uint32

	//
	IsSDNCtlrPrimaryNode bool

	//
	LogLevel int8

	//
	LogLocation string

	//
	LogSizeLimit uint32

	//
	LogTimeLimit uint32

	//
	Password string

	//
	UserName string
}

func NewDiagnosticInfoEx1(instance *cim.WmiInstance) (newInstance *DiagnosticInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DiagnosticInfo{
		WmiInstance: tmp,
	}
	return
}

func NewDiagnosticInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DiagnosticInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DiagnosticInfo{
		WmiInstance: tmp,
	}
	return
}

// SetDeviceType sets the value of DeviceType for the instance
func (instance *DiagnosticInfo) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *DiagnosticInfo) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

// SetIsSDNCtlrPrimaryNode sets the value of IsSDNCtlrPrimaryNode for the instance
func (instance *DiagnosticInfo) SetPropertyIsSDNCtlrPrimaryNode(value bool) (err error) {
	return instance.SetProperty("IsSDNCtlrPrimaryNode", (value))
}

// GetIsSDNCtlrPrimaryNode gets the value of IsSDNCtlrPrimaryNode for the instance
func (instance *DiagnosticInfo) GetPropertyIsSDNCtlrPrimaryNode() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSDNCtlrPrimaryNode")
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

// SetLogLevel sets the value of LogLevel for the instance
func (instance *DiagnosticInfo) SetPropertyLogLevel(value int8) (err error) {
	return instance.SetProperty("LogLevel", (value))
}

// GetLogLevel gets the value of LogLevel for the instance
func (instance *DiagnosticInfo) GetPropertyLogLevel() (value int8, err error) {
	retValue, err := instance.GetProperty("LogLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetLogLocation sets the value of LogLocation for the instance
func (instance *DiagnosticInfo) SetPropertyLogLocation(value string) (err error) {
	return instance.SetProperty("LogLocation", (value))
}

// GetLogLocation gets the value of LogLocation for the instance
func (instance *DiagnosticInfo) GetPropertyLogLocation() (value string, err error) {
	retValue, err := instance.GetProperty("LogLocation")
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

// SetLogSizeLimit sets the value of LogSizeLimit for the instance
func (instance *DiagnosticInfo) SetPropertyLogSizeLimit(value uint32) (err error) {
	return instance.SetProperty("LogSizeLimit", (value))
}

// GetLogSizeLimit gets the value of LogSizeLimit for the instance
func (instance *DiagnosticInfo) GetPropertyLogSizeLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogSizeLimit")
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

// SetLogTimeLimit sets the value of LogTimeLimit for the instance
func (instance *DiagnosticInfo) SetPropertyLogTimeLimit(value uint32) (err error) {
	return instance.SetProperty("LogTimeLimit", (value))
}

// GetLogTimeLimit gets the value of LogTimeLimit for the instance
func (instance *DiagnosticInfo) GetPropertyLogTimeLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogTimeLimit")
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
func (instance *DiagnosticInfo) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *DiagnosticInfo) GetPropertyPassword() (value string, err error) {
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

// SetUserName sets the value of UserName for the instance
func (instance *DiagnosticInfo) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *DiagnosticInfo) GetPropertyUserName() (value string, err error) {
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
