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

// Msvm_EthernetSwitchPortRdmaSettingData struct
type Msvm_EthernetSwitchPortRdmaSettingData struct {
	Msvm_EthernetSwitchPortFeatureSettingData

	//
	RdmaOffloadWeight uint32
}

// SetRdmaOffloadWeight sets the value of RdmaOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) SetPropertyRdmaOffloadWeight(value uint32) (err error) {
	return instance.SetProperty("RdmaOffloadWeight", value)
}

// GetRdmaOffloadWeight gets the value of RdmaOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) GetPropertyRdmaOffloadWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("RdmaOffloadWeight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
