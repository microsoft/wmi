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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions struct { 
	*Win32_PerfRawData

	// 
	FreeSpaceintempdbKB uint64

	// 
	LongestTransactionRunningTime uint64

	// 
	NonSnapshotVersionTransactions uint64

	// 
	SnapshotTransactions uint64

	// 
	Transactions uint64

	// 
	Updateconflictratio uint64

	// 
	Updateconflictratio_Base uint32

	// 
	UpdateSnapshotTransactions uint64

	// 
	VersionCleanuprateKBPers uint64

	// 
	VersionGenerationrateKBPers uint64

	// 
	VersionStoreSizeKB uint64

	// 
	VersionStoreunitcount uint64

	// 
	VersionStoreunitcreation uint64

	// 
	VersionStoreunittruncation uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactionsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetFreeSpaceintempdbKB sets the value of FreeSpaceintempdbKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyFreeSpaceintempdbKB(value uint64) (err error) { 
    return instance.SetProperty("FreeSpaceintempdbKB", (value))
}

// GetFreeSpaceintempdbKB gets the value of FreeSpaceintempdbKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyFreeSpaceintempdbKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FreeSpaceintempdbKB")
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

// SetLongestTransactionRunningTime sets the value of LongestTransactionRunningTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyLongestTransactionRunningTime(value uint64) (err error) { 
    return instance.SetProperty("LongestTransactionRunningTime", (value))
}

// GetLongestTransactionRunningTime gets the value of LongestTransactionRunningTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyLongestTransactionRunningTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LongestTransactionRunningTime")
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

// SetNonSnapshotVersionTransactions sets the value of NonSnapshotVersionTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyNonSnapshotVersionTransactions(value uint64) (err error) { 
    return instance.SetProperty("NonSnapshotVersionTransactions", (value))
}

// GetNonSnapshotVersionTransactions gets the value of NonSnapshotVersionTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyNonSnapshotVersionTransactions()(value uint64, err error) { 
    retValue, err := instance.GetProperty("NonSnapshotVersionTransactions")
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

// SetSnapshotTransactions sets the value of SnapshotTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertySnapshotTransactions(value uint64) (err error) { 
    return instance.SetProperty("SnapshotTransactions", (value))
}

// GetSnapshotTransactions gets the value of SnapshotTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertySnapshotTransactions()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SnapshotTransactions")
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

// SetTransactions sets the value of Transactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyTransactions(value uint64) (err error) { 
    return instance.SetProperty("Transactions", (value))
}

// GetTransactions gets the value of Transactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyTransactions()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Transactions")
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

// SetUpdateconflictratio sets the value of Updateconflictratio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyUpdateconflictratio(value uint64) (err error) { 
    return instance.SetProperty("Updateconflictratio", (value))
}

// GetUpdateconflictratio gets the value of Updateconflictratio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyUpdateconflictratio()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Updateconflictratio")
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

// SetUpdateconflictratio_Base sets the value of Updateconflictratio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyUpdateconflictratio_Base(value uint32) (err error) { 
    return instance.SetProperty("Updateconflictratio_Base", (value))
}

// GetUpdateconflictratio_Base gets the value of Updateconflictratio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyUpdateconflictratio_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Updateconflictratio_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetUpdateSnapshotTransactions sets the value of UpdateSnapshotTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyUpdateSnapshotTransactions(value uint64) (err error) { 
    return instance.SetProperty("UpdateSnapshotTransactions", (value))
}

// GetUpdateSnapshotTransactions gets the value of UpdateSnapshotTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyUpdateSnapshotTransactions()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UpdateSnapshotTransactions")
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

// SetVersionCleanuprateKBPers sets the value of VersionCleanuprateKBPers for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionCleanuprateKBPers(value uint64) (err error) { 
    return instance.SetProperty("VersionCleanuprateKBPers", (value))
}

// GetVersionCleanuprateKBPers gets the value of VersionCleanuprateKBPers for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionCleanuprateKBPers()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionCleanuprateKBPers")
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

// SetVersionGenerationrateKBPers sets the value of VersionGenerationrateKBPers for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionGenerationrateKBPers(value uint64) (err error) { 
    return instance.SetProperty("VersionGenerationrateKBPers", (value))
}

// GetVersionGenerationrateKBPers gets the value of VersionGenerationrateKBPers for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionGenerationrateKBPers()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionGenerationrateKBPers")
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

// SetVersionStoreSizeKB sets the value of VersionStoreSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionStoreSizeKB(value uint64) (err error) { 
    return instance.SetProperty("VersionStoreSizeKB", (value))
}

// GetVersionStoreSizeKB gets the value of VersionStoreSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionStoreSizeKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionStoreSizeKB")
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

// SetVersionStoreunitcount sets the value of VersionStoreunitcount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionStoreunitcount(value uint64) (err error) { 
    return instance.SetProperty("VersionStoreunitcount", (value))
}

// GetVersionStoreunitcount gets the value of VersionStoreunitcount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionStoreunitcount()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionStoreunitcount")
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

// SetVersionStoreunitcreation sets the value of VersionStoreunitcreation for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionStoreunitcreation(value uint64) (err error) { 
    return instance.SetProperty("VersionStoreunitcreation", (value))
}

// GetVersionStoreunitcreation gets the value of VersionStoreunitcreation for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionStoreunitcreation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionStoreunitcreation")
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

// SetVersionStoreunittruncation sets the value of VersionStoreunittruncation for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) SetPropertyVersionStoreunittruncation(value uint64) (err error) { 
    return instance.SetProperty("VersionStoreunittruncation", (value))
}

// GetVersionStoreunittruncation gets the value of VersionStoreunittruncation for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDTransactions) GetPropertyVersionStoreunittruncation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VersionStoreunittruncation")
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

