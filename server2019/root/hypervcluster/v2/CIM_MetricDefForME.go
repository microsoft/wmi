// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_MetricDefForME struct
type CIM_MetricDefForME struct {
	CIM_Dependency

	// MetricCollectionEnabled indicates whether the metric defined by the referenced CIM_BaseMetricDefinition is being collected for the referenced CIM_ManagedElement. A value of 2 "Enabled" shall indicate the metric is being collected. A value of 3 "Disabled" shall indicate the metric is not being collected. When collection of a metric is re-enabled, the metric is re-initialized such that any values for a current access metric reflect data collected after the time at which collection was re-enabled.
	MetricCollectionEnabled MetricDefForME_MetricCollectionEnabled
}

// SetMetricCollectionEnabled sets the value of MetricCollectionEnabled for the instance
func (instance *CIM_MetricDefForME) SetPropertyMetricCollectionEnabled(value MetricDefForME_MetricCollectionEnabled) (err error) {
	return instance.SetProperty("MetricCollectionEnabled", value)
}

// GetMetricCollectionEnabled gets the value of MetricCollectionEnabled for the instance
func (instance *CIM_MetricDefForME) GetPropertyMetricCollectionEnabled() (value MetricDefForME_MetricCollectionEnabled, err error) {
	retValue, err := instance.GetProperty("MetricCollectionEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(MetricDefForME_MetricCollectionEnabled)
	if !ok {
		// TODO: Set an error
	}
	return
}
