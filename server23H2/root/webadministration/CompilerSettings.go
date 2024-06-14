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

// CompilerSettings struct
type CompilerSettings struct {
	*EmbeddedObject

	//
	Compilers []CompilerElement
}

func NewCompilerSettingsEx1(instance *cim.WmiInstance) (newInstance *CompilerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CompilerSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewCompilerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CompilerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CompilerSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetCompilers sets the value of Compilers for the instance
func (instance *CompilerSettings) SetPropertyCompilers(value []CompilerElement) (err error) {
	return instance.SetProperty("Compilers", (value))
}

// GetCompilers gets the value of Compilers for the instance
func (instance *CompilerSettings) GetPropertyCompilers() (value []CompilerElement, err error) {
	retValue, err := instance.GetProperty("Compilers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(CompilerElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " CompilerElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, CompilerElement(valuetmp))
	}

	return
}
