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

// AssemblySettings struct
type AssemblySettings struct {
	*EmbeddedObject

	//
	Assemblies []AssemblyElement
}

func NewAssemblySettingsEx1(instance *cim.WmiInstance) (newInstance *AssemblySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AssemblySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewAssemblySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AssemblySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AssemblySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAssemblies sets the value of Assemblies for the instance
func (instance *AssemblySettings) SetPropertyAssemblies(value []AssemblyElement) (err error) {
	return instance.SetProperty("Assemblies", (value))
}

// GetAssemblies gets the value of Assemblies for the instance
func (instance *AssemblySettings) GetPropertyAssemblies() (value []AssemblyElement, err error) {
	retValue, err := instance.GetProperty("Assemblies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(AssemblyElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " AssemblyElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AssemblyElement(valuetmp))
	}

	return
}
