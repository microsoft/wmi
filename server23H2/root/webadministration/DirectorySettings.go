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

// DirectorySettings struct
type DirectorySettings struct {
	*EmbeddedObject

	//
	CodeSubDirectories []DirectoryElement
}

func NewDirectorySettingsEx1(instance *cim.WmiInstance) (newInstance *DirectorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DirectorySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDirectorySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DirectorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DirectorySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetCodeSubDirectories sets the value of CodeSubDirectories for the instance
func (instance *DirectorySettings) SetPropertyCodeSubDirectories(value []DirectoryElement) (err error) {
	return instance.SetProperty("CodeSubDirectories", (value))
}

// GetCodeSubDirectories gets the value of CodeSubDirectories for the instance
func (instance *DirectorySettings) GetPropertyCodeSubDirectories() (value []DirectoryElement, err error) {
	retValue, err := instance.GetProperty("CodeSubDirectories")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DirectoryElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DirectoryElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DirectoryElement(valuetmp))
	}

	return
}
