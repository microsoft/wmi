// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_MTProcessorSummary struct
type MSFT_MTProcessorSummary struct { 
	*CIM_ManagedElement

	// 
	AverageSpeed float32

	// 
	Cores uint32

	// 
	CurrentIndex uint16

	// 
	Handles uint32

	// 
	IntervalSeconds uint16

	// 
	L1Cache uint32

	// 
	L2Cache uint32

	// 
	L3Cache uint32

	// 
	LogicalProcessors uint32

	// 
	MaximumSpeed float32

	// 
	Name string

	// 
	NumaNodes uint16

	// 
	Privileged []float32

	// 
	Processes uint32

	// 
	Sockets uint32

	// 
	Threads uint32

	// 
	Uptime uint64

	// 
	Utilization []float32

	// 
	Virtualization uint16
}

	func NewMSFT_MTProcessorSummaryEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTProcessorSummary, err error) {tmp, err := NewCIM_ManagedElementEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_MTProcessorSummary {
	CIM_ManagedElement: tmp,
	}
	return
	}
	

	func NewMSFT_MTProcessorSummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_MTProcessorSummary, err error) {tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_MTProcessorSummary {
	CIM_ManagedElement: tmp,
	}
	return
	}
	

// SetAverageSpeed sets the value of AverageSpeed for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyAverageSpeed(value float32) (err error) { 
    return instance.SetProperty("AverageSpeed", (value))
}

// GetAverageSpeed gets the value of AverageSpeed for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyAverageSpeed()(value float32, err error) { 
    retValue, err := instance.GetProperty("AverageSpeed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(float32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = float32(valuetmp)

    return
}

// SetCores sets the value of Cores for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyCores(value uint32) (err error) { 
    return instance.SetProperty("Cores", (value))
}

// GetCores gets the value of Cores for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyCores()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Cores")
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

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyCurrentIndex(value uint16) (err error) { 
    return instance.SetProperty("CurrentIndex", (value))
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyCurrentIndex()(value uint16, err error) { 
    retValue, err := instance.GetProperty("CurrentIndex")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHandles sets the value of Handles for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyHandles(value uint32) (err error) { 
    return instance.SetProperty("Handles", (value))
}

// GetHandles gets the value of Handles for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyHandles()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Handles")
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

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyIntervalSeconds(value uint16) (err error) { 
    return instance.SetProperty("IntervalSeconds", (value))
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyIntervalSeconds()(value uint16, err error) { 
    retValue, err := instance.GetProperty("IntervalSeconds")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetL1Cache sets the value of L1Cache for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyL1Cache(value uint32) (err error) { 
    return instance.SetProperty("L1Cache", (value))
}

// GetL1Cache gets the value of L1Cache for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyL1Cache()(value uint32, err error) { 
    retValue, err := instance.GetProperty("L1Cache")
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

// SetL2Cache sets the value of L2Cache for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyL2Cache(value uint32) (err error) { 
    return instance.SetProperty("L2Cache", (value))
}

// GetL2Cache gets the value of L2Cache for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyL2Cache()(value uint32, err error) { 
    retValue, err := instance.GetProperty("L2Cache")
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

// SetL3Cache sets the value of L3Cache for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyL3Cache(value uint32) (err error) { 
    return instance.SetProperty("L3Cache", (value))
}

// GetL3Cache gets the value of L3Cache for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyL3Cache()(value uint32, err error) { 
    retValue, err := instance.GetProperty("L3Cache")
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

// SetLogicalProcessors sets the value of LogicalProcessors for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyLogicalProcessors(value uint32) (err error) { 
    return instance.SetProperty("LogicalProcessors", (value))
}

// GetLogicalProcessors gets the value of LogicalProcessors for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyLogicalProcessors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogicalProcessors")
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

// SetMaximumSpeed sets the value of MaximumSpeed for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyMaximumSpeed(value float32) (err error) { 
    return instance.SetProperty("MaximumSpeed", (value))
}

// GetMaximumSpeed gets the value of MaximumSpeed for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyMaximumSpeed()(value float32, err error) { 
    retValue, err := instance.GetProperty("MaximumSpeed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(float32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = float32(valuetmp)

    return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetNumaNodes sets the value of NumaNodes for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyNumaNodes(value uint16) (err error) { 
    return instance.SetProperty("NumaNodes", (value))
}

// GetNumaNodes gets the value of NumaNodes for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyNumaNodes()(value uint16, err error) { 
    retValue, err := instance.GetProperty("NumaNodes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetPrivileged sets the value of Privileged for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyPrivileged(value []float32) (err error) { 
    return instance.SetProperty("Privileged", (value))
}

// GetPrivileged gets the value of Privileged for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyPrivileged()(value []float32, err error) { 
    retValue, err := instance.GetProperty("Privileged")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(float32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, float32(valuetmp))
    }

    return
}

// SetProcesses sets the value of Processes for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyProcesses(value uint32) (err error) { 
    return instance.SetProperty("Processes", (value))
}

// GetProcesses gets the value of Processes for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyProcesses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Processes")
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

// SetSockets sets the value of Sockets for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertySockets(value uint32) (err error) { 
    return instance.SetProperty("Sockets", (value))
}

// GetSockets gets the value of Sockets for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertySockets()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Sockets")
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

// SetThreads sets the value of Threads for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyThreads(value uint32) (err error) { 
    return instance.SetProperty("Threads", (value))
}

// GetThreads gets the value of Threads for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyThreads()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Threads")
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

// SetUptime sets the value of Uptime for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyUptime(value uint64) (err error) { 
    return instance.SetProperty("Uptime", (value))
}

// GetUptime gets the value of Uptime for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyUptime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Uptime")
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

// SetUtilization sets the value of Utilization for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyUtilization(value []float32) (err error) { 
    return instance.SetProperty("Utilization", (value))
}

// GetUtilization gets the value of Utilization for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyUtilization()(value []float32, err error) { 
    retValue, err := instance.GetProperty("Utilization")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(float32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, float32(valuetmp))
    }

    return
}

// SetVirtualization sets the value of Virtualization for the instance
func (instance *MSFT_MTProcessorSummary) SetPropertyVirtualization(value uint16) (err error) { 
    return instance.SetProperty("Virtualization", (value))
}

// GetVirtualization gets the value of Virtualization for the instance
func (instance *MSFT_MTProcessorSummary) GetPropertyVirtualization()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Virtualization")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

