// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CAU_Update_Installer_Result struct
type MSFT_CAU_Update_Installer_Result struct {
	cim.WmiInstance

	//
	LaunchInstallerHResult uint32

	//
	ReturnCode int32
}

// SetLaunchInstallerHResult sets the value of LaunchInstallerHResult for the instance
func (instance *MSFT_CAU_Update_Installer_Result) SetPropertyLaunchInstallerHResult(value uint32) (err error) {
	return instance.SetProperty("LaunchInstallerHResult", value)
}

// GetLaunchInstallerHResult gets the value of LaunchInstallerHResult for the instance
func (instance *MSFT_CAU_Update_Installer_Result) GetPropertyLaunchInstallerHResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("LaunchInstallerHResult")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReturnCode sets the value of ReturnCode for the instance
func (instance *MSFT_CAU_Update_Installer_Result) SetPropertyReturnCode(value int32) (err error) {
	return instance.SetProperty("ReturnCode", value)
}

// GetReturnCode gets the value of ReturnCode for the instance
func (instance *MSFT_CAU_Update_Installer_Result) GetPropertyReturnCode() (value int32, err error) {
	retValue, err := instance.GetProperty("ReturnCode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
