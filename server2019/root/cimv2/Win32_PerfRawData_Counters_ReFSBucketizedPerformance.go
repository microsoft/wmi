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

// Win32_PerfRawData_Counters_ReFSBucketizedPerformance struct
type Win32_PerfRawData_Counters_ReFSBucketizedPerformance struct { 
	*Win32_PerfRawData

	// 
	Value1TotalOperationsCount uint64

	// 
	Value2Totaltimensforalloperations uint64

	// 
	Value3Totalbytesforalloperations uint64

	// 
	Value4Countbucket01128µs uint64

	// 
	Value4Countbucket02256µs uint64

	// 
	Value4Countbucket03512µs uint64

	// 
	Value4Countbucket041ms uint64

	// 
	Value4Countbucket054ms uint64

	// 
	Value4Countbucket0616ms uint64

	// 
	Value4Countbucket0764ms uint64

	// 
	Value4Countbucket08128ms uint64

	// 
	Value4Countbucket09256ms uint64

	// 
	Value4Countbucket10512ms uint64

	// 
	Value4Countbucket111s uint64

	// 
	Value4Countbucket122s uint64

	// 
	Value4Countbucket1310s uint64

	// 
	Value4Countbucket1420s uint64

	// 
	Value4Countbucket1530s uint64

	// 
	Value4Countbucket1630s uint64

	// 
	Value5Timensbucket01128µs uint64

	// 
	Value5Timensbucket02256µs uint64

	// 
	Value5Timensbucket03512µs uint64

	// 
	Value5Timensbucket041ms uint64

	// 
	Value5Timensbucket054ms uint64

	// 
	Value5Timensbucket0616ms uint64

	// 
	Value5Timensbucket0764ms uint64

	// 
	Value5Timensbucket08128ms uint64

	// 
	Value5Timensbucket09256ms uint64

	// 
	Value5Timensbucket10512ms uint64

	// 
	Value5Timensbucket111s uint64

	// 
	Value5Timensbucket122s uint64

	// 
	Value5Timensbucket1310s uint64

	// 
	Value5Timensbucket1420s uint64

	// 
	Value5Timensbucket1530s uint64

	// 
	Value5Timensbucket1630s uint64

	// 
	Value6Bytesbucket01128µs uint64

	// 
	Value6Bytesbucket02256µs uint64

	// 
	Value6Bytesbucket03512µs uint64

	// 
	Value6Bytesbucket041ms uint64

	// 
	Value6Bytesbucket054ms uint64

	// 
	Value6Bytesbucket0616ms uint64

	// 
	Value6Bytesbucket0764ms uint64

	// 
	Value6Bytesbucket08128ms uint64

	// 
	Value6Bytesbucket09256ms uint64

	// 
	Value6Bytesbucket10512ms uint64

	// 
	Value6Bytesbucket111s uint64

	// 
	Value6Bytesbucket122s uint64

	// 
	Value6Bytesbucket1310s uint64

	// 
	Value6Bytesbucket1420s uint64

	// 
	Value6Bytesbucket1530s uint64

	// 
	Value6Bytesbucket1630s uint64
}

	func NewWin32_PerfRawData_Counters_ReFSBucketizedPerformanceEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Counters_ReFSBucketizedPerformance {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_Counters_ReFSBucketizedPerformanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Counters_ReFSBucketizedPerformance {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetValue1TotalOperationsCount sets the value of Value1TotalOperationsCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue1TotalOperationsCount(value uint64) (err error) { 
    return instance.SetProperty("Value1TotalOperationsCount", (value))
}

// GetValue1TotalOperationsCount gets the value of Value1TotalOperationsCount for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue1TotalOperationsCount()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value1TotalOperationsCount")
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

// SetValue2Totaltimensforalloperations sets the value of Value2Totaltimensforalloperations for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue2Totaltimensforalloperations(value uint64) (err error) { 
    return instance.SetProperty("Value2Totaltimensforalloperations", (value))
}

// GetValue2Totaltimensforalloperations gets the value of Value2Totaltimensforalloperations for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue2Totaltimensforalloperations()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value2Totaltimensforalloperations")
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

// SetValue3Totalbytesforalloperations sets the value of Value3Totalbytesforalloperations for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue3Totalbytesforalloperations(value uint64) (err error) { 
    return instance.SetProperty("Value3Totalbytesforalloperations", (value))
}

