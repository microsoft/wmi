// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_FcEndpoint struct
type Msvm_FcEndpoint struct {
	*CIM_ProtocolEndpoint

	//
	Connected bool

	//
	WWNN string

	//
	WWPN string
}

func NewMsvm_FcEndpointEx1(instance *cim.WmiInstance) (newInstance *Msvm_FcEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewMsvm_FcEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FcEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetConnected sets the value of Connected for the instance
func (instance *Msvm_FcEndpoint) SetPropertyConnected(value bool) (err error) {
	return instance.SetProperty("Connected", value)
}

// GetConnected gets the value of Connected for the instance
func (instance *Msvm_FcEndpoint) GetPropertyConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("Connected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWWNN sets the value of WWNN for the instance
func (instance *Msvm_FcEndpoint) SetPropertyWWNN(value string) (err error) {
	return instance.SetProperty("WWNN", value)
}

// GetWWNN gets the value of WWNN for the instance
func (instance *Msvm_FcEndpoint) GetPropertyWWNN() (value string, err error) {
	retValue, err := instance.GetProperty("WWNN")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWWPN sets the value of WWPN for the instance
func (instance *Msvm_FcEndpoint) SetPropertyWWPN(value string) (err error) {
	return instance.SetProperty("WWPN", value)
}

// GetWWPN gets the value of WWPN for the instance
func (instance *Msvm_FcEndpoint) GetPropertyWWPN() (value string, err error) {
	retValue, err := instance.GetProperty("WWPN")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
