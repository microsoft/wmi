// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_WebApplication struct
type MDM_WebApplication struct {
	*cim.WmiInstance

	//
	PackageName string

	//
	PackageUrl string

	//
	ShortcutFilename string

	//
	ShortcutFolder string
}

func NewMDM_WebApplicationEx1(instance *cim.WmiInstance) (newInstance *MDM_WebApplication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WebApplication{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WebApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WebApplication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WebApplication{
		WmiInstance: tmp,
	}
	return
}

// SetPackageName sets the value of PackageName for the instance
func (instance *MDM_WebApplication) SetPropertyPackageName(value string) (err error) {
	return instance.SetProperty("PackageName", value)
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_WebApplication) GetPropertyPackageName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageUrl sets the value of PackageUrl for the instance
func (instance *MDM_WebApplication) SetPropertyPackageUrl(value string) (err error) {
	return instance.SetProperty("PackageUrl", value)
}

// GetPackageUrl gets the value of PackageUrl for the instance
func (instance *MDM_WebApplication) GetPropertyPackageUrl() (value string, err error) {
	retValue, err := instance.GetProperty("PackageUrl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShortcutFilename sets the value of ShortcutFilename for the instance
func (instance *MDM_WebApplication) SetPropertyShortcutFilename(value string) (err error) {
	return instance.SetProperty("ShortcutFilename", value)
}

// GetShortcutFilename gets the value of ShortcutFilename for the instance
func (instance *MDM_WebApplication) GetPropertyShortcutFilename() (value string, err error) {
	retValue, err := instance.GetProperty("ShortcutFilename")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShortcutFolder sets the value of ShortcutFolder for the instance
func (instance *MDM_WebApplication) SetPropertyShortcutFolder(value string) (err error) {
	return instance.SetProperty("ShortcutFolder", value)
}

// GetShortcutFolder gets the value of ShortcutFolder for the instance
func (instance *MDM_WebApplication) GetPropertyShortcutFolder() (value string, err error) {
	retValue, err := instance.GetProperty("ShortcutFolder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
