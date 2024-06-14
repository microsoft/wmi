// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetConnectionProfile struct
type MSFT_NetConnectionProfile struct {
	*MSFT_NetSettingData

	//
	DomainAuthenticationKind uint32

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	IPv4Connectivity uint32

	//
	IPv6Connectivity uint32

	//
	Name string

	//
	NetworkCategory uint32
}

func NewMSFT_NetConnectionProfileEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConnectionProfile, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConnectionProfile{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetConnectionProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConnectionProfile, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConnectionProfile{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetDomainAuthenticationKind sets the value of DomainAuthenticationKind for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyDomainAuthenticationKind(value uint32) (err error) {
	return instance.SetProperty("DomainAuthenticationKind", (value))
}

// GetDomainAuthenticationKind gets the value of DomainAuthenticationKind for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyDomainAuthenticationKind() (value uint32, err error) {
	retValue, err := instance.GetProperty("DomainAuthenticationKind")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyInterfaceAlias() (value string, err error) {
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
func (instance *MSFT_NetConnectionProfile) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyInterfaceIndex() (value uint32, err error) {
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

// SetIPv4Connectivity sets the value of IPv4Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyIPv4Connectivity(value uint32) (err error) {
	return instance.SetProperty("IPv4Connectivity", (value))
}

// GetIPv4Connectivity gets the value of IPv4Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyIPv4Connectivity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4Connectivity")
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

// SetIPv6Connectivity sets the value of IPv6Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyIPv6Connectivity(value uint32) (err error) {
	return instance.SetProperty("IPv6Connectivity", (value))
}

// GetIPv6Connectivity gets the value of IPv6Connectivity for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyIPv6Connectivity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6Connectivity")
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
func (instance *MSFT_NetConnectionProfile) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyName() (value string, err error) {
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

// SetNetworkCategory sets the value of NetworkCategory for the instance
func (instance *MSFT_NetConnectionProfile) SetPropertyNetworkCategory(value uint32) (err error) {
	return instance.SetProperty("NetworkCategory", (value))
}

// GetNetworkCategory gets the value of NetworkCategory for the instance
func (instance *MSFT_NetConnectionProfile) GetPropertyNetworkCategory() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkCategory")
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
