// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_GlobalSaclSetting struct
type RSOP_GlobalSaclSetting struct {
	*RSOP_PolicySetting

	//
	SettingType string

	//
	SettingValue string
}

func NewRSOP_GlobalSaclSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_GlobalSaclSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_GlobalSaclSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_GlobalSaclSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_GlobalSaclSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_GlobalSaclSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetSettingType sets the value of SettingType for the instance
func (instance *RSOP_GlobalSaclSetting) SetPropertySettingType(value string) (err error) {
	return instance.SetProperty("SettingType", value)
}

// GetSettingType gets the value of SettingType for the instance
func (instance *RSOP_GlobalSaclSetting) GetPropertySettingType() (value string, err error) {
	retValue, err := instance.GetProperty("SettingType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *RSOP_GlobalSaclSetting) SetPropertySettingValue(value string) (err error) {
	return instance.SetProperty("SettingValue", value)
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *RSOP_GlobalSaclSetting) GetPropertySettingValue() (value string, err error) {
	retValue, err := instance.GetProperty("SettingValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
