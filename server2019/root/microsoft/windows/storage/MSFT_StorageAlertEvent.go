// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_StorageAlertEvent struct
type MSFT_StorageAlertEvent struct {
	MSFT_StorageEvent

	//
	AlertType uint16
}

// SetAlertType sets the value of AlertType for the instance
func (instance *MSFT_StorageAlertEvent) SetPropertyAlertType(value uint16) (err error) {
	return instance.SetProperty("AlertType", value)
}

// GetAlertType gets the value of AlertType for the instance
func (instance *MSFT_StorageAlertEvent) GetPropertyAlertType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AlertType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
