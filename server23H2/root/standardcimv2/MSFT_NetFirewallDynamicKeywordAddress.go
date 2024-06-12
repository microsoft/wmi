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

// MSFT_NetFirewallDynamicKeywordAddress struct
type MSFT_NetFirewallDynamicKeywordAddress struct {
	*CIM_ManagedElement

	//
	Addresses string

	//
	AutoResolve bool

	//
	Id string

	//
	Keyword string

	//
	PolicyStoreSource string

	//
	PolicyStoreSourceType uint16
}

func NewMSFT_NetFirewallDynamicKeywordAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallDynamicKeywordAddress, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallDynamicKeywordAddress{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallDynamicKeywordAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallDynamicKeywordAddress, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallDynamicKeywordAddress{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAddresses sets the value of Addresses for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyAddresses(value string) (err error) {
	return instance.SetProperty("Addresses", (value))
}

// GetAddresses gets the value of Addresses for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyAddresses() (value string, err error) {
	retValue, err := instance.GetProperty("Addresses")
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

// SetAutoResolve sets the value of AutoResolve for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyAutoResolve(value bool) (err error) {
	return instance.SetProperty("AutoResolve", (value))
}

// GetAutoResolve gets the value of AutoResolve for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyAutoResolve() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoResolve")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetKeyword sets the value of Keyword for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyKeyword(value string) (err error) {
	return instance.SetProperty("Keyword", (value))
}

// GetKeyword gets the value of Keyword for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyKeyword() (value string, err error) {
	retValue, err := instance.GetProperty("Keyword")
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

// SetPolicyStoreSource sets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyPolicyStoreSource(value string) (err error) {
	return instance.SetProperty("PolicyStoreSource", (value))
}

// GetPolicyStoreSource gets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyPolicyStoreSource() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSource")
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

// SetPolicyStoreSourceType sets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", (value))
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetFirewallDynamicKeywordAddress) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSourceType")
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

//

// <param name="Addresses" type="string "></param>
// <param name="Append" type="bool "></param>
// <param name="Id" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallDynamicKeywordAddress) UpdateDynamicKeywordAddress( /* IN */ Id string,
	/* IN */ Addresses string,
	/* IN */ Append bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateDynamicKeywordAddress", Id, Addresses, Append)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
