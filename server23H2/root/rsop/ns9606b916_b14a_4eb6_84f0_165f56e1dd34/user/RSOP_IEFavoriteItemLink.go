// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.User
//
// ////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEFavoriteItemLink struct
type RSOP_IEFavoriteItemLink struct {
	*cim.WmiInstance

	//
	favoriteItem RSOP_IEFavoriteItem

	//
	policySetting RSOP_IEAKPolicySetting
}

func NewRSOP_IEFavoriteItemLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEFavoriteItemLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteItemLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEFavoriteItemLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEFavoriteItemLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteItemLink{
		WmiInstance: tmp,
	}
	return
}

// SetfavoriteItem sets the value of favoriteItem for the instance
func (instance *RSOP_IEFavoriteItemLink) SetPropertyfavoriteItem(value RSOP_IEFavoriteItem) (err error) {
	return instance.SetProperty("favoriteItem", (value))
}

// GetfavoriteItem gets the value of favoriteItem for the instance
func (instance *RSOP_IEFavoriteItemLink) GetPropertyfavoriteItem() (value RSOP_IEFavoriteItem, err error) {
	retValue, err := instance.GetProperty("favoriteItem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEFavoriteItem)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEFavoriteItem is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEFavoriteItem(valuetmp)

	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEFavoriteItemLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEFavoriteItemLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEAKPolicySetting is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEAKPolicySetting(valuetmp)

	return
}
