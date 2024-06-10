// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnMultiTenancy struct
type VpnMultiTenancy struct {
	*RemoteAccessCommon

	//
	AuthenticationPolicy VpnAuth

	//
	CapacityKbps uint64

	//
	VpnMultiTenancyStatus string
}

func NewVpnMultiTenancyEx1(instance *cim.WmiInstance) (newInstance *VpnMultiTenancy, err error) {
	tmp, err := NewRemoteAccessCommonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnMultiTenancy{
		RemoteAccessCommon: tmp,
	}
	return
}

func NewVpnMultiTenancyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnMultiTenancy, err error) {
	tmp, err := NewRemoteAccessCommonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnMultiTenancy{
		RemoteAccessCommon: tmp,
	}
	return
}

// SetAuthenticationPolicy sets the value of AuthenticationPolicy for the instance
func (instance *VpnMultiTenancy) SetPropertyAuthenticationPolicy(value VpnAuth) (err error) {
	return instance.SetProperty("AuthenticationPolicy", (value))
}

// GetAuthenticationPolicy gets the value of AuthenticationPolicy for the instance
func (instance *VpnMultiTenancy) GetPropertyAuthenticationPolicy() (value VpnAuth, err error) {
	retValue, err := instance.GetProperty("AuthenticationPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VpnAuth)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VpnAuth is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VpnAuth(valuetmp)

	return
}

// SetCapacityKbps sets the value of CapacityKbps for the instance
func (instance *VpnMultiTenancy) SetPropertyCapacityKbps(value uint64) (err error) {
	return instance.SetProperty("CapacityKbps", (value))
}

// GetCapacityKbps gets the value of CapacityKbps for the instance
func (instance *VpnMultiTenancy) GetPropertyCapacityKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("CapacityKbps")
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

// SetVpnMultiTenancyStatus sets the value of VpnMultiTenancyStatus for the instance
func (instance *VpnMultiTenancy) SetPropertyVpnMultiTenancyStatus(value string) (err error) {
	return instance.SetProperty("VpnMultiTenancyStatus", (value))
}

// GetVpnMultiTenancyStatus gets the value of VpnMultiTenancyStatus for the instance
func (instance *VpnMultiTenancy) GetPropertyVpnMultiTenancyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("VpnMultiTenancyStatus")
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
