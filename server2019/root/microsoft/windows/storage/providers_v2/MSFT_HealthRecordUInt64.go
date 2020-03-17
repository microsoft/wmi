// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_HealthRecordUInt64 struct
type MSFT_HealthRecordUInt64 struct {
	MSFT_HealthRecord

	//
	Value uint64
}

// SetValue sets the value of Value for the instance
func (instance *MSFT_HealthRecordUInt64) SetPropertyValue(value uint64) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_HealthRecordUInt64) GetPropertyValue() (value uint64, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
