// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.subscription
//
// ////////////////////////////////////////////
package subscription

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SMTPEventConsumer struct
type SMTPEventConsumer struct {
	*__EventConsumer

	//
	BccLine string

	//
	CcLine string

	//
	FromLine string

	//
	HeaderFields []string

	//
	Message string

	//
	Name string

	//
	ReplyToLine string

	//
	SMTPServer string

	//
	Subject string

	//
	ToLine string
}

func NewSMTPEventConsumerEx1(instance *cim.WmiInstance) (newInstance *SMTPEventConsumer, err error) {
	tmp, err := New__EventConsumerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SMTPEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

func NewSMTPEventConsumerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SMTPEventConsumer, err error) {
	tmp, err := New__EventConsumerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SMTPEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

// SetBccLine sets the value of BccLine for the instance
func (instance *SMTPEventConsumer) SetPropertyBccLine(value string) (err error) {
	return instance.SetProperty("BccLine", (value))
}

// GetBccLine gets the value of BccLine for the instance
func (instance *SMTPEventConsumer) GetPropertyBccLine() (value string, err error) {
	retValue, err := instance.GetProperty("BccLine")
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

// SetCcLine sets the value of CcLine for the instance
func (instance *SMTPEventConsumer) SetPropertyCcLine(value string) (err error) {
	return instance.SetProperty("CcLine", (value))
}

// GetCcLine gets the value of CcLine for the instance
func (instance *SMTPEventConsumer) GetPropertyCcLine() (value string, err error) {
	retValue, err := instance.GetProperty("CcLine")
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

// SetFromLine sets the value of FromLine for the instance
func (instance *SMTPEventConsumer) SetPropertyFromLine(value string) (err error) {
	return instance.SetProperty("FromLine", (value))
}

// GetFromLine gets the value of FromLine for the instance
func (instance *SMTPEventConsumer) GetPropertyFromLine() (value string, err error) {
	retValue, err := instance.GetProperty("FromLine")
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

// SetHeaderFields sets the value of HeaderFields for the instance
func (instance *SMTPEventConsumer) SetPropertyHeaderFields(value []string) (err error) {
	return instance.SetProperty("HeaderFields", (value))
}

// GetHeaderFields gets the value of HeaderFields for the instance
func (instance *SMTPEventConsumer) GetPropertyHeaderFields() (value []string, err error) {
	retValue, err := instance.GetProperty("HeaderFields")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMessage sets the value of Message for the instance
func (instance *SMTPEventConsumer) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", (value))
}

// GetMessage gets the value of Message for the instance
func (instance *SMTPEventConsumer) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
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

// SetName sets the value of Name for the instance
func (instance *SMTPEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SMTPEventConsumer) GetPropertyName() (value string, err error) {
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

// SetReplyToLine sets the value of ReplyToLine for the instance
func (instance *SMTPEventConsumer) SetPropertyReplyToLine(value string) (err error) {
	return instance.SetProperty("ReplyToLine", (value))
}

// GetReplyToLine gets the value of ReplyToLine for the instance
func (instance *SMTPEventConsumer) GetPropertyReplyToLine() (value string, err error) {
	retValue, err := instance.GetProperty("ReplyToLine")
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

// SetSMTPServer sets the value of SMTPServer for the instance
func (instance *SMTPEventConsumer) SetPropertySMTPServer(value string) (err error) {
	return instance.SetProperty("SMTPServer", (value))
}

// GetSMTPServer gets the value of SMTPServer for the instance
func (instance *SMTPEventConsumer) GetPropertySMTPServer() (value string, err error) {
	retValue, err := instance.GetProperty("SMTPServer")
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

// SetSubject sets the value of Subject for the instance
func (instance *SMTPEventConsumer) SetPropertySubject(value string) (err error) {
	return instance.SetProperty("Subject", (value))
}

// GetSubject gets the value of Subject for the instance
func (instance *SMTPEventConsumer) GetPropertySubject() (value string, err error) {
	retValue, err := instance.GetProperty("Subject")
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

// SetToLine sets the value of ToLine for the instance
func (instance *SMTPEventConsumer) SetPropertyToLine(value string) (err error) {
	return instance.SetProperty("ToLine", (value))
}

// GetToLine gets the value of ToLine for the instance
func (instance *SMTPEventConsumer) GetPropertyToLine() (value string, err error) {
	retValue, err := instance.GetProperty("ToLine")
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
