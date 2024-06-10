// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetVirtualizationGlobalSettingData struct
type MSFT_NetVirtualizationGlobalSettingData struct {
	*MSFT_NetSettingData

	// 26
	UseExternalRouter bool
}

func NewMSFT_NetVirtualizationGlobalSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationGlobalSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationGlobalSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationGlobalSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationGlobalSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationGlobalSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetUseExternalRouter sets the value of UseExternalRouter for the instance
func (instance *MSFT_NetVirtualizationGlobalSettingData) SetPropertyUseExternalRouter(value bool) (err error) {
	return instance.SetProperty("UseExternalRouter", (value))
}

// GetUseExternalRouter gets the value of UseExternalRouter for the instance
func (instance *MSFT_NetVirtualizationGlobalSettingData) GetPropertyUseExternalRouter() (value bool, err error) {
	retValue, err := instance.GetProperty("UseExternalRouter")
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
