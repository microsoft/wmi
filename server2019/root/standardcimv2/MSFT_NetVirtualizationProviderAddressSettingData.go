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

// MSFT_NetVirtualizationProviderAddressSettingData struct
type MSFT_NetVirtualizationProviderAddressSettingData struct {
	*MSFT_NetSettingData

	// 15
	AddressState NetVirtualizationProviderAddressSettingData_AddressState

	// 12
	InterfaceIndex uint32

	// 21
	MACAddress string

	// 22
	ManagedByCluster bool

	// 13
	PrefixLength uint8

	// 11
	ProviderAddress string

	// 14
	VlanID uint16
}

func NewMSFT_NetVirtualizationProviderAddressSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationProviderAddressSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationProviderAddressSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationProviderAddressSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationProviderAddressSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationProviderAddressSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAddressState sets the value of AddressState for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyAddressState(value NetVirtualizationProviderAddressSettingData_AddressState) (err error) {
	return instance.SetProperty("AddressState", (value))
}

// GetAddressState gets the value of AddressState for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyAddressState() (value NetVirtualizationProviderAddressSettingData_AddressState, err error) {
	retValue, err := instance.GetProperty("AddressState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = NetVirtualizationProviderAddressSettingData_AddressState(valuetmp)

	return
}

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyInterfaceIndex() (value uint32, err error) {
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

// SetMACAddress sets the value of MACAddress for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", (value))
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MACAddress")
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

// SetManagedByCluster sets the value of ManagedByCluster for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyManagedByCluster(value bool) (err error) {
	return instance.SetProperty("ManagedByCluster", (value))
}

// GetManagedByCluster gets the value of ManagedByCluster for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyManagedByCluster() (value bool, err error) {
	retValue, err := instance.GetProperty("ManagedByCluster")
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

// SetPrefixLength sets the value of PrefixLength for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyPrefixLength(value uint8) (err error) {
	return instance.SetProperty("PrefixLength", (value))
}

// GetPrefixLength gets the value of PrefixLength for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyPrefixLength() (value uint8, err error) {
	retValue, err := instance.GetProperty("PrefixLength")
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

// SetProviderAddress sets the value of ProviderAddress for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyProviderAddress(value string) (err error) {
	return instance.SetProperty("ProviderAddress", (value))
}

// GetProviderAddress gets the value of ProviderAddress for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyProviderAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderAddress")
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

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) SetPropertyVlanID(value uint16) (err error) {
	return instance.SetProperty("VlanID", (value))
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetVirtualizationProviderAddressSettingData) GetPropertyVlanID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanID")
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
