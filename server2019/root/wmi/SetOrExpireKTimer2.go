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

// SetOrExpireKTimer2 struct
type SetOrExpireKTimer2 struct { 
	*PerfInfo_V2

	// 
	Callback uint32

	// 
	CallbackContext uint32

	// 
	DueTime uint64

	// 
	MaximumDueTime uint64

	// 
	Period uint64

	// 
	Timer uint32

	// 
	TimerFlags uint8
}

	func NewSetOrExpireKTimer2Ex1(instance *cim.WmiInstance) (newInstance *SetOrExpireKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SetOrExpireKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewSetOrExpireKTimer2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SetOrExpireKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SetOrExpireKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetCallback sets the value of Callback for the instance
func (instance *SetOrExpireKTimer2) SetPropertyCallback(value uint32) (err error) { 
    return instance.SetProperty("Callback", (value))
}

// GetCallback gets the value of Callback for the instance
func (instance *SetOrExpireKTimer2) GetPropertyCallback()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Callback")
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

// SetCallbackContext sets the value of CallbackContext for the instance
func (instance *SetOrExpireKTimer2) SetPropertyCallbackContext(value uint32) (err error) { 
    return instance.SetProperty("CallbackContext", (value))
}

// GetCallbackContext gets the value of CallbackContext for the instance
func (instance *SetOrExpireKTimer2) GetPropertyCallbackContext()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CallbackContext")
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

// SetDueTime sets the value of DueTime for the instance
func (instance *SetOrExpireKTimer2) SetPropertyDueTime(value uint64) (err error) { 
    return instance.SetProperty("DueTime", (value))
}

// GetDueTime gets the value of DueTime for the instance
func (instance *SetOrExpireKTimer2) GetPropertyDueTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DueTime")
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

// SetMaximumDueTime sets the value of MaximumDueTime for the instance
func (instance *SetOrExpireKTimer2) SetPropertyMaximumDueTime(value uint64) (err error) { 
    return instance.SetProperty("MaximumDueTime", (value))
}

// GetMaximumDueTime gets the value of MaximumDueTime for the instance
func (instance *SetOrExpireKTimer2) GetPropertyMaximumDueTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaximumDueTime")
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

// SetPeriod sets the value of Period for the instance
func (instance *SetOrExpireKTimer2) SetPropertyPeriod(value uint64) (err error) { 
    return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *SetOrExpireKTimer2) GetPropertyPeriod()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Period")
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

// SetTimer sets the value of Timer for the instance
func (instance *SetOrExpireKTimer2) SetPropertyTimer(value uint32) (err error) { 
    return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *SetOrExpireKTimer2) GetPropertyTimer()(value uint32, err error) { 
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

// SetTimerFlags sets the value of TimerFlags for the instance
func (instance *SetOrExpireKTimer2) SetPropertyTimerFlags(value uint8) (err error) { 
    return instance.SetProperty("TimerFlags", (value))
}

// GetTimerFlags gets the value of TimerFlags for the instance
func (instance *SetOrExpireKTimer2) GetPropertyTimerFlags()(value uint8, err error) { 
    retValue, err := instance.GetProperty("TimerFlags")
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

