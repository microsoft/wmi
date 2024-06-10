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

// AssemblyCollectionElement struct
type AssemblyCollectionElement struct {
	*CollectionElement

	//
	AssemblyName string

	//
	PublicKey string

	//
	Version string
}

func NewAssemblyCollectionElementEx1(instance *cim.WmiInstance) (newInstance *AssemblyCollectionElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AssemblyCollectionElement{
		CollectionElement: tmp,
	}
	return
}

func NewAssemblyCollectionElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AssemblyCollectionElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AssemblyCollectionElement{
		CollectionElement: tmp,
	}
	return
}

// SetAssemblyName sets the value of AssemblyName for the instance
func (instance *AssemblyCollectionElement) SetPropertyAssemblyName(value string) (err error) {
	return instance.SetProperty("AssemblyName", (value))
}

// GetAssemblyName gets the value of AssemblyName for the instance
func (instance *AssemblyCollectionElement) GetPropertyAssemblyName() (value string, err error) {
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

// SetPublicKey sets the value of PublicKey for the instance
func (instance *AssemblyCollectionElement) SetPropertyPublicKey(value string) (err error) {
	return instance.SetProperty("PublicKey", (value))
}

// GetPublicKey gets the value of PublicKey for the instance
func (instance *AssemblyCollectionElement) GetPropertyPublicKey() (value string, err error) {
	retValue, err := instance.GetProperty("PublicKey")
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

// SetVersion sets the value of Version for the instance
func (instance *AssemblyCollectionElement) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *AssemblyCollectionElement) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
