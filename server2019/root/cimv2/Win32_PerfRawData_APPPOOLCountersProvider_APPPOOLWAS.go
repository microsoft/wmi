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

// Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS struct
type Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS struct { 
	*Win32_PerfRawData

	// 
	CurrentApplicationPoolState uint32

	// 
	CurrentApplicationPoolUptime uint64

	// 
	CurrentWorkerProcesses uint32

	// 
	MaximumWorkerProcesses uint32

	// 
	RecentWorkerProcessFailures uint32

	// 
	TimeSinceLastWorkerProcessFailure uint64

	// 
	TotalApplicationPoolRecycles uint32

	// 
	TotalApplicationPoolUptime uint64

	// 
	TotalWorkerProcessesCreated uint32

	// 
	TotalWorkerProcessFailures uint32

	// 
	TotalWorkerProcessPingFailures uint32

	// 
	TotalWorkerProcessShutdownFailures uint32

	// 
	TotalWorkerProcessStartupFailures uint32
}

	func NewWin32_PerfRawData_APPPOOLCountersProvider_APPPOOLWASEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_APPPOOLCountersProvider_APPPOOLWASEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetCurrentApplicationPoolState sets the value of CurrentApplicationPoolState for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyCurrentApplicationPoolState(value uint32) (err error) { 
    return instance.SetProperty("CurrentApplicationPoolState", (value))
}

// GetCurrentApplicationPoolState gets the value of CurrentApplicationPoolState for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyCurrentApplicationPoolState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentApplicationPoolState")
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

// SetCurrentApplicationPoolUptime sets the value of CurrentApplicationPoolUptime for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyCurrentApplicationPoolUptime(value uint64) (err error) { 
    return instance.SetProperty("CurrentApplicationPoolUptime", (value))
}

// GetCurrentApplicationPoolUptime gets the value of CurrentApplicationPoolUptime for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyCurrentApplicationPoolUptime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CurrentApplicationPoolUptime")
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

// SetCurrentWorkerProcesses sets the value of CurrentWorkerProcesses for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyCurrentWorkerProcesses(value uint32) (err error) { 
    return instance.SetProperty("CurrentWorkerProcesses", (value))
}

// GetCurrentWorkerProcesses gets the value of CurrentWorkerProcesses for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyCurrentWorkerProcesses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentWorkerProcesses")
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

// SetMaximumWorkerProcesses sets the value of MaximumWorkerProcesses for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyMaximumWorkerProcesses(value uint32) (err error) { 
    return instance.SetProperty("MaximumWorkerProcesses", (value))
}

// GetMaximumWorkerProcesses gets the value of MaximumWorkerProcesses for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyMaximumWorkerProcesses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumWorkerProcesses")
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

// SetRecentWorkerProcessFailures sets the value of RecentWorkerProcessFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyRecentWorkerProcessFailures(value uint32) (err error) { 
    return instance.SetProperty("RecentWorkerProcessFailures", (value))
}

// GetRecentWorkerProcessFailures gets the value of RecentWorkerProcessFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyRecentWorkerProcessFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RecentWorkerProcessFailures")
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

// SetTimeSinceLastWorkerProcessFailure sets the value of TimeSinceLastWorkerProcessFailure for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTimeSinceLastWorkerProcessFailure(value uint64) (err error) { 
    return instance.SetProperty("TimeSinceLastWorkerProcessFailure", (value))
}

// GetTimeSinceLastWorkerProcessFailure gets the value of TimeSinceLastWorkerProcessFailure for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTimeSinceLastWorkerProcessFailure()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TimeSinceLastWorkerProcessFailure")
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

// SetTotalApplicationPoolRecycles sets the value of TotalApplicationPoolRecycles for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalApplicationPoolRecycles(value uint32) (err error) { 
    return instance.SetProperty("TotalApplicationPoolRecycles", (value))
}

// GetTotalApplicationPoolRecycles gets the value of TotalApplicationPoolRecycles for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalApplicationPoolRecycles()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalApplicationPoolRecycles")
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

// SetTotalApplicationPoolUptime sets the value of TotalApplicationPoolUptime for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalApplicationPoolUptime(value uint64) (err error) { 
    return instance.SetProperty("TotalApplicationPoolUptime", (value))
}

// GetTotalApplicationPoolUptime gets the value of TotalApplicationPoolUptime for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalApplicationPoolUptime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalApplicationPoolUptime")
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

// SetTotalWorkerProcessesCreated sets the value of TotalWorkerProcessesCreated for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalWorkerProcessesCreated(value uint32) (err error) { 
    return instance.SetProperty("TotalWorkerProcessesCreated", (value))
}

// GetTotalWorkerProcessesCreated gets the value of TotalWorkerProcessesCreated for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalWorkerProcessesCreated()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalWorkerProcessesCreated")
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

// SetTotalWorkerProcessFailures sets the value of TotalWorkerProcessFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalWorkerProcessFailures(value uint32) (err error) { 
    return instance.SetProperty("TotalWorkerProcessFailures", (value))
}

// GetTotalWorkerProcessFailures gets the value of TotalWorkerProcessFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalWorkerProcessFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalWorkerProcessFailures")
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

// SetTotalWorkerProcessPingFailures sets the value of TotalWorkerProcessPingFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalWorkerProcessPingFailures(value uint32) (err error) { 
    return instance.SetProperty("TotalWorkerProcessPingFailures", (value))
}

// GetTotalWorkerProcessPingFailures gets the value of TotalWorkerProcessPingFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalWorkerProcessPingFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalWorkerProcessPingFailures")
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

// SetTotalWorkerProcessShutdownFailures sets the value of TotalWorkerProcessShutdownFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalWorkerProcessShutdownFailures(value uint32) (err error) { 
    return instance.SetProperty("TotalWorkerProcessShutdownFailures", (value))
}

// GetTotalWorkerProcessShutdownFailures gets the value of TotalWorkerProcessShutdownFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalWorkerProcessShutdownFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalWorkerProcessShutdownFailures")
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

// SetTotalWorkerProcessStartupFailures sets the value of TotalWorkerProcessStartupFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) SetPropertyTotalWorkerProcessStartupFailures(value uint32) (err error) { 
    return instance.SetProperty("TotalWorkerProcessStartupFailures", (value))
}

// GetTotalWorkerProcessStartupFailures gets the value of TotalWorkerProcessStartupFailures for the instance
func (instance *Win32_PerfRawData_APPPOOLCountersProvider_APPPOOLWAS) GetPropertyTotalWorkerProcessStartupFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalWorkerProcessStartupFailures")
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