// GetValue3Totalbytesforalloperations gets the value of Value3Totalbytesforalloperations for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue3Totalbytesforalloperations()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value3Totalbytesforalloperations")
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

// SetValue4Countbucket01128µs sets the value of Value4Countbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket01128µs(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket01128µs", (value))
}

// GetValue4Countbucket01128µs gets the value of Value4Countbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket01128µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket01128µs")
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

// SetValue4Countbucket02256µs sets the value of Value4Countbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket02256µs(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket02256µs", (value))
}

// GetValue4Countbucket02256µs gets the value of Value4Countbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket02256µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket02256µs")
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

// SetValue4Countbucket03512µs sets the value of Value4Countbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket03512µs(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket03512µs", (value))
}

// GetValue4Countbucket03512µs gets the value of Value4Countbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket03512µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket03512µs")
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

// SetValue4Countbucket041ms sets the value of Value4Countbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket041ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket041ms", (value))
}

// GetValue4Countbucket041ms gets the value of Value4Countbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket041ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket041ms")
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

// SetValue4Countbucket054ms sets the value of Value4Countbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket054ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket054ms", (value))
}

// GetValue4Countbucket054ms gets the value of Value4Countbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket054ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket054ms")
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

// SetValue4Countbucket0616ms sets the value of Value4Countbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket0616ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket0616ms", (value))
}

// GetValue4Countbucket0616ms gets the value of Value4Countbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket0616ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket0616ms")
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

// SetValue4Countbucket0764ms sets the value of Value4Countbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket0764ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket0764ms", (value))
}

// GetValue4Countbucket0764ms gets the value of Value4Countbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket0764ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket0764ms")
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

// SetValue4Countbucket08128ms sets the value of Value4Countbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket08128ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket08128ms", (value))
}

// GetValue4Countbucket08128ms gets the value of Value4Countbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket08128ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket08128ms")
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

// SetValue4Countbucket09256ms sets the value of Value4Countbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket09256ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket09256ms", (value))
}

// GetValue4Countbucket09256ms gets the value of Value4Countbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket09256ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket09256ms")
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

// SetValue4Countbucket10512ms sets the value of Value4Countbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket10512ms(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket10512ms", (value))
}

// GetValue4Countbucket10512ms gets the value of Value4Countbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket10512ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket10512ms")
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

// SetValue4Countbucket111s sets the value of Value4Countbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket111s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket111s", (value))
}

// GetValue4Countbucket111s gets the value of Value4Countbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket111s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket111s")
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

// SetValue4Countbucket122s sets the value of Value4Countbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket122s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket122s", (value))
}

// GetValue4Countbucket122s gets the value of Value4Countbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket122s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket122s")
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

// SetValue4Countbucket1310s sets the value of Value4Countbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket1310s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket1310s", (value))
}

// GetValue4Countbucket1310s gets the value of Value4Countbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket1310s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket1310s")
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

// SetValue4Countbucket1420s sets the value of Value4Countbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket1420s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket1420s", (value))
}

// GetValue4Countbucket1420s gets the value of Value4Countbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket1420s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket1420s")
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

// SetValue4Countbucket1530s sets the value of Value4Countbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket1530s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket1530s", (value))
}

// GetValue4Countbucket1530s gets the value of Value4Countbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket1530s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket1530s")
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

// SetValue4Countbucket1630s sets the value of Value4Countbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue4Countbucket1630s(value uint64) (err error) { 
    return instance.SetProperty("Value4Countbucket1630s", (value))
}

