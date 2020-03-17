// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_ExtensionEventSourceLink struct
type RSOP_ExtensionEventSourceLink struct {
	cim.WmiInstance

	//
	eventSource RSOP_ExtensionEventSource

	//
	extensionStatus RSOP_ExtensionStatus
}

// SeteventSource sets the value of eventSource for the instance
func (instance *RSOP_ExtensionEventSourceLink) SetPropertyeventSource(value RSOP_ExtensionEventSource) (err error) {
	return instance.SetProperty("eventSource", value)
}

// GeteventSource gets the value of eventSource for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyeventSource() (value RSOP_ExtensionEventSource, err error) {
	retValue, err := instance.GetProperty("eventSource")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_ExtensionEventSource)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetextensionStatus sets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) SetPropertyextensionStatus(value RSOP_ExtensionStatus) (err error) {
	return instance.SetProperty("extensionStatus", value)
}

// GetextensionStatus gets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyextensionStatus() (value RSOP_ExtensionStatus, err error) {
	retValue, err := instance.GetProperty("extensionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_ExtensionStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}
