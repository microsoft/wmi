// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LoaderNewDllEvent struct
type LoaderNewDllEvent struct {
	*Image

	//
	FilePath string

	//
	LoadReason uint32

	//
	NewDllBaseAddress uint32

	//
	ParentDllBaseAddress uint32
}

func NewLoaderNewDllEventEx1(instance *cim.WmiInstance) (newInstance *LoaderNewDllEvent, err error) {
	tmp, err := NewImageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderNewDllEvent{
		Image: tmp,
	}
	return
}

func NewLoaderNewDllEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderNewDllEvent, err error) {
	tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderNewDllEvent{
		Image: tmp,
	}
	return
}

// SetFilePath sets the value of FilePath for the instance
func (instance *LoaderNewDllEvent) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", (value))
}

// GetFilePath gets the value of FilePath for the instance
func (instance *LoaderNewDllEvent) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
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

// SetLoadReason sets the value of LoadReason for the instance
func (instance *LoaderNewDllEvent) SetPropertyLoadReason(value uint32) (err error) {
	return instance.SetProperty("LoadReason", (value))
}

// GetLoadReason gets the value of LoadReason for the instance
func (instance *LoaderNewDllEvent) GetPropertyLoadReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadReason")
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

// SetNewDllBaseAddress sets the value of NewDllBaseAddress for the instance
func (instance *LoaderNewDllEvent) SetPropertyNewDllBaseAddress(value uint32) (err error) {
	return instance.SetProperty("NewDllBaseAddress", (value))
}

// GetNewDllBaseAddress gets the value of NewDllBaseAddress for the instance
func (instance *LoaderNewDllEvent) GetPropertyNewDllBaseAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewDllBaseAddress")
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

// SetParentDllBaseAddress sets the value of ParentDllBaseAddress for the instance
func (instance *LoaderNewDllEvent) SetPropertyParentDllBaseAddress(value uint32) (err error) {
	return instance.SetProperty("ParentDllBaseAddress", (value))
}

// GetParentDllBaseAddress gets the value of ParentDllBaseAddress for the instance
func (instance *LoaderNewDllEvent) GetPropertyParentDllBaseAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParentDllBaseAddress")
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