// GetValue4Countbucket1630s gets the value of Value4Countbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue4Countbucket1630s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value4Countbucket1630s")
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

// SetValue5Timensbucket01128µs sets the value of Value5Timensbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket01128µs(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket01128µs", (value))
}

// GetValue5Timensbucket01128µs gets the value of Value5Timensbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket01128µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket01128µs")
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

// SetValue5Timensbucket02256µs sets the value of Value5Timensbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket02256µs(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket02256µs", (value))
}

// GetValue5Timensbucket02256µs gets the value of Value5Timensbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket02256µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket02256µs")
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

// SetValue5Timensbucket03512µs sets the value of Value5Timensbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket03512µs(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket03512µs", (value))
}

// GetValue5Timensbucket03512µs gets the value of Value5Timensbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket03512µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket03512µs")
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

// SetValue5Timensbucket041ms sets the value of Value5Timensbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket041ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket041ms", (value))
}

// GetValue5Timensbucket041ms gets the value of Value5Timensbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket041ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket041ms")
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

// SetValue5Timensbucket054ms sets the value of Value5Timensbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket054ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket054ms", (value))
}

// GetValue5Timensbucket054ms gets the value of Value5Timensbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket054ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket054ms")
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

// SetValue5Timensbucket0616ms sets the value of Value5Timensbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket0616ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket0616ms", (value))
}

// GetValue5Timensbucket0616ms gets the value of Value5Timensbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket0616ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket0616ms")
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

// SetValue5Timensbucket0764ms sets the value of Value5Timensbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket0764ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket0764ms", (value))
}

// GetValue5Timensbucket0764ms gets the value of Value5Timensbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket0764ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket0764ms")
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

// SetValue5Timensbucket08128ms sets the value of Value5Timensbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket08128ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket08128ms", (value))
}

// GetValue5Timensbucket08128ms gets the value of Value5Timensbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket08128ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket08128ms")
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

// SetValue5Timensbucket09256ms sets the value of Value5Timensbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket09256ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket09256ms", (value))
}

// GetValue5Timensbucket09256ms gets the value of Value5Timensbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket09256ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket09256ms")
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

// SetValue5Timensbucket10512ms sets the value of Value5Timensbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket10512ms(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket10512ms", (value))
}

// GetValue5Timensbucket10512ms gets the value of Value5Timensbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket10512ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket10512ms")
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

// SetValue5Timensbucket111s sets the value of Value5Timensbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket111s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket111s", (value))
}

// GetValue5Timensbucket111s gets the value of Value5Timensbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket111s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket111s")
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

// SetValue5Timensbucket122s sets the value of Value5Timensbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket122s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket122s", (value))
}

// GetValue5Timensbucket122s gets the value of Value5Timensbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket122s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket122s")
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

// SetValue5Timensbucket1310s sets the value of Value5Timensbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket1310s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket1310s", (value))
}

// GetValue5Timensbucket1310s gets the value of Value5Timensbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket1310s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket1310s")
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

// SetValue5Timensbucket1420s sets the value of Value5Timensbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket1420s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket1420s", (value))
}

// GetValue5Timensbucket1420s gets the value of Value5Timensbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket1420s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket1420s")
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

// SetValue5Timensbucket1530s sets the value of Value5Timensbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket1530s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket1530s", (value))
}

// GetValue5Timensbucket1530s gets the value of Value5Timensbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket1530s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket1530s")
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

// SetValue5Timensbucket1630s sets the value of Value5Timensbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue5Timensbucket1630s(value uint64) (err error) { 
    return instance.SetProperty("Value5Timensbucket1630s", (value))
}

// GetValue5Timensbucket1630s gets the value of Value5Timensbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue5Timensbucket1630s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value5Timensbucket1630s")
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

// SetValue6Bytesbucket01128µs sets the value of Value6Bytesbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket01128µs(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket01128µs", (value))
}

// GetValue6Bytesbucket01128µs gets the value of Value6Bytesbucket01128µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket01128µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket01128µs")
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

