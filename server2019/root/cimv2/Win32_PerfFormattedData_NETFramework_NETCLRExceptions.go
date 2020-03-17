// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NETFramework_NETCLRExceptions struct
type Win32_PerfFormattedData_NETFramework_NETCLRExceptions struct {
	Win32_PerfFormattedData

	//
	NumberofExcepsThrown uint32

	//
	NumberofExcepsThrownPersec uint32

	//
	NumberofFiltersPersec uint32

	//
	NumberofFinallysPersec uint32

	//
	ThrowToCatchDepthPersec uint32
}

// SetNumberofExcepsThrown sets the value of NumberofExcepsThrown for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) SetPropertyNumberofExcepsThrown(value uint32) (err error) {
	return instance.SetProperty("NumberofExcepsThrown", value)
}

// GetNumberofExcepsThrown gets the value of NumberofExcepsThrown for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) GetPropertyNumberofExcepsThrown() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofExcepsThrown")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofExcepsThrownPersec sets the value of NumberofExcepsThrownPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) SetPropertyNumberofExcepsThrownPersec(value uint32) (err error) {
	return instance.SetProperty("NumberofExcepsThrownPersec", value)
}

// GetNumberofExcepsThrownPersec gets the value of NumberofExcepsThrownPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) GetPropertyNumberofExcepsThrownPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofExcepsThrownPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofFiltersPersec sets the value of NumberofFiltersPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) SetPropertyNumberofFiltersPersec(value uint32) (err error) {
	return instance.SetProperty("NumberofFiltersPersec", value)
}

// GetNumberofFiltersPersec gets the value of NumberofFiltersPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) GetPropertyNumberofFiltersPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofFiltersPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofFinallysPersec sets the value of NumberofFinallysPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) SetPropertyNumberofFinallysPersec(value uint32) (err error) {
	return instance.SetProperty("NumberofFinallysPersec", value)
}

// GetNumberofFinallysPersec gets the value of NumberofFinallysPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) GetPropertyNumberofFinallysPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofFinallysPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThrowToCatchDepthPersec sets the value of ThrowToCatchDepthPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) SetPropertyThrowToCatchDepthPersec(value uint32) (err error) {
	return instance.SetProperty("ThrowToCatchDepthPersec", value)
}

// GetThrowToCatchDepthPersec gets the value of ThrowToCatchDepthPersec for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRExceptions) GetPropertyThrowToCatchDepthPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThrowToCatchDepthPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
