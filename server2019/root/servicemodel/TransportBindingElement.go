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

// TransportBindingElement struct
type TransportBindingElement struct {
	*BindingElement

	// A Boolean value that specifies if the user wants to take control of message addressing.
	ManualAddressing bool

	// The maximum buffer pool size for the binding.
	MaxBufferPoolSize int64

	// The maximum size for a message that is processed by this binding.
	MaxReceivedMessageSize int64

	// The URI scheme for the transport.
	Scheme string
}

func NewTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *TransportBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransportBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransportBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransportBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetManualAddressing sets the value of ManualAddressing for the instance
func (instance *TransportBindingElement) SetPropertyManualAddressing(value bool) (err error) {
	return instance.SetProperty("ManualAddressing", (value))
}

// GetManualAddressing gets the value of ManualAddressing for the instance
func (instance *TransportBindingElement) GetPropertyManualAddressing() (value bool, err error) {
	retValue, err := instance.GetProperty("ManualAddressing")
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

// SetMaxBufferPoolSize sets the value of MaxBufferPoolSize for the instance
func (instance *TransportBindingElement) SetPropertyMaxBufferPoolSize(value int64) (err error) {
	return instance.SetProperty("MaxBufferPoolSize", (value))
}

// GetMaxBufferPoolSize gets the value of MaxBufferPoolSize for the instance
func (instance *TransportBindingElement) GetPropertyMaxBufferPoolSize() (value int64, err error) {
	retValue, err := instance.GetProperty("MaxBufferPoolSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetMaxReceivedMessageSize sets the value of MaxReceivedMessageSize for the instance
func (instance *TransportBindingElement) SetPropertyMaxReceivedMessageSize(value int64) (err error) {
	return instance.SetProperty("MaxReceivedMessageSize", (value))
}

// GetMaxReceivedMessageSize gets the value of MaxReceivedMessageSize for the instance
func (instance *TransportBindingElement) GetPropertyMaxReceivedMessageSize() (value int64, err error) {
	retValue, err := instance.GetProperty("MaxReceivedMessageSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetScheme sets the value of Scheme for the instance
func (instance *TransportBindingElement) SetPropertyScheme(value string) (err error) {
	return instance.SetProperty("Scheme", (value))
}

// GetScheme gets the value of Scheme for the instance
func (instance *TransportBindingElement) GetPropertyScheme() (value string, err error) {
	retValue, err := instance.GetProperty("Scheme")
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
