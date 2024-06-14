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

// RequestLimitsElement struct
type RequestLimitsElement struct {
	*EmbeddedObject

	//
	HeaderLimits HeaderLimitsSettings

	//
	MaxAllowedContentLength uint32

	//
	MaxQueryString uint32

	//
	MaxUrl uint32
}

func NewRequestLimitsElementEx1(instance *cim.WmiInstance) (newInstance *RequestLimitsElement, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RequestLimitsElement{
		EmbeddedObject: tmp,
	}
	return
}

func NewRequestLimitsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RequestLimitsElement, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RequestLimitsElement{
		EmbeddedObject: tmp,
	}
	return
}

// SetHeaderLimits sets the value of HeaderLimits for the instance
func (instance *RequestLimitsElement) SetPropertyHeaderLimits(value HeaderLimitsSettings) (err error) {
	return instance.SetProperty("HeaderLimits", (value))
}

// GetHeaderLimits gets the value of HeaderLimits for the instance
func (instance *RequestLimitsElement) GetPropertyHeaderLimits() (value HeaderLimitsSettings, err error) {
	retValue, err := instance.GetProperty("HeaderLimits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HeaderLimitsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HeaderLimitsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HeaderLimitsSettings(valuetmp)

	return
}

// SetMaxAllowedContentLength sets the value of MaxAllowedContentLength for the instance
func (instance *RequestLimitsElement) SetPropertyMaxAllowedContentLength(value uint32) (err error) {
	return instance.SetProperty("MaxAllowedContentLength", (value))
}

// GetMaxAllowedContentLength gets the value of MaxAllowedContentLength for the instance
func (instance *RequestLimitsElement) GetPropertyMaxAllowedContentLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxAllowedContentLength")
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

// SetMaxQueryString sets the value of MaxQueryString for the instance
func (instance *RequestLimitsElement) SetPropertyMaxQueryString(value uint32) (err error) {
	return instance.SetProperty("MaxQueryString", (value))
}

// GetMaxQueryString gets the value of MaxQueryString for the instance
func (instance *RequestLimitsElement) GetPropertyMaxQueryString() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxQueryString")
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

// SetMaxUrl sets the value of MaxUrl for the instance
func (instance *RequestLimitsElement) SetPropertyMaxUrl(value uint32) (err error) {
	return instance.SetProperty("MaxUrl", (value))
}

// GetMaxUrl gets the value of MaxUrl for the instance
func (instance *RequestLimitsElement) GetPropertyMaxUrl() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxUrl")
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
