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

// RSOP_ExtensionStatus struct
type RSOP_ExtensionStatus struct {
	cim.WmiInstance

	//
	beginTime string

	//
	displayName string

	//
	endTime string

	//
	error uint32

	//
	extensionGuid string

	//
	loggingStatus uint32
}

// SetbeginTime sets the value of beginTime for the instance
func (instance *RSOP_ExtensionStatus) SetPropertybeginTime(value string) (err error) {
	return instance.SetProperty("beginTime", value)
}

// GetbeginTime gets the value of beginTime for the instance
func (instance *RSOP_ExtensionStatus) GetPropertybeginTime() (value string, err error) {
	retValue, err := instance.GetProperty("beginTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetdisplayName sets the value of displayName for the instance
func (instance *RSOP_ExtensionStatus) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", value)
}

// GetdisplayName gets the value of displayName for the instance
func (instance *RSOP_ExtensionStatus) GetPropertydisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("displayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetendTime sets the value of endTime for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyendTime(value string) (err error) {
	return instance.SetProperty("endTime", value)
}

// GetendTime gets the value of endTime for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyendTime() (value string, err error) {
	retValue, err := instance.GetProperty("endTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Seterror sets the value of error for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyerror(value uint32) (err error) {
	return instance.SetProperty("error", value)
}

// Geterror gets the value of error for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyerror() (value uint32, err error) {
	retValue, err := instance.GetProperty("error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetextensionGuid sets the value of extensionGuid for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyextensionGuid(value string) (err error) {
	return instance.SetProperty("extensionGuid", value)
}

// GetextensionGuid gets the value of extensionGuid for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyextensionGuid() (value string, err error) {
	retValue, err := instance.GetProperty("extensionGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetloggingStatus sets the value of loggingStatus for the instance
func (instance *RSOP_ExtensionStatus) SetPropertyloggingStatus(value uint32) (err error) {
	return instance.SetProperty("loggingStatus", value)
}

// GetloggingStatus gets the value of loggingStatus for the instance
func (instance *RSOP_ExtensionStatus) GetPropertyloggingStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("loggingStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
