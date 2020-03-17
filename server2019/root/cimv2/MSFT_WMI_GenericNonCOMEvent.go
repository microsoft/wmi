// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_WMI_GenericNonCOMEvent struct
type MSFT_WMI_GenericNonCOMEvent struct {
	__ExtrinsicEvent

	//
	ProcessId uint32

	//
	PropertyNames []string

	//
	PropertyValues []string

	//
	ProviderName string
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", value)
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPropertyNames sets the value of PropertyNames for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) SetPropertyPropertyNames(value []string) (err error) {
	return instance.SetProperty("PropertyNames", value)
}

// GetPropertyNames gets the value of PropertyNames for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) GetPropertyPropertyNames() (value []string, err error) {
	retValue, err := instance.GetProperty("PropertyNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPropertyValues sets the value of PropertyValues for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) SetPropertyPropertyValues(value []string) (err error) {
	return instance.SetProperty("PropertyValues", value)
}

// GetPropertyValues gets the value of PropertyValues for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) GetPropertyPropertyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("PropertyValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *MSFT_WMI_GenericNonCOMEvent) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
