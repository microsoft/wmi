// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_ServiceAffectsElement struct
type MLNX_ServiceAffectsElement struct {
	*CIM_ServiceAffectsElement
}

func NewMLNX_ServiceAffectsElementEx1(instance *cim.WmiInstance) (newInstance *MLNX_ServiceAffectsElement, err error) {
	tmp, err := NewCIM_ServiceAffectsElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_ServiceAffectsElement{
		CIM_ServiceAffectsElement: tmp,
	}
	return
}

func NewMLNX_ServiceAffectsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_ServiceAffectsElement, err error) {
	tmp, err := NewCIM_ServiceAffectsElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_ServiceAffectsElement{
		CIM_ServiceAffectsElement: tmp,
	}
	return
}
