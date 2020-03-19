// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageAlertEvent struct
type MSFT_StorageAlertEvent struct {
	*MSFT_StorageEvent

	// This field describes the type of alert being received.
	AlertType StorageAlertEvent_AlertType
}

func NewMSFT_StorageAlertEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageAlertEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageAlertEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

func NewMSFT_StorageAlertEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageAlertEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageAlertEvent{
		MSFT_StorageEvent: tmp,
	}
	return
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
