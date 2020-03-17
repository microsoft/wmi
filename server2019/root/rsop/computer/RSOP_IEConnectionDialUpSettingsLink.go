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

// RSOP_IEConnectionDialUpSettingsLink struct
type RSOP_IEConnectionDialUpSettingsLink struct {
	cim.WmiInstance

	//
	dialUpSettings RSOP_IEConnectionDialUpSettings

	//
	policySetting RSOP_IEAKPolicySetting
}

// SetdialUpSettings sets the value of dialUpSettings for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) SetPropertydialUpSettings(value RSOP_IEConnectionDialUpSettings) (err error) {
	return instance.SetProperty("dialUpSettings", value)
}

// GetdialUpSettings gets the value of dialUpSettings for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) GetPropertydialUpSettings() (value RSOP_IEConnectionDialUpSettings, err error) {
	retValue, err := instance.GetProperty("dialUpSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEConnectionDialUpSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", value)
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpSettingsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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
