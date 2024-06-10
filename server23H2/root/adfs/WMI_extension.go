// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ADFS
//
// ////////////////////////////////////////////
package adfs

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WMI_extension struct
type WMI_extension struct {
	*__Win32Provider

	//
	AssemblyName string

	//
	AssemblyPath string

	//
	CLRVersion string
}

func NewWMI_extensionEx1(instance *cim.WmiInstance) (newInstance *WMI_extension, err error) {
	tmp, err := New__Win32ProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WMI_extension{
		__Win32Provider: tmp,
	}
	return
}

func NewWMI_extensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WMI_extension, err error) {
	tmp, err := New__Win32ProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WMI_extension{
		__Win32Provider: tmp,
	}
	return
}

// SetAssemblyName sets the value of AssemblyName for the instance
func (instance *WMI_extension) SetPropertyAssemblyName(value string) (err error) {
	return instance.SetProperty("AssemblyName", (value))
}

// GetAssemblyName gets the value of AssemblyName for the instance
func (instance *WMI_extension) GetPropertyAssemblyName() (value string, err error) {
	retValue, err := instance.GetProperty("AssemblyName")
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

// SetAssemblyPath sets the value of AssemblyPath for the instance
func (instance *WMI_extension) SetPropertyAssemblyPath(value string) (err error) {
	return instance.SetProperty("AssemblyPath", (value))
}

// GetAssemblyPath gets the value of AssemblyPath for the instance
func (instance *WMI_extension) GetPropertyAssemblyPath() (value string, err error) {
	retValue, err := instance.GetProperty("AssemblyPath")
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

// SetCLRVersion sets the value of CLRVersion for the instance
func (instance *WMI_extension) SetPropertyCLRVersion(value string) (err error) {
	return instance.SetProperty("CLRVersion", (value))
}

// GetCLRVersion gets the value of CLRVersion for the instance
func (instance *WMI_extension) GetPropertyCLRVersion() (value string, err error) {
	retValue, err := instance.GetProperty("CLRVersion")
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
