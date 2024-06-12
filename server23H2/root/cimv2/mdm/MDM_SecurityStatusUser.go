// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_SecurityStatusUser struct
type MDM_SecurityStatusUser struct {
	*cim.WmiInstance

	//
	ConnectedAccountPolicy uint32

	//
	DeviceEncryptionPolicy uint32

	//
	EncryptionStatus uint32

	//
	HasConnectedAccount bool

	//
	key uint32

	//
	PasswordStatus uint32
}

func NewMDM_SecurityStatusUserEx1(instance *cim.WmiInstance) (newInstance *MDM_SecurityStatusUser, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_SecurityStatusUser{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_SecurityStatusUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_SecurityStatusUser, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_SecurityStatusUser{
		WmiInstance: tmp,
	}
	return
}

// SetConnectedAccountPolicy sets the value of ConnectedAccountPolicy for the instance
func (instance *MDM_SecurityStatusUser) SetPropertyConnectedAccountPolicy(value uint32) (err error) {
	return instance.SetProperty("ConnectedAccountPolicy", (value))
}

// GetConnectedAccountPolicy gets the value of ConnectedAccountPolicy for the instance
func (instance *MDM_SecurityStatusUser) GetPropertyConnectedAccountPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectedAccountPolicy")
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

// SetDeviceEncryptionPolicy sets the value of DeviceEncryptionPolicy for the instance
func (instance *MDM_SecurityStatusUser) SetPropertyDeviceEncryptionPolicy(value uint32) (err error) {
	return instance.SetProperty("DeviceEncryptionPolicy", (value))
}

// GetDeviceEncryptionPolicy gets the value of DeviceEncryptionPolicy for the instance
func (instance *MDM_SecurityStatusUser) GetPropertyDeviceEncryptionPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceEncryptionPolicy")
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

// SetEncryptionStatus sets the value of EncryptionStatus for the instance
func (instance *MDM_SecurityStatusUser) SetPropertyEncryptionStatus(value uint32) (err error) {
	return instance.SetProperty("EncryptionStatus", (value))
}

// GetEncryptionStatus gets the value of EncryptionStatus for the instance
func (instance *MDM_SecurityStatusUser) GetPropertyEncryptionStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncryptionStatus")
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

// SetHasConnectedAccount sets the value of HasConnectedAccount for the instance
func (instance *MDM_SecurityStatusUser) SetPropertyHasConnectedAccount(value bool) (err error) {
	return instance.SetProperty("HasConnectedAccount", (value))
}

// GetHasConnectedAccount gets the value of HasConnectedAccount for the instance
func (instance *MDM_SecurityStatusUser) GetPropertyHasConnectedAccount() (value bool, err error) {
	retValue, err := instance.GetProperty("HasConnectedAccount")
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

// Setkey sets the value of key for the instance
func (instance *MDM_SecurityStatusUser) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_SecurityStatusUser) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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

// SetPasswordStatus sets the value of PasswordStatus for the instance
func (instance *MDM_SecurityStatusUser) SetPropertyPasswordStatus(value uint32) (err error) {
	return instance.SetProperty("PasswordStatus", (value))
}

// GetPasswordStatus gets the value of PasswordStatus for the instance
func (instance *MDM_SecurityStatusUser) GetPropertyPasswordStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("PasswordStatus")
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
