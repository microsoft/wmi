// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsmqBindingElementBase struct
type MsmqBindingElementBase struct {
	*TransportBindingElement

	// A URI that contains the location of the per-application dead letter queue, where messages that have expired or that have failed transfer or delivery are placed.
	CustomDeadLetterQueue string

	// An enumeration value that indicates the type of dead letter queue to use.
	DeadLetterQueue string

	// A value that indicates whether the messages processed by this binding are durable or volatile.
	Durable bool

	// A Boolean value that indicates whether messages processed by this binding will be received exactly once.
	ExactlyOnce bool

	// The maximum number of retry cycles to attempt delivery of messages to the receiving application.
	MaxRetryCycles int32

	// A boolean value that indicates whether messages received by this binding should be locked when received.
	ReceiveContextEnabled bool

	// The settings for poison message handling.
	ReceiveErrorHandling string

	// The maximum number of immediate retry attempts on a message that is read from the application queue.
	ReceiveRetryCount int32

	// A value that indicates the time delay between retry cycles when attempting to deliver a message that could not be delivered immediately.
	RetryCycleDelay string

	// The interval of time that indicates how long the messages processed by this binding can be in the queue before they expire.
	TimeToLive string

	// A Boolean value that indicates whether messages processed by this binding should be traced.
	UseMsmqTracing bool

	// A Boolean value that indicates whether copies of messages processed by this binding should be stored in the source journal queue.
	UseSourceJournal bool

	// Gets or sets the interval of time before a message locked by ReceiveContext is unlocked and returned to the Queue.
	ValidityDuration string
}

func NewMsmqBindingElementBaseEx1(instance *cim.WmiInstance) (newInstance *MsmqBindingElementBase, err error) {
	tmp, err := NewTransportBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsmqBindingElementBase{
		TransportBindingElement: tmp,
	}
	return
}

func NewMsmqBindingElementBaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsmqBindingElementBase, err error) {
	tmp, err := NewTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsmqBindingElementBase{
		TransportBindingElement: tmp,
	}
	return
}

// SetCustomDeadLetterQueue sets the value of CustomDeadLetterQueue for the instance
func (instance *MsmqBindingElementBase) SetPropertyCustomDeadLetterQueue(value string) (err error) {
	return instance.SetProperty("CustomDeadLetterQueue", (value))
}

