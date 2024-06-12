// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSBEA911FB_A2AD_4EA4_BB08_1252FC88818C.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEEE80211GroupPolicySetting struct
type RSOP_IEEE80211GroupPolicySetting struct {
	*RSOP_PolicySetting

	//
	description string

	//
	msieee80211PolicyData string

	//
	msieee80211PolicyReserved []uint8

	//
	whenChanged uint32
}

func NewRSOP_IEEE80211GroupPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEEE80211GroupPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE80211GroupPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_IEEE80211GroupPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEEE80211GroupPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE80211GroupPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// Setdescription sets the value of description for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", (value))
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) GetPropertydescription() (value string, err error) {
	retValue, err := instance.GetProperty("description")
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

// Setmsieee80211PolicyData sets the value of msieee80211PolicyData for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) SetPropertymsieee80211PolicyData(value string) (err error) {
	return instance.SetProperty("msieee80211PolicyData", (value))
}

// Getmsieee80211PolicyData gets the value of msieee80211PolicyData for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) GetPropertymsieee80211PolicyData() (value string, err error) {
	retValue, err := instance.GetProperty("msieee80211PolicyData")
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

// Setmsieee80211PolicyReserved sets the value of msieee80211PolicyReserved for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) SetPropertymsieee80211PolicyReserved(value []uint8) (err error) {
	return instance.SetProperty("msieee80211PolicyReserved", (value))
}

// Getmsieee80211PolicyReserved gets the value of msieee80211PolicyReserved for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) GetPropertymsieee80211PolicyReserved() (value []uint8, err error) {
	retValue, err := instance.GetProperty("msieee80211PolicyReserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetwhenChanged sets the value of whenChanged for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) SetPropertywhenChanged(value uint32) (err error) {
	return instance.SetProperty("whenChanged", (value))
}

// GetwhenChanged gets the value of whenChanged for the instance
func (instance *RSOP_IEEE80211GroupPolicySetting) GetPropertywhenChanged() (value uint32, err error) {
	retValue, err := instance.GetProperty("whenChanged")
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
