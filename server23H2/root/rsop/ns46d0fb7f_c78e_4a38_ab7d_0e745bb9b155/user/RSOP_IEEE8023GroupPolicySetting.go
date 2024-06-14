// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS46D0FB7F_C78E_4A38_AB7D_0E745BB9B155.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEEE8023GroupPolicySetting struct
type RSOP_IEEE8023GroupPolicySetting struct {
	*RSOP_PolicySetting

	//
	description string

	//
	msieee8023PolicyData string

	//
	msieee8023PolicyReserved []uint8

	//
	whenChanged uint32
}

func NewRSOP_IEEE8023GroupPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEEE8023GroupPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE8023GroupPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_IEEE8023GroupPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEEE8023GroupPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE8023GroupPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// Setdescription sets the value of description for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", (value))
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertydescription() (value string, err error) {
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

// Setmsieee8023PolicyData sets the value of msieee8023PolicyData for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertymsieee8023PolicyData(value string) (err error) {
	return instance.SetProperty("msieee8023PolicyData", (value))
}

// Getmsieee8023PolicyData gets the value of msieee8023PolicyData for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertymsieee8023PolicyData() (value string, err error) {
	retValue, err := instance.GetProperty("msieee8023PolicyData")
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

// Setmsieee8023PolicyReserved sets the value of msieee8023PolicyReserved for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertymsieee8023PolicyReserved(value []uint8) (err error) {
	return instance.SetProperty("msieee8023PolicyReserved", (value))
}

// Getmsieee8023PolicyReserved gets the value of msieee8023PolicyReserved for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertymsieee8023PolicyReserved() (value []uint8, err error) {
	retValue, err := instance.GetProperty("msieee8023PolicyReserved")
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
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertywhenChanged(value uint32) (err error) {
	return instance.SetProperty("whenChanged", (value))
}

// GetwhenChanged gets the value of whenChanged for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertywhenChanged() (value uint32, err error) {
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
