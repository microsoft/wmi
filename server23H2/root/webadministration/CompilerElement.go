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

// CompilerElement struct
type CompilerElement struct {
	*CollectionElement

	//
	CompilerOptions string

	//
	Extension string

	//
	Language string

	//
	Type string

	//
	WarningLevel int32
}

func NewCompilerElementEx1(instance *cim.WmiInstance) (newInstance *CompilerElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CompilerElement{
		CollectionElement: tmp,
	}
	return
}

func NewCompilerElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CompilerElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CompilerElement{
		CollectionElement: tmp,
	}
	return
}

// SetCompilerOptions sets the value of CompilerOptions for the instance
func (instance *CompilerElement) SetPropertyCompilerOptions(value string) (err error) {
	return instance.SetProperty("CompilerOptions", (value))
}

// GetCompilerOptions gets the value of CompilerOptions for the instance
func (instance *CompilerElement) GetPropertyCompilerOptions() (value string, err error) {
	retValue, err := instance.GetProperty("CompilerOptions")
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

// SetExtension sets the value of Extension for the instance
func (instance *CompilerElement) SetPropertyExtension(value string) (err error) {
	return instance.SetProperty("Extension", (value))
}

// GetExtension gets the value of Extension for the instance
func (instance *CompilerElement) GetPropertyExtension() (value string, err error) {
	retValue, err := instance.GetProperty("Extension")
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

// SetLanguage sets the value of Language for the instance
func (instance *CompilerElement) SetPropertyLanguage(value string) (err error) {
	return instance.SetProperty("Language", (value))
}

// GetLanguage gets the value of Language for the instance
func (instance *CompilerElement) GetPropertyLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("Language")
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

// SetType sets the value of Type for the instance
func (instance *CompilerElement) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *CompilerElement) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetWarningLevel sets the value of WarningLevel for the instance
func (instance *CompilerElement) SetPropertyWarningLevel(value int32) (err error) {
	return instance.SetProperty("WarningLevel", (value))
}

// GetWarningLevel gets the value of WarningLevel for the instance
func (instance *CompilerElement) GetPropertyWarningLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("WarningLevel")
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
