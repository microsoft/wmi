// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_BaseServiceSetIdentifier struct
type MSNdis_80211_BaseServiceSetIdentifier struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211MacAddress []uint8
}

func NewMSNdis_80211_BaseServiceSetIdentifierEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_BaseServiceSetIdentifier, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_BaseServiceSetIdentifier{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_BaseServiceSetIdentifierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_BaseServiceSetIdentifier, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_BaseServiceSetIdentifier{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_BaseServiceSetIdentifier) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_BaseServiceSetIdentifier) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_BaseServiceSetIdentifier) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_BaseServiceSetIdentifier) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211MacAddress sets the value of Ndis80211MacAddress for the instance
func (instance *MSNdis_80211_BaseServiceSetIdentifier) SetPropertyNdis80211MacAddress(value []uint8) (err error) {
	return instance.SetProperty("Ndis80211MacAddress", (value))
}

// GetNdis80211MacAddress gets the value of Ndis80211MacAddress for the instance
func (instance *MSNdis_80211_BaseServiceSetIdentifier) GetPropertyNdis80211MacAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ndis80211MacAddress")
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
