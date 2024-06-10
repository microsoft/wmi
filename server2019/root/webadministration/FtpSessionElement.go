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

// FtpSessionElement struct
type FtpSessionElement struct {
	*CollectionElement

	//
	BytesReceived string

	//
	BytesSent string

	//
	ClientIp string

	//
	CommandStartTime string

	//
	CurrentCommand string

	//
	LastErrorStatus uint32

	//
	PreviousCommand string

	//
	SessionId string

	//
	SessionStartTime string

	//
	UserName string
}

func NewFtpSessionElementEx1(instance *cim.WmiInstance) (newInstance *FtpSessionElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpSessionElement{
		CollectionElement: tmp,
	}
	return
}

func NewFtpSessionElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpSessionElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpSessionElement{
		CollectionElement: tmp,
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *FtpSessionElement) SetPropertyBytesReceived(value string) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *FtpSessionElement) GetPropertyBytesReceived() (value string, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
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

// SetBytesSent sets the value of BytesSent for the instance
func (instance *FtpSessionElement) SetPropertyBytesSent(value string) (err error) {
	return instance.SetProperty("BytesSent", (value))
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *FtpSessionElement) GetPropertyBytesSent() (value string, err error) {
	retValue, err := instance.GetProperty("BytesSent")
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

// SetClientIp sets the value of ClientIp for the instance
func (instance *FtpSessionElement) SetPropertyClientIp(value string) (err error) {
	return instance.SetProperty("ClientIp", (value))
}

// GetClientIp gets the value of ClientIp for the instance
func (instance *FtpSessionElement) GetPropertyClientIp() (value string, err error) {
	retValue, err := instance.GetProperty("ClientIp")
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

// SetCommandStartTime sets the value of CommandStartTime for the instance
func (instance *FtpSessionElement) SetPropertyCommandStartTime(value string) (err error) {
	return instance.SetProperty("CommandStartTime", (value))
}

// GetCommandStartTime gets the value of CommandStartTime for the instance
func (instance *FtpSessionElement) GetPropertyCommandStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("CommandStartTime")
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

// SetCurrentCommand sets the value of CurrentCommand for the instance
func (instance *FtpSessionElement) SetPropertyCurrentCommand(value string) (err error) {
	return instance.SetProperty("CurrentCommand", (value))
}

// GetCurrentCommand gets the value of CurrentCommand for the instance
func (instance *FtpSessionElement) GetPropertyCurrentCommand() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentCommand")
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

// SetLastErrorStatus sets the value of LastErrorStatus for the instance
func (instance *FtpSessionElement) SetPropertyLastErrorStatus(value uint32) (err error) {
	return instance.SetProperty("LastErrorStatus", (value))
}

// GetLastErrorStatus gets the value of LastErrorStatus for the instance
func (instance *FtpSessionElement) GetPropertyLastErrorStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastErrorStatus")
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

// SetPreviousCommand sets the value of PreviousCommand for the instance
func (instance *FtpSessionElement) SetPropertyPreviousCommand(value string) (err error) {
	return instance.SetProperty("PreviousCommand", (value))
}

// GetPreviousCommand gets the value of PreviousCommand for the instance
func (instance *FtpSessionElement) GetPropertyPreviousCommand() (value string, err error) {
	retValue, err := instance.GetProperty("PreviousCommand")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *FtpSessionElement) SetPropertySessionId(value string) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *FtpSessionElement) GetPropertySessionId() (value string, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetSessionStartTime sets the value of SessionStartTime for the instance
func (instance *FtpSessionElement) SetPropertySessionStartTime(value string) (err error) {
	return instance.SetProperty("SessionStartTime", (value))
}

// GetSessionStartTime gets the value of SessionStartTime for the instance
func (instance *FtpSessionElement) GetPropertySessionStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("SessionStartTime")
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
func (instance *FtpSessionElement) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *FtpSessionElement) GetPropertyUserName() (value string, err error) {
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
