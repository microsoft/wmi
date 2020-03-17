// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_ProtocolController struct
type CIM_ProtocolController struct {
	CIM_LogicalDevice

	// Maximum number of Units that can be controlled by or accessed through this ProtocolController.
	MaxUnitsControlled uint32
}

// SetMaxUnitsControlled sets the value of MaxUnitsControlled for the instance
func (instance *CIM_ProtocolController) SetPropertyMaxUnitsControlled(value uint32) (err error) {
	return instance.SetProperty("MaxUnitsControlled", value)
}

// GetMaxUnitsControlled gets the value of MaxUnitsControlled for the instance
func (instance *CIM_ProtocolController) GetPropertyMaxUnitsControlled() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxUnitsControlled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
