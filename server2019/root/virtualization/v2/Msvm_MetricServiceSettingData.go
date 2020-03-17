// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MetricServiceSettingData struct
type Msvm_MetricServiceSettingData struct {
	CIM_SettingData

	//
	MetricsFlushInterval string
}

// SetMetricsFlushInterval sets the value of MetricsFlushInterval for the instance
func (instance *Msvm_MetricServiceSettingData) SetPropertyMetricsFlushInterval(value string) (err error) {
	return instance.SetProperty("MetricsFlushInterval", value)
}

// GetMetricsFlushInterval gets the value of MetricsFlushInterval for the instance
func (instance *Msvm_MetricServiceSettingData) GetPropertyMetricsFlushInterval() (value string, err error) {
	retValue, err := instance.GetProperty("MetricsFlushInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_MetricServiceSettingData) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
