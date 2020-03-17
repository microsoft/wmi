// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_VpnConnection struct
type MSFT_VpnConnection struct {
	cim.WmiInstance

	//
	AllUserConnection bool

	//
	Name string

	//
	Profile string
}

// SetAllUserConnection sets the value of AllUserConnection for the instance
func (instance *MSFT_VpnConnection) SetPropertyAllUserConnection(value bool) (err error) {
	return instance.SetProperty("AllUserConnection", value)
}

// GetAllUserConnection gets the value of AllUserConnection for the instance
func (instance *MSFT_VpnConnection) GetPropertyAllUserConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllUserConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_VpnConnection) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_VpnConnection) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_VpnConnection) SetPropertyProfile(value string) (err error) {
	return instance.SetProperty("Profile", value)
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_VpnConnection) GetPropertyProfile() (value string, err error) {
	retValue, err := instance.GetProperty("Profile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
