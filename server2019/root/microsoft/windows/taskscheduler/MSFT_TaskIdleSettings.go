// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_TaskIdleSettings struct
type MSFT_TaskIdleSettings struct { 
	*cim.WmiInstance

	// 
	IdleDuration string

	// 
	RestartOnIdle bool

	// 
	StopOnIdleEnd bool

	// 
	WaitTimeout string
}

	func NewMSFT_TaskIdleSettingsEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskIdleSettings, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_TaskIdleSettings {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_TaskIdleSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_TaskIdleSettings, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_TaskIdleSettings {
	WmiInstance: tmp,
	}
	return
	}
	

// SetIdleDuration sets the value of IdleDuration for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyIdleDuration(value string) (err error) { 
    return instance.SetProperty("IdleDuration", (value))
}

// GetIdleDuration gets the value of IdleDuration for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyIdleDuration()(value string, err error) { 
    retValue, err := instance.GetProperty("IdleDuration")
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

// SetRestartOnIdle sets the value of RestartOnIdle for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyRestartOnIdle(value bool) (err error) { 
    return instance.SetProperty("RestartOnIdle", (value))
}

// GetRestartOnIdle gets the value of RestartOnIdle for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyRestartOnIdle()(value bool, err error) { 
    retValue, err := instance.GetProperty("RestartOnIdle")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetStopOnIdleEnd sets the value of StopOnIdleEnd for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyStopOnIdleEnd(value bool) (err error) { 
    return instance.SetProperty("StopOnIdleEnd", (value))
}

// GetStopOnIdleEnd gets the value of StopOnIdleEnd for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyStopOnIdleEnd()(value bool, err error) { 
    retValue, err := instance.GetProperty("StopOnIdleEnd")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetWaitTimeout sets the value of WaitTimeout for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyWaitTimeout(value string) (err error) { 
    return instance.SetProperty("WaitTimeout", (value))
}

// GetWaitTimeout gets the value of WaitTimeout for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyWaitTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("WaitTimeout")
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

