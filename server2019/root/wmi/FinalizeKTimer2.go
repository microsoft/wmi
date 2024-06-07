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

// FinalizeKTimer2 struct
type FinalizeKTimer2 struct { 
	*PerfInfo_V2

	// 
	DisableCallback uint32

	// 
	DisableContext uint32

	// 
	Timer uint32
}

	func NewFinalizeKTimer2Ex1(instance *cim.WmiInstance) (newInstance *FinalizeKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &FinalizeKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewFinalizeKTimer2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FinalizeKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FinalizeKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetDisableCallback sets the value of DisableCallback for the instance
func (instance *FinalizeKTimer2) SetPropertyDisableCallback(value uint32) (err error) { 
    return instance.SetProperty("DisableCallback", (value))
}

// GetDisableCallback gets the value of DisableCallback for the instance
func (instance *FinalizeKTimer2) GetPropertyDisableCallback()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DisableCallback")
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

// SetDisableContext sets the value of DisableContext for the instance
func (instance *FinalizeKTimer2) SetPropertyDisableContext(value uint32) (err error) { 
    return instance.SetProperty("DisableContext", (value))
}

// GetDisableContext gets the value of DisableContext for the instance
func (instance *FinalizeKTimer2) GetPropertyDisableContext()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DisableContext")
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

// SetTimer sets the value of Timer for the instance
func (instance *FinalizeKTimer2) SetPropertyTimer(value uint32) (err error) { 
    return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *FinalizeKTimer2) GetPropertyTimer()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Timer")
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

