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

// Msvm_BaseMetricDefinition struct
type Msvm_BaseMetricDefinition struct {
	CIM_BaseMetricDefinition
}

func (instance *Msvm_BaseMetricDefinition) GetRelatedBaseMetricDefinition() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_BaseMetricDefinition")
}

func (instance *Msvm_BaseMetricDefinition) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