// SetValue6Bytesbucket02256µs sets the value of Value6Bytesbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket02256µs(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket02256µs", (value))
}

// GetValue6Bytesbucket02256µs gets the value of Value6Bytesbucket02256µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket02256µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket02256µs")
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

// SetValue6Bytesbucket03512µs sets the value of Value6Bytesbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket03512µs(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket03512µs", (value))
}

// GetValue6Bytesbucket03512µs gets the value of Value6Bytesbucket03512µs for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket03512µs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket03512µs")
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

// SetValue6Bytesbucket041ms sets the value of Value6Bytesbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket041ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket041ms", (value))
}

// GetValue6Bytesbucket041ms gets the value of Value6Bytesbucket041ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket041ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket041ms")
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

// SetValue6Bytesbucket054ms sets the value of Value6Bytesbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket054ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket054ms", (value))
}

// GetValue6Bytesbucket054ms gets the value of Value6Bytesbucket054ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket054ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket054ms")
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

// SetValue6Bytesbucket0616ms sets the value of Value6Bytesbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket0616ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket0616ms", (value))
}

// GetValue6Bytesbucket0616ms gets the value of Value6Bytesbucket0616ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket0616ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket0616ms")
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

// SetValue6Bytesbucket0764ms sets the value of Value6Bytesbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket0764ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket0764ms", (value))
}

// GetValue6Bytesbucket0764ms gets the value of Value6Bytesbucket0764ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket0764ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket0764ms")
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

// SetValue6Bytesbucket08128ms sets the value of Value6Bytesbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket08128ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket08128ms", (value))
}

// GetValue6Bytesbucket08128ms gets the value of Value6Bytesbucket08128ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket08128ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket08128ms")
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

// SetValue6Bytesbucket09256ms sets the value of Value6Bytesbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket09256ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket09256ms", (value))
}

// GetValue6Bytesbucket09256ms gets the value of Value6Bytesbucket09256ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket09256ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket09256ms")
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

// SetValue6Bytesbucket10512ms sets the value of Value6Bytesbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket10512ms(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket10512ms", (value))
}

// GetValue6Bytesbucket10512ms gets the value of Value6Bytesbucket10512ms for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket10512ms()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket10512ms")
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

// SetValue6Bytesbucket111s sets the value of Value6Bytesbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket111s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket111s", (value))
}

// GetValue6Bytesbucket111s gets the value of Value6Bytesbucket111s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket111s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket111s")
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

// SetValue6Bytesbucket122s sets the value of Value6Bytesbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket122s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket122s", (value))
}

// GetValue6Bytesbucket122s gets the value of Value6Bytesbucket122s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket122s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket122s")
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

// SetValue6Bytesbucket1310s sets the value of Value6Bytesbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket1310s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket1310s", (value))
}

// GetValue6Bytesbucket1310s gets the value of Value6Bytesbucket1310s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket1310s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket1310s")
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

// SetValue6Bytesbucket1420s sets the value of Value6Bytesbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket1420s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket1420s", (value))
}

// GetValue6Bytesbucket1420s gets the value of Value6Bytesbucket1420s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket1420s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket1420s")
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

// SetValue6Bytesbucket1530s sets the value of Value6Bytesbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket1530s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket1530s", (value))
}

// GetValue6Bytesbucket1530s gets the value of Value6Bytesbucket1530s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket1530s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket1530s")
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

// SetValue6Bytesbucket1630s sets the value of Value6Bytesbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) SetPropertyValue6Bytesbucket1630s(value uint64) (err error) { 
    return instance.SetProperty("Value6Bytesbucket1630s", (value))
}

// GetValue6Bytesbucket1630s gets the value of Value6Bytesbucket1630s for the instance
func (instance *Win32_PerfRawData_Counters_ReFSBucketizedPerformance) GetPropertyValue6Bytesbucket1630s()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Value6Bytesbucket1630s")
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

