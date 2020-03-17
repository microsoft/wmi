// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_QosCapabilities struct
type MSFT_NetAdapter_QosCapabilities struct {
	cim.WmiInstance

	//
	CeeDcbxSupported bool

	//
	IeeeDcbxSupported bool

	//
	MacSecBypassSupported bool

	//
	NumberOfEtsCapableTrafficClasses uint8

	//
	NumberOfPfcEnabledTrafficClasses uint8

	//
	NumberOfTrafficClasses uint8
}

// SetCeeDcbxSupported sets the value of CeeDcbxSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyCeeDcbxSupported(value bool) (err error) {
	return instance.SetProperty("CeeDcbxSupported", value)
}

// GetCeeDcbxSupported gets the value of CeeDcbxSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyCeeDcbxSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("CeeDcbxSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIeeeDcbxSupported sets the value of IeeeDcbxSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyIeeeDcbxSupported(value bool) (err error) {
	return instance.SetProperty("IeeeDcbxSupported", value)
}

// GetIeeeDcbxSupported gets the value of IeeeDcbxSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyIeeeDcbxSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IeeeDcbxSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMacSecBypassSupported sets the value of MacSecBypassSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyMacSecBypassSupported(value bool) (err error) {
	return instance.SetProperty("MacSecBypassSupported", value)
}

// GetMacSecBypassSupported gets the value of MacSecBypassSupported for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyMacSecBypassSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("MacSecBypassSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfEtsCapableTrafficClasses sets the value of NumberOfEtsCapableTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyNumberOfEtsCapableTrafficClasses(value uint8) (err error) {
	return instance.SetProperty("NumberOfEtsCapableTrafficClasses", value)
}

// GetNumberOfEtsCapableTrafficClasses gets the value of NumberOfEtsCapableTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyNumberOfEtsCapableTrafficClasses() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfEtsCapableTrafficClasses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfPfcEnabledTrafficClasses sets the value of NumberOfPfcEnabledTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyNumberOfPfcEnabledTrafficClasses(value uint8) (err error) {
	return instance.SetProperty("NumberOfPfcEnabledTrafficClasses", value)
}

// GetNumberOfPfcEnabledTrafficClasses gets the value of NumberOfPfcEnabledTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyNumberOfPfcEnabledTrafficClasses() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfPfcEnabledTrafficClasses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfTrafficClasses sets the value of NumberOfTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) SetPropertyNumberOfTrafficClasses(value uint8) (err error) {
	return instance.SetProperty("NumberOfTrafficClasses", value)
}

// GetNumberOfTrafficClasses gets the value of NumberOfTrafficClasses for the instance
func (instance *MSFT_NetAdapter_QosCapabilities) GetPropertyNumberOfTrafficClasses() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfTrafficClasses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
