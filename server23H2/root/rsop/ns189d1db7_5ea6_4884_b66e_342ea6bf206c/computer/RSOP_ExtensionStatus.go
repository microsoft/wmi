// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS189D1DB7_5EA6_4884_B66E_342EA6BF206C.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_ExtensionStatus struct
type RSOP_ExtensionStatus struct {
	*cim.WmiInstance

	//
	beginTime string

	//
	displayName string

	//
	endTime string

	//
	error uint32

	//
	extensionGuid string

	//
	loggingStatus uint32
}

func NewRSOP_ExtensionStatusEx1(instance *cim.WmiInstance) (newInstance *RSOP_ExtensionStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionStatus{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_ExtensionStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ExtensionStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionStatus{
		WmiInstance: tmp,
	}
	return
}

// SetbeginTime sets the value of beginTime for the instance
func (instance *RSOP_ExtensionStatus) SetPropertybeginTime(value string) (err error) {
	return instance.SetProperty("beginTime", (value))
}

// GetbeginTime gets the value of beginTime for the instance
func (instance *RSOP_ExtensionStatus) GetPropertybeginTime() (value string, err error) {
	retValue, err := instance.GetProperty("beginTime")
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

// SetdisplayName sets the value of displayName for the instance
func (instance *RSOP_ExtensionStatus) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", (value))
}

// GetdisplayName gets the value of displayName for the instance
func (instance *RSOP_ExtensionStatus) GetPropertydisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("displayName")
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

// SetendTime sets the value of endTime for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyendTime(value string) (err error) {
	return instance.SetProperty("endTime", (value))
}

// GetendTime gets the value of endTime for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyendTime() (value string, err error) {
	retValue, err := instance.GetProperty("endTime")
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

// Seterror sets the value of error for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyerror(value uint32) (err error) {
	return instance.SetProperty("error", (value))
}

// Geterror gets the value of error for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyerror() (value uint32, err error) {
	retValue, err := instance.GetProperty("error")
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

// SetextensionGuid sets the value of extensionGuid for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyextensionGuid(value string) (err error) {
	return instance.SetProperty("extensionGuid", (value))
}

// GetextensionGuid gets the value of extensionGuid for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyextensionGuid() (value string, err error) {
	retValue, err := instance.GetProperty("extensionGuid")
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

// SetloggingStatus sets the value of loggingStatus for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyloggingStatus(value uint32) (err error) {
	return instance.SetProperty("loggingStatus", (value))
}

// GetloggingStatus gets the value of loggingStatus for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyloggingStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("loggingStatus")
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
