// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbShareChangeEvent struct
type MSFT_SmbShareChangeEvent struct {
	*cim.WmiInstance

	//
	EventType SmbShareChangeEvent_EventType

	//
	Share MSFT_SmbShare
}

func NewMSFT_SmbShareChangeEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbShareChangeEvent, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbShareChangeEvent{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbShareChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbShareChangeEvent, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbShareChangeEvent{
		WmiInstance: tmp,
	}
	return
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
