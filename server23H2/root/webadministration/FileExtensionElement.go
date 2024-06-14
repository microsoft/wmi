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

// FileExtensionElement struct
type FileExtensionElement struct {
	*CollectionElement

	//
	Allowed bool

	//
	FileExtension string
}

func NewFileExtensionElementEx1(instance *cim.WmiInstance) (newInstance *FileExtensionElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileExtensionElement{
		CollectionElement: tmp,
	}
	return
}

func NewFileExtensionElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileExtensionElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileExtensionElement{
		CollectionElement: tmp,
	}
	return
}

// SetAllowed sets the value of Allowed for the instance
func (instance *FileExtensionElement) SetPropertyAllowed(value bool) (err error) {
	return instance.SetProperty("Allowed", (value))
}

// GetAllowed gets the value of Allowed for the instance
func (instance *FileExtensionElement) GetPropertyAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("Allowed")
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

// SetFileExtension sets the value of FileExtension for the instance
func (instance *FileExtensionElement) SetPropertyFileExtension(value string) (err error) {
	return instance.SetProperty("FileExtension", (value))
}

// GetFileExtension gets the value of FileExtension for the instance
func (instance *FileExtensionElement) GetPropertyFileExtension() (value string, err error) {
	retValue, err := instance.GetProperty("FileExtension")
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
