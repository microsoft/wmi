// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbShareChangeEvent struct
type MSFT_SmbShareChangeEvent struct {
	cim.WmiInstance

	//
	EventType SmbShareChangeEvent_EventType

	//
	Share MSFT_SmbShare
}

// SetEventType sets the value of EventType for the instance
func (instance *MSFT_SmbShareChangeEvent) SetPropertyEventType(value SmbShareChangeEvent_EventType) (err error) {
	return instance.SetProperty("EventType", value)
}

// GetEventType gets the value of EventType for the instance
func (instance *MSFT_SmbShareChangeEvent) GetPropertyEventType() (value SmbShareChangeEvent_EventType, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShareChangeEvent_EventType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShare sets the value of Share for the instance
func (instance *MSFT_SmbShareChangeEvent) SetPropertyShare(value MSFT_SmbShare) (err error) {
	return instance.SetProperty("Share", value)
}

// GetShare gets the value of Share for the instance
func (instance *MSFT_SmbShareChangeEvent) GetPropertyShare() (value MSFT_SmbShare, err error) {
	retValue, err := instance.GetProperty("Share")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_SmbShare)
	if !ok {
		// TODO: Set an error
	}
	return
}
