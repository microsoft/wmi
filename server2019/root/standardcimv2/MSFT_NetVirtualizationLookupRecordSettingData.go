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

// MSFT_NetVirtualizationLookupRecordSettingData struct
type MSFT_NetVirtualizationLookupRecordSettingData struct {
	*MSFT_NetSettingData

	// 33
	Context string

	// 28
	CustomerAddress string

	// 32
	CustomerID string

	// 30
	MACAddress string

	// 29
	ProviderAddress string

	// 34
	Rule NetVirtualizationLookupRecordSettingData_Rule

	// 40
	Type NetVirtualizationLookupRecordSettingData_Type

	// 43
	Unsynchronized bool

	// 42
	Unusable bool

	// 39
	UseVmMACAddress bool

	// 31
	VirtualSubnetID uint32

	// 38
	VMName string
}

func NewMSFT_NetVirtualizationLookupRecordSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationLookupRecordSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationLookupRecordSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationLookupRecordSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationLookupRecordSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationLookupRecordSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetContext sets the value of Context for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyContext(value string) (err error) {
	return instance.SetProperty("Context", (value))
}

// GetContext gets the value of Context for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyContext() (value string, err error) {
	retValue, err := instance.GetProperty("Context")
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

// SetCustomerAddress sets the value of CustomerAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyCustomerAddress(value string) (err error) {
	return instance.SetProperty("CustomerAddress", (value))
}

// GetCustomerAddress gets the value of CustomerAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyCustomerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("CustomerAddress")
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

// SetCustomerID sets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyCustomerID(value string) (err error) {
	return instance.SetProperty("CustomerID", (value))
}

// GetCustomerID gets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyCustomerID() (value string, err error) {
	retValue, err := instance.GetProperty("CustomerID")
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

// SetMACAddress sets the value of MACAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", (value))
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyMACAddress() (value string, err error) {
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

// SetProviderAddress sets the value of ProviderAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyProviderAddress(value string) (err error) {
	return instance.SetProperty("ProviderAddress", (value))
}

// GetProviderAddress gets the value of ProviderAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyProviderAddress() (value string, err error) {
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

// SetRule sets the value of Rule for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyRule(value NetVirtualizationLookupRecordSettingData_Rule) (err error) {
	return instance.SetProperty("Rule", (value))
}

// GetRule gets the value of Rule for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyRule() (value NetVirtualizationLookupRecordSettingData_Rule, err error) {
	retValue, err := instance.GetProperty("Rule")
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

	value = NetVirtualizationLookupRecordSettingData_Rule(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyType(value NetVirtualizationLookupRecordSettingData_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyType() (value NetVirtualizationLookupRecordSettingData_Type, err error) {
	retValue, err := instance.GetProperty("Type")
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

	value = NetVirtualizationLookupRecordSettingData_Type(valuetmp)

	return
}

// SetUnsynchronized sets the value of Unsynchronized for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyUnsynchronized(value bool) (err error) {
	return instance.SetProperty("Unsynchronized", (value))
}

// GetUnsynchronized gets the value of Unsynchronized for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyUnsynchronized() (value bool, err error) {
	retValue, err := instance.GetProperty("Unsynchronized")
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

// SetUnusable sets the value of Unusable for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyUnusable(value bool) (err error) {
	return instance.SetProperty("Unusable", (value))
}

// GetUnusable gets the value of Unusable for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyUnusable() (value bool, err error) {
	retValue, err := instance.GetProperty("Unusable")
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

// SetUseVmMACAddress sets the value of UseVmMACAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyUseVmMACAddress(value bool) (err error) {
	return instance.SetProperty("UseVmMACAddress", (value))
}

// GetUseVmMACAddress gets the value of UseVmMACAddress for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyUseVmMACAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("UseVmMACAddress")
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

// SetVirtualSubnetID sets the value of VirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyVirtualSubnetID(value uint32) (err error) {
	return instance.SetProperty("VirtualSubnetID", (value))
}

// GetVirtualSubnetID gets the value of VirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyVirtualSubnetID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSubnetID")
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

// SetVMName sets the value of VMName for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) SetPropertyVMName(value string) (err error) {
	return instance.SetProperty("VMName", (value))
}

// GetVMName gets the value of VMName for the instance
func (instance *MSFT_NetVirtualizationLookupRecordSettingData) GetPropertyVMName() (value string, err error) {
	retValue, err := instance.GetProperty("VMName")
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
