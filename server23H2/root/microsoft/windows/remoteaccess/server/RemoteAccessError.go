// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessError struct
type RemoteAccessError struct {
	*MSFT_WmiError

	//
	RemoteAccessErrorCode uint32
}

func NewRemoteAccessErrorEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessError, err error) {
	tmp, err := NewMSFT_WmiErrorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessError{
		MSFT_WmiError: tmp,
	}
	return
}

func NewRemoteAccessErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessError, err error) {
	tmp, err := NewMSFT_WmiErrorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessError{
		MSFT_WmiError: tmp,
	}
	return
}

// SetRemoteAccessErrorCode sets the value of RemoteAccessErrorCode for the instance
func (instance *RemoteAccessError) SetPropertyRemoteAccessErrorCode(value uint32) (err error) {
	return instance.SetProperty("RemoteAccessErrorCode", (value))
}

// GetRemoteAccessErrorCode gets the value of RemoteAccessErrorCode for the instance
func (instance *RemoteAccessError) GetPropertyRemoteAccessErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemoteAccessErrorCode")
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
