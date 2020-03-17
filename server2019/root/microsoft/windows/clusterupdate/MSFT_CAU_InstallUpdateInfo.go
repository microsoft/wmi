// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

// MSFT_CAU_InstallUpdateInfo struct
type MSFT_CAU_InstallUpdateInfo struct {
	MSFT_CAU_DownloadUpdateInfo

	//
	LongRebootHint bool

	//
	RebootRequired bool
}

// SetLongRebootHint sets the value of LongRebootHint for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) SetPropertyLongRebootHint(value bool) (err error) {
	return instance.SetProperty("LongRebootHint", value)
}

// GetLongRebootHint gets the value of LongRebootHint for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) GetPropertyLongRebootHint() (value bool, err error) {
	retValue, err := instance.GetProperty("LongRebootHint")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) SetPropertyRebootRequired(value bool) (err error) {
	return instance.SetProperty("RebootRequired", value)
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) GetPropertyRebootRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
