// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetNatGlobal struct
type MSFT_NetNatGlobal struct {
	*MSFT_NetSettingData

	//
	InterRoutingDomainHairpinningMode uint32
}

func NewMSFT_NetNatGlobalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatGlobal, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatGlobal{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatGlobalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatGlobal, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatGlobal{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetInterRoutingDomainHairpinningMode sets the value of InterRoutingDomainHairpinningMode for the instance
func (instance *MSFT_NetNatGlobal) SetPropertyInterRoutingDomainHairpinningMode(value uint32) (err error) {
	return instance.SetProperty("InterRoutingDomainHairpinningMode", (value))
}

// GetInterRoutingDomainHairpinningMode gets the value of InterRoutingDomainHairpinningMode for the instance
func (instance *MSFT_NetNatGlobal) GetPropertyInterRoutingDomainHairpinningMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterRoutingDomainHairpinningMode")
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
