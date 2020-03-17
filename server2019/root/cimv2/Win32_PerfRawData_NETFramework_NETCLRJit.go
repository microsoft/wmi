// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETFramework_NETCLRJit struct
type Win32_PerfRawData_NETFramework_NETCLRJit struct {
	Win32_PerfRawData

	//
	ILBytesJittedPersec uint32

	//
	NumberofILBytesJitted uint32

	//
	NumberofMethodsJitted uint32

	//
	PercentTimeinJit uint32

	//
	PercentTimeinJit_Base uint32

	//
	StandardJitFailures uint32

	//
	TotalNumberofILBytesJitted uint32
}

// SetILBytesJittedPersec sets the value of ILBytesJittedPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyILBytesJittedPersec(value uint32) (err error) {
	return instance.SetProperty("ILBytesJittedPersec", value)
}

// GetILBytesJittedPersec gets the value of ILBytesJittedPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyILBytesJittedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ILBytesJittedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofILBytesJitted sets the value of NumberofILBytesJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyNumberofILBytesJitted(value uint32) (err error) {
	return instance.SetProperty("NumberofILBytesJitted", value)
}

// GetNumberofILBytesJitted gets the value of NumberofILBytesJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyNumberofILBytesJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofILBytesJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofMethodsJitted sets the value of NumberofMethodsJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyNumberofMethodsJitted(value uint32) (err error) {
	return instance.SetProperty("NumberofMethodsJitted", value)
}

// GetNumberofMethodsJitted gets the value of NumberofMethodsJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyNumberofMethodsJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofMethodsJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeinJit sets the value of PercentTimeinJit for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyPercentTimeinJit(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinJit", value)
}

// GetPercentTimeinJit gets the value of PercentTimeinJit for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyPercentTimeinJit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentTimeinJit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeinJit_Base sets the value of PercentTimeinJit_Base for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyPercentTimeinJit_Base(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinJit_Base", value)
}

// GetPercentTimeinJit_Base gets the value of PercentTimeinJit_Base for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyPercentTimeinJit_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentTimeinJit_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardJitFailures sets the value of StandardJitFailures for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyStandardJitFailures(value uint32) (err error) {
	return instance.SetProperty("StandardJitFailures", value)
}

// GetStandardJitFailures gets the value of StandardJitFailures for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyStandardJitFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardJitFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofILBytesJitted sets the value of TotalNumberofILBytesJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) SetPropertyTotalNumberofILBytesJitted(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofILBytesJitted", value)
}

// GetTotalNumberofILBytesJitted gets the value of TotalNumberofILBytesJitted for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRJit) GetPropertyTotalNumberofILBytesJitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofILBytesJitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
