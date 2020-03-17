// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NCProvClientConnected struct
type MSFT_NCProvClientConnected struct {
	MSFT_NCProvEvent

	//
	Inproc bool
}

// SetInproc sets the value of Inproc for the instance
func (instance *MSFT_NCProvClientConnected) SetPropertyInproc(value bool) (err error) {
	return instance.SetProperty("Inproc", value)
}

// GetInproc gets the value of Inproc for the instance
func (instance *MSFT_NCProvClientConnected) GetPropertyInproc() (value bool, err error) {
	retValue, err := instance.GetProperty("Inproc")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
