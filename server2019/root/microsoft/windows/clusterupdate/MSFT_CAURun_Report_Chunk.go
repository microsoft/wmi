// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

// MSFT_CAURun_Report_Chunk struct
type MSFT_CAURun_Report_Chunk struct {
	MSFT_CAURun_Report_ID

	//
	Data string

	//
	SequenceNumber uint32
}

// SetData sets the value of Data for the instance
func (instance *MSFT_CAURun_Report_Chunk) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_CAURun_Report_Chunk) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *MSFT_CAURun_Report_Chunk) SetPropertySequenceNumber(value uint32) (err error) {
	return instance.SetProperty("SequenceNumber", value)
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *MSFT_CAURun_Report_Chunk) GetPropertySequenceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
