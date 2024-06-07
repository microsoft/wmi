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

// Win32_PerfFormattedData_Counters_StorportUnitWrites struct
type Win32_PerfFormattedData_Counters_StorportUnitWrites struct { 
	*Win32_PerfFormattedData

	// 
	SuccessfulWritesPersecBucket014K uint64

	// 
	SuccessfulWritesPersecBucket028K uint64

	// 
	SuccessfulWritesPersecBucket0316K uint64

	// 
	SuccessfulWritesPersecBucket0432K uint64

	// 
	SuccessfulWritesPersecBucket0564K uint64

	// 
	SuccessfulWritesPersecBucket06128K uint64

	// 
	SuccessfulWritesPersecBucket07256K uint64

	// 
	SuccessfulWritesPersecBucket081M uint64

	// 
	SuccessfulWritesPersecBucket091M uint64

	// 
	WriteBytesAverage uint64

	// 
	WriteBytesPersec uint64

	// 
	WriteLatency uint64

	// 
	WriteLatencyBucket01128us uint64

	// 
	WriteLatencyBucket02256us uint64

	// 
	WriteLatencyBucket03512us uint64

	// 
	WriteLatencyBucket041ms uint64

	// 
	WriteLatencyBucket054ms uint64

	// 
	WriteLatencyBucket0616ms uint64

	// 
	WriteLatencyBucket0764ms uint64

	// 
	WriteLatencyBucket08128ms uint64

	// 
	WriteLatencyBucket09256ms uint64

	// 
	WriteLatencyBucket10512ms uint64

	// 
	WriteLatencyBucket111s uint64

	// 
	WriteLatencyBucket122s uint64

	// 
	WriteLatencyBucket1310s uint64

	// 
	WriteLatencyBucket1410s uint64

	// 
	WritesPersec uint64

	// 
	WritesPersecBucket01128us uint64

	// 
	WritesPersecBucket02256us uint64

	// 
	WritesPersecBucket03512us uint64

	// 
	WritesPersecBucket041ms uint64

	// 
	WritesPersecBucket054ms uint64

	// 
	WritesPersecBucket0616ms uint64

	// 
	WritesPersecBucket0764ms uint64

	// 
	WritesPersecBucket08128ms uint64

	// 
	WritesPersecBucket09256ms uint64

	// 
	WritesPersecBucket10512ms uint64

	// 
	WritesPersecBucket111s uint64

	// 
	WritesPersecBucket122s uint64

	// 
	WritesPersecBucket1310s uint64

	// 
	WritesPersecBucket1410s uint64
}

	func NewWin32_PerfFormattedData_Counters_StorportUnitWritesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitWrites, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitWrites {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Counters_StorportUnitWritesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitWrites, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitWrites {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetSuccessfulWritesPersecBucket014K sets the value of SuccessfulWritesPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket014K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket014K", (value))
}

// GetSuccessfulWritesPersecBucket014K gets the value of SuccessfulWritesPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket014K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket014K")
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

// SetSuccessfulWritesPersecBucket028K sets the value of SuccessfulWritesPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket028K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket028K", (value))
}

// GetSuccessfulWritesPersecBucket028K gets the value of SuccessfulWritesPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket028K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket028K")
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

// SetSuccessfulWritesPersecBucket0316K sets the value of SuccessfulWritesPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket0316K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket0316K", (value))
}

// GetSuccessfulWritesPersecBucket0316K gets the value of SuccessfulWritesPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket0316K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket0316K")
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

// SetSuccessfulWritesPersecBucket0432K sets the value of SuccessfulWritesPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket0432K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket0432K", (value))
}

// GetSuccessfulWritesPersecBucket0432K gets the value of SuccessfulWritesPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket0432K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket0432K")
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

// SetSuccessfulWritesPersecBucket0564K sets the value of SuccessfulWritesPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket0564K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket0564K", (value))
}

// GetSuccessfulWritesPersecBucket0564K gets the value of SuccessfulWritesPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket0564K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket0564K")
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

