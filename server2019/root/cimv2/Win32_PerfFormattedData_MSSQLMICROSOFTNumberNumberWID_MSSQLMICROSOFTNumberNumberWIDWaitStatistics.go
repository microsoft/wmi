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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics struct { 
	*Win32_PerfFormattedData

	// 
	Lockwaits uint64

	// 
	Logbufferwaits uint64

	// 
	Logwritewaits uint64

	// 
	Memorygrantqueuewaits uint64

	// 
	NetworkIOwaits uint64

	// 
	NonPagelatchwaits uint64

	// 
	PageIOlatchwaits uint64

	// 
	Pagelatchwaits uint64

	// 
	Threadsafememoryobjectswaits uint64

	// 
	Transactionownershipwaits uint64

	// 
	Waitfortheworker uint64

	// 
	Workspacesynchronizationwaits uint64
}

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetLockwaits sets the value of Lockwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyLockwaits(value uint64) (err error) { 
    return instance.SetProperty("Lockwaits", (value))
}

// GetLockwaits gets the value of Lockwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyLockwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Lockwaits")
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

// SetLogbufferwaits sets the value of Logbufferwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyLogbufferwaits(value uint64) (err error) { 
    return instance.SetProperty("Logbufferwaits", (value))
}

// GetLogbufferwaits gets the value of Logbufferwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyLogbufferwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Logbufferwaits")
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

// SetLogwritewaits sets the value of Logwritewaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyLogwritewaits(value uint64) (err error) { 
    return instance.SetProperty("Logwritewaits", (value))
}

// GetLogwritewaits gets the value of Logwritewaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyLogwritewaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Logwritewaits")
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

// SetMemorygrantqueuewaits sets the value of Memorygrantqueuewaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyMemorygrantqueuewaits(value uint64) (err error) { 
    return instance.SetProperty("Memorygrantqueuewaits", (value))
}

// GetMemorygrantqueuewaits gets the value of Memorygrantqueuewaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyMemorygrantqueuewaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Memorygrantqueuewaits")
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

// SetNetworkIOwaits sets the value of NetworkIOwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyNetworkIOwaits(value uint64) (err error) { 
    return instance.SetProperty("NetworkIOwaits", (value))
}

// GetNetworkIOwaits gets the value of NetworkIOwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyNetworkIOwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("NetworkIOwaits")
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

// SetNonPagelatchwaits sets the value of NonPagelatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyNonPagelatchwaits(value uint64) (err error) { 
    return instance.SetProperty("NonPagelatchwaits", (value))
}

// GetNonPagelatchwaits gets the value of NonPagelatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyNonPagelatchwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("NonPagelatchwaits")
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

// SetPageIOlatchwaits sets the value of PageIOlatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyPageIOlatchwaits(value uint64) (err error) { 
    return instance.SetProperty("PageIOlatchwaits", (value))
}

// GetPageIOlatchwaits gets the value of PageIOlatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyPageIOlatchwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PageIOlatchwaits")
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

// SetPagelatchwaits sets the value of Pagelatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyPagelatchwaits(value uint64) (err error) { 
    return instance.SetProperty("Pagelatchwaits", (value))
}

// GetPagelatchwaits gets the value of Pagelatchwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyPagelatchwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Pagelatchwaits")
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

// SetThreadsafememoryobjectswaits sets the value of Threadsafememoryobjectswaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyThreadsafememoryobjectswaits(value uint64) (err error) { 
    return instance.SetProperty("Threadsafememoryobjectswaits", (value))
}

// GetThreadsafememoryobjectswaits gets the value of Threadsafememoryobjectswaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyThreadsafememoryobjectswaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Threadsafememoryobjectswaits")
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

// SetTransactionownershipwaits sets the value of Transactionownershipwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyTransactionownershipwaits(value uint64) (err error) { 
    return instance.SetProperty("Transactionownershipwaits", (value))
}

// GetTransactionownershipwaits gets the value of Transactionownershipwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyTransactionownershipwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Transactionownershipwaits")
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

// SetWaitfortheworker sets the value of Waitfortheworker for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyWaitfortheworker(value uint64) (err error) { 
    return instance.SetProperty("Waitfortheworker", (value))
}

// GetWaitfortheworker gets the value of Waitfortheworker for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyWaitfortheworker()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Waitfortheworker")
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

// SetWorkspacesynchronizationwaits sets the value of Workspacesynchronizationwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) SetPropertyWorkspacesynchronizationwaits(value uint64) (err error) { 
    return instance.SetProperty("Workspacesynchronizationwaits", (value))
}

// GetWorkspacesynchronizationwaits gets the value of Workspacesynchronizationwaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWaitStatistics) GetPropertyWorkspacesynchronizationwaits()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Workspacesynchronizationwaits")
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

