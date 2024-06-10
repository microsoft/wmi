// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Client
//
// ////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnConnectionIPsecConfiguration struct
type VpnConnectionIPsecConfiguration struct {
	*cim.WmiInstance

	//
	AuthenticationTransformConstants uint32

	//
	CipherTransformConstants uint32

	//
	DHGroup uint32

	//
	EncryptionMethod uint32

	//
	IntegrityCheckMethod uint32

	//
	PfsGroup uint32
}

func NewVpnConnectionIPsecConfigurationEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionIPsecConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnConnectionIPsecConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewVpnConnectionIPsecConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnectionIPsecConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnectionIPsecConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAuthenticationTransformConstants sets the value of AuthenticationTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyAuthenticationTransformConstants(value uint32) (err error) {
	return instance.SetProperty("AuthenticationTransformConstants", (value))
}

// GetAuthenticationTransformConstants gets the value of AuthenticationTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyAuthenticationTransformConstants() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthenticationTransformConstants")
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

// SetCipherTransformConstants sets the value of CipherTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyCipherTransformConstants(value uint32) (err error) {
	return instance.SetProperty("CipherTransformConstants", (value))
}

// GetCipherTransformConstants gets the value of CipherTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyCipherTransformConstants() (value uint32, err error) {
	retValue, err := instance.GetProperty("CipherTransformConstants")
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

// SetDHGroup sets the value of DHGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyDHGroup(value uint32) (err error) {
	return instance.SetProperty("DHGroup", (value))
}

// GetDHGroup gets the value of DHGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyDHGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("DHGroup")
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

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyEncryptionMethod(value uint32) (err error) {
	return instance.SetProperty("EncryptionMethod", (value))
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyEncryptionMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncryptionMethod")
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

// SetIntegrityCheckMethod sets the value of IntegrityCheckMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyIntegrityCheckMethod(value uint32) (err error) {
	return instance.SetProperty("IntegrityCheckMethod", (value))
}

// GetIntegrityCheckMethod gets the value of IntegrityCheckMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyIntegrityCheckMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntegrityCheckMethod")
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

// SetPfsGroup sets the value of PfsGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyPfsGroup(value uint32) (err error) {
	return instance.SetProperty("PfsGroup", (value))
}

// GetPfsGroup gets the value of PfsGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyPfsGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("PfsGroup")
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
