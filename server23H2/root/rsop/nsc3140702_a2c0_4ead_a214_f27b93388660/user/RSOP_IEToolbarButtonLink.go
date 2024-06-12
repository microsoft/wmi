// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSC3140702_A2C0_4EAD_A214_F27B93388660.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEToolbarButtonLink struct
type RSOP_IEToolbarButtonLink struct {
	*cim.WmiInstance

	//
	policySetting RSOP_IEAKPolicySetting

	//
	toolbarButton RSOP_IEToolbarButton
}

func NewRSOP_IEToolbarButtonLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEToolbarButtonLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEToolbarButtonLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEToolbarButtonLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEToolbarButtonLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEToolbarButtonLink{
		WmiInstance: tmp,
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEToolbarButtonLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEToolbarButtonLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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

// SettoolbarButton sets the value of toolbarButton for the instance
func (instance *RSOP_IEToolbarButtonLink) SetPropertytoolbarButton(value RSOP_IEToolbarButton) (err error) {
	return instance.SetProperty("toolbarButton", (value))
}

// GettoolbarButton gets the value of toolbarButton for the instance
func (instance *RSOP_IEToolbarButtonLink) GetPropertytoolbarButton() (value RSOP_IEToolbarButton, err error) {
	retValue, err := instance.GetProperty("toolbarButton")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEToolbarButton)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEToolbarButton is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEToolbarButton(valuetmp)

	return
}
