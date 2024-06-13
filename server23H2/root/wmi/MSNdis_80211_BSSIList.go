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

// MSNdis_80211_BSSIList struct
type MSNdis_80211_BSSIList struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211BSSIList []MSNdis_80211_WLanBssId

	//
	NumberOfItems uint32
}

func NewMSNdis_80211_BSSIListEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_BSSIList, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_BSSIList{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_BSSIListEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_BSSIList, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_BSSIList{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_BSSIList) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_BSSIList) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_BSSIList) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_BSSIList) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211BSSIList sets the value of Ndis80211BSSIList for the instance
func (instance *MSNdis_80211_BSSIList) SetPropertyNdis80211BSSIList(value []MSNdis_80211_WLanBssId) (err error) {
	return instance.SetProperty("Ndis80211BSSIList", (value))
}

// GetNdis80211BSSIList gets the value of Ndis80211BSSIList for the instance
func (instance *MSNdis_80211_BSSIList) GetPropertyNdis80211BSSIList() (value []MSNdis_80211_WLanBssId, err error) {
	retValue, err := instance.GetProperty("Ndis80211BSSIList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSNdis_80211_WLanBssId)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_WLanBssId is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSNdis_80211_WLanBssId(valuetmp))
	}

	return
}

// SetNumberOfItems sets the value of NumberOfItems for the instance
func (instance *MSNdis_80211_BSSIList) SetPropertyNumberOfItems(value uint32) (err error) {
	return instance.SetProperty("NumberOfItems", (value))
}

// GetNumberOfItems gets the value of NumberOfItems for the instance
func (instance *MSNdis_80211_BSSIList) GetPropertyNumberOfItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfItems")
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
