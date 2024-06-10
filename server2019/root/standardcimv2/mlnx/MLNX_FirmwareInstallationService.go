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

// MLNX_FirmwareInstallationService struct
type MLNX_FirmwareInstallationService struct {
	*CIM_SoftwareInstallationService
}

func NewMLNX_FirmwareInstallationServiceEx1(instance *cim.WmiInstance) (newInstance *MLNX_FirmwareInstallationService, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareInstallationService{
		CIM_SoftwareInstallationService: tmp,
	}
	return
}

func NewMLNX_FirmwareInstallationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_FirmwareInstallationService, err error) {
	tmp, err := NewCIM_SoftwareInstallationServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareInstallationService{
		CIM_SoftwareInstallationService: tmp,
	}
	return
}

//

// <param name="Device" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="Log" type="string "></param>
// <param name="Reboot" type="bool "></param>
// <param name="Source" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_FirmwareInstallationService) Update( /* IN */ Source string,
	/* IN */ Device string,
	/* IN */ Reboot bool,
	/* IN */ Force bool,
	/* IN */ Log string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Update", Source, Device, Reboot, Force, Log)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
