// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS5B3DCE03_EB70_4EA2_8019_E0415C3C4C7B.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionWinINetSettingsLink struct
type RSOP_IEConnectionWinINetSettingsLink struct {
	*cim.WmiInstance

	//
	policySetting RSOP_IEAKPolicySetting

	//
	winINetSettings RSOP_IEConnectionWinINetSettings
}

func NewRSOP_IEConnectionWinINetSettingsLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionWinINetSettingsLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionWinINetSettingsLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionWinINetSettingsLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionWinINetSettingsLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionWinINetSettingsLink{
		WmiInstance: tmp,
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEAKPolicySetting is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEAKPolicySetting(valuetmp)

	return
}

// SetwinINetSettings sets the value of winINetSettings for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) SetPropertywinINetSettings(value RSOP_IEConnectionWinINetSettings) (err error) {
	return instance.SetProperty("winINetSettings", (value))
}

// GetwinINetSettings gets the value of winINetSettings for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) GetPropertywinINetSettings() (value RSOP_IEConnectionWinINetSettings, err error) {
	retValue, err := instance.GetProperty("winINetSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEConnectionWinINetSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEConnectionWinINetSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEConnectionWinINetSettings(valuetmp)

	return
}
