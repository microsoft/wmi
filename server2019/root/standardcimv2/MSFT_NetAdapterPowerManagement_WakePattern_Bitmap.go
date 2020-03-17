// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterPowerManagement_WakePattern_Bitmap struct
type MSFT_NetAdapterPowerManagement_WakePattern_Bitmap struct {
	MSFT_NetAdapterPowerManagement_WakePattern

	//
	Mask []uint8

	//
	Pattern []uint8
}

// SetMask sets the value of Mask for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) SetPropertyMask(value []uint8) (err error) {
	return instance.SetProperty("Mask", value)
}

// GetMask gets the value of Mask for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) GetPropertyMask() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Mask")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPattern sets the value of Pattern for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) SetPropertyPattern(value []uint8) (err error) {
	return instance.SetProperty("Pattern", value)
}

// GetPattern gets the value of Pattern for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) GetPropertyPattern() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Pattern")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
