// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("UpdateErrorCode", (value))
}

// GetUpdateErrorCode gets the value of UpdateErrorCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateErrorCode() (value int32, err error) {
	retValue, err := instance.GetProperty("UpdateErrorCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetUpdateResultCode sets the value of UpdateResultCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) SetPropertyUpdateResultCode(value uint32) (err error) {
	return instance.SetProperty("UpdateResultCode", (value))
}

// GetUpdateResultCode gets the value of UpdateResultCode for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateResultCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdateResultCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUpdateTimestamp sets the value of UpdateTimestamp for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) SetPropertyUpdateTimestamp(value string) (err error) {
	return instance.SetProperty("UpdateTimestamp", (value))
}

// GetUpdateTimestamp gets the value of UpdateTimestamp for the instance
func (instance *MSFT_CAU_DownloadUpdateInfo) GetPropertyUpdateTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateTimestamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
