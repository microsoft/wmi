// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_BaseMetricValue struct
type Msvm_BaseMetricValue struct {
	*CIM_BaseMetricValue
}

func NewMsvm_BaseMetricValueEx1(instance *cim.WmiInstance) (newInstance *Msvm_BaseMetricValue, err error) {
	tmp, err := NewCIM_BaseMetricValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BaseMetricValue{
		CIM_BaseMetricValue: tmp,
	}
	return
}

func NewMsvm_BaseMetricValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BaseMetricValue, err error) {
	tmp, err := NewCIM_BaseMetricValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BaseMetricValue{
		CIM_BaseMetricValue: tmp,
	}
	return
}
