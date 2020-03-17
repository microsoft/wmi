// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetServiceFilter struct
type MSFT_NetServiceFilter struct {
	CIM_FilterEntryBase

	//
	ServiceName string
}

// SetServiceName sets the value of ServiceName for the instance
func (instance *MSFT_NetServiceFilter) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", value)
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *MSFT_NetServiceFilter) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
