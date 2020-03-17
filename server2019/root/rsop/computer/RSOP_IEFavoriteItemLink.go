// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEFavoriteItemLink struct
type RSOP_IEFavoriteItemLink struct {
	cim.WmiInstance

	//
	favoriteItem RSOP_IEFavoriteItem

	//
	policySetting RSOP_IEAKPolicySetting
}

// SetfavoriteItem sets the value of favoriteItem for the instance
func (instance *RSOP_IEFavoriteItemLink) SetPropertyfavoriteItem(value RSOP_IEFavoriteItem) (err error) {
	return instance.SetProperty("favoriteItem", value)
}

// GetfavoriteItem gets the value of favoriteItem for the instance
func (instance *RSOP_IEFavoriteItemLink) GetPropertyfavoriteItem() (value RSOP_IEFavoriteItem, err error) {
	retValue, err := instance.GetProperty("favoriteItem")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEFavoriteItem)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEFavoriteItemLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", value)
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEFavoriteItemLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}
