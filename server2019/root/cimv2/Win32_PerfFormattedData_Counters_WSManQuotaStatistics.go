// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_WSManQuotaStatistics struct
type Win32_PerfFormattedData_Counters_WSManQuotaStatistics struct {
	Win32_PerfFormattedData

	//
	ActiveOperations uint32

	//
	ActiveShells uint32

	//
	ActiveUsers uint32

	//
	ProcessID uint32

	//
	SystemQuotaViolationsPerSecond uint32

	//
	TotalRequestsPerSecond uint32

	//
	UserQuotaViolationsPerSecond uint32
}

// SetActiveOperations sets the value of ActiveOperations for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyActiveOperations(value uint32) (err error) {
	return instance.SetProperty("ActiveOperations", value)
}

// GetActiveOperations gets the value of ActiveOperations for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyActiveOperations() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveOperations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveShells sets the value of ActiveShells for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyActiveShells(value uint32) (err error) {
	return instance.SetProperty("ActiveShells", value)
}

// GetActiveShells gets the value of ActiveShells for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyActiveShells() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveShells")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveUsers sets the value of ActiveUsers for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyActiveUsers(value uint32) (err error) {
	return instance.SetProperty("ActiveUsers", value)
}

// GetActiveUsers gets the value of ActiveUsers for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyActiveUsers() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", value)
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemQuotaViolationsPerSecond sets the value of SystemQuotaViolationsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertySystemQuotaViolationsPerSecond(value uint32) (err error) {
	return instance.SetProperty("SystemQuotaViolationsPerSecond", value)
}

// GetSystemQuotaViolationsPerSecond gets the value of SystemQuotaViolationsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertySystemQuotaViolationsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("SystemQuotaViolationsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalRequestsPerSecond sets the value of TotalRequestsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyTotalRequestsPerSecond(value uint32) (err error) {
	return instance.SetProperty("TotalRequestsPerSecond", value)
}

// GetTotalRequestsPerSecond gets the value of TotalRequestsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyTotalRequestsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalRequestsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserQuotaViolationsPerSecond sets the value of UserQuotaViolationsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) SetPropertyUserQuotaViolationsPerSecond(value uint32) (err error) {
	return instance.SetProperty("UserQuotaViolationsPerSecond", value)
}

// GetUserQuotaViolationsPerSecond gets the value of UserQuotaViolationsPerSecond for the instance
func (instance *Win32_PerfFormattedData_Counters_WSManQuotaStatistics) GetPropertyUserQuotaViolationsPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserQuotaViolationsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
