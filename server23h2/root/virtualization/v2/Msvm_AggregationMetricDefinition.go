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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AggregationMetricDefinition struct
type Msvm_AggregationMetricDefinition struct {
	*CIM_AggregationMetricDefinition
}

func NewMsvm_AggregationMetricDefinitionEx1(instance *cim.WmiInstance) (newInstance *Msvm_AggregationMetricDefinition, err error) {
	tmp, err := NewCIM_AggregationMetricDefinitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AggregationMetricDefinition{
		CIM_AggregationMetricDefinition: tmp,
	}
	return
}

func NewMsvm_AggregationMetricDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AggregationMetricDefinition, err error) {
	tmp, err := NewCIM_AggregationMetricDefinitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AggregationMetricDefinition{
		CIM_AggregationMetricDefinition: tmp,
	}
	return
}

func (instance *Msvm_AggregationMetricDefinition) GetRelatedAggregationMetricDefinition() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AggregationMetricDefinition")
}

func (instance *Msvm_AggregationMetricDefinition) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
