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

// HostingEnvironmentSection struct
type HostingEnvironmentSection struct {
	*ConfigurationSection

	//
	IdleTimeout string

	//
	ShadowCopyBinAssemblies bool

	//
	ShutdownTimeout string
}

func NewHostingEnvironmentSectionEx1(instance *cim.WmiInstance) (newInstance *HostingEnvironmentSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HostingEnvironmentSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewHostingEnvironmentSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HostingEnvironmentSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HostingEnvironmentSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *HostingEnvironmentSection) SetPropertyIdleTimeout(value string) (err error) {
	return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *HostingEnvironmentSection) GetPropertyIdleTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("IdleTimeout")
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

// SetShadowCopyBinAssemblies sets the value of ShadowCopyBinAssemblies for the instance
func (instance *HostingEnvironmentSection) SetPropertyShadowCopyBinAssemblies(value bool) (err error) {
	return instance.SetProperty("ShadowCopyBinAssemblies", (value))
}

// GetShadowCopyBinAssemblies gets the value of ShadowCopyBinAssemblies for the instance
func (instance *HostingEnvironmentSection) GetPropertyShadowCopyBinAssemblies() (value bool, err error) {
	retValue, err := instance.GetProperty("ShadowCopyBinAssemblies")
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

// SetShutdownTimeout sets the value of ShutdownTimeout for the instance
func (instance *HostingEnvironmentSection) SetPropertyShutdownTimeout(value string) (err error) {
	return instance.SetProperty("ShutdownTimeout", (value))
}

// GetShutdownTimeout gets the value of ShutdownTimeout for the instance
func (instance *HostingEnvironmentSection) GetPropertyShutdownTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ShutdownTimeout")
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
