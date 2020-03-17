// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ESENT_DatabaseTableClasses struct
type Win32_PerfFormattedData_ESENT_DatabaseTableClasses struct {
	Win32_PerfFormattedData

	//
	DatabaseCacheMissAttachedAverageLatency uint64

	//
	DatabaseCacheMissesPersec uint32

	//
	DatabaseCachePercentHit uint32

	//
	DatabaseCachePercentHitUnique uint32

	//
	DatabaseCacheRequestsPersec uint32

	//
	DatabaseCacheRequestsPersecUnique uint32

	//
	DatabaseCacheSize uint64

	//
	DatabaseCacheSizeMB uint64
}

// SetDatabaseCacheMissAttachedAverageLatency sets the value of DatabaseCacheMissAttachedAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheMissAttachedAverageLatency(value uint64) (err error) {
	return instance.SetProperty("DatabaseCacheMissAttachedAverageLatency", value)
}

// GetDatabaseCacheMissAttachedAverageLatency gets the value of DatabaseCacheMissAttachedAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheMissAttachedAverageLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheMissAttachedAverageLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheMissesPersec sets the value of DatabaseCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheMissesPersec(value uint32) (err error) {
	return instance.SetProperty("DatabaseCacheMissesPersec", value)
}

// GetDatabaseCacheMissesPersec gets the value of DatabaseCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheMissesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheMissesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCachePercentHit sets the value of DatabaseCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCachePercentHit(value uint32) (err error) {
	return instance.SetProperty("DatabaseCachePercentHit", value)
}

// GetDatabaseCachePercentHit gets the value of DatabaseCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCachePercentHit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCachePercentHit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCachePercentHitUnique sets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCachePercentHitUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCachePercentHitUnique", value)
}

// GetDatabaseCachePercentHitUnique gets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCachePercentHitUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCachePercentHitUnique")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheRequestsPersec sets the value of DatabaseCacheRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("DatabaseCacheRequestsPersec", value)
}

// GetDatabaseCacheRequestsPersec gets the value of DatabaseCacheRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheRequestsPersecUnique sets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheRequestsPersecUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCacheRequestsPersecUnique", value)
}

// GetDatabaseCacheRequestsPersecUnique gets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheRequestsPersecUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheRequestsPersecUnique")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheSize sets the value of DatabaseCacheSize for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheSize(value uint64) (err error) {
	return instance.SetProperty("DatabaseCacheSize", value)
}

// GetDatabaseCacheSize gets the value of DatabaseCacheSize for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheSizeMB sets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) SetPropertyDatabaseCacheSizeMB(value uint64) (err error) {
	return instance.SetProperty("DatabaseCacheSizeMB", value)
}

// GetDatabaseCacheSizeMB gets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseTableClasses) GetPropertyDatabaseCacheSizeMB() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheSizeMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
