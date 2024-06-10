// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessAccountingConnection struct
type RemoteAccessAccountingConnection struct {
	*RemoteAccessConnection

	//
	SessionId uint64

	//
	UserName string
}

func NewRemoteAccessAccountingConnectionEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessAccountingConnection, err error) {
	tmp, err := NewRemoteAccessConnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessAccountingConnection{
		RemoteAccessConnection: tmp,
	}
	return
}

func NewRemoteAccessAccountingConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessAccountingConnection, err error) {
	tmp, err := NewRemoteAccessConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessAccountingConnection{
		RemoteAccessConnection: tmp,
	}
	return
}

// SetSessionId sets the value of SessionId for the instance
func (instance *RemoteAccessAccountingConnection) SetPropertySessionId(value uint64) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *RemoteAccessAccountingConnection) GetPropertySessionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("SessionId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUserName sets the value of UserName for the instance
func (instance *RemoteAccessAccountingConnection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *RemoteAccessAccountingConnection) GetPropertyUserName() (value string, err error) {
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
