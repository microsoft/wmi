// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_AggregationMetricDefinition struct
type CIM_AggregationMetricDefinition struct {
	*CIM_BaseMetricDefinition

	// The SimpleFunction property identifies the basic computation performed on an underlying metric to arrive at the value of this derived metric. This property shall be NULL when the ChangeType property has a value other than 5 "Simple Function".
	///2="Minimum" indicates that the metric reports the lowest value detected for the associated monitored entity. This is also known as a low watermark.
	///.3="Maximum" indicates that the metric reports the maximum value detected for the associated monitored entity. This is also known as a high watermark.
	///4="Average" indicates the metric reports the average value of the underlying metric values.
	///5="Median" indicates the metric reports the median value of the underlying metric values.
	///6="Mode" indicates the metric reports the modal value of the underlying metric values.
	SimpleFunction AggregationMetricDefinition_SimpleFunction
}

func NewCIM_AggregationMetricDefinitionEx1(instance *cim.WmiInstance) (newInstance *CIM_AggregationMetricDefinition, err error) {
	tmp, err := NewCIM_BaseMetricDefinitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AggregationMetricDefinition{
		CIM_BaseMetricDefinition: tmp,
	}
	return
}

func NewCIM_AggregationMetricDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AggregationMetricDefinition, err error) {
	tmp, err := NewCIM_BaseMetricDefinitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AggregationMetricDefinition{
		CIM_BaseMetricDefinition: tmp,
	}
	return
}

// SetSimpleFunction sets the value of SimpleFunction for the instance
func (instance *CIM_AggregationMetricDefinition) SetPropertySimpleFunction(value AggregationMetricDefinition_SimpleFunction) (err error) {
	return instance.SetProperty("SimpleFunction", (value))
}

// GetSimpleFunction gets the value of SimpleFunction for the instance
func (instance *CIM_AggregationMetricDefinition) GetPropertySimpleFunction() (value AggregationMetricDefinition_SimpleFunction, err error) {
	retValue, err := instance.GetProperty("SimpleFunction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AggregationMetricDefinition_SimpleFunction(valuetmp)

	return
}
