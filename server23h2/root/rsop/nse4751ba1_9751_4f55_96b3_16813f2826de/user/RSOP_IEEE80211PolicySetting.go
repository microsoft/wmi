// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSE4751BA1_9751_4F55_96B3_16813F2826DE.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEEE80211PolicySetting struct
type RSOP_IEEE80211PolicySetting struct {
	*RSOP_PolicySetting

	//
	ClassName string

	//
	description string

	//
	msieee80211Data []uint8

	//
	msieee80211DataType uint32

	//
	msieee80211ID string

	//
	msieee80211Name string

	//
	whenChanged uint32
}

func NewRSOP_IEEE80211PolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEEE80211PolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE80211PolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_IEEE80211PolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEEE80211PolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEEE80211PolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetClassName sets the value of ClassName for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertyClassName(value string) (err error) {
	return instance.SetProperty("ClassName", (value))
}

// GetClassName gets the value of ClassName for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertyClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ClassName")
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

// Setdescription sets the value of description for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", (value))
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertydescription() (value string, err error) {
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

// Setmsieee80211Data sets the value of msieee80211Data for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertymsieee80211Data(value []uint8) (err error) {
	return instance.SetProperty("msieee80211Data", (value))
}

// Getmsieee80211Data gets the value of msieee80211Data for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertymsieee80211Data() (value []uint8, err error) {
	retValue, err := instance.GetProperty("msieee80211Data")
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

// Setmsieee80211DataType sets the value of msieee80211DataType for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertymsieee80211DataType(value uint32) (err error) {
	return instance.SetProperty("msieee80211DataType", (value))
}

// Getmsieee80211DataType gets the value of msieee80211DataType for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertymsieee80211DataType() (value uint32, err error) {
	retValue, err := instance.GetProperty("msieee80211DataType")
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

// Setmsieee80211ID sets the value of msieee80211ID for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertymsieee80211ID(value string) (err error) {
	return instance.SetProperty("msieee80211ID", (value))
}

// Getmsieee80211ID gets the value of msieee80211ID for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertymsieee80211ID() (value string, err error) {
	retValue, err := instance.GetProperty("msieee80211ID")
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

// Setmsieee80211Name sets the value of msieee80211Name for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertymsieee80211Name(value string) (err error) {
	return instance.SetProperty("msieee80211Name", (value))
}

// Getmsieee80211Name gets the value of msieee80211Name for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertymsieee80211Name() (value string, err error) {
	retValue, err := instance.GetProperty("msieee80211Name")
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

// SetwhenChanged sets the value of whenChanged for the instance
func (instance *RSOP_IEEE80211PolicySetting) SetPropertywhenChanged(value uint32) (err error) {
	return instance.SetProperty("whenChanged", (value))
}

// GetwhenChanged gets the value of whenChanged for the instance
func (instance *RSOP_IEEE80211PolicySetting) GetPropertywhenChanged() (value uint32, err error) {
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
