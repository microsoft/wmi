// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MetricServiceSettingData struct
type Msvm_MetricServiceSettingData struct {
	*CIM_SettingData

	//
	MetricsFlushInterval string
}

func NewMsvm_MetricServiceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_MetricServiceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MetricServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
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
