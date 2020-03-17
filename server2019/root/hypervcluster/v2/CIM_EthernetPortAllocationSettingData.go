// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_EthernetPortAllocationSettingData struct
type CIM_EthernetPortAllocationSettingData struct {
	CIM_ResourceAllocationSettingData

	// The desired VLAN mode that is requested for use. This property is usedto set the initial OperationalEndpointMode property value in theinstance of CIM_VLANEndpoint associated with the targeted Ethernet Port.Refer to the description for the property OperationalEndpointMode inCIM_VLANEndpoint for a description of the values
	DesiredVLANEndpointMode EthernetPortAllocationSettingData_DesiredVLANEndpointMode

	// A string describing the type of VLAN endpoint model that is supported by this VLANEndpoint, when the value of the mode property is set to 1 (i.e., "Other"). This property should be set to NULL when the mode property is any value other than 1.
	OtherEndpointMode string
}

// SetDesiredVLANEndpointMode sets the value of DesiredVLANEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) SetPropertyDesiredVLANEndpointMode(value EthernetPortAllocationSettingData_DesiredVLANEndpointMode) (err error) {
	return instance.SetProperty("DesiredVLANEndpointMode", value)
}

// GetDesiredVLANEndpointMode gets the value of DesiredVLANEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) GetPropertyDesiredVLANEndpointMode() (value EthernetPortAllocationSettingData_DesiredVLANEndpointMode, err error) {
	retValue, err := instance.GetProperty("DesiredVLANEndpointMode")
	if err != nil {
		return
	}
	value, ok := retValue.(EthernetPortAllocationSettingData_DesiredVLANEndpointMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherEndpointMode sets the value of OtherEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) SetPropertyOtherEndpointMode(value string) (err error) {
	return instance.SetProperty("OtherEndpointMode", value)
}

// GetOtherEndpointMode gets the value of OtherEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) GetPropertyOtherEndpointMode() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEndpointMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
