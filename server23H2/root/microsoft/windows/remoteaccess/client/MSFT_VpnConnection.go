// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Client
//
// ////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_VpnConnection struct
type MSFT_VpnConnection struct {
	*cim.WmiInstance

	//
	AllUserConnection bool

	//
	Name string

	//
	Profile string
}

func NewMSFT_VpnConnectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_VpnConnection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_VpnConnection{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_VpnConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_VpnConnection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_VpnConnection{
		WmiInstance: tmp,
	}
	return
}

// SetAllUserConnection sets the value of AllUserConnection for the instance
func (instance *MSFT_VpnConnection) SetPropertyAllUserConnection(value bool) (err error) {
	return instance.SetProperty("AllUserConnection", (value))
}

// GetAllUserConnection gets the value of AllUserConnection for the instance
func (instance *MSFT_VpnConnection) GetPropertyAllUserConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllUserConnection")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_VpnConnection) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_VpnConnection) GetPropertyName() (value string, err error) {
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

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_VpnConnection) SetPropertyProfile(value string) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_VpnConnection) GetPropertyProfile() (value string, err error) {
	retValue, err := instance.GetProperty("Profile")
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

//

// <param name="Profile" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VpnConnection) Set( /* IN */ Profile string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", Profile)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
