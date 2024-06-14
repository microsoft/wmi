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

// ExtendedProtectionSettings struct
type ExtendedProtectionSettings struct {
	*EmbeddedObject

	//
	ExtendedProtection []NameElement

	//
	Flags int32

	//
	TokenChecking int32
}

func NewExtendedProtectionSettingsEx1(instance *cim.WmiInstance) (newInstance *ExtendedProtectionSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ExtendedProtectionSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewExtendedProtectionSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ExtendedProtectionSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ExtendedProtectionSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetExtendedProtection sets the value of ExtendedProtection for the instance
func (instance *ExtendedProtectionSettings) SetPropertyExtendedProtection(value []NameElement) (err error) {
	return instance.SetProperty("ExtendedProtection", (value))
}

// GetExtendedProtection gets the value of ExtendedProtection for the instance
func (instance *ExtendedProtectionSettings) GetPropertyExtendedProtection() (value []NameElement, err error) {
	retValue, err := instance.GetProperty("ExtendedProtection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NameElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NameElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NameElement(valuetmp))
	}

	return
}

// SetFlags sets the value of Flags for the instance
func (instance *ExtendedProtectionSettings) SetPropertyFlags(value int32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *ExtendedProtectionSettings) GetPropertyFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetTokenChecking sets the value of TokenChecking for the instance
func (instance *ExtendedProtectionSettings) SetPropertyTokenChecking(value int32) (err error) {
	return instance.SetProperty("TokenChecking", (value))
}

// GetTokenChecking gets the value of TokenChecking for the instance
func (instance *ExtendedProtectionSettings) GetPropertyTokenChecking() (value int32, err error) {
	retValue, err := instance.GetProperty("TokenChecking")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
