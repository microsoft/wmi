// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETFramework_NETCLRInterop struct
type Win32_PerfRawData_NETFramework_NETCLRInterop struct {
	Win32_PerfRawData

	//
	NumberofCCWs uint32

	//
	Numberofmarshalling uint32

	//
	NumberofStubs uint32

	//
	NumberofTLBexportsPersec uint32

	//
	NumberofTLBimportsPersec uint32
}

// SetNumberofCCWs sets the value of NumberofCCWs for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) SetPropertyNumberofCCWs(value uint32) (err error) {
	return instance.SetProperty("NumberofCCWs", value)
}

// GetNumberofCCWs gets the value of NumberofCCWs for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) GetPropertyNumberofCCWs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofCCWs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofmarshalling sets the value of Numberofmarshalling for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) SetPropertyNumberofmarshalling(value uint32) (err error) {
	return instance.SetProperty("Numberofmarshalling", value)
}

// GetNumberofmarshalling gets the value of Numberofmarshalling for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) GetPropertyNumberofmarshalling() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofmarshalling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofStubs sets the value of NumberofStubs for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) SetPropertyNumberofStubs(value uint32) (err error) {
	return instance.SetProperty("NumberofStubs", value)
}

// GetNumberofStubs gets the value of NumberofStubs for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) GetPropertyNumberofStubs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofStubs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofTLBexportsPersec sets the value of NumberofTLBexportsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) SetPropertyNumberofTLBexportsPersec(value uint32) (err error) {
	return instance.SetProperty("NumberofTLBexportsPersec", value)
}

// GetNumberofTLBexportsPersec gets the value of NumberofTLBexportsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) GetPropertyNumberofTLBexportsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofTLBexportsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofTLBimportsPersec sets the value of NumberofTLBimportsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) SetPropertyNumberofTLBimportsPersec(value uint32) (err error) {
	return instance.SetProperty("NumberofTLBimportsPersec", value)
}

// GetNumberofTLBimportsPersec gets the value of NumberofTLBimportsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRInterop) GetPropertyNumberofTLBimportsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofTLBimportsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
