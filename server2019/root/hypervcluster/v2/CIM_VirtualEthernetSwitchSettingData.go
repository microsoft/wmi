// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_VirtualEthernetSwitchSettingData struct
type CIM_VirtualEthernetSwitchSettingData struct {
	CIM_VirtualSystemSettingData

	// A list of host resource pools to be associated or that are currently associated with the Ethernet Switch for the purpose of the allocation of Ethernet connections between a virtual machine and an Ethernet switch. Each non-Null value of the AssociatedResourcePool property shall conform to the production WBEM_URI_UntypedInstancePath as defined in DSP0207.
	AssociatedResourcePool []string

	// This property specifies the number of unique MAC addresses that can be learned by the switch to support MAC Address Learning, as defined in the IEEE 802.1 standard.
	MaxNumMACAddress uint32

	// A list of VLAN Ids that this switch can access.
	VLANConnection []string
}

// SetAssociatedResourcePool sets the value of AssociatedResourcePool for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyAssociatedResourcePool(value []string) (err error) {
	return instance.SetProperty("AssociatedResourcePool", value)
}

// GetAssociatedResourcePool gets the value of AssociatedResourcePool for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyAssociatedResourcePool() (value []string, err error) {
	retValue, err := instance.GetProperty("AssociatedResourcePool")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxNumMACAddress sets the value of MaxNumMACAddress for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyMaxNumMACAddress(value uint32) (err error) {
	return instance.SetProperty("MaxNumMACAddress", value)
}

// GetMaxNumMACAddress gets the value of MaxNumMACAddress for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyMaxNumMACAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumMACAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVLANConnection sets the value of VLANConnection for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) SetPropertyVLANConnection(value []string) (err error) {
	return instance.SetProperty("VLANConnection", value)
}

// GetVLANConnection gets the value of VLANConnection for the instance
func (instance *CIM_VirtualEthernetSwitchSettingData) GetPropertyVLANConnection() (value []string, err error) {
	retValue, err := instance.GetProperty("VLANConnection")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
