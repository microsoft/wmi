// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager struct { 
	*Win32_PerfFormattedData

	// 
	ConnectionMemoryKB uint64

	// 
	DatabaseCacheMemoryKB uint64

	// 
	Externalbenefitofmemory uint64

	// 
	FreeMemoryKB uint64

	// 
	GrantedWorkspaceMemoryKB uint64

	// 
	LockBlocks uint64

	// 
	LockBlocksAllocated uint64

	// 
	LockMemoryKB uint64

	// 
	LockOwnerBlocks uint64

	// 
	LockOwnerBlocksAllocated uint64

	// 
	LogPoolMemoryKB uint64

	// 
	MaximumWorkspaceMemoryKB uint64

	// 
	MemoryGrantsOutstanding uint64

	// 
	MemoryGrantsPending uint64

	// 
	OptimizerMemoryKB uint64

	// 
	ReservedServerMemoryKB uint64

	// 
	SQLCacheMemoryKB uint64

	// 
	StolenServerMemoryKB uint64

	// 
	TargetServerMemoryKB uint64

	// 
	TotalServerMemoryKB uint64
}

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManagerEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetConnectionMemoryKB sets the value of ConnectionMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyConnectionMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("ConnectionMemoryKB", (value))
}

// GetConnectionMemoryKB gets the value of ConnectionMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyConnectionMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ConnectionMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetDatabaseCacheMemoryKB sets the value of DatabaseCacheMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyDatabaseCacheMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("DatabaseCacheMemoryKB", (value))
}

// GetDatabaseCacheMemoryKB gets the value of DatabaseCacheMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyDatabaseCacheMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DatabaseCacheMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetExternalbenefitofmemory sets the value of Externalbenefitofmemory for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyExternalbenefitofmemory(value uint64) (err error) { 
    return instance.SetProperty("Externalbenefitofmemory", (value))
}

// GetExternalbenefitofmemory gets the value of Externalbenefitofmemory for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyExternalbenefitofmemory()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Externalbenefitofmemory")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetFreeMemoryKB sets the value of FreeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyFreeMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("FreeMemoryKB", (value))
}

// GetFreeMemoryKB gets the value of FreeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyFreeMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FreeMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetGrantedWorkspaceMemoryKB sets the value of GrantedWorkspaceMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyGrantedWorkspaceMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("GrantedWorkspaceMemoryKB", (value))
}

// GetGrantedWorkspaceMemoryKB gets the value of GrantedWorkspaceMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyGrantedWorkspaceMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("GrantedWorkspaceMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLockBlocks sets the value of LockBlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLockBlocks(value uint64) (err error) { 
    return instance.SetProperty("LockBlocks", (value))
}

// GetLockBlocks gets the value of LockBlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLockBlocks()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LockBlocks")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLockBlocksAllocated sets the value of LockBlocksAllocated for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLockBlocksAllocated(value uint64) (err error) { 
    return instance.SetProperty("LockBlocksAllocated", (value))
}

// GetLockBlocksAllocated gets the value of LockBlocksAllocated for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLockBlocksAllocated()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LockBlocksAllocated")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLockMemoryKB sets the value of LockMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLockMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("LockMemoryKB", (value))
}

// GetLockMemoryKB gets the value of LockMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLockMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LockMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLockOwnerBlocks sets the value of LockOwnerBlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLockOwnerBlocks(value uint64) (err error) { 
    return instance.SetProperty("LockOwnerBlocks", (value))
}

// GetLockOwnerBlocks gets the value of LockOwnerBlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLockOwnerBlocks()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LockOwnerBlocks")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLockOwnerBlocksAllocated sets the value of LockOwnerBlocksAllocated for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLockOwnerBlocksAllocated(value uint64) (err error) { 
    return instance.SetProperty("LockOwnerBlocksAllocated", (value))
}

// GetLockOwnerBlocksAllocated gets the value of LockOwnerBlocksAllocated for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLockOwnerBlocksAllocated()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LockOwnerBlocksAllocated")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetLogPoolMemoryKB sets the value of LogPoolMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyLogPoolMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("LogPoolMemoryKB", (value))
}

// GetLogPoolMemoryKB gets the value of LogPoolMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyLogPoolMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LogPoolMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetMaximumWorkspaceMemoryKB sets the value of MaximumWorkspaceMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyMaximumWorkspaceMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("MaximumWorkspaceMemoryKB", (value))
}

// GetMaximumWorkspaceMemoryKB gets the value of MaximumWorkspaceMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyMaximumWorkspaceMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaximumWorkspaceMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetMemoryGrantsOutstanding sets the value of MemoryGrantsOutstanding for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyMemoryGrantsOutstanding(value uint64) (err error) { 
    return instance.SetProperty("MemoryGrantsOutstanding", (value))
}

// GetMemoryGrantsOutstanding gets the value of MemoryGrantsOutstanding for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyMemoryGrantsOutstanding()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MemoryGrantsOutstanding")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetMemoryGrantsPending sets the value of MemoryGrantsPending for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyMemoryGrantsPending(value uint64) (err error) { 
    return instance.SetProperty("MemoryGrantsPending", (value))
}

// GetMemoryGrantsPending gets the value of MemoryGrantsPending for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyMemoryGrantsPending()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MemoryGrantsPending")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetOptimizerMemoryKB sets the value of OptimizerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyOptimizerMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("OptimizerMemoryKB", (value))
}

// GetOptimizerMemoryKB gets the value of OptimizerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyOptimizerMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("OptimizerMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetReservedServerMemoryKB sets the value of ReservedServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyReservedServerMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("ReservedServerMemoryKB", (value))
}

// GetReservedServerMemoryKB gets the value of ReservedServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyReservedServerMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReservedServerMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetSQLCacheMemoryKB sets the value of SQLCacheMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertySQLCacheMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("SQLCacheMemoryKB", (value))
}

// GetSQLCacheMemoryKB gets the value of SQLCacheMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertySQLCacheMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SQLCacheMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetStolenServerMemoryKB sets the value of StolenServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyStolenServerMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("StolenServerMemoryKB", (value))
}

// GetStolenServerMemoryKB gets the value of StolenServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyStolenServerMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("StolenServerMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetTargetServerMemoryKB sets the value of TargetServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyTargetServerMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("TargetServerMemoryKB", (value))
}

// GetTargetServerMemoryKB gets the value of TargetServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyTargetServerMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TargetServerMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetTotalServerMemoryKB sets the value of TotalServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) SetPropertyTotalServerMemoryKB(value uint64) (err error) { 
    return instance.SetProperty("TotalServerMemoryKB", (value))
}

// GetTotalServerMemoryKB gets the value of TotalServerMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryManager) GetPropertyTotalServerMemoryKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalServerMemoryKB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

