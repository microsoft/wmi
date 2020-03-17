// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_BindsToLANEndpoint struct
type CIM_BindsToLANEndpoint struct {
	CIM_BindsTo

	// This describes the framing method for the upper layer SAP or Endpoint that is bound to the LANEndpoint. Note: "Raw802.3" is only known to be used with the IPX protocol.
	FrameType BindsToLANEndpoint_FrameType
}

// SetFrameType sets the value of FrameType for the instance
func (instance *CIM_BindsToLANEndpoint) SetPropertyFrameType(value BindsToLANEndpoint_FrameType) (err error) {
	return instance.SetProperty("FrameType", value)
}

// GetFrameType gets the value of FrameType for the instance
func (instance *CIM_BindsToLANEndpoint) GetPropertyFrameType() (value BindsToLANEndpoint_FrameType, err error) {
	retValue, err := instance.GetProperty("FrameType")
	if err != nil {
		return
	}
	value, ok := retValue.(BindsToLANEndpoint_FrameType)
	if !ok {
		// TODO: Set an error
	}
	return
}
