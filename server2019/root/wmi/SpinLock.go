// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SpinLock struct
type SpinLock struct { 
	*Thread_V2

	// 
	AcquireDepth uint8

	// 
	AcquireTime uint64

	// 
	CallerAddress uint32

	// 
	Flag uint8

	// 
	InterruptCount uint32

	// 
	Irql uint8

	// 
	ReleaseTime uint64

	// 
	Reserved []uint8

	// 
	SpinCount uint32

	// 
	SpinLockAddress uint32

	// 
	ThreadId uint32

	// 
	WaitTimeInCycles uint32
}

	func NewSpinLockEx1(instance *cim.WmiInstance) (newInstance *SpinLock, err error) {tmp, err := NewThread_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SpinLock {
	Thread_V2: tmp,
	}
	return
	}
	

	func NewSpinLockEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SpinLock, err error) {tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SpinLock {
	Thread_V2: tmp,
	}
	return
	}
	

// SetAcquireDepth sets the value of AcquireDepth for the instance
func (instance *SpinLock) SetPropertyAcquireDepth(value uint8) (err error) { 
    return instance.SetProperty("AcquireDepth", (value))
}

// GetAcquireDepth gets the value of AcquireDepth for the instance
func (instance *SpinLock) GetPropertyAcquireDepth()(value uint8, err error) { 
    retValue, err := instance.GetProperty("AcquireDepth")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetAcquireTime sets the value of AcquireTime for the instance
func (instance *SpinLock) SetPropertyAcquireTime(value uint64) (err error) { 
    return instance.SetProperty("AcquireTime", (value))
}

// GetAcquireTime gets the value of AcquireTime for the instance
func (instance *SpinLock) GetPropertyAcquireTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AcquireTime")
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

// SetCallerAddress sets the value of CallerAddress for the instance
func (instance *SpinLock) SetPropertyCallerAddress(value uint32) (err error) { 
    return instance.SetProperty("CallerAddress", (value))
}

// GetCallerAddress gets the value of CallerAddress for the instance
func (instance *SpinLock) GetPropertyCallerAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CallerAddress")
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

// SetFlag sets the value of Flag for the instance
func (instance *SpinLock) SetPropertyFlag(value uint8) (err error) { 
    return instance.SetProperty("Flag", (value))
}

// GetFlag gets the value of Flag for the instance
func (instance *SpinLock) GetPropertyFlag()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Flag")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetInterruptCount sets the value of InterruptCount for the instance
func (instance *SpinLock) SetPropertyInterruptCount(value uint32) (err error) { 
    return instance.SetProperty("InterruptCount", (value))
}

// GetInterruptCount gets the value of InterruptCount for the instance
func (instance *SpinLock) GetPropertyInterruptCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InterruptCount")
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

// SetIrql sets the value of Irql for the instance
func (instance *SpinLock) SetPropertyIrql(value uint8) (err error) { 
    return instance.SetProperty("Irql", (value))
}

// GetIrql gets the value of Irql for the instance
func (instance *SpinLock) GetPropertyIrql()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Irql")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetReleaseTime sets the value of ReleaseTime for the instance
func (instance *SpinLock) SetPropertyReleaseTime(value uint64) (err error) { 
    return instance.SetProperty("ReleaseTime", (value))
}

// GetReleaseTime gets the value of ReleaseTime for the instance
func (instance *SpinLock) GetPropertyReleaseTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReleaseTime")
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

// SetReserved sets the value of Reserved for the instance
func (instance *SpinLock) SetPropertyReserved(value []uint8) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *SpinLock) GetPropertyReserved()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Reserved")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetSpinCount sets the value of SpinCount for the instance
func (instance *SpinLock) SetPropertySpinCount(value uint32) (err error) { 
    return instance.SetProperty("SpinCount", (value))
}

// GetSpinCount gets the value of SpinCount for the instance
func (instance *SpinLock) GetPropertySpinCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpinCount")
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

// SetSpinLockAddress sets the value of SpinLockAddress for the instance
func (instance *SpinLock) SetPropertySpinLockAddress(value uint32) (err error) { 
    return instance.SetProperty("SpinLockAddress", (value))
}

// GetSpinLockAddress gets the value of SpinLockAddress for the instance
func (instance *SpinLock) GetPropertySpinLockAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpinLockAddress")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *SpinLock) SetPropertyThreadId(value uint32) (err error) { 
    return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *SpinLock) GetPropertyThreadId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ThreadId")
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

// SetWaitTimeInCycles sets the value of WaitTimeInCycles for the instance
func (instance *SpinLock) SetPropertyWaitTimeInCycles(value uint32) (err error) { 
    return instance.SetProperty("WaitTimeInCycles", (value))
}

// GetWaitTimeInCycles gets the value of WaitTimeInCycles for the instance
func (instance *SpinLock) GetPropertyWaitTimeInCycles()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WaitTimeInCycles")
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

