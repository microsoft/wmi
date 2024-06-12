// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetQosTrafficClassSettingData struct
type MSFT_NetQosTrafficClassSettingData struct {
	*MSFT_NetSettingData

	//
	Algorithm uint8

	//
	BandwidthPercentage uint8

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	Name string

	//
	PolicySet uint8

	//
	Priority []uint8
}

func NewMSFT_NetQosTrafficClassSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetQosTrafficClassSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQosTrafficClassSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetQosTrafficClassSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetQosTrafficClassSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQosTrafficClassSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAlgorithm sets the value of Algorithm for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyAlgorithm(value uint8) (err error) {
	return instance.SetProperty("Algorithm", (value))
}

// GetAlgorithm gets the value of Algorithm for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyAlgorithm() (value uint8, err error) {
	retValue, err := instance.GetProperty("Algorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetBandwidthPercentage sets the value of BandwidthPercentage for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyBandwidthPercentage(value uint8) (err error) {
	return instance.SetProperty("BandwidthPercentage", (value))
}

// GetBandwidthPercentage gets the value of BandwidthPercentage for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyBandwidthPercentage() (value uint8, err error) {
	retValue, err := instance.GetProperty("BandwidthPercentage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetPolicySet sets the value of PolicySet for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyPolicySet(value uint8) (err error) {
	return instance.SetProperty("PolicySet", (value))
}

// GetPolicySet gets the value of PolicySet for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyPolicySet() (value uint8, err error) {
	retValue, err := instance.GetProperty("PolicySet")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) SetPropertyPriority(value []uint8) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetQosTrafficClassSettingData) GetPropertyPriority() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

//

// <param name="InterfaceAlias" type="string "></param>
// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="PolicySet" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetQosTrafficClassSettingData) SwitchPolicySet( /* IN */ PolicySet uint8,
	/* IN */ InterfaceAlias string,
	/* IN */ InterfaceIndex uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SwitchPolicySet", PolicySet, InterfaceAlias, InterfaceIndex)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
