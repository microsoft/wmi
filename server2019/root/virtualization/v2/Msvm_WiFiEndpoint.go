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

// Msvm_WiFiEndpoint struct
type Msvm_WiFiEndpoint struct {
	*CIM_WiFiEndpoint

	//
	Connected bool
}

func NewMsvm_WiFiEndpointEx1(instance *cim.WmiInstance) (newInstance *Msvm_WiFiEndpoint, err error) {
	tmp, err := NewCIM_WiFiEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiEndpoint{
		CIM_WiFiEndpoint: tmp,
	}
	return
}

func NewMsvm_WiFiEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_WiFiEndpoint, err error) {
	tmp, err := NewCIM_WiFiEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiEndpoint{
		CIM_WiFiEndpoint: tmp,
	}
	return
}

// SetConnected sets the value of Connected for the instance
func (instance *Msvm_WiFiEndpoint) SetPropertyConnected(value bool) (err error) {
	return instance.SetProperty("Connected", value)
}

// GetConnected gets the value of Connected for the instance
func (instance *Msvm_WiFiEndpoint) GetPropertyConnected() (value bool, err error) {
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
