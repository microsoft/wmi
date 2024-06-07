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

// PmcCounterProfile struct
type PmcCounterProfile struct { 
	*PerfInfo_V2

	// 
	InstructionPointer uint32

	// 
	ProfileSource uint16

	// 
	Reserved uint16

	// 
	ThreadId uint32
}

	func NewPmcCounterProfileEx1(instance *cim.WmiInstance) (newInstance *PmcCounterProfile, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &PmcCounterProfile {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewPmcCounterProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PmcCounterProfile, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PmcCounterProfile {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetInstructionPointer sets the value of InstructionPointer for the instance
func (instance *PmcCounterProfile) SetPropertyInstructionPointer(value uint32) (err error) { 
    return instance.SetProperty("InstructionPointer", (value))
}

// GetInstructionPointer gets the value of InstructionPointer for the instance
func (instance *PmcCounterProfile) GetPropertyInstructionPointer()(value uint32, err error) { 
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

// SetProfileSource sets the value of ProfileSource for the instance
func (instance *PmcCounterProfile) SetPropertyProfileSource(value uint16) (err error) { 
    return instance.SetProperty("ProfileSource", (value))
}

// GetProfileSource gets the value of ProfileSource for the instance
func (instance *PmcCounterProfile) GetPropertyProfileSource()(value uint16, err error) { 
    retValue, err := instance.GetProperty("ProfileSource")
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

// SetReserved sets the value of Reserved for the instance
func (instance *PmcCounterProfile) SetPropertyReserved(value uint16) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *PmcCounterProfile) GetPropertyReserved()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Reserved")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *PmcCounterProfile) SetPropertyThreadId(value uint32) (err error) { 
    return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *PmcCounterProfile) GetPropertyThreadId()(value uint32, err error) { 
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

