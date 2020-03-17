// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_LANEndpoint struct
type Msvm_LANEndpoint struct {
	CIM_LANEndpoint

	//
	Connected bool
}

// SetConnected sets the value of Connected for the instance
func (instance *Msvm_LANEndpoint) SetPropertyConnected(value bool) (err error) {
	return instance.SetProperty("Connected", value)
}

// GetConnected gets the value of Connected for the instance
func (instance *Msvm_LANEndpoint) GetPropertyConnected() (value bool, err error) {
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
func (instance *Msvm_LANEndpoint) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_LANEndpoint) GetRelatedInternalEthernetPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_InternalEthernetPort")
}
