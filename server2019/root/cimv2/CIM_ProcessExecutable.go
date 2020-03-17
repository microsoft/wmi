// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ProcessExecutable struct
type CIM_ProcessExecutable struct {
	CIM_Dependency

	//
	BaseAddress uint64

	//
	GlobalProcessCount uint32

	//
	ModuleInstance uint32

	//
	ProcessCount uint32
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *CIM_ProcessExecutable) SetPropertyBaseAddress(value uint64) (err error) {
	return instance.SetProperty("BaseAddress", value)
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *CIM_ProcessExecutable) GetPropertyBaseAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("BaseAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGlobalProcessCount sets the value of GlobalProcessCount for the instance
func (instance *CIM_ProcessExecutable) SetPropertyGlobalProcessCount(value uint32) (err error) {
	return instance.SetProperty("GlobalProcessCount", value)
}

// GetGlobalProcessCount gets the value of GlobalProcessCount for the instance
func (instance *CIM_ProcessExecutable) GetPropertyGlobalProcessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("GlobalProcessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModuleInstance sets the value of ModuleInstance for the instance
func (instance *CIM_ProcessExecutable) SetPropertyModuleInstance(value uint32) (err error) {
	return instance.SetProperty("ModuleInstance", value)
}

// GetModuleInstance gets the value of ModuleInstance for the instance
func (instance *CIM_ProcessExecutable) GetPropertyModuleInstance() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModuleInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessCount sets the value of ProcessCount for the instance
func (instance *CIM_ProcessExecutable) SetPropertyProcessCount(value uint32) (err error) {
	return instance.SetProperty("ProcessCount", value)
}

// GetProcessCount gets the value of ProcessCount for the instance
func (instance *CIM_ProcessExecutable) GetPropertyProcessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
