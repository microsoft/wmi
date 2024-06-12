// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ReplicationAuthorizationSettingData struct
type Msvm_ReplicationAuthorizationSettingData struct {
	*CIM_SettingData

	//
	AllowedPrimaryHostSystem string

	//
	ReplicaStorageLocation string

	//
	TrustGroup string
}

func NewMsvm_ReplicationAuthorizationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationAuthorizationSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationAuthorizationSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_ReplicationAuthorizationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationAuthorizationSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationAuthorizationSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAllowedPrimaryHostSystem sets the value of AllowedPrimaryHostSystem for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyAllowedPrimaryHostSystem(value string) (err error) {
	return instance.SetProperty("AllowedPrimaryHostSystem", (value))
}

// GetAllowedPrimaryHostSystem gets the value of AllowedPrimaryHostSystem for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyAllowedPrimaryHostSystem() (value string, err error) {
	retValue, err := instance.GetProperty("AllowedPrimaryHostSystem")
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

// SetReplicaStorageLocation sets the value of ReplicaStorageLocation for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyReplicaStorageLocation(value string) (err error) {
	return instance.SetProperty("ReplicaStorageLocation", (value))
}

// GetReplicaStorageLocation gets the value of ReplicaStorageLocation for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyReplicaStorageLocation() (value string, err error) {
	retValue, err := instance.GetProperty("ReplicaStorageLocation")
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

// SetTrustGroup sets the value of TrustGroup for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyTrustGroup(value string) (err error) {
	return instance.SetProperty("TrustGroup", (value))
}

// GetTrustGroup gets the value of TrustGroup for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyTrustGroup() (value string, err error) {
	retValue, err := instance.GetProperty("TrustGroup")
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
