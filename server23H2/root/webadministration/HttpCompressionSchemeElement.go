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

// HttpCompressionSchemeElement struct
type HttpCompressionSchemeElement struct {
	*CollectionElement

	//
	Dll string

	//
	DoDynamicCompression bool

	//
	DoStaticCompression bool

	//
	DynamicCompressionLevel uint32

	//
	Name string

	//
	StaticCompressionLevel uint32
}

func NewHttpCompressionSchemeElementEx1(instance *cim.WmiInstance) (newInstance *HttpCompressionSchemeElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpCompressionSchemeElement{
		CollectionElement: tmp,
	}
	return
}

func NewHttpCompressionSchemeElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpCompressionSchemeElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpCompressionSchemeElement{
		CollectionElement: tmp,
	}
	return
}

// SetDll sets the value of Dll for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyDll(value string) (err error) {
	return instance.SetProperty("Dll", (value))
}

// GetDll gets the value of Dll for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyDll() (value string, err error) {
	retValue, err := instance.GetProperty("Dll")
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

// SetDoDynamicCompression sets the value of DoDynamicCompression for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyDoDynamicCompression(value bool) (err error) {
	return instance.SetProperty("DoDynamicCompression", (value))
}

// GetDoDynamicCompression gets the value of DoDynamicCompression for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyDoDynamicCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("DoDynamicCompression")
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

// SetDoStaticCompression sets the value of DoStaticCompression for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyDoStaticCompression(value bool) (err error) {
	return instance.SetProperty("DoStaticCompression", (value))
}

// GetDoStaticCompression gets the value of DoStaticCompression for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyDoStaticCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("DoStaticCompression")
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

// SetDynamicCompressionLevel sets the value of DynamicCompressionLevel for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyDynamicCompressionLevel(value uint32) (err error) {
	return instance.SetProperty("DynamicCompressionLevel", (value))
}

// GetDynamicCompressionLevel gets the value of DynamicCompressionLevel for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyDynamicCompressionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicCompressionLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetStaticCompressionLevel sets the value of StaticCompressionLevel for the instance
func (instance *HttpCompressionSchemeElement) SetPropertyStaticCompressionLevel(value uint32) (err error) {
	return instance.SetProperty("StaticCompressionLevel", (value))
}

// GetStaticCompressionLevel gets the value of StaticCompressionLevel for the instance
func (instance *HttpCompressionSchemeElement) GetPropertyStaticCompressionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("StaticCompressionLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
