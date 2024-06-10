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

// TextMessageEncodingBindingElement struct
type TextMessageEncodingBindingElement struct {
	*MessageEncodingBindingElement

	// The character set encoding to be used for emitting messages on the binding.
	Encoding string

	// An integer that defines how many messages can be read simultaneously without allocating new readers.
	MaxReadPoolSize int32

	// An integer that defines how many messages can be sent simultaneously without allocating new writers.
	MaxWritePoolSize int32

	// The quotas of the readers.
	ReaderQuotas XmlDictionaryReaderQuotas
}

func NewTextMessageEncodingBindingElementEx1(instance *cim.WmiInstance) (newInstance *TextMessageEncodingBindingElement, err error) {
	tmp, err := NewMessageEncodingBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TextMessageEncodingBindingElement{
		MessageEncodingBindingElement: tmp,
	}
	return
}

func NewTextMessageEncodingBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TextMessageEncodingBindingElement, err error) {
	tmp, err := NewMessageEncodingBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TextMessageEncodingBindingElement{
		MessageEncodingBindingElement: tmp,
	}
	return
}

// SetEncoding sets the value of Encoding for the instance
func (instance *TextMessageEncodingBindingElement) SetPropertyEncoding(value string) (err error) {
	return instance.SetProperty("Encoding", (value))
}

// GetEncoding gets the value of Encoding for the instance
func (instance *TextMessageEncodingBindingElement) GetPropertyEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("Encoding")
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

// SetMaxReadPoolSize sets the value of MaxReadPoolSize for the instance
func (instance *TextMessageEncodingBindingElement) SetPropertyMaxReadPoolSize(value int32) (err error) {
	return instance.SetProperty("MaxReadPoolSize", (value))
}

// GetMaxReadPoolSize gets the value of MaxReadPoolSize for the instance
func (instance *TextMessageEncodingBindingElement) GetPropertyMaxReadPoolSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxReadPoolSize")
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

// SetMaxWritePoolSize sets the value of MaxWritePoolSize for the instance
func (instance *TextMessageEncodingBindingElement) SetPropertyMaxWritePoolSize(value int32) (err error) {
	return instance.SetProperty("MaxWritePoolSize", (value))
}

// GetMaxWritePoolSize gets the value of MaxWritePoolSize for the instance
func (instance *TextMessageEncodingBindingElement) GetPropertyMaxWritePoolSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxWritePoolSize")
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

// SetReaderQuotas sets the value of ReaderQuotas for the instance
func (instance *TextMessageEncodingBindingElement) SetPropertyReaderQuotas(value XmlDictionaryReaderQuotas) (err error) {
	return instance.SetProperty("ReaderQuotas", (value))
}

// GetReaderQuotas gets the value of ReaderQuotas for the instance
func (instance *TextMessageEncodingBindingElement) GetPropertyReaderQuotas() (value XmlDictionaryReaderQuotas, err error) {
	retValue, err := instance.GetProperty("ReaderQuotas")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(XmlDictionaryReaderQuotas)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " XmlDictionaryReaderQuotas is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = XmlDictionaryReaderQuotas(valuetmp)

	return
}
