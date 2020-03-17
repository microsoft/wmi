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

// Msvm_DynamicForwardingEntry struct
type Msvm_DynamicForwardingEntry struct {
	CIM_DynamicForwardingEntry

	//
	VlanId uint16
}

// SetVlanId sets the value of VlanId for the instance
func (instance *Msvm_DynamicForwardingEntry) SetPropertyVlanId(value uint16) (err error) {
	return instance.SetProperty("VlanId", value)
}

// GetVlanId gets the value of VlanId for the instance
func (instance *Msvm_DynamicForwardingEntry) GetPropertyVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_DynamicForwardingEntry) GetRelatedTransparentBridgingService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TransparentBridgingService")
}

func (instance *Msvm_DynamicForwardingEntry) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPort")
}
