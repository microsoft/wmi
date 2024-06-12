// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WebHttpBehavior struct
type WebHttpBehavior struct {
	*Behavior

	// Specifies the web message body style.
	DefaultBodyStyle string

	// Specifies the default web message format for outgoing requests if the operation does not specify one.
	DefaultOutgoingRequestFormat string

	// Specifies the default web message format for outgoing responses if the operation does not specify one.
	DefaultOutgoingResponseFormat string
}

func NewWebHttpBehaviorEx1(instance *cim.WmiInstance) (newInstance *WebHttpBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebHttpBehavior{
		Behavior: tmp,
	}
	return
}

func NewWebHttpBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebHttpBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebHttpBehavior{
		Behavior: tmp,
	}
	return
}

// SetDefaultBodyStyle sets the value of DefaultBodyStyle for the instance
func (instance *WebHttpBehavior) SetPropertyDefaultBodyStyle(value string) (err error) {
	return instance.SetProperty("DefaultBodyStyle", (value))
}

// GetDefaultBodyStyle gets the value of DefaultBodyStyle for the instance
func (instance *WebHttpBehavior) GetPropertyDefaultBodyStyle() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultBodyStyle")
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

// SetDefaultOutgoingRequestFormat sets the value of DefaultOutgoingRequestFormat for the instance
func (instance *WebHttpBehavior) SetPropertyDefaultOutgoingRequestFormat(value string) (err error) {
	return instance.SetProperty("DefaultOutgoingRequestFormat", (value))
}

// GetDefaultOutgoingRequestFormat gets the value of DefaultOutgoingRequestFormat for the instance
func (instance *WebHttpBehavior) GetPropertyDefaultOutgoingRequestFormat() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultOutgoingRequestFormat")
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

// SetDefaultOutgoingResponseFormat sets the value of DefaultOutgoingResponseFormat for the instance
func (instance *WebHttpBehavior) SetPropertyDefaultOutgoingResponseFormat(value string) (err error) {
	return instance.SetProperty("DefaultOutgoingResponseFormat", (value))
}

// GetDefaultOutgoingResponseFormat gets the value of DefaultOutgoingResponseFormat for the instance
func (instance *WebHttpBehavior) GetPropertyDefaultOutgoingResponseFormat() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultOutgoingResponseFormat")
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
