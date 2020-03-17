// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_HealthMetric struct
type MSCluster_HealthMetric struct {
	cim.WmiInstance

	//
	MetricId string

	//
	Records []MSCluster_MetricRecord
}

// SetMetricId sets the value of MetricId for the instance
func (instance *MSCluster_HealthMetric) SetPropertyMetricId(value string) (err error) {
	return instance.SetProperty("MetricId", value)
}

// GetMetricId gets the value of MetricId for the instance
func (instance *MSCluster_HealthMetric) GetPropertyMetricId() (value string, err error) {
	retValue, err := instance.GetProperty("MetricId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecords sets the value of Records for the instance
func (instance *MSCluster_HealthMetric) SetPropertyRecords(value []MSCluster_MetricRecord) (err error) {
	return instance.SetProperty("Records", value)
}

// GetRecords gets the value of Records for the instance
func (instance *MSCluster_HealthMetric) GetPropertyRecords() (value []MSCluster_MetricRecord, err error) {
	retValue, err := instance.GetProperty("Records")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSCluster_MetricRecord)
	if !ok {
		// TODO: Set an error
	}
	return
}
