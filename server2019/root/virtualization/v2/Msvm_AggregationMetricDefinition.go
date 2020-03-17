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

// Msvm_AggregationMetricDefinition struct
type Msvm_AggregationMetricDefinition struct {
	CIM_AggregationMetricDefinition
}

func (instance *Msvm_AggregationMetricDefinition) GetRelatedAggregationMetricDefinition() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AggregationMetricDefinition")
}

func (instance *Msvm_AggregationMetricDefinition) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
