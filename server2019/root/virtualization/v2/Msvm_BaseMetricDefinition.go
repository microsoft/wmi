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

// Msvm_BaseMetricDefinition struct
type Msvm_BaseMetricDefinition struct {
	*CIM_BaseMetricDefinition
}

func NewMsvm_BaseMetricDefinitionEx1(instance *cim.WmiInstance) (newInstance *Msvm_BaseMetricDefinition, err error) {
	tmp, err := NewCIM_BaseMetricDefinitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BaseMetricDefinition{
		CIM_BaseMetricDefinition: tmp,
	}
	return
}

func NewMsvm_BaseMetricDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BaseMetricDefinition, err error) {
	tmp, err := NewCIM_BaseMetricDefinitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BaseMetricDefinition{
		CIM_BaseMetricDefinition: tmp,
	}
	return
}

func (instance *Msvm_BaseMetricDefinition) GetRelatedBaseMetricDefinition() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_BaseMetricDefinition")
}

func (instance *Msvm_BaseMetricDefinition) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