// SetSuccessfulWritesPersecBucket06128K sets the value of SuccessfulWritesPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket06128K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket06128K", (value))
}

// GetSuccessfulWritesPersecBucket06128K gets the value of SuccessfulWritesPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket06128K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket06128K")
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

// SetSuccessfulWritesPersecBucket07256K sets the value of SuccessfulWritesPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket07256K(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket07256K", (value))
}

// GetSuccessfulWritesPersecBucket07256K gets the value of SuccessfulWritesPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket07256K()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket07256K")
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

// SetSuccessfulWritesPersecBucket081M sets the value of SuccessfulWritesPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket081M(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket081M", (value))
}

// GetSuccessfulWritesPersecBucket081M gets the value of SuccessfulWritesPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket081M()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket081M")
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

// SetSuccessfulWritesPersecBucket091M sets the value of SuccessfulWritesPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertySuccessfulWritesPersecBucket091M(value uint64) (err error) { 
    return instance.SetProperty("SuccessfulWritesPersecBucket091M", (value))
}

// GetSuccessfulWritesPersecBucket091M gets the value of SuccessfulWritesPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertySuccessfulWritesPersecBucket091M()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuccessfulWritesPersecBucket091M")
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

// SetWriteBytesAverage sets the value of WriteBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteBytesAverage(value uint64) (err error) { 
    return instance.SetProperty("WriteBytesAverage", (value))
}

// GetWriteBytesAverage gets the value of WriteBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteBytesAverage()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteBytesAverage")
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

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("WriteBytesPersec", (value))
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteBytesPersec()(value uint64, err error) { 
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

// SetWriteLatency sets the value of WriteLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatency(value uint64) (err error) { 
    return instance.SetProperty("WriteLatency", (value))
}

// GetWriteLatency gets the value of WriteLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatency()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatency")
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

// SetWriteLatencyBucket01128us sets the value of WriteLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket01128us(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket01128us", (value))
}

// GetWriteLatencyBucket01128us gets the value of WriteLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket01128us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket01128us")
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

// SetWriteLatencyBucket02256us sets the value of WriteLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket02256us(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket02256us", (value))
}

// GetWriteLatencyBucket02256us gets the value of WriteLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket02256us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket02256us")
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

// SetWriteLatencyBucket03512us sets the value of WriteLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket03512us(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket03512us", (value))
}

// GetWriteLatencyBucket03512us gets the value of WriteLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket03512us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket03512us")
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

// SetWriteLatencyBucket041ms sets the value of WriteLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket041ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket041ms", (value))
}

// GetWriteLatencyBucket041ms gets the value of WriteLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket041ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket041ms")
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

// SetWriteLatencyBucket054ms sets the value of WriteLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket054ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket054ms", (value))
}

// GetWriteLatencyBucket054ms gets the value of WriteLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket054ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket054ms")
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

// SetWriteLatencyBucket0616ms sets the value of WriteLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket0616ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket0616ms", (value))
}

// GetWriteLatencyBucket0616ms gets the value of WriteLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket0616ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket0616ms")
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

// SetWriteLatencyBucket0764ms sets the value of WriteLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket0764ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket0764ms", (value))
}

// GetWriteLatencyBucket0764ms gets the value of WriteLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket0764ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket0764ms")
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

// SetWriteLatencyBucket08128ms sets the value of WriteLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket08128ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket08128ms", (value))
}

// GetWriteLatencyBucket08128ms gets the value of WriteLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket08128ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket08128ms")
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

// SetWriteLatencyBucket09256ms sets the value of WriteLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket09256ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket09256ms", (value))
}

// GetWriteLatencyBucket09256ms gets the value of WriteLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket09256ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket09256ms")
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

// SetWriteLatencyBucket10512ms sets the value of WriteLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket10512ms(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket10512ms", (value))
}

// GetWriteLatencyBucket10512ms gets the value of WriteLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket10512ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket10512ms")
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

// SetWriteLatencyBucket111s sets the value of WriteLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket111s(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket111s", (value))
}

// GetWriteLatencyBucket111s gets the value of WriteLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket111s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket111s")
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

