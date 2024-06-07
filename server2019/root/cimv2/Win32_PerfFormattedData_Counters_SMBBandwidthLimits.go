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

// Win32_PerfFormattedData_Counters_SMBBandwidthLimits struct
type Win32_PerfFormattedData_Counters_SMBBandwidthLimits struct { 
	*Win32_PerfFormattedData

	// 
	AvgBytesPerRead uint64

	// 
	AvgBytesPerWrite uint64

	// 
	AvgDataBytesPerRequest uint64

	// 
	AvgDataQueueLength uint64

	// 
	AvgReadQueueLength uint64

	// 
	AvgsecPerDataRequest uint32

	// 
	AvgsecPerRead uint32

	// 
	AvgsecPerWrite uint32

	// 
	AvgWriteQueueLength uint64

	// 
	CurrentDataQueueLength uint32

	// 
	DataBytesPersec uint64

	// 
	DataRequestsPersec uint32

	// 
	ReadBytesPersec uint64

	// 
	ReadRequestsPersec uint32

	// 
	WriteBytesPersec uint64

	// 
	WriteRequestsPersec uint32
}

	func NewWin32_PerfFormattedData_Counters_SMBBandwidthLimitsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_SMBBandwidthLimits {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Counters_SMBBandwidthLimitsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_SMBBandwidthLimits {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgBytesPerRead(value uint64) (err error) { 
    return instance.SetProperty("AvgBytesPerRead", (value))
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgBytesPerRead()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerRead")
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

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgBytesPerWrite(value uint64) (err error) { 
    return instance.SetProperty("AvgBytesPerWrite", (value))
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgBytesPerWrite()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerWrite")
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

// SetAvgDataBytesPerRequest sets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgDataBytesPerRequest(value uint64) (err error) { 
    return instance.SetProperty("AvgDataBytesPerRequest", (value))
}

// GetAvgDataBytesPerRequest gets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgDataBytesPerRequest()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgDataBytesPerRequest")
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

// SetAvgDataQueueLength sets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgDataQueueLength(value uint64) (err error) { 
    return instance.SetProperty("AvgDataQueueLength", (value))
}

// GetAvgDataQueueLength gets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgDataQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgDataQueueLength")
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

// SetAvgReadQueueLength sets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgReadQueueLength(value uint64) (err error) { 
    return instance.SetProperty("AvgReadQueueLength", (value))
}

// GetAvgReadQueueLength gets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgReadQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgReadQueueLength")
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

// SetAvgsecPerDataRequest sets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgsecPerDataRequest(value uint32) (err error) { 
    return instance.SetProperty("AvgsecPerDataRequest", (value))
}

// GetAvgsecPerDataRequest gets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgsecPerDataRequest()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgsecPerDataRequest")
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

// SetAvgsecPerRead sets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgsecPerRead(value uint32) (err error) { 
    return instance.SetProperty("AvgsecPerRead", (value))
}

// GetAvgsecPerRead gets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgsecPerRead()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgsecPerRead")
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

// SetAvgsecPerWrite sets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgsecPerWrite(value uint32) (err error) { 
    return instance.SetProperty("AvgsecPerWrite", (value))
}

// GetAvgsecPerWrite gets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgsecPerWrite()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgsecPerWrite")
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

// SetAvgWriteQueueLength sets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyAvgWriteQueueLength(value uint64) (err error) { 
    return instance.SetProperty("AvgWriteQueueLength", (value))
}

// GetAvgWriteQueueLength gets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyAvgWriteQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgWriteQueueLength")
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

// SetCurrentDataQueueLength sets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyCurrentDataQueueLength(value uint32) (err error) { 
    return instance.SetProperty("CurrentDataQueueLength", (value))
}

// GetCurrentDataQueueLength gets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyCurrentDataQueueLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentDataQueueLength")
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

// SetDataBytesPersec sets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyDataBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("DataBytesPersec", (value))
}

// GetDataBytesPersec gets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyDataBytesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DataBytesPersec")
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

// SetDataRequestsPersec sets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyDataRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("DataRequestsPersec", (value))
}

// GetDataRequestsPersec gets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyDataRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DataRequestsPersec")
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

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyReadBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyReadBytesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReadBytesPersec")
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

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyReadRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("ReadRequestsPersec", (value))
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyReadRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadRequestsPersec")
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

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyWriteBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyWriteBytesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteBytesPersec")
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

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) SetPropertyWriteRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("WriteRequestsPersec", (value))
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBBandwidthLimits) GetPropertyWriteRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteRequestsPersec")
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

