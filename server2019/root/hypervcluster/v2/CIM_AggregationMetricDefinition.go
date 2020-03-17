// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_AggregationMetricDefinition struct
type CIM_AggregationMetricDefinition struct {
	CIM_BaseMetricDefinition

	// The SimpleFunction property identifies the basic computation performed on an underlying metric to arrive at the value of this derived metric. This property shall be NULL when the ChangeType property has a value other than 5 "Simple Function".
	///2="Minimum" indicates that the metric reports the lowest value detected for the associated monitored entity. This is also known as a low watermark.
	///.3="Maximum" indicates that the metric reports the maximum value detected for the associated monitored entity. This is also known as a high watermark.
	///4="Average" indicates the metric reports the average value of the underlying metric values.
	///5="Median" indicates the metric reports the median value of the underlying metric values.
	///6="Mode" indicates the metric reports the modal value of the underlying metric values.
	SimpleFunction AggregationMetricDefinition_SimpleFunction
}

// SetSimpleFunction sets the value of SimpleFunction for the instance
func (instance *CIM_AggregationMetricDefinition) SetPropertySimpleFunction(value AggregationMetricDefinition_SimpleFunction) (err error) {
	return instance.SetProperty("SimpleFunction", value)
}

// GetSimpleFunction gets the value of SimpleFunction for the instance
func (instance *CIM_AggregationMetricDefinition) GetPropertySimpleFunction() (value AggregationMetricDefinition_SimpleFunction, err error) {
	retValue, err := instance.GetProperty("SimpleFunction")
	if err != nil {
		return
	}
	value, ok := retValue.(AggregationMetricDefinition_SimpleFunction)
	if !ok {
		// TODO: Set an error
	}
	return
}