// SetWriteLatencyBucket122s sets the value of WriteLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket122s(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket122s", (value))
}

// GetWriteLatencyBucket122s gets the value of WriteLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket122s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket122s")
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

// SetWriteLatencyBucket1310s sets the value of WriteLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket1310s(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket1310s", (value))
}

// GetWriteLatencyBucket1310s gets the value of WriteLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket1310s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket1310s")
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

// SetWriteLatencyBucket1410s sets the value of WriteLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWriteLatencyBucket1410s(value uint64) (err error) { 
    return instance.SetProperty("WriteLatencyBucket1410s", (value))
}

// GetWriteLatencyBucket1410s gets the value of WriteLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWriteLatencyBucket1410s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteLatencyBucket1410s")
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

// SetWritesPersec sets the value of WritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersec(value uint64) (err error) { 
    return instance.SetProperty("WritesPersec", (value))
}

// GetWritesPersec gets the value of WritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersec")
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

// SetWritesPersecBucket01128us sets the value of WritesPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket01128us(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket01128us", (value))
}

// GetWritesPersecBucket01128us gets the value of WritesPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket01128us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket01128us")
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

// SetWritesPersecBucket02256us sets the value of WritesPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket02256us(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket02256us", (value))
}

// GetWritesPersecBucket02256us gets the value of WritesPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket02256us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket02256us")
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

// SetWritesPersecBucket03512us sets the value of WritesPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket03512us(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket03512us", (value))
}

// GetWritesPersecBucket03512us gets the value of WritesPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket03512us()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket03512us")
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

// SetWritesPersecBucket041ms sets the value of WritesPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket041ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket041ms", (value))
}

// GetWritesPersecBucket041ms gets the value of WritesPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket041ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket041ms")
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

// SetWritesPersecBucket054ms sets the value of WritesPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket054ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket054ms", (value))
}

// GetWritesPersecBucket054ms gets the value of WritesPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket054ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket054ms")
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

// SetWritesPersecBucket0616ms sets the value of WritesPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket0616ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket0616ms", (value))
}

// GetWritesPersecBucket0616ms gets the value of WritesPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket0616ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket0616ms")
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

// SetWritesPersecBucket0764ms sets the value of WritesPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket0764ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket0764ms", (value))
}

// GetWritesPersecBucket0764ms gets the value of WritesPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket0764ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket0764ms")
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

// SetWritesPersecBucket08128ms sets the value of WritesPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket08128ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket08128ms", (value))
}

// GetWritesPersecBucket08128ms gets the value of WritesPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket08128ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket08128ms")
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

// SetWritesPersecBucket09256ms sets the value of WritesPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket09256ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket09256ms", (value))
}

// GetWritesPersecBucket09256ms gets the value of WritesPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket09256ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket09256ms")
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

// SetWritesPersecBucket10512ms sets the value of WritesPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket10512ms(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket10512ms", (value))
}

// GetWritesPersecBucket10512ms gets the value of WritesPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket10512ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket10512ms")
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

// SetWritesPersecBucket111s sets the value of WritesPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket111s(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket111s", (value))
}

// GetWritesPersecBucket111s gets the value of WritesPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket111s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket111s")
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

// SetWritesPersecBucket122s sets the value of WritesPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket122s(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket122s", (value))
}

// GetWritesPersecBucket122s gets the value of WritesPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket122s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket122s")
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

// SetWritesPersecBucket1310s sets the value of WritesPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket1310s(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket1310s", (value))
}

// GetWritesPersecBucket1310s gets the value of WritesPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket1310s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket1310s")
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

// SetWritesPersecBucket1410s sets the value of WritesPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) SetPropertyWritesPersecBucket1410s(value uint64) (err error) { 
    return instance.SetProperty("WritesPersecBucket1410s", (value))
}

// GetWritesPersecBucket1410s gets the value of WritesPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitWrites) GetPropertyWritesPersecBucket1410s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPersecBucket1410s")
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

