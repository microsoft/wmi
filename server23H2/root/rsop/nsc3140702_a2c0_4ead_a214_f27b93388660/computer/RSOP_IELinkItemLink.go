// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSC3140702_A2C0_4EAD_A214_F27B93388660.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IELinkItemLink struct
type RSOP_IELinkItemLink struct {
	*cim.WmiInstance

	//
	linkItem RSOP_IELinkItem

	//
	policySetting RSOP_IEAKPolicySetting
}

func NewRSOP_IELinkItemLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IELinkItemLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItemLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IELinkItemLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IELinkItemLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItemLink{
		WmiInstance: tmp,
	}
	return
}

// SetlinkItem sets the value of linkItem for the instance
func (instance *RSOP_IELinkItemLink) SetPropertylinkItem(value RSOP_IELinkItem) (err error) {
	return instance.SetProperty("linkItem", (value))
}

// GetlinkItem gets the value of linkItem for the instance
func (instance *RSOP_IELinkItemLink) GetPropertylinkItem() (value RSOP_IELinkItem, err error) {
	retValue, err := instance.GetProperty("linkItem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IELinkItem)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IELinkItem is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IELinkItem(valuetmp)

	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IELinkItemLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IELinkItemLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
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
