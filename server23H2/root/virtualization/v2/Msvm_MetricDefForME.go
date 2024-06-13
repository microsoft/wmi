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

// Msvm_MetricDefForME struct
type Msvm_MetricDefForME struct {
	*CIM_MetricDefForME
}

func NewMsvm_MetricDefForMEEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricDefForME, err error) {
	tmp, err := NewCIM_MetricDefForMEEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricDefForME{
		CIM_MetricDefForME: tmp,
	}
	return
}

func NewMsvm_MetricDefForMEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MetricDefForME, err error) {
	tmp, err := NewCIM_MetricDefForMEEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MetricDefForME{
		CIM_MetricDefForME: tmp,
	}
	return
}
