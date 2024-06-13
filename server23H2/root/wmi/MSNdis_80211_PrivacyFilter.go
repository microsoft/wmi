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

// MSNdis_80211_PrivacyFilter struct
type MSNdis_80211_PrivacyFilter struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211PrivacyFilter PrivacyFilter_Ndis80211PrivacyFilter
}

func NewMSNdis_80211_PrivacyFilterEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_PrivacyFilter, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_PrivacyFilter{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_PrivacyFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_PrivacyFilter, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_PrivacyFilter{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_PrivacyFilter) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_PrivacyFilter) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_PrivacyFilter) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_PrivacyFilter) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211PrivacyFilter sets the value of Ndis80211PrivacyFilter for the instance
func (instance *MSNdis_80211_PrivacyFilter) SetPropertyNdis80211PrivacyFilter(value PrivacyFilter_Ndis80211PrivacyFilter) (err error) {
	return instance.SetProperty("Ndis80211PrivacyFilter", (value))
}

// GetNdis80211PrivacyFilter gets the value of Ndis80211PrivacyFilter for the instance
func (instance *MSNdis_80211_PrivacyFilter) GetPropertyNdis80211PrivacyFilter() (value PrivacyFilter_Ndis80211PrivacyFilter, err error) {
	retValue, err := instance.GetProperty("Ndis80211PrivacyFilter")
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

	value = PrivacyFilter_Ndis80211PrivacyFilter(valuetmp)

	return
}
