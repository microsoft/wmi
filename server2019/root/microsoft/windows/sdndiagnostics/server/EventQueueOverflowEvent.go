// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SDNDiagnostics.Server
//////////////////////////////////////////////
package server

// __EventQueueOverflowEvent struct
type __EventQueueOverflowEvent struct {
	__EventDroppedEvent

	//
	CurrentQueueSize uint32
}

// SetCurrentQueueSize sets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) SetPropertyCurrentQueueSize(value uint32) (err error) {
	return instance.SetProperty("CurrentQueueSize", value)
}

// GetCurrentQueueSize gets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) GetPropertyCurrentQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentQueueSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
