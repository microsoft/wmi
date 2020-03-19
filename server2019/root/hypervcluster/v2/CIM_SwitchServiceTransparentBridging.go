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

// CIM_SwitchServiceTransparentBridging struct
type CIM_SwitchServiceTransparentBridging struct {
	*CIM_ServiceComponent
}

func NewCIM_SwitchServiceTransparentBridgingEx1(instance *cim.WmiInstance) (newInstance *CIM_SwitchServiceTransparentBridging, err error) {
	tmp, err := NewCIM_ServiceComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchServiceTransparentBridging{
		CIM_ServiceComponent: tmp,
	}
	return
}

func NewCIM_SwitchServiceTransparentBridgingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SwitchServiceTransparentBridging, err error) {
	tmp, err := NewCIM_ServiceComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SwitchServiceTransparentBridging{
		CIM_ServiceComponent: tmp,
	}
	return
}
