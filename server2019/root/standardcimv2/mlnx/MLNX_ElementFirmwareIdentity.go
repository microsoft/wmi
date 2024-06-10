// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_ElementFirmwareIdentity struct
type MLNX_ElementFirmwareIdentity struct {
	*CIM_ElementSoftwareIdentity
}

func NewMLNX_ElementFirmwareIdentityEx1(instance *cim.WmiInstance) (newInstance *MLNX_ElementFirmwareIdentity, err error) {
	tmp, err := NewCIM_ElementSoftwareIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_ElementFirmwareIdentity{
		CIM_ElementSoftwareIdentity: tmp,
	}
	return
}

func NewMLNX_ElementFirmwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_ElementFirmwareIdentity, err error) {
	tmp, err := NewCIM_ElementSoftwareIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_ElementFirmwareIdentity{
		CIM_ElementSoftwareIdentity: tmp,
	}
	return
}
