// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WebGetAttribute struct
type WebGetAttribute struct {
	*Behavior

	// Specifies the web message body style.
	BodyStyle string

	// Specifies whether a body style has been explicitly specified.
	IsBodyStyleSetExplicitly bool

	// Specifies whether the outgoing request's web message format has been explicitly specified.
	IsRequestFormatSetExplicitly bool

	// Specifies whether the outgoing response's web message format has been explicitly specified.
	IsResponseFormatSetExplicitly bool

	// Specifies the outgoing request's web message format.
	RequestFormat string

	// Specifies the outgoing response's web message format.
	ResponseFormat string

	// Specifies the URI template for the request.
	UriTemplate string
}

func NewWebGetAttributeEx1(instance *cim.WmiInstance) (newInstance *WebGetAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebGetAttribute{
		Behavior: tmp,
	}
	return
}

func NewWebGetAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebGetAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebGetAttribute{
		Behavior: tmp,
	}
	return
}

// SetBodyStyle sets the value of BodyStyle for the instance
func (instance *WebGetAttribute) SetPropertyBodyStyle(value string) (err error) {
	return instance.SetProperty("BodyStyle", (value))
}

// GetBodyStyle gets the value of BodyStyle for the instance
func (instance *WebGetAttribute) GetPropertyBodyStyle() (value string, err error) {
	retValue, err := instance.GetProperty("BodyStyle")
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

// SetIsBodyStyleSetExplicitly sets the value of IsBodyStyleSetExplicitly for the instance
func (instance *WebGetAttribute) SetPropertyIsBodyStyleSetExplicitly(value bool) (err error) {
	return instance.SetProperty("IsBodyStyleSetExplicitly", (value))
}

// GetIsBodyStyleSetExplicitly gets the value of IsBodyStyleSetExplicitly for the instance
func (instance *WebGetAttribute) GetPropertyIsBodyStyleSetExplicitly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBodyStyleSetExplicitly")
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

// SetIsRequestFormatSetExplicitly sets the value of IsRequestFormatSetExplicitly for the instance
func (instance *WebGetAttribute) SetPropertyIsRequestFormatSetExplicitly(value bool) (err error) {
	return instance.SetProperty("IsRequestFormatSetExplicitly", (value))
}

// GetIsRequestFormatSetExplicitly gets the value of IsRequestFormatSetExplicitly for the instance
func (instance *WebGetAttribute) GetPropertyIsRequestFormatSetExplicitly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsRequestFormatSetExplicitly")
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

// SetIsResponseFormatSetExplicitly sets the value of IsResponseFormatSetExplicitly for the instance
func (instance *WebGetAttribute) SetPropertyIsResponseFormatSetExplicitly(value bool) (err error) {
	return instance.SetProperty("IsResponseFormatSetExplicitly", (value))
}

// GetIsResponseFormatSetExplicitly gets the value of IsResponseFormatSetExplicitly for the instance
func (instance *WebGetAttribute) GetPropertyIsResponseFormatSetExplicitly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsResponseFormatSetExplicitly")
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

// SetRequestFormat sets the value of RequestFormat for the instance
func (instance *WebGetAttribute) SetPropertyRequestFormat(value string) (err error) {
	return instance.SetProperty("RequestFormat", (value))
}

// GetRequestFormat gets the value of RequestFormat for the instance
func (instance *WebGetAttribute) GetPropertyRequestFormat() (value string, err error) {
	retValue, err := instance.GetProperty("RequestFormat")
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

// SetResponseFormat sets the value of ResponseFormat for the instance
func (instance *WebGetAttribute) SetPropertyResponseFormat(value string) (err error) {
	return instance.SetProperty("ResponseFormat", (value))
}

// GetResponseFormat gets the value of ResponseFormat for the instance
func (instance *WebGetAttribute) GetPropertyResponseFormat() (value string, err error) {
	retValue, err := instance.GetProperty("ResponseFormat")
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

// SetUriTemplate sets the value of UriTemplate for the instance
func (instance *WebGetAttribute) SetPropertyUriTemplate(value string) (err error) {
	return instance.SetProperty("UriTemplate", (value))
}

// GetUriTemplate gets the value of UriTemplate for the instance
func (instance *WebGetAttribute) GetPropertyUriTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("UriTemplate")
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
