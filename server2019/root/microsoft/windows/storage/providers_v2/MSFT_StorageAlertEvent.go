// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_StorageAlertEvent struct
type MSFT_StorageAlertEvent struct {
	MSFT_StorageEvent

	// This field describes the type of alert being received.
	AlertType StorageAlertEvent_AlertType
}

// SetAlertType sets the value of AlertType for the instance
func (instance *MSFT_StorageAlertEvent) SetPropertyAlertType(value StorageAlertEvent_AlertType) (err error) {
	return instance.SetProperty("AlertType", value)
}

// GetAlertType gets the value of AlertType for the instance
func (instance *MSFT_StorageAlertEvent) GetPropertyAlertType() (value StorageAlertEvent_AlertType, err error) {
	retValue, err := instance.GetProperty("AlertType")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageAlertEvent_AlertType)
	if !ok {
		// TODO: Set an error
	}
	return
}
