// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_FirewallRulesbyStore struct
type Win32_PerfFormattedData_Counters_FirewallRulesbyStore struct {
	*Win32_PerfFormattedData

	//
	AppIsolationStore uint64

	//
	DefaultsStore uint64

	//
	DynamicStore uint64

	//
	GPOStore uint64

	//
	GroupPolicyRSoPStore uint64

	//
	InterfaceIsolationDynamicStore uint64

	//
	InterfaceIsolationStore uint64

	//
	LocalStore uint64

	//
	MDMStore uint64

	//
	TenantRestrictionsStore uint64

	//
	Total uint64

	//
	WSHConfigurableStore uint64

	//
	WSHStaticStore uint64
}

func NewWin32_PerfFormattedData_Counters_FirewallRulesbyStoreEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_FirewallRulesbyStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_FirewallRulesbyStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_FirewallRulesbyStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAppIsolationStore sets the value of AppIsolationStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyAppIsolationStore(value uint64) (err error) {
	return instance.SetProperty("AppIsolationStore", (value))
}

// GetAppIsolationStore gets the value of AppIsolationStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyAppIsolationStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("AppIsolationStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDefaultsStore sets the value of DefaultsStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyDefaultsStore(value uint64) (err error) {
	return instance.SetProperty("DefaultsStore", (value))
}

// GetDefaultsStore gets the value of DefaultsStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyDefaultsStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultsStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDynamicStore sets the value of DynamicStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyDynamicStore(value uint64) (err error) {
	return instance.SetProperty("DynamicStore", (value))
}

// GetDynamicStore gets the value of DynamicStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyDynamicStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("DynamicStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetGPOStore sets the value of GPOStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyGPOStore(value uint64) (err error) {
	return instance.SetProperty("GPOStore", (value))
}

// GetGPOStore gets the value of GPOStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyGPOStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("GPOStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetGroupPolicyRSoPStore sets the value of GroupPolicyRSoPStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyGroupPolicyRSoPStore(value uint64) (err error) {
	return instance.SetProperty("GroupPolicyRSoPStore", (value))
}

// GetGroupPolicyRSoPStore gets the value of GroupPolicyRSoPStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyGroupPolicyRSoPStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupPolicyRSoPStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInterfaceIsolationDynamicStore sets the value of InterfaceIsolationDynamicStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyInterfaceIsolationDynamicStore(value uint64) (err error) {
	return instance.SetProperty("InterfaceIsolationDynamicStore", (value))
}

// GetInterfaceIsolationDynamicStore gets the value of InterfaceIsolationDynamicStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyInterfaceIsolationDynamicStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterfaceIsolationDynamicStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInterfaceIsolationStore sets the value of InterfaceIsolationStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyInterfaceIsolationStore(value uint64) (err error) {
	return instance.SetProperty("InterfaceIsolationStore", (value))
}

// GetInterfaceIsolationStore gets the value of InterfaceIsolationStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyInterfaceIsolationStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterfaceIsolationStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLocalStore sets the value of LocalStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyLocalStore(value uint64) (err error) {
	return instance.SetProperty("LocalStore", (value))
}

// GetLocalStore gets the value of LocalStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyLocalStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMDMStore sets the value of MDMStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyMDMStore(value uint64) (err error) {
	return instance.SetProperty("MDMStore", (value))
}

// GetMDMStore gets the value of MDMStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyMDMStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("MDMStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTenantRestrictionsStore sets the value of TenantRestrictionsStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyTenantRestrictionsStore(value uint64) (err error) {
	return instance.SetProperty("TenantRestrictionsStore", (value))
}

// GetTenantRestrictionsStore gets the value of TenantRestrictionsStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyTenantRestrictionsStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("TenantRestrictionsStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTotal sets the value of Total for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyTotal(value uint64) (err error) {
	return instance.SetProperty("Total", (value))
}

// GetTotal gets the value of Total for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("Total")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWSHConfigurableStore sets the value of WSHConfigurableStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyWSHConfigurableStore(value uint64) (err error) {
	return instance.SetProperty("WSHConfigurableStore", (value))
}

// GetWSHConfigurableStore gets the value of WSHConfigurableStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyWSHConfigurableStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("WSHConfigurableStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWSHStaticStore sets the value of WSHStaticStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) SetPropertyWSHStaticStore(value uint64) (err error) {
	return instance.SetProperty("WSHStaticStore", (value))
}

// GetWSHStaticStore gets the value of WSHStaticStore for the instance
func (instance *Win32_PerfFormattedData_Counters_FirewallRulesbyStore) GetPropertyWSHStaticStore() (value uint64, err error) {
	retValue, err := instance.GetProperty("WSHStaticStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
