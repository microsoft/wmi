// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager struct {
	Win32_PerfFormattedData

	//
	CryptoCheckpointsRestored uint64

	//
	CryptoCheckpointsRestoredPersec uint64

	//
	CryptoCheckpointsSaved uint64

	//
	CryptoCheckpointsSavedPersec uint64

	//
	RegistryCheckpointsRestored uint64

	//
	RegistryCheckpointsRestoredPersec uint64

	//
	RegistryCheckpointsSaved uint64

	//
	RegistryCheckpointsSavedPersec uint64
}

// SetCryptoCheckpointsRestored sets the value of CryptoCheckpointsRestored for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyCryptoCheckpointsRestored(value uint64) (err error) {
	return instance.SetProperty("CryptoCheckpointsRestored", value)
}

// GetCryptoCheckpointsRestored gets the value of CryptoCheckpointsRestored for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyCryptoCheckpointsRestored() (value uint64, err error) {
	retValue, err := instance.GetProperty("CryptoCheckpointsRestored")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCryptoCheckpointsRestoredPersec sets the value of CryptoCheckpointsRestoredPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyCryptoCheckpointsRestoredPersec(value uint64) (err error) {
	return instance.SetProperty("CryptoCheckpointsRestoredPersec", value)
}

// GetCryptoCheckpointsRestoredPersec gets the value of CryptoCheckpointsRestoredPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyCryptoCheckpointsRestoredPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CryptoCheckpointsRestoredPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCryptoCheckpointsSaved sets the value of CryptoCheckpointsSaved for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyCryptoCheckpointsSaved(value uint64) (err error) {
	return instance.SetProperty("CryptoCheckpointsSaved", value)
}

// GetCryptoCheckpointsSaved gets the value of CryptoCheckpointsSaved for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyCryptoCheckpointsSaved() (value uint64, err error) {
	retValue, err := instance.GetProperty("CryptoCheckpointsSaved")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCryptoCheckpointsSavedPersec sets the value of CryptoCheckpointsSavedPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyCryptoCheckpointsSavedPersec(value uint64) (err error) {
	return instance.SetProperty("CryptoCheckpointsSavedPersec", value)
}

// GetCryptoCheckpointsSavedPersec gets the value of CryptoCheckpointsSavedPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyCryptoCheckpointsSavedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CryptoCheckpointsSavedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryCheckpointsRestored sets the value of RegistryCheckpointsRestored for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyRegistryCheckpointsRestored(value uint64) (err error) {
	return instance.SetProperty("RegistryCheckpointsRestored", value)
}

// GetRegistryCheckpointsRestored gets the value of RegistryCheckpointsRestored for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyRegistryCheckpointsRestored() (value uint64, err error) {
	retValue, err := instance.GetProperty("RegistryCheckpointsRestored")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryCheckpointsRestoredPersec sets the value of RegistryCheckpointsRestoredPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyRegistryCheckpointsRestoredPersec(value uint64) (err error) {
	return instance.SetProperty("RegistryCheckpointsRestoredPersec", value)
}

// GetRegistryCheckpointsRestoredPersec gets the value of RegistryCheckpointsRestoredPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyRegistryCheckpointsRestoredPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RegistryCheckpointsRestoredPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryCheckpointsSaved sets the value of RegistryCheckpointsSaved for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyRegistryCheckpointsSaved(value uint64) (err error) {
	return instance.SetProperty("RegistryCheckpointsSaved", value)
}

// GetRegistryCheckpointsSaved gets the value of RegistryCheckpointsSaved for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyRegistryCheckpointsSaved() (value uint64, err error) {
	retValue, err := instance.GetProperty("RegistryCheckpointsSaved")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryCheckpointsSavedPersec sets the value of RegistryCheckpointsSavedPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) SetPropertyRegistryCheckpointsSavedPersec(value uint64) (err error) {
	return instance.SetProperty("RegistryCheckpointsSavedPersec", value)
}

// GetRegistryCheckpointsSavedPersec gets the value of RegistryCheckpointsSavedPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterCheckpointManager) GetPropertyRegistryCheckpointsSavedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RegistryCheckpointsSavedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
