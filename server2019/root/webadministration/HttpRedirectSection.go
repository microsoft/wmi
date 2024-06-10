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

// HttpRedirectSection struct
type HttpRedirectSection struct {
	*ConfigurationSectionWithCollection

	//
	ChildOnly bool

	//
	Destination string

	//
	Enabled bool

	//
	ExactDestination bool

	//
	HttpRedirect []WildcardRedirectElement

	//
	HttpResponseStatus int32
}

func NewHttpRedirectSectionEx1(instance *cim.WmiInstance) (newInstance *HttpRedirectSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpRedirectSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHttpRedirectSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpRedirectSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpRedirectSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetChildOnly sets the value of ChildOnly for the instance
func (instance *HttpRedirectSection) SetPropertyChildOnly(value bool) (err error) {
	return instance.SetProperty("ChildOnly", (value))
}

// GetChildOnly gets the value of ChildOnly for the instance
func (instance *HttpRedirectSection) GetPropertyChildOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("ChildOnly")
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

// SetDestination sets the value of Destination for the instance
func (instance *HttpRedirectSection) SetPropertyDestination(value string) (err error) {
	return instance.SetProperty("Destination", (value))
}

// GetDestination gets the value of Destination for the instance
func (instance *HttpRedirectSection) GetPropertyDestination() (value string, err error) {
	retValue, err := instance.GetProperty("Destination")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *HttpRedirectSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HttpRedirectSection) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetExactDestination sets the value of ExactDestination for the instance
func (instance *HttpRedirectSection) SetPropertyExactDestination(value bool) (err error) {
	return instance.SetProperty("ExactDestination", (value))
}

// GetExactDestination gets the value of ExactDestination for the instance
func (instance *HttpRedirectSection) GetPropertyExactDestination() (value bool, err error) {
	retValue, err := instance.GetProperty("ExactDestination")
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

// SetHttpRedirect sets the value of HttpRedirect for the instance
func (instance *HttpRedirectSection) SetPropertyHttpRedirect(value []WildcardRedirectElement) (err error) {
	return instance.SetProperty("HttpRedirect", (value))
}

// GetHttpRedirect gets the value of HttpRedirect for the instance
func (instance *HttpRedirectSection) GetPropertyHttpRedirect() (value []WildcardRedirectElement, err error) {
	retValue, err := instance.GetProperty("HttpRedirect")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WildcardRedirectElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WildcardRedirectElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WildcardRedirectElement(valuetmp))
	}

	return
}

// SetHttpResponseStatus sets the value of HttpResponseStatus for the instance
func (instance *HttpRedirectSection) SetPropertyHttpResponseStatus(value int32) (err error) {
	return instance.SetProperty("HttpResponseStatus", (value))
}

// GetHttpResponseStatus gets the value of HttpResponseStatus for the instance
func (instance *HttpRedirectSection) GetPropertyHttpResponseStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("HttpResponseStatus")
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
