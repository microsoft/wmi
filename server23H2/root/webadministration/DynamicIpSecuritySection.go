// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DynamicIpSecuritySection struct
type DynamicIpSecuritySection struct {
	*ConfigurationSectionWithCollection

	//
	DenyAction int32

	//
	DenyByConcurrentRequests DenyByConcurrentRequestsSettings

	//
	DenyByRequestRate DenyByRequestRateSettings

	//
	EnableLoggingOnlyMode bool

	//
	EnableProxyMode bool
}

func NewDynamicIpSecuritySectionEx1(instance *cim.WmiInstance) (newInstance *DynamicIpSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DynamicIpSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewDynamicIpSecuritySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DynamicIpSecuritySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DynamicIpSecuritySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetDenyAction sets the value of DenyAction for the instance
func (instance *DynamicIpSecuritySection) SetPropertyDenyAction(value int32) (err error) {
	return instance.SetProperty("DenyAction", (value))
}

// GetDenyAction gets the value of DenyAction for the instance
func (instance *DynamicIpSecuritySection) GetPropertyDenyAction() (value int32, err error) {
	retValue, err := instance.GetProperty("DenyAction")
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

	value = int32(valuetmp)

	return
}

// SetDenyByConcurrentRequests sets the value of DenyByConcurrentRequests for the instance
func (instance *DynamicIpSecuritySection) SetPropertyDenyByConcurrentRequests(value DenyByConcurrentRequestsSettings) (err error) {
	return instance.SetProperty("DenyByConcurrentRequests", (value))
}

// GetDenyByConcurrentRequests gets the value of DenyByConcurrentRequests for the instance
func (instance *DynamicIpSecuritySection) GetPropertyDenyByConcurrentRequests() (value DenyByConcurrentRequestsSettings, err error) {
	retValue, err := instance.GetProperty("DenyByConcurrentRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DenyByConcurrentRequestsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DenyByConcurrentRequestsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DenyByConcurrentRequestsSettings(valuetmp)

	return
}

// SetDenyByRequestRate sets the value of DenyByRequestRate for the instance
func (instance *DynamicIpSecuritySection) SetPropertyDenyByRequestRate(value DenyByRequestRateSettings) (err error) {
	return instance.SetProperty("DenyByRequestRate", (value))
}

// GetDenyByRequestRate gets the value of DenyByRequestRate for the instance
func (instance *DynamicIpSecuritySection) GetPropertyDenyByRequestRate() (value DenyByRequestRateSettings, err error) {
	retValue, err := instance.GetProperty("DenyByRequestRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DenyByRequestRateSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DenyByRequestRateSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DenyByRequestRateSettings(valuetmp)

	return
}

// SetEnableLoggingOnlyMode sets the value of EnableLoggingOnlyMode for the instance
func (instance *DynamicIpSecuritySection) SetPropertyEnableLoggingOnlyMode(value bool) (err error) {
	return instance.SetProperty("EnableLoggingOnlyMode", (value))
}

// GetEnableLoggingOnlyMode gets the value of EnableLoggingOnlyMode for the instance
func (instance *DynamicIpSecuritySection) GetPropertyEnableLoggingOnlyMode() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLoggingOnlyMode")
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

// SetEnableProxyMode sets the value of EnableProxyMode for the instance
func (instance *DynamicIpSecuritySection) SetPropertyEnableProxyMode(value bool) (err error) {
	return instance.SetProperty("EnableProxyMode", (value))
}

// GetEnableProxyMode gets the value of EnableProxyMode for the instance
func (instance *DynamicIpSecuritySection) GetPropertyEnableProxyMode() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableProxyMode")
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
