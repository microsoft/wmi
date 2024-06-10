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

// FileSettings struct
type FileSettings struct {
	*EmbeddedObject

	//
	Files []StringElement
}

func NewFileSettingsEx1(instance *cim.WmiInstance) (newInstance *FileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFileSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetFiles sets the value of Files for the instance
func (instance *FileSettings) SetPropertyFiles(value []StringElement) (err error) {
	return instance.SetProperty("Files", (value))
}

// GetFiles gets the value of Files for the instance
func (instance *FileSettings) GetPropertyFiles() (value []StringElement, err error) {
	retValue, err := instance.GetProperty("Files")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(StringElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " StringElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StringElement(valuetmp))
	}

	return
}
