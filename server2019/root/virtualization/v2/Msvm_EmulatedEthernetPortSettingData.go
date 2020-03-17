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

// Msvm_EmulatedEthernetPortSettingData struct
type Msvm_EmulatedEthernetPortSettingData struct {
	CIM_EthernetPortAllocationSettingData

	//
	ClusterMonitored bool

	//
	StaticMacAddress bool
}

// SetClusterMonitored sets the value of ClusterMonitored for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) SetPropertyClusterMonitored(value bool) (err error) {
	return instance.SetProperty("ClusterMonitored", value)
}

// GetClusterMonitored gets the value of ClusterMonitored for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) GetPropertyClusterMonitored() (value bool, err error) {
	retValue, err := instance.GetProperty("ClusterMonitored")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStaticMacAddress sets the value of StaticMacAddress for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) SetPropertyStaticMacAddress(value bool) (err error) {
	return instance.SetProperty("StaticMacAddress", value)
}

// GetStaticMacAddress gets the value of StaticMacAddress for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) GetPropertyStaticMacAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("StaticMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EmulatedEthernetPortSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
