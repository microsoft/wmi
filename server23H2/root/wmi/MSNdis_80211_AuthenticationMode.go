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

// MSNdis_80211_AuthenticationMode struct
type MSNdis_80211_AuthenticationMode struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211AuthenticationMode AuthenticationMode_Ndis80211AuthenticationMode
}

func NewMSNdis_80211_AuthenticationModeEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_AuthenticationMode, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_AuthenticationMode{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_AuthenticationModeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_AuthenticationMode, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_AuthenticationMode{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_AuthenticationMode) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_AuthenticationMode) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_AuthenticationMode) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_AuthenticationMode) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211AuthenticationMode sets the value of Ndis80211AuthenticationMode for the instance
func (instance *MSNdis_80211_AuthenticationMode) SetPropertyNdis80211AuthenticationMode(value AuthenticationMode_Ndis80211AuthenticationMode) (err error) {
	return instance.SetProperty("Ndis80211AuthenticationMode", (value))
}

// GetNdis80211AuthenticationMode gets the value of Ndis80211AuthenticationMode for the instance
func (instance *MSNdis_80211_AuthenticationMode) GetPropertyNdis80211AuthenticationMode() (value AuthenticationMode_Ndis80211AuthenticationMode, err error) {
	retValue, err := instance.GetProperty("Ndis80211AuthenticationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AuthenticationMode_Ndis80211AuthenticationMode(valuetmp)

	return
}
