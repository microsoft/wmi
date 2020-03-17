// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

// __EventDroppedEvent struct
type __EventDroppedEvent struct {
	__SystemEvent

	//
	Event __Event

	//
	IntendedConsumer __EventConsumer
}

// SetEvent sets the value of Event for the instance
func (instance *__EventDroppedEvent) SetPropertyEvent(value __Event) (err error) {
	return instance.SetProperty("Event", value)
}

// GetEvent gets the value of Event for the instance
func (instance *__EventDroppedEvent) GetPropertyEvent() (value __Event, err error) {
	retValue, err := instance.GetProperty("Event")
	if err != nil {
		return
	}
	value, ok := retValue.(__Event)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntendedConsumer sets the value of IntendedConsumer for the instance
func (instance *__EventDroppedEvent) SetPropertyIntendedConsumer(value __EventConsumer) (err error) {
	return instance.SetProperty("IntendedConsumer", value)
}

// GetIntendedConsumer gets the value of IntendedConsumer for the instance
func (instance *__EventDroppedEvent) GetPropertyIntendedConsumer() (value __EventConsumer, err error) {
	retValue, err := instance.GetProperty("IntendedConsumer")
	if err != nil {
		return
	}
	value, ok := retValue.(__EventConsumer)
	if !ok {
		// TODO: Set an error
	}
	return
}
