// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Channel struct
type Channel struct {
	*cim.WmiInstance

	// The local endpoint for the channel.
	LocalAddress string

	// The remote address associated with the channel.
	RemoteAddress string

	// A reference to the endpoint the channel connects to.
	RemoteEndpoint Endpoint

	// The current session Id, if any.
	SessionId string

	// The type of the channel.
	Type string
}

func NewChannelEx1(instance *cim.WmiInstance) (newInstance *Channel, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Channel{
		WmiInstance: tmp,
	}
	return
}

func NewChannelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Channel, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Channel{
		WmiInstance: tmp,
	}
	return
}

// SetLocalAddress sets the value of LocalAddress for the instance
func (instance *Channel) SetPropertyLocalAddress(value string) (err error) {
	return instance.SetProperty("LocalAddress", (value))
}

// GetLocalAddress gets the value of LocalAddress for the instance
func (instance *Channel) GetPropertyLocalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("LocalAddress")
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

// SetRemoteAddress sets the value of RemoteAddress for the instance
func (instance *Channel) SetPropertyRemoteAddress(value string) (err error) {
	return instance.SetProperty("RemoteAddress", (value))
}

// GetRemoteAddress gets the value of RemoteAddress for the instance
func (instance *Channel) GetPropertyRemoteAddress() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteAddress")
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

// SetRemoteEndpoint sets the value of RemoteEndpoint for the instance
func (instance *Channel) SetPropertyRemoteEndpoint(value Endpoint) (err error) {
	return instance.SetProperty("RemoteEndpoint", (value))
}

// GetRemoteEndpoint gets the value of RemoteEndpoint for the instance
func (instance *Channel) GetPropertyRemoteEndpoint() (value Endpoint, err error) {
	retValue, err := instance.GetProperty("RemoteEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Endpoint)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Endpoint is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Endpoint(valuetmp)

	return
}

// SetSessionId sets the value of SessionId for the instance
func (instance *Channel) SetPropertySessionId(value string) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *Channel) GetPropertySessionId() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *Channel) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Channel) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
