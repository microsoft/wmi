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

// Registry_HiveDirty struct
type Registry_HiveDirty struct {
	*Registry

	//
	DirtyReason uint32

	//
	Hive uint32

	//
	LinkPath string
}

func NewRegistry_HiveDirtyEx1(instance *cim.WmiInstance) (newInstance *Registry_HiveDirty, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveDirty{
		Registry: tmp,
	}
	return
}

func NewRegistry_HiveDirtyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_HiveDirty, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveDirty{
		Registry: tmp,
	}
	return
}

// SetDirtyReason sets the value of DirtyReason for the instance
func (instance *Registry_HiveDirty) SetPropertyDirtyReason(value uint32) (err error) {
	return instance.SetProperty("DirtyReason", (value))
}

// GetDirtyReason gets the value of DirtyReason for the instance
func (instance *Registry_HiveDirty) GetPropertyDirtyReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirtyReason")
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

// SetHive sets the value of Hive for the instance
func (instance *Registry_HiveDirty) SetPropertyHive(value uint32) (err error) {
	return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_HiveDirty) GetPropertyHive() (value uint32, err error) {
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

// SetLinkPath sets the value of LinkPath for the instance
func (instance *Registry_HiveDirty) SetPropertyLinkPath(value string) (err error) {
	return instance.SetProperty("LinkPath", (value))
}

// GetLinkPath gets the value of LinkPath for the instance
func (instance *Registry_HiveDirty) GetPropertyLinkPath() (value string, err error) {
	retValue, err := instance.GetProperty("LinkPath")
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