// GetCustomDeadLetterQueue gets the value of CustomDeadLetterQueue for the instance
func (instance *MsmqBindingElementBase) GetPropertyCustomDeadLetterQueue() (value string, err error) {
	retValue, err := instance.GetProperty("CustomDeadLetterQueue")
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

// SetDeadLetterQueue sets the value of DeadLetterQueue for the instance
func (instance *MsmqBindingElementBase) SetPropertyDeadLetterQueue(value string) (err error) {
	return instance.SetProperty("DeadLetterQueue", (value))
}

// GetDeadLetterQueue gets the value of DeadLetterQueue for the instance
func (instance *MsmqBindingElementBase) GetPropertyDeadLetterQueue() (value string, err error) {
	retValue, err := instance.GetProperty("DeadLetterQueue")
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

// SetDurable sets the value of Durable for the instance
func (instance *MsmqBindingElementBase) SetPropertyDurable(value bool) (err error) {
	return instance.SetProperty("Durable", (value))
}

// GetDurable gets the value of Durable for the instance
func (instance *MsmqBindingElementBase) GetPropertyDurable() (value bool, err error) {
	retValue, err := instance.GetProperty("Durable")
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

// SetExactlyOnce sets the value of ExactlyOnce for the instance
func (instance *MsmqBindingElementBase) SetPropertyExactlyOnce(value bool) (err error) {
	return instance.SetProperty("ExactlyOnce", (value))
}

// GetExactlyOnce gets the value of ExactlyOnce for the instance
func (instance *MsmqBindingElementBase) GetPropertyExactlyOnce() (value bool, err error) {
	retValue, err := instance.GetProperty("ExactlyOnce")
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

// SetMaxRetryCycles sets the value of MaxRetryCycles for the instance
func (instance *MsmqBindingElementBase) SetPropertyMaxRetryCycles(value int32) (err error) {
	return instance.SetProperty("MaxRetryCycles", (value))
}

// GetMaxRetryCycles gets the value of MaxRetryCycles for the instance
func (instance *MsmqBindingElementBase) GetPropertyMaxRetryCycles() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxRetryCycles")
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

// SetReceiveContextEnabled sets the value of ReceiveContextEnabled for the instance
func (instance *MsmqBindingElementBase) SetPropertyReceiveContextEnabled(value bool) (err error) {
	return instance.SetProperty("ReceiveContextEnabled", (value))
}

// GetReceiveContextEnabled gets the value of ReceiveContextEnabled for the instance
func (instance *MsmqBindingElementBase) GetPropertyReceiveContextEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ReceiveContextEnabled")
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

// SetReceiveErrorHandling sets the value of ReceiveErrorHandling for the instance
func (instance *MsmqBindingElementBase) SetPropertyReceiveErrorHandling(value string) (err error) {
	return instance.SetProperty("ReceiveErrorHandling", (value))
}

// GetReceiveErrorHandling gets the value of ReceiveErrorHandling for the instance
func (instance *MsmqBindingElementBase) GetPropertyReceiveErrorHandling() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiveErrorHandling")
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

// SetReceiveRetryCount sets the value of ReceiveRetryCount for the instance
func (instance *MsmqBindingElementBase) SetPropertyReceiveRetryCount(value int32) (err error) {
	return instance.SetProperty("ReceiveRetryCount", (value))
}

// GetReceiveRetryCount gets the value of ReceiveRetryCount for the instance
func (instance *MsmqBindingElementBase) GetPropertyReceiveRetryCount() (value int32, err error) {
	retValue, err := instance.GetProperty("ReceiveRetryCount")
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

// SetRetryCycleDelay sets the value of RetryCycleDelay for the instance
func (instance *MsmqBindingElementBase) SetPropertyRetryCycleDelay(value string) (err error) {
	return instance.SetProperty("RetryCycleDelay", (value))
}

// GetRetryCycleDelay gets the value of RetryCycleDelay for the instance
func (instance *MsmqBindingElementBase) GetPropertyRetryCycleDelay() (value string, err error) {
	retValue, err := instance.GetProperty("RetryCycleDelay")
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

// SetTimeToLive sets the value of TimeToLive for the instance
func (instance *MsmqBindingElementBase) SetPropertyTimeToLive(value string) (err error) {
	return instance.SetProperty("TimeToLive", (value))
}

// GetTimeToLive gets the value of TimeToLive for the instance
func (instance *MsmqBindingElementBase) GetPropertyTimeToLive() (value string, err error) {
	retValue, err := instance.GetProperty("TimeToLive")
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

// SetUseMsmqTracing sets the value of UseMsmqTracing for the instance
func (instance *MsmqBindingElementBase) SetPropertyUseMsmqTracing(value bool) (err error) {
	return instance.SetProperty("UseMsmqTracing", (value))
}

// GetUseMsmqTracing gets the value of UseMsmqTracing for the instance
func (instance *MsmqBindingElementBase) GetPropertyUseMsmqTracing() (value bool, err error) {
	retValue, err := instance.GetProperty("UseMsmqTracing")
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

// SetUseSourceJournal sets the value of UseSourceJournal for the instance
func (instance *MsmqBindingElementBase) SetPropertyUseSourceJournal(value bool) (err error) {
	return instance.SetProperty("UseSourceJournal", (value))
}

// GetUseSourceJournal gets the value of UseSourceJournal for the instance
func (instance *MsmqBindingElementBase) GetPropertyUseSourceJournal() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSourceJournal")
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

// SetValidityDuration sets the value of ValidityDuration for the instance
func (instance *MsmqBindingElementBase) SetPropertyValidityDuration(value string) (err error) {
	return instance.SetProperty("ValidityDuration", (value))
}

// GetValidityDuration gets the value of ValidityDuration for the instance
func (instance *MsmqBindingElementBase) GetPropertyValidityDuration() (value string, err error) {
	retValue, err := instance.GetProperty("ValidityDuration")
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
