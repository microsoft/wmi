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

// CIM_MetricForME struct
type CIM_MetricForME struct {
	*CIM_Dependency
}

func NewCIM_MetricForMEEx1(instance *cim.WmiInstance) (newInstance *CIM_MetricForME, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricForME{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_MetricForMEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MetricForME, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricForME{
		CIM_Dependency: tmp,
	}
	return
}
