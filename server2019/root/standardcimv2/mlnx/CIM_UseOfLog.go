// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_UseOfLog struct
type CIM_UseOfLog struct {
	CIM_Dependency

	//
	RecordedData string
}

// SetRecordedData sets the value of RecordedData for the instance
func (instance *CIM_UseOfLog) SetPropertyRecordedData(value string) (err error) {
	return instance.SetProperty("RecordedData", value)
}

// GetRecordedData gets the value of RecordedData for the instance
func (instance *CIM_UseOfLog) GetPropertyRecordedData() (value string, err error) {
	retValue, err := instance.GetProperty("RecordedData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
