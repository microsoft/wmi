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

// AssemblyElement struct
type AssemblyElement struct {
	*CollectionElement

	//
	Assembly string
}

func NewAssemblyElementEx1(instance *cim.WmiInstance) (newInstance *AssemblyElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AssemblyElement{
		CollectionElement: tmp,
	}
	return
}

func NewAssemblyElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AssemblyElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AssemblyElement{
		CollectionElement: tmp,
	}
	return
}

// SetAssembly sets the value of Assembly for the instance
func (instance *AssemblyElement) SetPropertyAssembly(value string) (err error) {
	return instance.SetProperty("Assembly", (value))
}

// GetAssembly gets the value of Assembly for the instance
func (instance *AssemblyElement) GetPropertyAssembly() (value string, err error) {
	retValue, err := instance.GetProperty("Assembly")
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
