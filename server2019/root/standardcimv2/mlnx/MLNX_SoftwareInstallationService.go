// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_SoftwareInstallationService struct
type MLNX_SoftwareInstallationService struct {
	CIM_SoftwareInstallationService
}

//

// <param name="Install" type="bool "></param>
// <param name="Log" type="string "></param>
// <param name="Reboot" type="bool "></param>
// <param name="Repair" type="bool "></param>
// <param name="Source" type="string "></param>
// <param name="Target" type="string "></param>
// <param name="Uninstall" type="bool "></param>
// <param name="Update" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_SoftwareInstallationService) Install( /* IN */ Source string,
	/* IN */ Target string,
	/* IN */ Install bool,
	/* IN */ Update bool,
	/* IN */ Repair bool,
	/* IN */ Uninstall bool,
	/* IN */ Log string,
	/* IN */ Reboot bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Install", Source, Target, Install, Update, Repair, Uninstall, Log, Reboot)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
