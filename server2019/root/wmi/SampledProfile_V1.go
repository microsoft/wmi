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

// SampledProfile_V1 struct
type SampledProfile_V1 struct { 
	*PerfInfo_V1

	// 
	Count uint16

	// 
	InstructionPointer uint32

	// 
	ThreadId uint32
}

	func NewSampledProfile_V1Ex1(instance *cim.WmiInstance) (newInstance *SampledProfile_V1, err error) {tmp, err := NewPerfInfo_V1Ex1(instance)
		
	if err != nil { return }
	newInstance = &SampledProfile_V1 {
	PerfInfo_V1: tmp,
	}
	return
	}
	

	func NewSampledProfile_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SampledProfile_V1, err error) {tmp, err := NewPerfInfo_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SampledProfile_V1 {
	PerfInfo_V1: tmp,
	}
	return
	}
	

// SetCount sets the value of Count for the instance
func (instance *SampledProfile_V1) SetPropertyCount(value uint16) (err error) { 
    return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *SampledProfile_V1) GetPropertyCount()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Count")
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

// SetInstructionPointer sets the value of InstructionPointer for the instance
func (instance *SampledProfile_V1) SetPropertyInstructionPointer(value uint32) (err error) { 
    return instance.SetProperty("InstructionPointer", (value))
}

// GetInstructionPointer gets the value of InstructionPointer for the instance
func (instance *SampledProfile_V1) GetPropertyInstructionPointer()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InstructionPointer")
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
func (instance *SampledProfile_V1) SetPropertyThreadId(value uint32) (err error) { 
    return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *SampledProfile_V1) GetPropertyThreadId()(value uint32, err error) { 
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

