// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PS_VpnServerAddress struct
type PS_VpnServerAddress struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	ServerAddress string
}

func NewPS_VpnServerAddressEx1(instance *cim.WmiInstance) (newInstance *PS_VpnServerAddress, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnServerAddress{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnServerAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnServerAddress, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnServerAddress{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *PS_VpnServerAddress) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *PS_VpnServerAddress) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *PS_VpnServerAddress) SetPropertyServerAddress(value string) (err error) {
	return instance.SetProperty("ServerAddress", (value))
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *PS_VpnServerAddress) GetPropertyServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerAddress")
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

// <param name="FriendlyName" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="ServerAddress" type="string "></param>

// <param name="cmdletOutput" type="VpnServerAddress "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnServerAddress) New( /* IN */ ServerAddress string,
	/* IN */ FriendlyName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput VpnServerAddress) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("New", ServerAddress, FriendlyName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
