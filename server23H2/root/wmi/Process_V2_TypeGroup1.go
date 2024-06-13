// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Process_V2_TypeGroup1 struct
type Process_V2_TypeGroup1 struct {
	*Process_V2

	//
	CommandLine string

	//
	ExitStatus int32

	//
	ImageFileName string

	//
	ParentId uint32

	//
	ProcessId uint32

	//
	SessionId uint32

	//
	UniqueProcessKey uint32

	//
	UserSID interface{}
}

func NewProcess_V2_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *Process_V2_TypeGroup1, err error) {
	tmp, err := NewProcess_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup1{
		Process_V2: tmp,
	}
	return
}

func NewProcess_V2_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Process_V2_TypeGroup1, err error) {
	tmp, err := NewProcess_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup1{
		Process_V2: tmp,
	}
	return
}

// SetCommandLine sets the value of CommandLine for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", (value))
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLine")
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

// SetExitStatus sets the value of ExitStatus for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyExitStatus(value int32) (err error) {
	return instance.SetProperty("ExitStatus", (value))
}

// GetExitStatus gets the value of ExitStatus for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyExitStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("ExitStatus")
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

// SetImageFileName sets the value of ImageFileName for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyImageFileName(value string) (err error) {
	return instance.SetProperty("ImageFileName", (value))
}

// GetImageFileName gets the value of ImageFileName for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyImageFileName() (value string, err error) {
	retValue, err := instance.GetProperty("ImageFileName")
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

// SetParentId sets the value of ParentId for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyParentId(value uint32) (err error) {
	return instance.SetProperty("ParentId", (value))
}

// GetParentId gets the value of ParentId for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyParentId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParentId")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *Process_V2_TypeGroup1) SetPropertySessionId(value uint32) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *Process_V2_TypeGroup1) GetPropertySessionId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetUniqueProcessKey sets the value of UniqueProcessKey for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyUniqueProcessKey(value uint32) (err error) {
	return instance.SetProperty("UniqueProcessKey", (value))
}

// GetUniqueProcessKey gets the value of UniqueProcessKey for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyUniqueProcessKey() (value uint32, err error) {
	retValue, err := instance.GetProperty("UniqueProcessKey")
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

// SetUserSID sets the value of UserSID for the instance
func (instance *Process_V2_TypeGroup1) SetPropertyUserSID(value interface{}) (err error) {
	return instance.SetProperty("UserSID", (value))
}

// GetUserSID gets the value of UserSID for the instance
func (instance *Process_V2_TypeGroup1) GetPropertyUserSID() (value interface{}, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
