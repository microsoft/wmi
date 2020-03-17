// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.HardwareManagement
//////////////////////////////////////////////
package hardwaremanagement

// MSFT_PCSVLogRecord struct
type MSFT_PCSVLogRecord struct {
	CIM_LogRecord

	//
	RawData []uint8
}

// SetRawData sets the value of RawData for the instance
func (instance *MSFT_PCSVLogRecord) SetPropertyRawData(value []uint8) (err error) {
	return instance.SetProperty("RawData", value)
}

// GetRawData gets the value of RawData for the instance
func (instance *MSFT_PCSVLogRecord) GetPropertyRawData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RawData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
