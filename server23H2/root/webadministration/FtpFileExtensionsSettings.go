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

// FtpFileExtensionsSettings struct
type FtpFileExtensionsSettings struct {
	*EmbeddedObject

	//
	AllowUnlisted bool

	//
	FileExtensions []FileExtensionElement
}

func NewFtpFileExtensionsSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpFileExtensionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpFileExtensionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpFileExtensionsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpFileExtensionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpFileExtensionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *FtpFileExtensionsSettings) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *FtpFileExtensionsSettings) GetPropertyAllowUnlisted() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnlisted")
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

// SetFileExtensions sets the value of FileExtensions for the instance
func (instance *FtpFileExtensionsSettings) SetPropertyFileExtensions(value []FileExtensionElement) (err error) {
	return instance.SetProperty("FileExtensions", (value))
}

// GetFileExtensions gets the value of FileExtensions for the instance
func (instance *FtpFileExtensionsSettings) GetPropertyFileExtensions() (value []FileExtensionElement, err error) {
	retValue, err := instance.GetProperty("FileExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FileExtensionElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FileExtensionElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FileExtensionElement(valuetmp))
	}

	return
}
