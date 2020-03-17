// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiCompareSuppression struct
type Msft_MiCompareSuppression struct {
	cim.WmiInstance

	//
	SuppressionSignal []interface{}

	//
	Timestamp string
}

// SetSuppressionSignal sets the value of SuppressionSignal for the instance
func (instance *Msft_MiCompareSuppression) SetPropertySuppressionSignal(value []interface{}) (err error) {
	return instance.SetProperty("SuppressionSignal", value)
}

// GetSuppressionSignal gets the value of SuppressionSignal for the instance
func (instance *Msft_MiCompareSuppression) GetPropertySuppressionSignal() (value []interface{}, err error) {
	retValue, err := instance.GetProperty("SuppressionSignal")
	if err != nil {
		return
	}
	value, ok := retValue.([]interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *Msft_MiCompareSuppression) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", value)
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *Msft_MiCompareSuppression) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
