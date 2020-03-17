// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_DAConnectionStatus struct
type MSFT_DAConnectionStatus struct {
	MSFT_NetSettingData

	//
	Status uint32

	//
	Substatus uint32
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DAConnectionStatus) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DAConnectionStatus) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubstatus sets the value of Substatus for the instance
func (instance *MSFT_DAConnectionStatus) SetPropertySubstatus(value uint32) (err error) {
	return instance.SetProperty("Substatus", value)
}

// GetSubstatus gets the value of Substatus for the instance
func (instance *MSFT_DAConnectionStatus) GetPropertySubstatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Substatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
