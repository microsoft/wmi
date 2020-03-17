// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_ClussvcPerfProvider_ClusterResources struct
type Win32_PerfRawData_ClussvcPerfProvider_ClusterResources struct {
	Win32_PerfRawData

	//
	ResourceControls uint64

	//
	ResourceControlsPersec uint64

	//
	ResourceFailure uint64

	//
	ResourceFailureAccessViolation uint64

	//
	ResourceFailureDeadlock uint64

	//
	ResourcesOnline uint64

	//
	ResourceTypeControls uint64

	//
	ResourceTypeControlsPersec uint64
}

// SetResourceControls sets the value of ResourceControls for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceControls(value uint64) (err error) {
	return instance.SetProperty("ResourceControls", value)
}

// GetResourceControls gets the value of ResourceControls for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceControls() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceControls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceControlsPersec sets the value of ResourceControlsPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceControlsPersec(value uint64) (err error) {
	return instance.SetProperty("ResourceControlsPersec", value)
}

// GetResourceControlsPersec gets the value of ResourceControlsPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceControlsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceControlsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceFailure sets the value of ResourceFailure for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceFailure(value uint64) (err error) {
	return instance.SetProperty("ResourceFailure", value)
}

// GetResourceFailure gets the value of ResourceFailure for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceFailure() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceFailure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceFailureAccessViolation sets the value of ResourceFailureAccessViolation for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceFailureAccessViolation(value uint64) (err error) {
	return instance.SetProperty("ResourceFailureAccessViolation", value)
}

// GetResourceFailureAccessViolation gets the value of ResourceFailureAccessViolation for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceFailureAccessViolation() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceFailureAccessViolation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceFailureDeadlock sets the value of ResourceFailureDeadlock for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceFailureDeadlock(value uint64) (err error) {
	return instance.SetProperty("ResourceFailureDeadlock", value)
}

// GetResourceFailureDeadlock gets the value of ResourceFailureDeadlock for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceFailureDeadlock() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceFailureDeadlock")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesOnline sets the value of ResourcesOnline for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourcesOnline(value uint64) (err error) {
	return instance.SetProperty("ResourcesOnline", value)
}

// GetResourcesOnline gets the value of ResourcesOnline for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourcesOnline() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourcesOnline")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceTypeControls sets the value of ResourceTypeControls for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceTypeControls(value uint64) (err error) {
	return instance.SetProperty("ResourceTypeControls", value)
}

// GetResourceTypeControls gets the value of ResourceTypeControls for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceTypeControls() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceTypeControls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceTypeControlsPersec sets the value of ResourceTypeControlsPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) SetPropertyResourceTypeControlsPersec(value uint64) (err error) {
	return instance.SetProperty("ResourceTypeControlsPersec", value)
}

// GetResourceTypeControlsPersec gets the value of ResourceTypeControlsPersec for the instance
func (instance *Win32_PerfRawData_ClussvcPerfProvider_ClusterResources) GetPropertyResourceTypeControlsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceTypeControlsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
