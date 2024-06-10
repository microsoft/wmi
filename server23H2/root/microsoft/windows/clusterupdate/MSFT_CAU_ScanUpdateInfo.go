// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ClusterUpdate
//
// ////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAU_ScanUpdateInfo struct
type MSFT_CAU_ScanUpdateInfo struct {
	*cim.WmiInstance

	//
	UpdateDesc string

	//
	UpdateId string

	//
	UpdateTitle string
}

func NewMSFT_CAU_ScanUpdateInfoEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_ScanUpdateInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_ScanUpdateInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAU_ScanUpdateInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_ScanUpdateInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_ScanUpdateInfo{
		WmiInstance: tmp,
	}
	return
}

// SetUpdateDesc sets the value of UpdateDesc for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateDesc(value string) (err error) {
	return instance.SetProperty("UpdateDesc", (value))
}

// GetUpdateDesc gets the value of UpdateDesc for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateDesc() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateDesc")
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

// SetUpdateId sets the value of UpdateId for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateId(value string) (err error) {
	return instance.SetProperty("UpdateId", (value))
}

// GetUpdateId gets the value of UpdateId for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateId() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateId")
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

// SetUpdateTitle sets the value of UpdateTitle for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) SetPropertyUpdateTitle(value string) (err error) {
	return instance.SetProperty("UpdateTitle", (value))
}

// GetUpdateTitle gets the value of UpdateTitle for the instance
func (instance *MSFT_CAU_ScanUpdateInfo) GetPropertyUpdateTitle() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateTitle")
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
