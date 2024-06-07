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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_TaskLogonTrigger struct
type MSFT_TaskLogonTrigger struct { 
	*MSFT_TaskTrigger

	// 
	Delay string

	// 
	UserId string
}

	func NewMSFT_TaskLogonTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskLogonTrigger, err error) {tmp, err := NewMSFT_TaskTriggerEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_TaskLogonTrigger {
	MSFT_TaskTrigger: tmp,
	}
	return
	}
	

	func NewMSFT_TaskLogonTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_TaskLogonTrigger, err error) {tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_TaskLogonTrigger {
	MSFT_TaskTrigger: tmp,
	}
	return
	}
	

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskLogonTrigger) SetPropertyDelay(value string) (err error) { 
    return instance.SetProperty("Delay", (value))
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskLogonTrigger) GetPropertyDelay()(value string, err error) { 
    retValue, err := instance.GetProperty("Delay")
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

// SetUserId sets the value of UserId for the instance
func (instance *MSFT_TaskLogonTrigger) SetPropertyUserId(value string) (err error) { 
    return instance.SetProperty("UserId", (value))
}

// GetUserId gets the value of UserId for the instance
func (instance *MSFT_TaskLogonTrigger) GetPropertyUserId()(value string, err error) { 
    retValue, err := instance.GetProperty("UserId")
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

