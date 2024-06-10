// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D.User
//
// ////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSoP_PolicySettingLink struct
type RSoP_PolicySettingLink struct {
	*cim.WmiInstance

	//
	setting RSOP_PolicySetting

	//
	status RSoP_PolicySettingStatus
}

func NewRSoP_PolicySettingLinkEx1(instance *cim.WmiInstance) (newInstance *RSoP_PolicySettingLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSoP_PolicySettingLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSoP_PolicySettingLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSoP_PolicySettingLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSoP_PolicySettingLink{
		WmiInstance: tmp,
	}
	return
}

// Setsetting sets the value of setting for the instance
func (instance *RSoP_PolicySettingLink) SetPropertysetting(value RSOP_PolicySetting) (err error) {
	return instance.SetProperty("setting", (value))
}

// Getsetting gets the value of setting for the instance
func (instance *RSoP_PolicySettingLink) GetPropertysetting() (value RSOP_PolicySetting, err error) {
	retValue, err := instance.GetProperty("setting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_PolicySetting)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_PolicySetting is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_PolicySetting(valuetmp)

	return
}

// Setstatus sets the value of status for the instance
func (instance *RSoP_PolicySettingLink) SetPropertystatus(value RSoP_PolicySettingStatus) (err error) {
	return instance.SetProperty("status", (value))
}

// Getstatus gets the value of status for the instance
func (instance *RSoP_PolicySettingLink) GetPropertystatus() (value RSoP_PolicySettingStatus, err error) {
	retValue, err := instance.GetProperty("status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSoP_PolicySettingStatus)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSoP_PolicySettingStatus is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSoP_PolicySettingStatus(valuetmp)

	return
}
