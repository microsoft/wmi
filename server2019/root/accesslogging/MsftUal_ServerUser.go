// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftUal_ServerUser struct
type MsftUal_ServerUser struct {
	*cim.WmiInstance

	// The incremental counter of client user accesses for a particular client user.
	ActivityCount uint32

	// Unit identification for the local server.
	ChassisSerialNumber string

	// The date and time when a client user name is first seen by a server.
	FirstSeen string

	// The date and time when a client user name is last seen by a server.
	LastSeen string

	// The client user name that accompanies the UAL payload from installed roles and products, if applicable.
	UserName string

	// SMBIOS reported universally unique identifier for this server unit.
	UUID string
}

func NewMsftUal_ServerUserEx1(instance *cim.WmiInstance) (newInstance *MsftUal_ServerUser, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_ServerUser{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_ServerUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_ServerUser, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_ServerUser{
		WmiInstance: tmp,
	}
	return
}

// SetActivityCount sets the value of ActivityCount for the instance
func (instance *MsftUal_ServerUser) SetPropertyActivityCount(value uint32) (err error) {
	return instance.SetProperty("ActivityCount", (value))
}

// GetActivityCount gets the value of ActivityCount for the instance
func (instance *MsftUal_ServerUser) GetPropertyActivityCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActivityCount")
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

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *MsftUal_ServerUser) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", (value))
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *MsftUal_ServerUser) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
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

// SetFirstSeen sets the value of FirstSeen for the instance
func (instance *MsftUal_ServerUser) SetPropertyFirstSeen(value string) (err error) {
	return instance.SetProperty("FirstSeen", (value))
}

// GetFirstSeen gets the value of FirstSeen for the instance
func (instance *MsftUal_ServerUser) GetPropertyFirstSeen() (value string, err error) {
	retValue, err := instance.GetProperty("FirstSeen")
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

// SetLastSeen sets the value of LastSeen for the instance
func (instance *MsftUal_ServerUser) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", (value))
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_ServerUser) GetPropertyLastSeen() (value string, err error) {
	retValue, err := instance.GetProperty("LastSeen")
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
func (instance *MsftUal_ServerUser) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *MsftUal_ServerUser) GetPropertyUserName() (value string, err error) {
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

// SetUUID sets the value of UUID for the instance
func (instance *MsftUal_ServerUser) SetPropertyUUID(value string) (err error) {
	return instance.SetProperty("UUID", (value))
}

// GetUUID gets the value of UUID for the instance
func (instance *MsftUal_ServerUser) GetPropertyUUID() (value string, err error) {
	retValue, err := instance.GetProperty("UUID")
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
