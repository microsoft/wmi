// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_FirmwareInstallationService struct
type MLNX_FirmwareInstallationService struct {
	CIM_SoftwareInstallationService
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
