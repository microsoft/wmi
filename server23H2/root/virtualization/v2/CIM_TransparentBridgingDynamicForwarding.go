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

// CIM_TransparentBridgingDynamicForwarding struct
type CIM_TransparentBridgingDynamicForwarding struct {
	*CIM_Dependency
}

func NewCIM_TransparentBridgingDynamicForwardingEx1(instance *cim.WmiInstance) (newInstance *CIM_TransparentBridgingDynamicForwarding, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_TransparentBridgingDynamicForwarding{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_TransparentBridgingDynamicForwardingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_TransparentBridgingDynamicForwarding, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_TransparentBridgingDynamicForwarding{
		CIM_Dependency: tmp,
	}
	return
}
