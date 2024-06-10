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

// MLNX_FirmwareInstallationServiceCapabilities struct
type MLNX_FirmwareInstallationServiceCapabilities struct {
	*CIM_SoftwareInstallationServiceCapabilities
}

func NewMLNX_FirmwareInstallationServiceCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_FirmwareInstallationServiceCapabilities, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareInstallationServiceCapabilities{
		CIM_SoftwareInstallationServiceCapabilities: tmp,
	}
	return
}

func NewMLNX_FirmwareInstallationServiceCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_FirmwareInstallationServiceCapabilities, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareInstallationServiceCapabilities{
		CIM_SoftwareInstallationServiceCapabilities: tmp,
	}
	return
}
