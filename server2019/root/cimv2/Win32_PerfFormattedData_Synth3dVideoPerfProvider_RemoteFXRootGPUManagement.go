// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement struct
type Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement struct {
	Win32_PerfFormattedData

	//
	ResourcesVMsrunningRemoteFX uint64

	//
	VRAMAvailableMBperGPU uint64

	//
	VRAMReservedPercentperGPU uint64
}

// SetResourcesVMsrunningRemoteFX sets the value of ResourcesVMsrunningRemoteFX for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyResourcesVMsrunningRemoteFX(value uint64) (err error) {
	return instance.SetProperty("ResourcesVMsrunningRemoteFX", value)
}

// GetResourcesVMsrunningRemoteFX gets the value of ResourcesVMsrunningRemoteFX for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyResourcesVMsrunningRemoteFX() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourcesVMsrunningRemoteFX")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVRAMAvailableMBperGPU sets the value of VRAMAvailableMBperGPU for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyVRAMAvailableMBperGPU(value uint64) (err error) {
	return instance.SetProperty("VRAMAvailableMBperGPU", value)
}

// GetVRAMAvailableMBperGPU gets the value of VRAMAvailableMBperGPU for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyVRAMAvailableMBperGPU() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMAvailableMBperGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVRAMReservedPercentperGPU sets the value of VRAMReservedPercentperGPU for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyVRAMReservedPercentperGPU(value uint64) (err error) {
	return instance.SetProperty("VRAMReservedPercentperGPU", value)
}

// GetVRAMReservedPercentperGPU gets the value of VRAMReservedPercentperGPU for the instance
func (instance *Win32_PerfFormattedData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyVRAMReservedPercentperGPU() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMReservedPercentperGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
