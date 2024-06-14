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

// FileExtensionsSettings struct
type FileExtensionsSettings struct {
	*EmbeddedObject

	//
	AllowUnlisted bool

	//
	ApplyToWebDAV bool

	//
	FileExtensions []FileExtensionElement
}

func NewFileExtensionsSettingsEx1(instance *cim.WmiInstance) (newInstance *FileExtensionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileExtensionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFileExtensionsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileExtensionsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileExtensionsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *FileExtensionsSettings) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *FileExtensionsSettings) GetPropertyAllowUnlisted() (value bool, err error) {
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

// SetApplyToWebDAV sets the value of ApplyToWebDAV for the instance
func (instance *FileExtensionsSettings) SetPropertyApplyToWebDAV(value bool) (err error) {
	return instance.SetProperty("ApplyToWebDAV", (value))
}

// GetApplyToWebDAV gets the value of ApplyToWebDAV for the instance
func (instance *FileExtensionsSettings) GetPropertyApplyToWebDAV() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplyToWebDAV")
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
func (instance *FileExtensionsSettings) SetPropertyFileExtensions(value []FileExtensionElement) (err error) {
	return instance.SetProperty("FileExtensions", (value))
}

// GetFileExtensions gets the value of FileExtensions for the instance
func (instance *FileExtensionsSettings) GetPropertyFileExtensions() (value []FileExtensionElement, err error) {
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
