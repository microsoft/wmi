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

// MSFT_NetISATAPConfiguration struct
type MSFT_NetISATAPConfiguration struct {
	*MSFT_NetSettingData

	//
	PolicyStore string

	//
	ResolutionInterval uint32

	//
	ResolutionState uint32

	//
	Router string

	//
	State uint32
}

func NewMSFT_NetISATAPConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetISATAPConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetISATAPConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetISATAPConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetISATAPConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetISATAPConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_NetISATAPConfiguration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", (value))
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NetISATAPConfiguration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
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

// SetResolutionInterval sets the value of ResolutionInterval for the instance
func (instance *MSFT_NetISATAPConfiguration) SetPropertyResolutionInterval(value uint32) (err error) {
	return instance.SetProperty("ResolutionInterval", (value))
}

// GetResolutionInterval gets the value of ResolutionInterval for the instance
func (instance *MSFT_NetISATAPConfiguration) GetPropertyResolutionInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResolutionInterval")
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

// SetResolutionState sets the value of ResolutionState for the instance
func (instance *MSFT_NetISATAPConfiguration) SetPropertyResolutionState(value uint32) (err error) {
	return instance.SetProperty("ResolutionState", (value))
}

// GetResolutionState gets the value of ResolutionState for the instance
func (instance *MSFT_NetISATAPConfiguration) GetPropertyResolutionState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResolutionState")
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

// SetRouter sets the value of Router for the instance
func (instance *MSFT_NetISATAPConfiguration) SetPropertyRouter(value string) (err error) {
	return instance.SetProperty("Router", (value))
}

// GetRouter gets the value of Router for the instance
func (instance *MSFT_NetISATAPConfiguration) GetPropertyRouter() (value string, err error) {
	retValue, err := instance.GetProperty("Router")
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

// SetState sets the value of State for the instance
func (instance *MSFT_NetISATAPConfiguration) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetISATAPConfiguration) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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

//

// <param name="PassThru" type="bool "></param>
// <param name="ResolutionInterval" type="bool "></param>
// <param name="ResolutionState" type="bool "></param>
// <param name="Router" type="bool "></param>
// <param name="State" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetISATAPConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetISATAPConfiguration) Reset( /* IN */ State bool,
	/* IN */ Router bool,
	/* IN */ ResolutionState bool,
	/* IN */ ResolutionInterval bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetISATAPConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", State, Router, ResolutionState, ResolutionInterval, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
