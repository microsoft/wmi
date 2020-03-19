// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RegistryValueChangeEvent struct
type RegistryValueChangeEvent struct {
	*RegistryEvent

	//
	Hive string

	//
	KeyPath string

	//
	ValueName string
}

func NewRegistryValueChangeEventEx1(instance *cim.WmiInstance) (newInstance *RegistryValueChangeEvent, err error) {
	tmp, err := NewRegistryEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RegistryValueChangeEvent{
		RegistryEvent: tmp,
	}
	return
}

func NewRegistryValueChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RegistryValueChangeEvent, err error) {
	tmp, err := NewRegistryEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RegistryValueChangeEvent{
		RegistryEvent: tmp,
	}
	return
}

// SetHive sets the value of Hive for the instance
func (instance *RegistryValueChangeEvent) SetPropertyHive(value string) (err error) {
	return instance.SetProperty("Hive", value)
}

// GetHive gets the value of Hive for the instance
func (instance *RegistryValueChangeEvent) GetPropertyHive() (value string, err error) {
	retValue, err := instance.GetProperty("Hive")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyPath sets the value of KeyPath for the instance
func (instance *RegistryValueChangeEvent) SetPropertyKeyPath(value string) (err error) {
	return instance.SetProperty("KeyPath", value)
}

// GetKeyPath gets the value of KeyPath for the instance
func (instance *RegistryValueChangeEvent) GetPropertyKeyPath() (value string, err error) {
	retValue, err := instance.GetProperty("KeyPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueName sets the value of ValueName for the instance
func (instance *RegistryValueChangeEvent) SetPropertyValueName(value string) (err error) {
	return instance.SetProperty("ValueName", value)
}

// GetValueName gets the value of ValueName for the instance
func (instance *RegistryValueChangeEvent) GetPropertyValueName() (value string, err error) {
	retValue, err := instance.GetProperty("ValueName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
