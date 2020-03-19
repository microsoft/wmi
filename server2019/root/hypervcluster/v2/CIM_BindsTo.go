// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BindsTo struct
type CIM_BindsTo struct {
	*CIM_SAPSAPDependency
}

func NewCIM_BindsToEx1(instance *cim.WmiInstance) (newInstance *CIM_BindsTo, err error) {
	tmp, err := NewCIM_SAPSAPDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BindsTo{
		CIM_SAPSAPDependency: tmp,
	}
	return
}

func NewCIM_BindsToEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BindsTo, err error) {
	tmp, err := NewCIM_SAPSAPDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BindsTo{
		CIM_SAPSAPDependency: tmp,
	}
	return
}
