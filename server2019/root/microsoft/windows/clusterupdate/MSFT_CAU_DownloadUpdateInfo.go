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

// MSFT_CAU_DownloadUpdateInfo struct
type MSFT_CAU_DownloadUpdateInfo struct {
	*MSFT_CAU_ScanUpdateInfo

	//
	UpdateErrorCode int32

	//
	UpdateResultCode uint32

	//
	UpdateTimestamp string
}

func NewMSFT_CAU_DownloadUpdateInfoEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_DownloadUpdateInfo, err error) {
	tmp, err := NewMSFT_CAU_ScanUpdateInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_DownloadUpdateInfo{
		MSFT_CAU_ScanUpdateInfo: tmp,
	}
	return
}

func NewMSFT_CAU_DownloadUpdateInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_DownloadUpdateInfo, err error) {
	tmp, err := NewMSFT_CAU_ScanUpdateInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_DownloadUpdateInfo{
		MSFT_CAU_ScanUpdateInfo: tmp,
	}
	return
}

// SetUpdateErrorCode sets the value of UpdateErrorCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) SetPropertyUpdateErrorCode(value int32) (err error) {
	return instance.SetProperty("UpdateErrorCode", value)
}

// GetUpdateErrorCode gets the value of UpdateErrorCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateErrorCode() (value int32, err error) {
	retValue, err := instance.GetProperty("UpdateErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateResultCode sets the value of UpdateResultCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) SetPropertyUpdateResultCode(value uint32) (err error) {
	return instance.SetProperty("UpdateResultCode", value)
}

// GetUpdateResultCode gets the value of UpdateResultCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateResultCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdateResultCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateTimestamp sets the value of UpdateTimestamp for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) SetPropertyUpdateTimestamp(value string) (err error) {
	return instance.SetProperty("UpdateTimestamp", value)
}

// GetUpdateTimestamp gets the value of UpdateTimestamp for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateTimestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
