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

// RemoteAccessRadiusConfig struct
type RemoteAccessRadiusConfig struct {
	*cim.WmiInstance

	//
	RadiusServerConfig interface{}

	//
	SharedSecret string
}

func NewRemoteAccessRadiusConfigEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessRadiusConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusConfig{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessRadiusConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessRadiusConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusConfig{
		WmiInstance: tmp,
	}
	return
}

// SetRadiusServerConfig sets the value of RadiusServerConfig for the instance
func (instance *RemoteAccessRadiusConfig) SetPropertyRadiusServerConfig(value interface{}) (err error) {
	return instance.SetProperty("RadiusServerConfig", (value))
}

// GetRadiusServerConfig gets the value of RadiusServerConfig for the instance
func (instance *RemoteAccessRadiusConfig) GetPropertyRadiusServerConfig() (value interface{}, err error) {
	retValue, err := instance.GetProperty("RadiusServerConfig")
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

// SetSharedSecret sets the value of SharedSecret for the instance
func (instance *RemoteAccessRadiusConfig) SetPropertySharedSecret(value string) (err error) {
	return instance.SetProperty("SharedSecret", (value))
}

// GetSharedSecret gets the value of SharedSecret for the instance
func (instance *RemoteAccessRadiusConfig) GetPropertySharedSecret() (value string, err error) {
	retValue, err := instance.GetProperty("SharedSecret")
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
