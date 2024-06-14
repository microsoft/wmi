// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AspCache struct
type AspCache struct {
	*EmbeddedObject

	//
	DiskTemplateCacheDirectory string

	//
	EnableTypelibCache bool

	//
	MaxDiskTemplateCacheFiles uint32

	//
	ScriptEngineCacheMax uint32

	//
	ScriptFileCacheSize uint32
}

func NewAspCacheEx1(instance *cim.WmiInstance) (newInstance *AspCache, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AspCache{
		EmbeddedObject: tmp,
	}
	return
}

func NewAspCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AspCache, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AspCache{
		EmbeddedObject: tmp,
	}
	return
}

// SetDiskTemplateCacheDirectory sets the value of DiskTemplateCacheDirectory for the instance
func (instance *AspCache) SetPropertyDiskTemplateCacheDirectory(value string) (err error) {
	return instance.SetProperty("DiskTemplateCacheDirectory", (value))
}

// GetDiskTemplateCacheDirectory gets the value of DiskTemplateCacheDirectory for the instance
func (instance *AspCache) GetPropertyDiskTemplateCacheDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DiskTemplateCacheDirectory")
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

// SetEnableTypelibCache sets the value of EnableTypelibCache for the instance
func (instance *AspCache) SetPropertyEnableTypelibCache(value bool) (err error) {
	return instance.SetProperty("EnableTypelibCache", (value))
}

// GetEnableTypelibCache gets the value of EnableTypelibCache for the instance
func (instance *AspCache) GetPropertyEnableTypelibCache() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableTypelibCache")
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

// SetMaxDiskTemplateCacheFiles sets the value of MaxDiskTemplateCacheFiles for the instance
func (instance *AspCache) SetPropertyMaxDiskTemplateCacheFiles(value uint32) (err error) {
	return instance.SetProperty("MaxDiskTemplateCacheFiles", (value))
}

// GetMaxDiskTemplateCacheFiles gets the value of MaxDiskTemplateCacheFiles for the instance
func (instance *AspCache) GetPropertyMaxDiskTemplateCacheFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDiskTemplateCacheFiles")
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

// SetScriptEngineCacheMax sets the value of ScriptEngineCacheMax for the instance
func (instance *AspCache) SetPropertyScriptEngineCacheMax(value uint32) (err error) {
	return instance.SetProperty("ScriptEngineCacheMax", (value))
}

// GetScriptEngineCacheMax gets the value of ScriptEngineCacheMax for the instance
func (instance *AspCache) GetPropertyScriptEngineCacheMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScriptEngineCacheMax")
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

// SetScriptFileCacheSize sets the value of ScriptFileCacheSize for the instance
func (instance *AspCache) SetPropertyScriptFileCacheSize(value uint32) (err error) {
	return instance.SetProperty("ScriptFileCacheSize", (value))
}

// GetScriptFileCacheSize gets the value of ScriptFileCacheSize for the instance
func (instance *AspCache) GetPropertyScriptFileCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScriptFileCacheSize")
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
