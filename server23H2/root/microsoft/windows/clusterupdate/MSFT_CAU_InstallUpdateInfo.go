// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAU_InstallUpdateInfo struct
type MSFT_CAU_InstallUpdateInfo struct {
	*MSFT_CAU_DownloadUpdateInfo

	//
	CommitRequired bool

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

// SetCommitRequired sets the value of CommitRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) SetPropertyCommitRequired(value bool) (err error) {
	return instance.SetProperty("CommitRequired", (value))
}

// GetCommitRequired gets the value of CommitRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) GetPropertyCommitRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("CommitRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLongRebootHint sets the value of LongRebootHint for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) SetPropertyLongRebootHint(value bool) (err error) {
	return instance.SetProperty("LongRebootHint", (value))
}

// GetLongRebootHint gets the value of LongRebootHint for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) GetPropertyLongRebootHint() (value bool, err error) {
	retValue, err := instance.GetProperty("LongRebootHint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) SetPropertyRebootRequired(value bool) (err error) {
	return instance.SetProperty("RebootRequired", (value))
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MSFT_CAU_InstallUpdateInfo) GetPropertyRebootRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
