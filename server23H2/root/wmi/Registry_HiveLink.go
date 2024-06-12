// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Registry_HiveLink struct
type Registry_HiveLink struct {
	*Registry

	//
	Hive uint32

	//
	Path string
}

func NewRegistry_HiveLinkEx1(instance *cim.WmiInstance) (newInstance *Registry_HiveLink, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveLink{
		Registry: tmp,
	}
	return
}

func NewRegistry_HiveLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_HiveLink, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveLink{
		Registry: tmp,
	}
	return
}

// SetHive sets the value of Hive for the instance
func (instance *Registry_HiveLink) SetPropertyHive(value uint32) (err error) {
	return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_HiveLink) GetPropertyHive() (value uint32, err error) {
	retValue, err := instance.GetProperty("Hive")
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

// SetPath sets the value of Path for the instance
func (instance *Registry_HiveLink) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Registry_HiveLink) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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
