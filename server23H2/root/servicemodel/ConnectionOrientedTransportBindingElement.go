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

// ConnectionOrientedTransportBindingElement struct
type ConnectionOrientedTransportBindingElement struct {
	*TransportBindingElement

	// The Timespan that specifies how long the channel initialization has to complete before timing out.
	ChannelInitializationTimeout string

	// The size of the buffer used to transmit a chunk of the serialized message on the wire from the client or service.
	ConnectionBufferSize int32

	// A value that indicates whether the hostname is used to reach the service when matching on the URI.
	HostNameComparisonMode string

	// The maximum size of the buffer to use.
	MaxBufferSize int32

	// The maximum interval of time that a chunk of a message or a full message can remain buffered in memory before being sent out.
	MaxOutputDelay string

	// The maximum number of pending asynchronous accept threads that are available for processing incoming connections on the service.
	MaxPendingAccepts int32

	// The maximum number of pending connections.
	MaxPendingConnections int32

	// A value that specifies whether the messages are buffered or streamed with the connection-oriented transport.
	TransferMode string
}

func NewConnectionOrientedTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *ConnectionOrientedTransportBindingElement, err error) {
	tmp, err := NewTransportBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConnectionOrientedTransportBindingElement{
		TransportBindingElement: tmp,
	}
	return
}

func NewConnectionOrientedTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConnectionOrientedTransportBindingElement, err error) {
	tmp, err := NewTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConnectionOrientedTransportBindingElement{
		TransportBindingElement: tmp,
	}
	return
}

// SetChannelInitializationTimeout sets the value of ChannelInitializationTimeout for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyChannelInitializationTimeout(value string) (err error) {
	return instance.SetProperty("ChannelInitializationTimeout", (value))
}

// GetChannelInitializationTimeout gets the value of ChannelInitializationTimeout for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyChannelInitializationTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ChannelInitializationTimeout")
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

// SetConnectionBufferSize sets the value of ConnectionBufferSize for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyConnectionBufferSize(value int32) (err error) {
	return instance.SetProperty("ConnectionBufferSize", (value))
}

// GetConnectionBufferSize gets the value of ConnectionBufferSize for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyConnectionBufferSize() (value int32, err error) {
	retValue, err := instance.GetProperty("ConnectionBufferSize")
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

// SetHostNameComparisonMode sets the value of HostNameComparisonMode for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyHostNameComparisonMode(value string) (err error) {
	return instance.SetProperty("HostNameComparisonMode", (value))
}

// GetHostNameComparisonMode gets the value of HostNameComparisonMode for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyHostNameComparisonMode() (value string, err error) {
	retValue, err := instance.GetProperty("HostNameComparisonMode")
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

// SetMaxBufferSize sets the value of MaxBufferSize for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyMaxBufferSize(value int32) (err error) {
	return instance.SetProperty("MaxBufferSize", (value))
}

// GetMaxBufferSize gets the value of MaxBufferSize for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyMaxBufferSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBufferSize")
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

// SetMaxOutputDelay sets the value of MaxOutputDelay for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyMaxOutputDelay(value string) (err error) {
	return instance.SetProperty("MaxOutputDelay", (value))
}

// GetMaxOutputDelay gets the value of MaxOutputDelay for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyMaxOutputDelay() (value string, err error) {
	retValue, err := instance.GetProperty("MaxOutputDelay")
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

// SetMaxPendingAccepts sets the value of MaxPendingAccepts for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyMaxPendingAccepts(value int32) (err error) {
	return instance.SetProperty("MaxPendingAccepts", (value))
}

// GetMaxPendingAccepts gets the value of MaxPendingAccepts for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyMaxPendingAccepts() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxPendingAccepts")
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

// SetMaxPendingConnections sets the value of MaxPendingConnections for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyMaxPendingConnections(value int32) (err error) {
	return instance.SetProperty("MaxPendingConnections", (value))
}

// GetMaxPendingConnections gets the value of MaxPendingConnections for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyMaxPendingConnections() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxPendingConnections")
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

// SetTransferMode sets the value of TransferMode for the instance
func (instance *ConnectionOrientedTransportBindingElement) SetPropertyTransferMode(value string) (err error) {
	return instance.SetProperty("TransferMode", (value))
}

// GetTransferMode gets the value of TransferMode for the instance
func (instance *ConnectionOrientedTransportBindingElement) GetPropertyTransferMode() (value string, err error) {
	retValue, err := instance.GetProperty("TransferMode")
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
