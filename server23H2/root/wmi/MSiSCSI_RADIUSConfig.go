// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_RADIUSConfig struct
type MSiSCSI_RADIUSConfig struct {
	*cim.WmiInstance

	//
	Active bool

	// Fixed address of backup RADIUS server
	BackupRADIUSServer ISCSI_IP_Address

	//
	InstanceName string

	// Fixed address of primary RADIUS server
	RADIUSServer ISCSI_IP_Address

	// Must be zero
	Reserved uint32

	// Shared secret for communicating with primary and backup RADIUS servers. This field is a variable length array.
	SharedSecret []uint8

	// Size in bytes of shared secret used to communicate with RADIUS servers
	SharedSecretSizeInBytes uint32

	// TRUE if adapter should use RADIUS for CHAP authentication
	UseRADIUSForCHAP bool
}

func NewMSiSCSI_RADIUSConfigEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_RADIUSConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RADIUSConfig{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_RADIUSConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_RADIUSConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RADIUSConfig{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetBackupRADIUSServer sets the value of BackupRADIUSServer for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyBackupRADIUSServer(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("BackupRADIUSServer", (value))
}

// GetBackupRADIUSServer gets the value of BackupRADIUSServer for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyBackupRADIUSServer() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("BackupRADIUSServer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetRADIUSServer sets the value of RADIUSServer for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyRADIUSServer(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("RADIUSServer", (value))
}

// GetRADIUSServer gets the value of RADIUSServer for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyRADIUSServer() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("RADIUSServer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetSharedSecret sets the value of SharedSecret for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertySharedSecret(value []uint8) (err error) {
	return instance.SetProperty("SharedSecret", (value))
}

// GetSharedSecret gets the value of SharedSecret for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertySharedSecret() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SharedSecret")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetSharedSecretSizeInBytes sets the value of SharedSecretSizeInBytes for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertySharedSecretSizeInBytes(value uint32) (err error) {
	return instance.SetProperty("SharedSecretSizeInBytes", (value))
}

// GetSharedSecretSizeInBytes gets the value of SharedSecretSizeInBytes for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertySharedSecretSizeInBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("SharedSecretSizeInBytes")
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

// SetUseRADIUSForCHAP sets the value of UseRADIUSForCHAP for the instance
func (instance *MSiSCSI_RADIUSConfig) SetPropertyUseRADIUSForCHAP(value bool) (err error) {
	return instance.SetProperty("UseRADIUSForCHAP", (value))
}

// GetUseRADIUSForCHAP gets the value of UseRADIUSForCHAP for the instance
func (instance *MSiSCSI_RADIUSConfig) GetPropertyUseRADIUSForCHAP() (value bool, err error) {
	retValue, err := instance.GetProperty("UseRADIUSForCHAP")
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
