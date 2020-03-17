// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_Volume_Repair_AsyncOutput struct
type MSFT_Volume_Repair_AsyncOutput struct {
	MSFT_StorageJobOutParams

	//
	Output uint32
}

// SetOutput sets the value of Output for the instance
func (instance *MSFT_Volume_Repair_AsyncOutput) SetPropertyOutput(value uint32) (err error) {
	return instance.SetProperty("Output", value)
}

// GetOutput gets the value of Output for the instance
func (instance *MSFT_Volume_Repair_AsyncOutput) GetPropertyOutput() (value uint32, err error) {
	retValue, err := instance.GetProperty("Output")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
