// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_MetricRecordReal64 struct
type MSCluster_MetricRecordReal64 struct {
	MSCluster_MetricRecord

	//
	Value float64
}

// SetValue sets the value of Value for the instance
func (instance *MSCluster_MetricRecordReal64) SetPropertyValue(value float64) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MSCluster_MetricRecordReal64) GetPropertyValue() (value float64, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}
