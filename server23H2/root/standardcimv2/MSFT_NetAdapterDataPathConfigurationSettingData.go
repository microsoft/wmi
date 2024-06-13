// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterDataPathConfigurationSettingData struct
type MSFT_NetAdapterDataPathConfigurationSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	Profile string

	//
	ProfileSource uint32
}

func NewMSFT_NetAdapterDataPathConfigurationSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterDataPathConfigurationSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterDataPathConfigurationSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterDataPathConfigurationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterDataPathConfigurationSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterDataPathConfigurationSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetAdapterDataPathConfigurationSettingData) SetPropertyProfile(value string) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetAdapterDataPathConfigurationSettingData) GetPropertyProfile() (value string, err error) {
	retValue, err := instance.GetProperty("Profile")
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

// SetProfileSource sets the value of ProfileSource for the instance
func (instance *MSFT_NetAdapterDataPathConfigurationSettingData) SetPropertyProfileSource(value uint32) (err error) {
	return instance.SetProperty("ProfileSource", (value))
}

// GetProfileSource gets the value of ProfileSource for the instance
func (instance *MSFT_NetAdapterDataPathConfigurationSettingData) GetPropertyProfileSource() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProfileSource")
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
