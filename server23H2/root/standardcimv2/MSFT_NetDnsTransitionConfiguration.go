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

// MSFT_NetDnsTransitionConfiguration struct
type MSFT_NetDnsTransitionConfiguration struct {
	*MSFT_NetSettingData

	//
	AcceptInterface []string

	//
	AlwaysSynthesize bool

	//
	ExclusionList []string

	//
	Latency uint32

	//
	OnlySendAQuery bool

	//
	PrefixMapping []string

	//
	SendInterface []string

	//
	State uint32
}

func NewMSFT_NetDnsTransitionConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetDnsTransitionConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetDnsTransitionConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetDnsTransitionConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetDnsTransitionConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetDnsTransitionConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAcceptInterface sets the value of AcceptInterface for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyAcceptInterface(value []string) (err error) {
	return instance.SetProperty("AcceptInterface", (value))
}

// GetAcceptInterface gets the value of AcceptInterface for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyAcceptInterface() (value []string, err error) {
	retValue, err := instance.GetProperty("AcceptInterface")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetAlwaysSynthesize sets the value of AlwaysSynthesize for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyAlwaysSynthesize(value bool) (err error) {
	return instance.SetProperty("AlwaysSynthesize", (value))
}

// GetAlwaysSynthesize gets the value of AlwaysSynthesize for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyAlwaysSynthesize() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysSynthesize")
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

// SetExclusionList sets the value of ExclusionList for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyExclusionList(value []string) (err error) {
	return instance.SetProperty("ExclusionList", (value))
}

// GetExclusionList gets the value of ExclusionList for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyExclusionList() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusionList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetLatency sets the value of Latency for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
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

// SetOnlySendAQuery sets the value of OnlySendAQuery for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyOnlySendAQuery(value bool) (err error) {
	return instance.SetProperty("OnlySendAQuery", (value))
}

// GetOnlySendAQuery gets the value of OnlySendAQuery for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyOnlySendAQuery() (value bool, err error) {
	retValue, err := instance.GetProperty("OnlySendAQuery")
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

// SetPrefixMapping sets the value of PrefixMapping for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyPrefixMapping(value []string) (err error) {
	return instance.SetProperty("PrefixMapping", (value))
}

// GetPrefixMapping gets the value of PrefixMapping for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyPrefixMapping() (value []string, err error) {
	retValue, err := instance.GetProperty("PrefixMapping")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSendInterface sets the value of SendInterface for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertySendInterface(value []string) (err error) {
	return instance.SetProperty("SendInterface", (value))
}

// GetSendInterface gets the value of SendInterface for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertySendInterface() (value []string, err error) {
	retValue, err := instance.GetProperty("SendInterface")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetDnsTransitionConfiguration) GetPropertyState() (value uint32, err error) {
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

// <param name="OutputObject" type="MSFT_NetDnsTransitionConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetDnsTransitionConfiguration) Enable( /* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetDnsTransitionConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PassThru" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetDnsTransitionConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetDnsTransitionConfiguration) Disable( /* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetDnsTransitionConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AcceptInterface" type="bool "></param>
// <param name="AlwaysSynthesize" type="bool "></param>
// <param name="ExclusionList" type="bool "></param>
// <param name="Latency" type="bool "></param>
// <param name="OnlySendAQuery" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PrefixMapping" type="bool "></param>
// <param name="SendInterface" type="bool "></param>
// <param name="State" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetDnsTransitionConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetDnsTransitionConfiguration) Reset( /* IN */ State bool,
	/* IN */ OnlySendAQuery bool,
	/* IN */ Latency bool,
	/* IN */ AlwaysSynthesize bool,
	/* IN */ PrefixMapping bool,
	/* IN */ ExclusionList bool,
	/* IN */ SendInterface bool,
	/* IN */ AcceptInterface bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetDnsTransitionConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", State, OnlySendAQuery, Latency, AlwaysSynthesize, PrefixMapping, ExclusionList, SendInterface, AcceptInterface, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
