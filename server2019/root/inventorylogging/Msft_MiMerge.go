// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// Msft_MiMerge struct
type Msft_MiMerge struct {
	Msft_MiStream

	//
	Inputs []Msft_MiStream
}

// SetInputs sets the value of Inputs for the instance
func (instance *Msft_MiMerge) SetPropertyInputs(value []Msft_MiStream) (err error) {
	return instance.SetProperty("Inputs", value)
}

// GetInputs gets the value of Inputs for the instance
func (instance *Msft_MiMerge) GetPropertyInputs() (value []Msft_MiStream, err error) {
	retValue, err := instance.GetProperty("Inputs")
	if err != nil {
		return
	}
	value, ok := retValue.([]Msft_MiStream)
	if !ok {
		// TODO: Set an error
	}
	return
}
