// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm
//
// ////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_RestrictionsUser struct
type MDM_RestrictionsUser struct {
	*cim.WmiInstance

	//
	key uint32

	//
	PCSettingsMeteredNetworkSyncEnabled bool

	//
	PCSettingsPasswordSyncEnabled bool

	//
	PCSettingsSyncEnabled bool
}

func NewMDM_RestrictionsUserEx1(instance *cim.WmiInstance) (newInstance *MDM_RestrictionsUser, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_RestrictionsUser{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_RestrictionsUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_RestrictionsUser, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_RestrictionsUser{
		WmiInstance: tmp,
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MDM_RestrictionsUser) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_RestrictionsUser) GetPropertykey() (value uint32, err error) {
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

// SetPCSettingsMeteredNetworkSyncEnabled sets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsMeteredNetworkSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsMeteredNetworkSyncEnabled", (value))
}

// GetPCSettingsMeteredNetworkSyncEnabled gets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsMeteredNetworkSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsMeteredNetworkSyncEnabled")
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

// SetPCSettingsPasswordSyncEnabled sets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsPasswordSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsPasswordSyncEnabled", (value))
}

// GetPCSettingsPasswordSyncEnabled gets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsPasswordSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsPasswordSyncEnabled")
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

// SetPCSettingsSyncEnabled sets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsSyncEnabled", (value))
}

// GetPCSettingsSyncEnabled gets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsSyncEnabled")
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
