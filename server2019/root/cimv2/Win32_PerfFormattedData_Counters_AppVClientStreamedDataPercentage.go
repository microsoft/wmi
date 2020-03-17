// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_AppVClientStreamedDataPercentage struct
type Win32_PerfFormattedData_Counters_AppVClientStreamedDataPercentage struct {
	Win32_PerfFormattedData

	//
	PrimaryFeaturePercentStreamed uint32
}

// SetPrimaryFeaturePercentStreamed sets the value of PrimaryFeaturePercentStreamed for the instance
func (instance *Win32_PerfFormattedData_Counters_AppVClientStreamedDataPercentage) SetPropertyPrimaryFeaturePercentStreamed(value uint32) (err error) {
	return instance.SetProperty("PrimaryFeaturePercentStreamed", value)
}

// GetPrimaryFeaturePercentStreamed gets the value of PrimaryFeaturePercentStreamed for the instance
func (instance *Win32_PerfFormattedData_Counters_AppVClientStreamedDataPercentage) GetPropertyPrimaryFeaturePercentStreamed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrimaryFeaturePercentStreamed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
