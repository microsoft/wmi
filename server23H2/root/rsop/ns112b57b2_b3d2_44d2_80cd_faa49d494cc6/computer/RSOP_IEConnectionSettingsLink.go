// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS112B57B2_B3D2_44D2_80CD_FAA49D494CC6.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionSettingsLink struct
type RSOP_IEConnectionSettingsLink struct {
	*cim.WmiInstance

	//
	connectionSettings RSOP_IEConnectionSettings

	//
	policySetting RSOP_IEAKPolicySetting
}

func NewRSOP_IEConnectionSettingsLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionSettingsLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionSettingsLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionSettingsLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionSettingsLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionSettingsLink{
		WmiInstance: tmp,
	}
	return
}

// SetconnectionSettings sets the value of connectionSettings for the instance
func (instance *RSOP_IEConnectionSettingsLink) SetPropertyconnectionSettings(value RSOP_IEConnectionSettings) (err error) {
	return instance.SetProperty("connectionSettings", (value))
}

// GetconnectionSettings gets the value of connectionSettings for the instance
func (instance *RSOP_IEConnectionSettingsLink) GetPropertyconnectionSettings() (value RSOP_IEConnectionSettings, err error) {
	retValue, err := instance.GetProperty("connectionSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEConnectionSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEConnectionSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEConnectionSettings(valuetmp)

	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionSettingsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionSettingsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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
