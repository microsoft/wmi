// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_NetworkPipe struct
type CIM_NetworkPipe struct {
	CIM_EnabledLogicalElement

	//
	AggregationBehavior uint16

	//
	Directionality uint16
}

// SetAggregationBehavior sets the value of AggregationBehavior for the instance
func (instance *CIM_NetworkPipe) SetPropertyAggregationBehavior(value uint16) (err error) {
	return instance.SetProperty("AggregationBehavior", value)
}

// GetAggregationBehavior gets the value of AggregationBehavior for the instance
func (instance *CIM_NetworkPipe) GetPropertyAggregationBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("AggregationBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectionality sets the value of Directionality for the instance
func (instance *CIM_NetworkPipe) SetPropertyDirectionality(value uint16) (err error) {
	return instance.SetProperty("Directionality", value)
}

// GetDirectionality gets the value of Directionality for the instance
func (instance *CIM_NetworkPipe) GetPropertyDirectionality() (value uint16, err error) {
	retValue, err := instance.GetProperty("Directionality")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
