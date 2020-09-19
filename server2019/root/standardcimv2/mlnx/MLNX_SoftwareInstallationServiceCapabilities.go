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

// MLNX_SoftwareInstallationServiceCapabilities struct
type MLNX_SoftwareInstallationServiceCapabilities struct {
	*CIM_SoftwareInstallationServiceCapabilities
}

func NewMLNX_SoftwareInstallationServiceCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_SoftwareInstallationServiceCapabilities, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareInstallationServiceCapabilities{
		CIM_SoftwareInstallationServiceCapabilities: tmp,
	}
	return
}

func NewMLNX_SoftwareInstallationServiceCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_SoftwareInstallationServiceCapabilities, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareInstallationServiceCapabilities{
		CIM_SoftwareInstallationServiceCapabilities: tmp,
	}
	return
}
