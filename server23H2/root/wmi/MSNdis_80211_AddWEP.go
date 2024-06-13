// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_AddWEP struct
type MSNdis_80211_AddWEP struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	KeyIndex uint32

	//
	KeyLength uint32

	//
	KeyMaterial []uint8

	//
	Length uint32
}

func NewMSNdis_80211_AddWEPEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_AddWEP, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_AddWEP{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_AddWEPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_AddWEP, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_AddWEP{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetKeyIndex sets the value of KeyIndex for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyKeyIndex(value uint32) (err error) {
	return instance.SetProperty("KeyIndex", (value))
}

// GetKeyIndex gets the value of KeyIndex for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyKeyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeyIndex")
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

// SetKeyLength sets the value of KeyLength for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyKeyLength(value uint32) (err error) {
	return instance.SetProperty("KeyLength", (value))
}

// GetKeyLength gets the value of KeyLength for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyKeyLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeyLength")
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

// SetKeyMaterial sets the value of KeyMaterial for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyKeyMaterial(value []uint8) (err error) {
	return instance.SetProperty("KeyMaterial", (value))
}

// GetKeyMaterial gets the value of KeyMaterial for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyKeyMaterial() (value []uint8, err error) {
	retValue, err := instance.GetProperty("KeyMaterial")
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

// SetLength sets the value of Length for the instance
func (instance *MSNdis_80211_AddWEP) SetPropertyLength(value uint32) (err error) {
	return instance.SetProperty("Length", (value))
}

// GetLength gets the value of Length for the instance
func (instance *MSNdis_80211_AddWEP) GetPropertyLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("Length")
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
