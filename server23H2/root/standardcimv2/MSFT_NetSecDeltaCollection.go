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

// MSFT_NetSecDeltaCollection struct
type MSFT_NetSecDeltaCollection struct {
	*CIM_SettingData

	//
	Action uint16

	//
	EndpointType uint16

	//
	IPsecRuleDisplayName string

	//
	IPsecRuleName string

	//
	IPv4Addresses []string

	//
	IPv6Addresses []string

	//
	NameResolutionFailures []string

	//
	PolicyStore string
}

func NewMSFT_NetSecDeltaCollectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSecDeltaCollection, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSecDeltaCollection{
		CIM_SettingData: tmp,
	}
	return
}

func NewMSFT_NetSecDeltaCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSecDeltaCollection, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSecDeltaCollection{
		CIM_SettingData: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyAction(value uint16) (err error) {
	return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetEndpointType sets the value of EndpointType for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyEndpointType(value uint16) (err error) {
	return instance.SetProperty("EndpointType", (value))
}

// GetEndpointType gets the value of EndpointType for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyEndpointType() (value uint16, err error) {
	retValue, err := instance.GetProperty("EndpointType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetIPsecRuleDisplayName sets the value of IPsecRuleDisplayName for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyIPsecRuleDisplayName(value string) (err error) {
	return instance.SetProperty("IPsecRuleDisplayName", (value))
}

// GetIPsecRuleDisplayName gets the value of IPsecRuleDisplayName for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyIPsecRuleDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("IPsecRuleDisplayName")
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

// SetIPsecRuleName sets the value of IPsecRuleName for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyIPsecRuleName(value string) (err error) {
	return instance.SetProperty("IPsecRuleName", (value))
}

// GetIPsecRuleName gets the value of IPsecRuleName for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyIPsecRuleName() (value string, err error) {
	retValue, err := instance.GetProperty("IPsecRuleName")
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

// SetIPv4Addresses sets the value of IPv4Addresses for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyIPv4Addresses(value []string) (err error) {
	return instance.SetProperty("IPv4Addresses", (value))
}

// GetIPv4Addresses gets the value of IPv4Addresses for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyIPv4Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4Addresses")
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

// SetIPv6Addresses sets the value of IPv6Addresses for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyIPv6Addresses(value []string) (err error) {
	return instance.SetProperty("IPv6Addresses", (value))
}

// GetIPv6Addresses gets the value of IPv6Addresses for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyIPv6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6Addresses")
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

// SetNameResolutionFailures sets the value of NameResolutionFailures for the instance
func (instance *MSFT_NetSecDeltaCollection) SetPropertyNameResolutionFailures(value []string) (err error) {
	return instance.SetProperty("NameResolutionFailures", (value))
}

// GetNameResolutionFailures gets the value of NameResolutionFailures for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyNameResolutionFailures() (value []string, err error) {
	retValue, err := instance.GetProperty("NameResolutionFailures")
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
func (instance *MSFT_NetSecDeltaCollection) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", (value))
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NetSecDeltaCollection) GetPropertyPolicyStore() (value string, err error) {
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
