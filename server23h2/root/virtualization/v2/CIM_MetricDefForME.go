// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_MetricDefForME struct
type CIM_MetricDefForME struct {
	*CIM_Dependency

	// MetricCollectionEnabled indicates whether the metric defined by the referenced CIM_BaseMetricDefinition is being collected for the referenced CIM_ManagedElement. A value of 2 "Enabled" shall indicate the metric is being collected. A value of 3 "Disabled" shall indicate the metric is not being collected. When collection of a metric is re-enabled, the metric is re-initialized such that any values for a current access metric reflect data collected after the time at which collection was re-enabled.
	MetricCollectionEnabled MetricDefForME_MetricCollectionEnabled
}

func NewCIM_MetricDefForMEEx1(instance *cim.WmiInstance) (newInstance *CIM_MetricDefForME, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricDefForME{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_MetricDefForMEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MetricDefForME, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricDefForME{
		CIM_Dependency: tmp,
	}
	return
}

// SetMetricCollectionEnabled sets the value of MetricCollectionEnabled for the instance
func (instance *CIM_MetricDefForME) SetPropertyMetricCollectionEnabled(value MetricDefForME_MetricCollectionEnabled) (err error) {
	return instance.SetProperty("MetricCollectionEnabled", (value))
}

// GetMetricCollectionEnabled gets the value of MetricCollectionEnabled for the instance
func (instance *CIM_MetricDefForME) GetPropertyMetricCollectionEnabled() (value MetricDefForME_MetricCollectionEnabled, err error) {
	retValue, err := instance.GetProperty("MetricCollectionEnabled")
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

	value = MetricDefForME_MetricCollectionEnabled(valuetmp)

	return
}
