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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MetricInstance struct
type Msvm_MetricInstance struct {
	*CIM_MetricInstance
}

func NewMsvm_MetricInstanceEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricInstance, err error) {
	tmp, err := NewCIM_MetricInstanceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricInstance{
		CIM_MetricInstance: tmp,
	}
	return
}

func NewMsvm_MetricInstanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MetricInstance, err error) {
	tmp, err := NewCIM_MetricInstanceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricInstance{
		CIM_MetricInstance: tmp,
	}
	return
}
