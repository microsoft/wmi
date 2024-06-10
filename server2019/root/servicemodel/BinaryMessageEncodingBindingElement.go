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

// BinaryMessageEncodingBindingElement struct
type BinaryMessageEncodingBindingElement struct {
	*MessageEncodingBindingElement

	// The compression format to apply to the messages.
	CompressionFormat string

	// An integer that defines how many messages can be read simultaneously without allocating new readers.
	MaxReadPoolSize int32

	// A value that specifies the size, in bytes, of the buffer used for encoding.
	MaxSessionSize int32

	// An integer that defines how many messages can be sent simultaneously without allocating new writers.
	MaxWritePoolSize int32

	// The quotas of the readers.
	ReaderQuotas XmlDictionaryReaderQuotas
}

func NewBinaryMessageEncodingBindingElementEx1(instance *cim.WmiInstance) (newInstance *BinaryMessageEncodingBindingElement, err error) {
	tmp, err := NewMessageEncodingBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BinaryMessageEncodingBindingElement{
		MessageEncodingBindingElement: tmp,
	}
	return
}

func NewBinaryMessageEncodingBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BinaryMessageEncodingBindingElement, err error) {
	tmp, err := NewMessageEncodingBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BinaryMessageEncodingBindingElement{
		MessageEncodingBindingElement: tmp,
	}
	return
}

// SetCompressionFormat sets the value of CompressionFormat for the instance
func (instance *BinaryMessageEncodingBindingElement) SetPropertyCompressionFormat(value string) (err error) {
	return instance.SetProperty("CompressionFormat", (value))
}

// GetCompressionFormat gets the value of CompressionFormat for the instance
func (instance *BinaryMessageEncodingBindingElement) GetPropertyCompressionFormat() (value string, err error) {
	retValue, err := instance.GetProperty("CompressionFormat")
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
func (instance *BinaryMessageEncodingBindingElement) SetPropertyMaxReadPoolSize(value int32) (err error) {
	return instance.SetProperty("MaxReadPoolSize", (value))
}

// GetMaxReadPoolSize gets the value of MaxReadPoolSize for the instance
func (instance *BinaryMessageEncodingBindingElement) GetPropertyMaxReadPoolSize() (value int32, err error) {
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

// SetMaxSessionSize sets the value of MaxSessionSize for the instance
func (instance *BinaryMessageEncodingBindingElement) SetPropertyMaxSessionSize(value int32) (err error) {
	return instance.SetProperty("MaxSessionSize", (value))
}

// GetMaxSessionSize gets the value of MaxSessionSize for the instance
func (instance *BinaryMessageEncodingBindingElement) GetPropertyMaxSessionSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxSessionSize")
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
func (instance *BinaryMessageEncodingBindingElement) SetPropertyMaxWritePoolSize(value int32) (err error) {
	return instance.SetProperty("MaxWritePoolSize", (value))
}

// GetMaxWritePoolSize gets the value of MaxWritePoolSize for the instance
func (instance *BinaryMessageEncodingBindingElement) GetPropertyMaxWritePoolSize() (value int32, err error) {
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
func (instance *BinaryMessageEncodingBindingElement) SetPropertyReaderQuotas(value XmlDictionaryReaderQuotas) (err error) {
	return instance.SetProperty("ReaderQuotas", (value))
}

// GetReaderQuotas gets the value of ReaderQuotas for the instance
func (instance *BinaryMessageEncodingBindingElement) GetPropertyReaderQuotas() (value XmlDictionaryReaderQuotas, err error) {
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
