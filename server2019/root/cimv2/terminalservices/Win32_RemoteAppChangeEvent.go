// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_RemoteAppChangeEvent struct
type Win32_RemoteAppChangeEvent struct {
	__ExtrinsicEvent

	// Type of operation corresponding to the event
	OperationType RemoteAppChangeEvent_OperationType

	// Object changed by the operation corresponding to the event
	TargetInstance interface{}
}

// SetOperationType sets the value of OperationType for the instance
func (instance *Win32_RemoteAppChangeEvent) SetPropertyOperationType(value RemoteAppChangeEvent_OperationType) (err error) {
	return instance.SetProperty("OperationType", value)
}

// GetOperationType gets the value of OperationType for the instance
func (instance *Win32_RemoteAppChangeEvent) GetPropertyOperationType() (value RemoteAppChangeEvent_OperationType, err error) {
	retValue, err := instance.GetProperty("OperationType")
	if err != nil {
		return
	}
	value, ok := retValue.(RemoteAppChangeEvent_OperationType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetInstance sets the value of TargetInstance for the instance
func (instance *Win32_RemoteAppChangeEvent) SetPropertyTargetInstance(value interface{}) (err error) {
	return instance.SetProperty("TargetInstance", value)
}

// GetTargetInstance gets the value of TargetInstance for the instance
func (instance *Win32_RemoteAppChangeEvent) GetPropertyTargetInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TargetInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
