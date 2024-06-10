// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WebDavAuthoringSection struct
type WebDavAuthoringSection struct {
	*ConfigurationSectionWithCollection

	//
	CompatFlags int32

	//
	Enabled bool

	//
	FileSystem WebDavFileSystemSettings

	//
	Locks WebDavLocksSettings

	//
	MaxAllowedXmlRequestLength uint32

	//
	Properties WebDavPropertiesSettings

	//
	RequireSsl bool
}

func NewWebDavAuthoringSectionEx1(instance *cim.WmiInstance) (newInstance *WebDavAuthoringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavAuthoringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewWebDavAuthoringSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavAuthoringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavAuthoringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetCompatFlags sets the value of CompatFlags for the instance
func (instance *WebDavAuthoringSection) SetPropertyCompatFlags(value int32) (err error) {
	return instance.SetProperty("CompatFlags", (value))
}

// GetCompatFlags gets the value of CompatFlags for the instance
func (instance *WebDavAuthoringSection) GetPropertyCompatFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("CompatFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *WebDavAuthoringSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *WebDavAuthoringSection) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetFileSystem sets the value of FileSystem for the instance
func (instance *WebDavAuthoringSection) SetPropertyFileSystem(value WebDavFileSystemSettings) (err error) {
	return instance.SetProperty("FileSystem", (value))
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *WebDavAuthoringSection) GetPropertyFileSystem() (value WebDavFileSystemSettings, err error) {
	retValue, err := instance.GetProperty("FileSystem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebDavFileSystemSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebDavFileSystemSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebDavFileSystemSettings(valuetmp)

	return
}

// SetLocks sets the value of Locks for the instance
func (instance *WebDavAuthoringSection) SetPropertyLocks(value WebDavLocksSettings) (err error) {
	return instance.SetProperty("Locks", (value))
}

// GetLocks gets the value of Locks for the instance
func (instance *WebDavAuthoringSection) GetPropertyLocks() (value WebDavLocksSettings, err error) {
	retValue, err := instance.GetProperty("Locks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebDavLocksSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebDavLocksSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebDavLocksSettings(valuetmp)

	return
}

// SetMaxAllowedXmlRequestLength sets the value of MaxAllowedXmlRequestLength for the instance
func (instance *WebDavAuthoringSection) SetPropertyMaxAllowedXmlRequestLength(value uint32) (err error) {
	return instance.SetProperty("MaxAllowedXmlRequestLength", (value))
}

// GetMaxAllowedXmlRequestLength gets the value of MaxAllowedXmlRequestLength for the instance
func (instance *WebDavAuthoringSection) GetPropertyMaxAllowedXmlRequestLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxAllowedXmlRequestLength")
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

// SetProperties sets the value of Properties for the instance
func (instance *WebDavAuthoringSection) SetPropertyProperties(value WebDavPropertiesSettings) (err error) {
	return instance.SetProperty("Properties", (value))
}

// GetProperties gets the value of Properties for the instance
func (instance *WebDavAuthoringSection) GetPropertyProperties() (value WebDavPropertiesSettings, err error) {
	retValue, err := instance.GetProperty("Properties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebDavPropertiesSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebDavPropertiesSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebDavPropertiesSettings(valuetmp)

	return
}

// SetRequireSsl sets the value of RequireSsl for the instance
func (instance *WebDavAuthoringSection) SetPropertyRequireSsl(value bool) (err error) {
	return instance.SetProperty("RequireSsl", (value))
}

// GetRequireSsl gets the value of RequireSsl for the instance
func (instance *WebDavAuthoringSection) GetPropertyRequireSsl() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSsl")
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
