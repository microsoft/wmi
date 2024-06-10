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

// SchemaImporterExtensionsSection struct
type SchemaImporterExtensionsSection struct {
	*ConfigurationSectionWithCollection

	//
	SchemaImporterExtensions []NameTypeElement
}

func NewSchemaImporterExtensionsSectionEx1(instance *cim.WmiInstance) (newInstance *SchemaImporterExtensionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SchemaImporterExtensionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSchemaImporterExtensionsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SchemaImporterExtensionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SchemaImporterExtensionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetSchemaImporterExtensions sets the value of SchemaImporterExtensions for the instance
func (instance *SchemaImporterExtensionsSection) SetPropertySchemaImporterExtensions(value []NameTypeElement) (err error) {
	return instance.SetProperty("SchemaImporterExtensions", (value))
}

// GetSchemaImporterExtensions gets the value of SchemaImporterExtensions for the instance
func (instance *SchemaImporterExtensionsSection) GetPropertySchemaImporterExtensions() (value []NameTypeElement, err error) {
	retValue, err := instance.GetProperty("SchemaImporterExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NameTypeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NameTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NameTypeElement(valuetmp))
	}

	return
}
