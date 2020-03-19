// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_InitiatorPortToiSCSISession struct
type MSFT_InitiatorPortToiSCSISession struct {
	*cim.WmiInstance

	//
	InitiatorPort MSFT_InitiatorPort

	//
	iSCSISession MSFT_iSCSISession
}

func NewMSFT_InitiatorPortToiSCSISessionEx1(instance *cim.WmiInstance) (newInstance *MSFT_InitiatorPortToiSCSISession, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorPortToiSCSISession{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_InitiatorPortToiSCSISessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_InitiatorPortToiSCSISession, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorPortToiSCSISession{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorPort sets the value of InitiatorPort for the instance
func (instance *MSFT_InitiatorPortToiSCSISession) SetPropertyInitiatorPort(value MSFT_InitiatorPort) (err error) {
	return instance.SetProperty("InitiatorPort", value)
}

// GetInitiatorPort gets the value of InitiatorPort for the instance
func (instance *MSFT_InitiatorPortToiSCSISession) GetPropertyInitiatorPort() (value MSFT_InitiatorPort, err error) {
	retValue, err := instance.GetProperty("InitiatorPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_InitiatorPort)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetiSCSISession sets the value of iSCSISession for the instance
func (instance *MSFT_InitiatorPortToiSCSISession) SetPropertyiSCSISession(value MSFT_iSCSISession) (err error) {
	return instance.SetProperty("iSCSISession", value)
}

// GetiSCSISession gets the value of iSCSISession for the instance
func (instance *MSFT_InitiatorPortToiSCSISession) GetPropertyiSCSISession() (value MSFT_iSCSISession, err error) {
	retValue, err := instance.GetProperty("iSCSISession")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_iSCSISession)
	if !ok {
		// TODO: Set an error
	}
	return
}
