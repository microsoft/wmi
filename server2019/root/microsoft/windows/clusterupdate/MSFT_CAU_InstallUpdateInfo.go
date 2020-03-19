// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CAU_InstallUpdateInfo struct
type MSFT_CAU_InstallUpdateInfo struct {
	*MSFT_CAU_DownloadUpdateInfo

	//
	LongRebootHint bool

	//
	RebootRequired bool
}

func NewMSFT_CAU_InstallUpdateInfoEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_InstallUpdateInfo, err error) {
	tmp, err := NewMSFT_CAU_DownloadUpdateInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_InstallUpdateInfo{
		MSFT_CAU_DownloadUpdateInfo: tmp,
	}
	return
}

func NewMSFT_CAU_InstallUpdateInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_InstallUpdateInfo, err error) {
	tmp, err := NewMSFT_CAU_DownloadUpdateInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_InstallUpdateInfo{
		MSFT_CAU_DownloadUpdateInfo: tmp,
	}
	return
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
