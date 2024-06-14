// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS73FA2896_7854_48DF_9C91_9A864C4C97C7.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionDialUpSettingsLink struct
type RSOP_IEConnectionDialUpSettingsLink struct {
	*cim.WmiInstance

	//
	dialUpSettings RSOP_IEConnectionDialUpSettings

	//
	policySetting RSOP_IEAKPolicySetting
}

func NewRSOP_IEConnectionDialUpSettingsLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionDialUpSettingsLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpSettingsLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionDialUpSettingsLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionDialUpSettingsLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpSettingsLink{
		WmiInstance: tmp,
	}
	return
}

// SetdialUpSettings sets the value of dialUpSettings for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) SetPropertydialUpSettings(value RSOP_IEConnectionDialUpSettings) (err error) {
	return instance.SetProperty("dialUpSettings", (value))
}

// GetdialUpSettings gets the value of dialUpSettings for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) GetPropertydialUpSettings() (value RSOP_IEConnectionDialUpSettings, err error) {
	retValue, err := instance.GetProperty("dialUpSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEConnectionDialUpSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEConnectionDialUpSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEConnectionDialUpSettings(valuetmp)

	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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
