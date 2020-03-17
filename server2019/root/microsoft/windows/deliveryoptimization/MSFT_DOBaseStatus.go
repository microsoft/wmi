// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DOBaseStatus struct
type MSFT_DOBaseStatus struct {
	cim.WmiInstance

	//
	Id uint8
}

// SetId sets the value of Id for the instance
func (instance *MSFT_DOBaseStatus) SetPropertyId(value uint8) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_DOBaseStatus) GetPropertyId() (value uint8, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
