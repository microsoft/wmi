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

// Msvm_MetricForME struct
type Msvm_MetricForME struct {
	*CIM_MetricForME
}

func NewMsvm_MetricForMEEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricForME, err error) {
	tmp, err := NewCIM_MetricForMEEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricForME{
		CIM_MetricForME: tmp,
	}
	return
}

func NewMsvm_MetricForMEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MetricForME, err error) {
	tmp, err := NewCIM_MetricForMEEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricForME{
		CIM_MetricForME: tmp,
	}
	return
}
