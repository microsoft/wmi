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

// Msvm_ExternalEthernetPort struct
type Msvm_ExternalEthernetPort struct {
	CIM_EthernetPort

	//
	IsBound bool
}

// SetIsBound sets the value of IsBound for the instance
func (instance *Msvm_ExternalEthernetPort) SetPropertyIsBound(value bool) (err error) {
	return instance.SetProperty("IsBound", value)
}

// GetIsBound gets the value of IsBound for the instance
func (instance *Msvm_ExternalEthernetPort) GetPropertyIsBound() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBound")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_ExternalEthernetPort) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ExternalEthernetPort) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_ExternalEthernetPort) GetRelatedExternalEthernetPortCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ExternalEthernetPortCapabilities")
}
