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

// Registry_HiveRundown struct
type Registry_HiveRundown struct {
	*Registry

	//
	FileName string

	//
	Hive uint32

	//
	LinkPath string

	//
	LoadedKeyCount uint32

	//
	Size uint64
}

func NewRegistry_HiveRundownEx1(instance *cim.WmiInstance) (newInstance *Registry_HiveRundown, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveRundown{
		Registry: tmp,
	}
	return
}

func NewRegistry_HiveRundownEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_HiveRundown, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveRundown{
		Registry: tmp,
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *Registry_HiveRundown) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Registry_HiveRundown) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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

// SetHive sets the value of Hive for the instance
func (instance *Registry_HiveRundown) SetPropertyHive(value uint32) (err error) {
	return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_HiveRundown) GetPropertyHive() (value uint32, err error) {
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
func (instance *Registry_HiveRundown) SetPropertyLinkPath(value string) (err error) {
	return instance.SetProperty("LinkPath", (value))
}

// GetLinkPath gets the value of LinkPath for the instance
func (instance *Registry_HiveRundown) GetPropertyLinkPath() (value string, err error) {
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

// SetLoadedKeyCount sets the value of LoadedKeyCount for the instance
func (instance *Registry_HiveRundown) SetPropertyLoadedKeyCount(value uint32) (err error) {
	return instance.SetProperty("LoadedKeyCount", (value))
}

// GetLoadedKeyCount gets the value of LoadedKeyCount for the instance
func (instance *Registry_HiveRundown) GetPropertyLoadedKeyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadedKeyCount")
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

// SetSize sets the value of Size for the instance
func (instance *Registry_HiveRundown) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *Registry_HiveRundown) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
