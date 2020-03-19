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

// Msvm_MetricServiceCapabilities struct
type Msvm_MetricServiceCapabilities struct {
	*CIM_MetricServiceCapabilities
}

func NewMsvm_MetricServiceCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricServiceCapabilities, err error) {
	tmp, err := NewCIM_MetricServiceCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricServiceCapabilities{
		CIM_MetricServiceCapabilities: tmp,
	}
	return
}

func NewMsvm_MetricServiceCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MetricServiceCapabilities, err error) {
	tmp, err := NewCIM_MetricServiceCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricServiceCapabilities{
		CIM_MetricServiceCapabilities: tmp,
	}
	return
}

func (instance *Msvm_MetricServiceCapabilities) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}
