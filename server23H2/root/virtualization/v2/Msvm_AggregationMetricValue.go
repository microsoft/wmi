// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AggregationMetricValue struct
type Msvm_AggregationMetricValue struct {
	*CIM_AggregationMetricValue
}

func NewMsvm_AggregationMetricValueEx1(instance *cim.WmiInstance) (newInstance *Msvm_AggregationMetricValue, err error) {
	tmp, err := NewCIM_AggregationMetricValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AggregationMetricValue{
		CIM_AggregationMetricValue: tmp,
	}
	return
}

func NewMsvm_AggregationMetricValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AggregationMetricValue, err error) {
	tmp, err := NewCIM_AggregationMetricValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AggregationMetricValue{
		CIM_AggregationMetricValue: tmp,
	}
	return
}
