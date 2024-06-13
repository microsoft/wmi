// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Registry_Config struct
type Registry_Config struct {
	*Registry

	//
	CurrentControlSet uint32
}

func NewRegistry_ConfigEx1(instance *cim.WmiInstance) (newInstance *Registry_Config, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_Config{
		Registry: tmp,
	}
	return
}

func NewRegistry_ConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_Config, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_Config{
		Registry: tmp,
	}
	return
}

// SetCurrentControlSet sets the value of CurrentControlSet for the instance
func (instance *Registry_Config) SetPropertyCurrentControlSet(value uint32) (err error) {
	return instance.SetProperty("CurrentControlSet", (value))
}

// GetCurrentControlSet gets the value of CurrentControlSet for the instance
func (instance *Registry_Config) GetPropertyCurrentControlSet() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentControlSet")
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
