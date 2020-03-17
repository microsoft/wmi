// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// MsftSil_WindowsUpdate struct
type MsftSil_WindowsUpdate struct {
	MsftSil_Data

	//
	ID string

	//
	InstallDate string
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
