// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetCompartment struct
type MSFT_NetCompartment struct {
	MSFT_NetSettingData

	//
	CompartmentDescription string

	//
	CompartmentGuid string

	//
	CompartmentId uint32

	//
	CompartmentType uint32
}

// SetCompartmentDescription sets the value of CompartmentDescription for the instance
func (instance *MSFT_NetCompartment) SetPropertyCompartmentDescription(value string) (err error) {
	return instance.SetProperty("CompartmentDescription", value)
}

// GetCompartmentDescription gets the value of CompartmentDescription for the instance
func (instance *MSFT_NetCompartment) GetPropertyCompartmentDescription() (value string, err error) {
	retValue, err := instance.GetProperty("CompartmentDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompartmentGuid sets the value of CompartmentGuid for the instance
func (instance *MSFT_NetCompartment) SetPropertyCompartmentGuid(value string) (err error) {
	return instance.SetProperty("CompartmentGuid", value)
}

// GetCompartmentGuid gets the value of CompartmentGuid for the instance
func (instance *MSFT_NetCompartment) GetPropertyCompartmentGuid() (value string, err error) {
	retValue, err := instance.GetProperty("CompartmentGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompartmentId sets the value of CompartmentId for the instance
func (instance *MSFT_NetCompartment) SetPropertyCompartmentId(value uint32) (err error) {
	return instance.SetProperty("CompartmentId", value)
}

// GetCompartmentId gets the value of CompartmentId for the instance
func (instance *MSFT_NetCompartment) GetPropertyCompartmentId() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompartmentId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompartmentType sets the value of CompartmentType for the instance
func (instance *MSFT_NetCompartment) SetPropertyCompartmentType(value uint32) (err error) {
	return instance.SetProperty("CompartmentType", value)
}

// GetCompartmentType gets the value of CompartmentType for the instance
func (instance *MSFT_NetCompartment) GetPropertyCompartmentType() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompartmentType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
