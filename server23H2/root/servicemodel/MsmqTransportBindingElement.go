// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsmqTransportBindingElement struct
type MsmqTransportBindingElement struct {
	*MsmqBindingElementBase

	// An integer that specifies the maximum size of the pool that contains internal MSMQ message objects.
	MaxPoolSize int32

	// An enumeration value that indicates the queued communication channel transport that this binding uses.
	QueueTransferProtocol string

	// Returns a Boolean value that indicates whether queue addresses should be converted using Active Directory.
	UseActiveDirectory bool
}

func NewMsmqTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *MsmqTransportBindingElement, err error) {
	tmp, err := NewMsmqBindingElementBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsmqTransportBindingElement{
		MsmqBindingElementBase: tmp,
	}
	return
}

func NewMsmqTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsmqTransportBindingElement, err error) {
	tmp, err := NewMsmqBindingElementBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsmqTransportBindingElement{
		MsmqBindingElementBase: tmp,
	}
	return
}

// SetMaxPoolSize sets the value of MaxPoolSize for the instance
func (instance *MsmqTransportBindingElement) SetPropertyMaxPoolSize(value int32) (err error) {
	return instance.SetProperty("MaxPoolSize", (value))
}

// GetMaxPoolSize gets the value of MaxPoolSize for the instance
func (instance *MsmqTransportBindingElement) GetPropertyMaxPoolSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxPoolSize")
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

// SetQueueTransferProtocol sets the value of QueueTransferProtocol for the instance
func (instance *MsmqTransportBindingElement) SetPropertyQueueTransferProtocol(value string) (err error) {
	return instance.SetProperty("QueueTransferProtocol", (value))
}

// GetQueueTransferProtocol gets the value of QueueTransferProtocol for the instance
func (instance *MsmqTransportBindingElement) GetPropertyQueueTransferProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("QueueTransferProtocol")
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

// SetUseActiveDirectory sets the value of UseActiveDirectory for the instance
func (instance *MsmqTransportBindingElement) SetPropertyUseActiveDirectory(value bool) (err error) {
	return instance.SetProperty("UseActiveDirectory", (value))
}

// GetUseActiveDirectory gets the value of UseActiveDirectory for the instance
func (instance *MsmqTransportBindingElement) GetPropertyUseActiveDirectory() (value bool, err error) {
	retValue, err := instance.GetProperty("UseActiveDirectory")
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
