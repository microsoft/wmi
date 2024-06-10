// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_NetworkType struct
type MSNdis_80211_NetworkType struct {
	*MSNdis

	//
	Ndis80211NetworkType uint32
}

func NewMSNdis_80211_NetworkTypeEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_NetworkType, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_NetworkType{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_NetworkTypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_NetworkType, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_NetworkType{
		MSNdis: tmp,
	}
	return
}

// SetNdis80211NetworkType sets the value of Ndis80211NetworkType for the instance
func (instance *MSNdis_80211_NetworkType) SetPropertyNdis80211NetworkType(value uint32) (err error) {
	return instance.SetProperty("Ndis80211NetworkType", (value))
}

// GetNdis80211NetworkType gets the value of Ndis80211NetworkType for the instance
func (instance *MSNdis_80211_NetworkType) GetPropertyNdis80211NetworkType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Ndis80211NetworkType")
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
