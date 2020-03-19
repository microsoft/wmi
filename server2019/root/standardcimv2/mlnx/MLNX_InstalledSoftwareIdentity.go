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

// MLNX_InstalledSoftwareIdentity struct
type MLNX_InstalledSoftwareIdentity struct {
	*CIM_InstalledSoftwareIdentity
}

func NewMLNX_InstalledSoftwareIdentityEx1(instance *cim.WmiInstance) (newInstance *MLNX_InstalledSoftwareIdentity, err error) {
	tmp, err := NewCIM_InstalledSoftwareIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_InstalledSoftwareIdentity{
		CIM_InstalledSoftwareIdentity: tmp,
	}
	return
}

func NewMLNX_InstalledSoftwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_InstalledSoftwareIdentity, err error) {
	tmp, err := NewCIM_InstalledSoftwareIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_InstalledSoftwareIdentity{
		CIM_InstalledSoftwareIdentity: tmp,
	}
	return
}
