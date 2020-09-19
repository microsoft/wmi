// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_AggregationMetricValue struct
type CIM_AggregationMetricValue struct {
	*CIM_BaseMetricValue

	// Property that represents the time duration over which the aggregation was computed. The start of a monitoring interval over which the aggregation function is applied is determined by subtracting the AggregationDuration from the AggregationTimestamp.
	AggregationDuration string

	// Identifies the time when the aggregation function was applied to determine the value of the metric instance. Note that this is different from the time when the instance is created. For a given CIM_AggregationMetricValue instance, the AggregationTimeStamp changes whenever the aggregation function is applied to calculate the value.
	AggregationTimeStamp string
}

func NewCIM_AggregationMetricValueEx1(instance *cim.WmiInstance) (newInstance *CIM_AggregationMetricValue, err error) {
	tmp, err := NewCIM_BaseMetricValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AggregationMetricValue{
		CIM_BaseMetricValue: tmp,
	}
	return
}

func NewCIM_AggregationMetricValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AggregationMetricValue, err error) {
	tmp, err := NewCIM_BaseMetricValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AggregationMetricValue{
		CIM_BaseMetricValue: tmp,
	}
	return
}

// SetAggregationDuration sets the value of AggregationDuration for the instance
func (instance *CIM_AggregationMetricValue) SetPropertyAggregationDuration(value string) (err error) {
	return instance.SetProperty("AggregationDuration", (value))
}

// GetAggregationDuration gets the value of AggregationDuration for the instance
func (instance *CIM_AggregationMetricValue) GetPropertyAggregationDuration() (value string, err error) {
	retValue, err := instance.GetProperty("AggregationDuration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetAggregationTimeStamp sets the value of AggregationTimeStamp for the instance
func (instance *CIM_AggregationMetricValue) SetPropertyAggregationTimeStamp(value string) (err error) {
	return instance.SetProperty("AggregationTimeStamp", (value))
}

// GetAggregationTimeStamp gets the value of AggregationTimeStamp for the instance
func (instance *CIM_AggregationMetricValue) GetPropertyAggregationTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("AggregationTimeStamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
