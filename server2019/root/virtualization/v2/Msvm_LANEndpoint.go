// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_LANEndpoint struct
type Msvm_LANEndpoint struct {
	*CIM_LANEndpoint

	//
	Connected bool
}

func NewMsvm_LANEndpointEx1(instance *cim.WmiInstance) (newInstance *Msvm_LANEndpoint, err error) {
	tmp, err := NewCIM_LANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_LANEndpoint{
		CIM_LANEndpoint: tmp,
	}
	return
}

func NewMsvm_LANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_LANEndpoint, err error) {
	tmp, err := NewCIM_LANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_LANEndpoint{
		CIM_LANEndpoint: tmp,
	}
	return
}

// SetConnected sets the value of Connected for the instance
func (instance *Msvm_LANEndpoint) SetPropertyConnected(value bool) (err error) {
	return instance.SetProperty("Connected", (value))
}

// GetConnected gets the value of Connected for the instance
func (instance *Msvm_LANEndpoint) GetPropertyConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("Connected")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
func (instance *Msvm_LANEndpoint) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_LANEndpoint) GetRelatedInternalEthernetPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_InternalEthernetPort")
}
