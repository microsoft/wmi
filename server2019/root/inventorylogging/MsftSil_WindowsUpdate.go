// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MsftSil_WindowsUpdate struct
type MsftSil_WindowsUpdate struct {
	*MsftSil_Data

	//
	ID string

	//
	InstallDate string
}

func NewMsftSil_WindowsUpdateEx1(instance *cim.WmiInstance) (newInstance *MsftSil_WindowsUpdate, err error) {
	tmp, err := NewMsftSil_DataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftSil_WindowsUpdate{
		MsftSil_Data: tmp,
	}
	return
}

func NewMsftSil_WindowsUpdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftSil_WindowsUpdate, err error) {
	tmp, err := NewMsftSil_DataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftSil_WindowsUpdate{
		MsftSil_Data: tmp,
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MsftSil_WindowsUpdate) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MsftSil_WindowsUpdate) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallDate sets the value of InstallDate for the instance
func (instance *MsftSil_WindowsUpdate) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", value)
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *MsftSil_WindowsUpdate) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
