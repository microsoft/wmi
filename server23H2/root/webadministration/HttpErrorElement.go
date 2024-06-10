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

// HttpErrorElement struct
type HttpErrorElement struct {
	*CollectionElement

	//
	Path string

	//
	PrefixLanguageFilePath string

	//
	ResponseMode int32

	//
	StatusCode uint32

	//
	SubStatusCode int32
}

func NewHttpErrorElementEx1(instance *cim.WmiInstance) (newInstance *HttpErrorElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpErrorElement{
		CollectionElement: tmp,
	}
	return
}

func NewHttpErrorElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpErrorElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpErrorElement{
		CollectionElement: tmp,
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *HttpErrorElement) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *HttpErrorElement) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPrefixLanguageFilePath sets the value of PrefixLanguageFilePath for the instance
func (instance *HttpErrorElement) SetPropertyPrefixLanguageFilePath(value string) (err error) {
	return instance.SetProperty("PrefixLanguageFilePath", (value))
}

// GetPrefixLanguageFilePath gets the value of PrefixLanguageFilePath for the instance
func (instance *HttpErrorElement) GetPropertyPrefixLanguageFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("PrefixLanguageFilePath")
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

// SetResponseMode sets the value of ResponseMode for the instance
func (instance *HttpErrorElement) SetPropertyResponseMode(value int32) (err error) {
	return instance.SetProperty("ResponseMode", (value))
}

// GetResponseMode gets the value of ResponseMode for the instance
func (instance *HttpErrorElement) GetPropertyResponseMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ResponseMode")
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

// SetStatusCode sets the value of StatusCode for the instance
func (instance *HttpErrorElement) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", (value))
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *HttpErrorElement) GetPropertyStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusCode")
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

// SetSubStatusCode sets the value of SubStatusCode for the instance
func (instance *HttpErrorElement) SetPropertySubStatusCode(value int32) (err error) {
	return instance.SetProperty("SubStatusCode", (value))
}

// GetSubStatusCode gets the value of SubStatusCode for the instance
func (instance *HttpErrorElement) GetPropertySubStatusCode() (value int32, err error) {
	retValue, err := instance.GetProperty("SubStatusCode")
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
