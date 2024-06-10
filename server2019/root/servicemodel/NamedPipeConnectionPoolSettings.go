// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NamedPipeConnectionPoolSettings struct
type NamedPipeConnectionPoolSettings struct {
	*cim.WmiInstance

	// The group name of the connection pool used by the binding element.
	GroupName string

	// The maximum time the connection can be idle before being disconnected.
	IdleTimeout string

	// The maximum number of outbound connections per endpoint on the client.
	MaxOutboundConnectionsPerEndpoint int32
}

func NewNamedPipeConnectionPoolSettingsEx1(instance *cim.WmiInstance) (newInstance *NamedPipeConnectionPoolSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &NamedPipeConnectionPoolSettings{
		WmiInstance: tmp,
	}
	return
}

func NewNamedPipeConnectionPoolSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NamedPipeConnectionPoolSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NamedPipeConnectionPoolSettings{
		WmiInstance: tmp,
	}
	return
}

// SetGroupName sets the value of GroupName for the instance
func (instance *NamedPipeConnectionPoolSettings) SetPropertyGroupName(value string) (err error) {
	return instance.SetProperty("GroupName", (value))
}

// GetGroupName gets the value of GroupName for the instance
func (instance *NamedPipeConnectionPoolSettings) GetPropertyGroupName() (value string, err error) {
	retValue, err := instance.GetProperty("GroupName")
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

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *NamedPipeConnectionPoolSettings) SetPropertyIdleTimeout(value string) (err error) {
	return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *NamedPipeConnectionPoolSettings) GetPropertyIdleTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("IdleTimeout")
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

// SetMaxOutboundConnectionsPerEndpoint sets the value of MaxOutboundConnectionsPerEndpoint for the instance
func (instance *NamedPipeConnectionPoolSettings) SetPropertyMaxOutboundConnectionsPerEndpoint(value int32) (err error) {
	return instance.SetProperty("MaxOutboundConnectionsPerEndpoint", (value))
}

// GetMaxOutboundConnectionsPerEndpoint gets the value of MaxOutboundConnectionsPerEndpoint for the instance
func (instance *NamedPipeConnectionPoolSettings) GetPropertyMaxOutboundConnectionsPerEndpoint() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxOutboundConnectionsPerEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
