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

// MSFT_TaskPrincipal struct
type MSFT_TaskPrincipal struct { 
	*cim.WmiInstance

	// 
	DisplayName string

	// 
	GroupId string

	// 
	Id string

	// 
	LogonType TaskPrincipal_LogonType

	// 
	RunLevel int32

	// 
	UserId string
}

	func NewMSFT_TaskPrincipalEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskPrincipal, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_TaskPrincipal {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_TaskPrincipalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_TaskPrincipal, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_TaskPrincipal {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyDisplayName(value string) (err error) { 
    return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyDisplayName()(value string, err error) { 
    retValue, err := instance.GetProperty("DisplayName")
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

// SetGroupId sets the value of GroupId for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyGroupId(value string) (err error) { 
    return instance.SetProperty("GroupId", (value))
}

// GetGroupId gets the value of GroupId for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyGroupId()(value string, err error) { 
    retValue, err := instance.GetProperty("GroupId")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyId(value string) (err error) { 
    return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyId()(value string, err error) { 
    retValue, err := instance.GetProperty("Id")
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

// SetLogonType sets the value of LogonType for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyLogonType(value TaskPrincipal_LogonType) (err error) { 
    return instance.SetProperty("LogonType", (value))
}

// GetLogonType gets the value of LogonType for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyLogonType()(value TaskPrincipal_LogonType, err error) { 
    retValue, err := instance.GetProperty("LogonType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = TaskPrincipal_LogonType(valuetmp)

    return
}

// SetRunLevel sets the value of RunLevel for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyRunLevel(value int32) (err error) { 
    return instance.SetProperty("RunLevel", (value))
}

// GetRunLevel gets the value of RunLevel for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyRunLevel()(value int32, err error) { 
    retValue, err := instance.GetProperty("RunLevel")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetUserId sets the value of UserId for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyUserId(value string) (err error) { 
    return instance.SetProperty("UserId", (value))
}

// GetUserId gets the value of UserId for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyUserId()(value string, err error) { 
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

