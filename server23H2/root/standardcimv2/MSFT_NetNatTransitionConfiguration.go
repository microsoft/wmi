// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetNatTransitionConfiguration struct
type MSFT_NetNatTransitionConfiguration struct {
	*MSFT_NetSettingData

	//
	InboundInterface []string

	//
	InstanceName string

	//
	IPv4AddressPortPool []string

	//
	OutboundInterface []string

	//
	PolicyStore uint32

	//
	PrefixMapping []string

	//
	State uint32

	//
	TcpMappingTimeout uint32
}

func NewMSFT_NetNatTransitionConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatTransitionConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatTransitionConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatTransitionConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatTransitionConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatTransitionConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetInboundInterface sets the value of InboundInterface for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyInboundInterface(value []string) (err error) {
	return instance.SetProperty("InboundInterface", (value))
}

// GetInboundInterface gets the value of InboundInterface for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyInboundInterface() (value []string, err error) {
	retValue, err := instance.GetProperty("InboundInterface")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetIPv4AddressPortPool sets the value of IPv4AddressPortPool for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyIPv4AddressPortPool(value []string) (err error) {
	return instance.SetProperty("IPv4AddressPortPool", (value))
}

// GetIPv4AddressPortPool gets the value of IPv4AddressPortPool for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyIPv4AddressPortPool() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4AddressPortPool")
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

// SetOutboundInterface sets the value of OutboundInterface for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyOutboundInterface(value []string) (err error) {
	return instance.SetProperty("OutboundInterface", (value))
}

// GetOutboundInterface gets the value of OutboundInterface for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyOutboundInterface() (value []string, err error) {
	retValue, err := instance.GetProperty("OutboundInterface")
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

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyPolicyStore(value uint32) (err error) {
	return instance.SetProperty("PolicyStore", (value))
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyPolicyStore() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
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

// SetPrefixMapping sets the value of PrefixMapping for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyPrefixMapping(value []string) (err error) {
	return instance.SetProperty("PrefixMapping", (value))
}

// GetPrefixMapping gets the value of PrefixMapping for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyPrefixMapping() (value []string, err error) {
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

// SetState sets the value of State for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyState() (value uint32, err error) {
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

// SetTcpMappingTimeout sets the value of TcpMappingTimeout for the instance
func (instance *MSFT_NetNatTransitionConfiguration) SetPropertyTcpMappingTimeout(value uint32) (err error) {
	return instance.SetProperty("TcpMappingTimeout", (value))
}

// GetTcpMappingTimeout gets the value of TcpMappingTimeout for the instance
func (instance *MSFT_NetNatTransitionConfiguration) GetPropertyTcpMappingTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpMappingTimeout")
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

// <param name="OutputObject" type="MSFT_NetNatTransitionConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetNatTransitionConfiguration) Enable( /* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetNatTransitionConfiguration) (result uint32, err error) {
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

// <param name="OutputObject" type="MSFT_NetNatTransitionConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetNatTransitionConfiguration) Disable( /* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetNatTransitionConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
