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

// Win32_PerfRawData_Counters_FirewallRulesbyProfile struct
type Win32_PerfRawData_Counters_FirewallRulesbyProfile struct {
	*Win32_PerfRawData

	//
	DomainProfile uint64

	//
	PrivateProfile uint64

	//
	PublicProfile uint64
}

func NewWin32_PerfRawData_Counters_FirewallRulesbyProfileEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_FirewallRulesbyProfile, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_FirewallRulesbyProfile{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_FirewallRulesbyProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_FirewallRulesbyProfile, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_FirewallRulesbyProfile{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetDomainProfile sets the value of DomainProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) SetPropertyDomainProfile(value uint64) (err error) {
	return instance.SetProperty("DomainProfile", (value))
}

// GetDomainProfile gets the value of DomainProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) GetPropertyDomainProfile() (value uint64, err error) {
	retValue, err := instance.GetProperty("DomainProfile")
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

// SetPrivateProfile sets the value of PrivateProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) SetPropertyPrivateProfile(value uint64) (err error) {
	return instance.SetProperty("PrivateProfile", (value))
}

// GetPrivateProfile gets the value of PrivateProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) GetPropertyPrivateProfile() (value uint64, err error) {
	retValue, err := instance.GetProperty("PrivateProfile")
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

// SetPublicProfile sets the value of PublicProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) SetPropertyPublicProfile(value uint64) (err error) {
	return instance.SetProperty("PublicProfile", (value))
}

// GetPublicProfile gets the value of PublicProfile for the instance
func (instance *Win32_PerfRawData_Counters_FirewallRulesbyProfile) GetPropertyPublicProfile() (value uint64, err error) {
	retValue, err := instance.GetProperty("PublicProfile")
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
