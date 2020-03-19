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

// RegistryTreeChangeEvent struct
type RegistryTreeChangeEvent struct {
	*RegistryEvent

	//
	Hive string

	//
	RootPath string
}

func NewRegistryTreeChangeEventEx1(instance *cim.WmiInstance) (newInstance *RegistryTreeChangeEvent, err error) {
	tmp, err := NewRegistryEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RegistryTreeChangeEvent{
		RegistryEvent: tmp,
	}
	return
}

func NewRegistryTreeChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RegistryTreeChangeEvent, err error) {
	tmp, err := NewRegistryEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RegistryTreeChangeEvent{
		RegistryEvent: tmp,
	}
	return
}

// SetHive sets the value of Hive for the instance
func (instance *RegistryTreeChangeEvent) SetPropertyHive(value string) (err error) {
	return instance.SetProperty("Hive", value)
}

// GetHive gets the value of Hive for the instance
func (instance *RegistryTreeChangeEvent) GetPropertyHive() (value string, err error) {
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

// SetRootPath sets the value of RootPath for the instance
func (instance *RegistryTreeChangeEvent) SetPropertyRootPath(value string) (err error) {
	return instance.SetProperty("RootPath", value)
}

// GetRootPath gets the value of RootPath for the instance
func (instance *RegistryTreeChangeEvent) GetPropertyRootPath() (value string, err error) {
	retValue, err := instance.GetProperty("RootPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
