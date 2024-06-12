// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_NetworkTypeInUse struct
type MSNdis_80211_NetworkTypeInUse struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	Ndis80211NetworkTypeInUse MSNdis_80211_NetworkType
}

func NewMSNdis_80211_NetworkTypeInUseEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_NetworkTypeInUse, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_NetworkTypeInUse{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_NetworkTypeInUseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_NetworkTypeInUse, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_NetworkTypeInUse{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_NetworkTypeInUse) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_NetworkTypeInUse) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_80211_NetworkTypeInUse) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_NetworkTypeInUse) GetPropertyInstanceName() (value string, err error) {
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

// SetNdis80211NetworkTypeInUse sets the value of Ndis80211NetworkTypeInUse for the instance
func (instance *MSNdis_80211_NetworkTypeInUse) SetPropertyNdis80211NetworkTypeInUse(value MSNdis_80211_NetworkType) (err error) {
	return instance.SetProperty("Ndis80211NetworkTypeInUse", (value))
}

// GetNdis80211NetworkTypeInUse gets the value of Ndis80211NetworkTypeInUse for the instance
func (instance *MSNdis_80211_NetworkTypeInUse) GetPropertyNdis80211NetworkTypeInUse() (value MSNdis_80211_NetworkType, err error) {
	retValue, err := instance.GetProperty("Ndis80211NetworkTypeInUse")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_80211_NetworkType)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_NetworkType is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_80211_NetworkType(valuetmp)

	return
}
