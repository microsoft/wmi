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

// MSFT_CAU_ScanUpdateInfo struct
type MSFT_CAU_ScanUpdateInfo struct {
	cim.WmiInstance

	//
	UpdateDesc string

	//
	UpdateId string

	//
	UpdateTitle string
}

// SetUpdateDesc sets the value of UpdateDesc for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateDesc(value string) (err error) {
	return instance.SetProperty("UpdateDesc", value)
}

// GetUpdateDesc gets the value of UpdateDesc for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateDesc() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateDesc")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateId sets the value of UpdateId for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateId(value string) (err error) {
	return instance.SetProperty("UpdateId", value)
}

// GetUpdateId gets the value of UpdateId for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateId() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateTitle sets the value of UpdateTitle for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateTitle(value string) (err error) {
	return instance.SetProperty("UpdateTitle", value)
}

// GetUpdateTitle gets the value of UpdateTitle for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateTitle() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateTitle")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
