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

// SpinLockConfig_V2 struct
type SpinLockConfig_V2 struct { 
	*PerfInfo_V2

	// 
	SpinLockAcquireSampleRate uint32

	// 
	SpinLockContentionSampleRate uint32

	// 
	SpinLockSpinThreshold uint32
}

	func NewSpinLockConfig_V2Ex1(instance *cim.WmiInstance) (newInstance *SpinLockConfig_V2, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SpinLockConfig_V2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewSpinLockConfig_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SpinLockConfig_V2, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SpinLockConfig_V2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetSpinLockAcquireSampleRate sets the value of SpinLockAcquireSampleRate for the instance
func (instance *SpinLockConfig_V2) SetPropertySpinLockAcquireSampleRate(value uint32) (err error) { 
    return instance.SetProperty("SpinLockAcquireSampleRate", (value))
}

// GetSpinLockAcquireSampleRate gets the value of SpinLockAcquireSampleRate for the instance
func (instance *SpinLockConfig_V2) GetPropertySpinLockAcquireSampleRate()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpinLockAcquireSampleRate")
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

// SetSpinLockContentionSampleRate sets the value of SpinLockContentionSampleRate for the instance
func (instance *SpinLockConfig_V2) SetPropertySpinLockContentionSampleRate(value uint32) (err error) { 
    return instance.SetProperty("SpinLockContentionSampleRate", (value))
}

// GetSpinLockContentionSampleRate gets the value of SpinLockContentionSampleRate for the instance
func (instance *SpinLockConfig_V2) GetPropertySpinLockContentionSampleRate()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpinLockContentionSampleRate")
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

// SetSpinLockSpinThreshold sets the value of SpinLockSpinThreshold for the instance
func (instance *SpinLockConfig_V2) SetPropertySpinLockSpinThreshold(value uint32) (err error) { 
    return instance.SetProperty("SpinLockSpinThreshold", (value))
}

// GetSpinLockSpinThreshold gets the value of SpinLockSpinThreshold for the instance
func (instance *SpinLockConfig_V2) GetPropertySpinLockSpinThreshold()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SpinLockSpinThreshold")
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

