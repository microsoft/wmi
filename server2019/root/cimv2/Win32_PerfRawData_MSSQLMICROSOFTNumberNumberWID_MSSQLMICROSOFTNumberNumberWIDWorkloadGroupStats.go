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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats struct { 
	*Win32_PerfRawData

	// 
	Activeparallelthreads uint64

	// 
	Activerequests uint64

	// 
	Blockedtasks uint64

	// 
	CPUusagePercent uint64

	// 
	CPUusagePercent_Base uint64

	// 
	Maxrequestcputimems uint64

	// 
	MaxrequestmemorygrantKB uint64

	// 
	QueryoptimizationsPersec uint64

	// 
	Queuedrequests uint64

	// 
	ReducedmemorygrantsPersec uint64

	// 
	RequestscompletedPersec uint64

	// 
	SuboptimalplansPersec uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStatsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStatsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetActiveparallelthreads sets the value of Activeparallelthreads for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyActiveparallelthreads(value uint64) (err error) { 
    return instance.SetProperty("Activeparallelthreads", (value))
}

// GetActiveparallelthreads gets the value of Activeparallelthreads for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyActiveparallelthreads()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Activeparallelthreads")
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

// SetActiverequests sets the value of Activerequests for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyActiverequests(value uint64) (err error) { 
    return instance.SetProperty("Activerequests", (value))
}

// GetActiverequests gets the value of Activerequests for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyActiverequests()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Activerequests")
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

// SetBlockedtasks sets the value of Blockedtasks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyBlockedtasks(value uint64) (err error) { 
    return instance.SetProperty("Blockedtasks", (value))
}

// GetBlockedtasks gets the value of Blockedtasks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyBlockedtasks()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Blockedtasks")
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

// SetCPUusagePercent sets the value of CPUusagePercent for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyCPUusagePercent(value uint64) (err error) { 
    return instance.SetProperty("CPUusagePercent", (value))
}

// GetCPUusagePercent gets the value of CPUusagePercent for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyCPUusagePercent()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CPUusagePercent")
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

// SetCPUusagePercent_Base sets the value of CPUusagePercent_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyCPUusagePercent_Base(value uint64) (err error) { 
    return instance.SetProperty("CPUusagePercent_Base", (value))
}

// GetCPUusagePercent_Base gets the value of CPUusagePercent_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyCPUusagePercent_Base()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CPUusagePercent_Base")
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

// SetMaxrequestcputimems sets the value of Maxrequestcputimems for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyMaxrequestcputimems(value uint64) (err error) { 
    return instance.SetProperty("Maxrequestcputimems", (value))
}

// GetMaxrequestcputimems gets the value of Maxrequestcputimems for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyMaxrequestcputimems()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Maxrequestcputimems")
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

// SetMaxrequestmemorygrantKB sets the value of MaxrequestmemorygrantKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyMaxrequestmemorygrantKB(value uint64) (err error) { 
    return instance.SetProperty("MaxrequestmemorygrantKB", (value))
}

// GetMaxrequestmemorygrantKB gets the value of MaxrequestmemorygrantKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyMaxrequestmemorygrantKB()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaxrequestmemorygrantKB")
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

// SetQueryoptimizationsPersec sets the value of QueryoptimizationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyQueryoptimizationsPersec(value uint64) (err error) { 
    return instance.SetProperty("QueryoptimizationsPersec", (value))
}

// GetQueryoptimizationsPersec gets the value of QueryoptimizationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyQueryoptimizationsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QueryoptimizationsPersec")
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

// SetQueuedrequests sets the value of Queuedrequests for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyQueuedrequests(value uint64) (err error) { 
    return instance.SetProperty("Queuedrequests", (value))
}

// GetQueuedrequests gets the value of Queuedrequests for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyQueuedrequests()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Queuedrequests")
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

// SetReducedmemorygrantsPersec sets the value of ReducedmemorygrantsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyReducedmemorygrantsPersec(value uint64) (err error) { 
    return instance.SetProperty("ReducedmemorygrantsPersec", (value))
}

// GetReducedmemorygrantsPersec gets the value of ReducedmemorygrantsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyReducedmemorygrantsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReducedmemorygrantsPersec")
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

// SetRequestscompletedPersec sets the value of RequestscompletedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertyRequestscompletedPersec(value uint64) (err error) { 
    return instance.SetProperty("RequestscompletedPersec", (value))
}

// GetRequestscompletedPersec gets the value of RequestscompletedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertyRequestscompletedPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("RequestscompletedPersec")
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

// SetSuboptimalplansPersec sets the value of SuboptimalplansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) SetPropertySuboptimalplansPersec(value uint64) (err error) { 
    return instance.SetProperty("SuboptimalplansPersec", (value))
}

// GetSuboptimalplansPersec gets the value of SuboptimalplansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDWorkloadGroupStats) GetPropertySuboptimalplansPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuboptimalplansPersec")
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

