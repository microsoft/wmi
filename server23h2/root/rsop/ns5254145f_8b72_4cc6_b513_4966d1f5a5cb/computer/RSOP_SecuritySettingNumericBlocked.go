// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS5254145F_8B72_4CC6_B513_4966D1F5A5CB.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_SecuritySettingNumericBlocked struct
type RSOP_SecuritySettingNumericBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	KeyName string

	//
	Setting uint32
}

func NewRSOP_SecuritySettingNumericBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_SecuritySettingNumericBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingNumericBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_SecuritySettingNumericBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SecuritySettingNumericBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingNumericBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingNumericBlocked) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", (value))
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingNumericBlocked) GetPropertyKeyName() (value string, err error) {
	retValue, err := instance.GetProperty("KeyName")
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

// SetSetting sets the value of Setting for the instance
func (instance *RSOP_SecuritySettingNumericBlocked) SetPropertySetting(value uint32) (err error) {
	return instance.SetProperty("Setting", (value))
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecuritySettingNumericBlocked) GetPropertySetting() (value uint32, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
