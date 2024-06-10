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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessServerAuthConfiguration struct
type RemoteAccessServerAuthConfiguration struct {
	*cim.WmiInstance

	//
	AuthConfig interface{}
}

func NewRemoteAccessServerAuthConfigurationEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessServerAuthConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessServerAuthConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessServerAuthConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessServerAuthConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessServerAuthConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAuthConfig sets the value of AuthConfig for the instance
func (instance *RemoteAccessServerAuthConfiguration) SetPropertyAuthConfig(value interface{}) (err error) {
	return instance.SetProperty("AuthConfig", (value))
}

// GetAuthConfig gets the value of AuthConfig for the instance
func (instance *RemoteAccessServerAuthConfiguration) GetPropertyAuthConfig() (value interface{}, err error) {
	retValue, err := instance.GetProperty("AuthConfig")
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
