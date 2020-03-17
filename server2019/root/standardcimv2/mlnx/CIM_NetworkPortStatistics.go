// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_NetworkPortStatistics struct
type CIM_NetworkPortStatistics struct {
	CIM_StatisticalData

	//
	BytesReceived uint64

	//
	BytesTransmitted uint64

	//
	PacketsReceived uint64

	//
	PacketsTransmitted uint64
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *CIM_NetworkPortStatistics) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *CIM_NetworkPortStatistics) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTransmitted sets the value of BytesTransmitted for the instance
func (instance *CIM_NetworkPortStatistics) SetPropertyBytesTransmitted(value uint64) (err error) {
	return instance.SetProperty("BytesTransmitted", value)
}

// GetBytesTransmitted gets the value of BytesTransmitted for the instance
func (instance *CIM_NetworkPortStatistics) GetPropertyBytesTransmitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceived sets the value of PacketsReceived for the instance
func (instance *CIM_NetworkPortStatistics) SetPropertyPacketsReceived(value uint64) (err error) {
	return instance.SetProperty("PacketsReceived", value)
}

// GetPacketsReceived gets the value of PacketsReceived for the instance
func (instance *CIM_NetworkPortStatistics) GetPropertyPacketsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsTransmitted sets the value of PacketsTransmitted for the instance
func (instance *CIM_NetworkPortStatistics) SetPropertyPacketsTransmitted(value uint64) (err error) {
	return instance.SetProperty("PacketsTransmitted", value)
}

// GetPacketsTransmitted gets the value of PacketsTransmitted for the instance
func (instance *CIM_NetworkPortStatistics) GetPropertyPacketsTransmitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsTransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
