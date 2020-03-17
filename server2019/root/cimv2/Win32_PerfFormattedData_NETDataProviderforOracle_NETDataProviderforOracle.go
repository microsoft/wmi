// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle struct
type Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle struct {
	Win32_PerfFormattedData

	//
	HardConnectsPerSecond uint32

	//
	HardDisconnectsPerSecond uint32

	//
	NumberOfActiveConnectionPoolGroups uint32

	//
	NumberOfActiveConnectionPools uint32

	//
	NumberOfActiveConnections uint32

	//
	NumberOfFreeConnections uint32

	//
	NumberOfInactiveConnectionPoolGroups uint32

	//
	NumberOfInactiveConnectionPools uint32

	//
	NumberOfNonPooledConnections uint32

	//
	NumberOfPooledConnections uint32

	//
	NumberOfReclaimedConnections uint32

	//
	NumberOfStasisConnections uint32

	//
	SoftConnectsPerSecond uint32

	//
	SoftDisconnectsPerSecond uint32
}

// SetHardConnectsPerSecond sets the value of HardConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyHardConnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("HardConnectsPerSecond", value)
}

// GetHardConnectsPerSecond gets the value of HardConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyHardConnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardConnectsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardDisconnectsPerSecond sets the value of HardDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyHardDisconnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("HardDisconnectsPerSecond", value)
}

// GetHardDisconnectsPerSecond gets the value of HardDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyHardDisconnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardDisconnectsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfActiveConnectionPoolGroups sets the value of NumberOfActiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnectionPoolGroups(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnectionPoolGroups", value)
}

// GetNumberOfActiveConnectionPoolGroups gets the value of NumberOfActiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnectionPoolGroups() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnectionPoolGroups")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfActiveConnectionPools sets the value of NumberOfActiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnectionPools(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnectionPools", value)
}

// GetNumberOfActiveConnectionPools gets the value of NumberOfActiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnectionPools() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnectionPools")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfActiveConnections sets the value of NumberOfActiveConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfActiveConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfActiveConnections", value)
}

// GetNumberOfActiveConnections gets the value of NumberOfActiveConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfActiveConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfActiveConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfFreeConnections sets the value of NumberOfFreeConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfFreeConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfFreeConnections", value)
}

// GetNumberOfFreeConnections gets the value of NumberOfFreeConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfFreeConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfFreeConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfInactiveConnectionPoolGroups sets the value of NumberOfInactiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfInactiveConnectionPoolGroups(value uint32) (err error) {
	return instance.SetProperty("NumberOfInactiveConnectionPoolGroups", value)
}

// GetNumberOfInactiveConnectionPoolGroups gets the value of NumberOfInactiveConnectionPoolGroups for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfInactiveConnectionPoolGroups() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInactiveConnectionPoolGroups")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfInactiveConnectionPools sets the value of NumberOfInactiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfInactiveConnectionPools(value uint32) (err error) {
	return instance.SetProperty("NumberOfInactiveConnectionPools", value)
}

// GetNumberOfInactiveConnectionPools gets the value of NumberOfInactiveConnectionPools for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfInactiveConnectionPools() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInactiveConnectionPools")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfNonPooledConnections sets the value of NumberOfNonPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfNonPooledConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfNonPooledConnections", value)
}

// GetNumberOfNonPooledConnections gets the value of NumberOfNonPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfNonPooledConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfNonPooledConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfPooledConnections sets the value of NumberOfPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfPooledConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfPooledConnections", value)
}

// GetNumberOfPooledConnections gets the value of NumberOfPooledConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfPooledConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPooledConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfReclaimedConnections sets the value of NumberOfReclaimedConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfReclaimedConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfReclaimedConnections", value)
}

// GetNumberOfReclaimedConnections gets the value of NumberOfReclaimedConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfReclaimedConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfReclaimedConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfStasisConnections sets the value of NumberOfStasisConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertyNumberOfStasisConnections(value uint32) (err error) {
	return instance.SetProperty("NumberOfStasisConnections", value)
}

// GetNumberOfStasisConnections gets the value of NumberOfStasisConnections for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertyNumberOfStasisConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfStasisConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftConnectsPerSecond sets the value of SoftConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertySoftConnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("SoftConnectsPerSecond", value)
}

// GetSoftConnectsPerSecond gets the value of SoftConnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertySoftConnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("SoftConnectsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftDisconnectsPerSecond sets the value of SoftDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) SetPropertySoftDisconnectsPerSecond(value uint32) (err error) {
	return instance.SetProperty("SoftDisconnectsPerSecond", value)
}

// GetSoftDisconnectsPerSecond gets the value of SoftDisconnectsPerSecond for the instance
func (instance *Win32_PerfFormattedData_NETDataProviderforOracle_NETDataProviderforOracle) GetPropertySoftDisconnectsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("SoftDisconnectsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
