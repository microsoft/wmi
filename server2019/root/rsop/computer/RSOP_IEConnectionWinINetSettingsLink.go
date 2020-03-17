// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEConnectionWinINetSettingsLink struct
type RSOP_IEConnectionWinINetSettingsLink struct {
	cim.WmiInstance

	//
	policySetting RSOP_IEAKPolicySetting

	//
	winINetSettings RSOP_IEConnectionWinINetSettings
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", value)
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetwinINetSettings sets the value of winINetSettings for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) SetPropertywinINetSettings(value RSOP_IEConnectionWinINetSettings) (err error) {
	return instance.SetProperty("winINetSettings", value)
}

// GetwinINetSettings gets the value of winINetSettings for the instance
func (instance *RSOP_IEConnectionWinINetSettingsLink) GetPropertywinINetSettings() (value RSOP_IEConnectionWinINetSettings, err error) {
	retValue, err := instance.GetProperty("winINetSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEConnectionWinINetSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}
