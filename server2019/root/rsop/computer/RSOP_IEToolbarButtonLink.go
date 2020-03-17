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

// RSOP_IEToolbarButtonLink struct
type RSOP_IEToolbarButtonLink struct {
	cim.WmiInstance

	//
	policySetting RSOP_IEAKPolicySetting

	//
	toolbarButton RSOP_IEToolbarButton
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEToolbarButtonLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", value)
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEToolbarButtonLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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

// SettoolbarButton sets the value of toolbarButton for the instance
func (instance *RSOP_IEToolbarButtonLink) SetPropertytoolbarButton(value RSOP_IEToolbarButton) (err error) {
	return instance.SetProperty("toolbarButton", value)
}

// GettoolbarButton gets the value of toolbarButton for the instance
func (instance *RSOP_IEToolbarButtonLink) GetPropertytoolbarButton() (value RSOP_IEToolbarButton, err error) {
	retValue, err := instance.GetProperty("toolbarButton")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEToolbarButton)
	if !ok {
		// TODO: Set an error
	}
	return
}
