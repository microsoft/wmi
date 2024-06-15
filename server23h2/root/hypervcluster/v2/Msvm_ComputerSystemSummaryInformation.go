// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ComputerSystemSummaryInformation struct
type Msvm_ComputerSystemSummaryInformation struct {
	*CIM_ElementView
}

func NewMsvm_ComputerSystemSummaryInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_ComputerSystemSummaryInformation, err error) {
	tmp, err := NewCIM_ElementViewEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ComputerSystemSummaryInformation{
		CIM_ElementView: tmp,
	}
	return
}

func NewMsvm_ComputerSystemSummaryInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ComputerSystemSummaryInformation, err error) {
	tmp, err := NewCIM_ElementViewEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ComputerSystemSummaryInformation{
		CIM_ElementView: tmp,
	}
	return
}
