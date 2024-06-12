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

// MSFT_NetNetworkLayerSecurityFilter struct
type MSFT_NetNetworkLayerSecurityFilter struct {
	*CIM_FilterEntryBase

	//
	Authentication uint16

	//
	Encryption uint16

	//
	LocalUsers string

	//
	OverrideBlockRules bool

	//
	RemoteMachines string

	//
	RemoteUsers string
}

func NewMSFT_NetNetworkLayerSecurityFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNetworkLayerSecurityFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNetworkLayerSecurityFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetNetworkLayerSecurityFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNetworkLayerSecurityFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNetworkLayerSecurityFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetAuthentication sets the value of Authentication for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyAuthentication(value uint16) (err error) {
	return instance.SetProperty("Authentication", (value))
}

// GetAuthentication gets the value of Authentication for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyAuthentication() (value uint16, err error) {
	retValue, err := instance.GetProperty("Authentication")
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

// SetEncryption sets the value of Encryption for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyEncryption(value uint16) (err error) {
	return instance.SetProperty("Encryption", (value))
}

// GetEncryption gets the value of Encryption for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyEncryption() (value uint16, err error) {
	retValue, err := instance.GetProperty("Encryption")
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

// SetLocalUsers sets the value of LocalUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyLocalUsers(value string) (err error) {
	return instance.SetProperty("LocalUsers", (value))
}

// GetLocalUsers gets the value of LocalUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyLocalUsers() (value string, err error) {
	retValue, err := instance.GetProperty("LocalUsers")
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

// SetOverrideBlockRules sets the value of OverrideBlockRules for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyOverrideBlockRules(value bool) (err error) {
	return instance.SetProperty("OverrideBlockRules", (value))
}

// GetOverrideBlockRules gets the value of OverrideBlockRules for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyOverrideBlockRules() (value bool, err error) {
	retValue, err := instance.GetProperty("OverrideBlockRules")
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

// SetRemoteMachines sets the value of RemoteMachines for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyRemoteMachines(value string) (err error) {
	return instance.SetProperty("RemoteMachines", (value))
}

// GetRemoteMachines gets the value of RemoteMachines for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyRemoteMachines() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachines")
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

// SetRemoteUsers sets the value of RemoteUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyRemoteUsers(value string) (err error) {
	return instance.SetProperty("RemoteUsers", (value))
}

// GetRemoteUsers gets the value of RemoteUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyRemoteUsers() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUsers")
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
