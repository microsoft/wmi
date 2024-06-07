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

// ThreadSetName struct
type ThreadSetName struct { 
	*Thread_V2

	// 
	ProcessId uint32

	// 
	ThreadId uint32

	// 
	ThreadName string
}

	func NewThreadSetNameEx1(instance *cim.WmiInstance) (newInstance *ThreadSetName, err error) {tmp, err := NewThread_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &ThreadSetName {
	Thread_V2: tmp,
	}
	return
	}
	

	func NewThreadSetNameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ThreadSetName, err error) {tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ThreadSetName {
	Thread_V2: tmp,
	}
	return
	}
	

// SetProcessId sets the value of ProcessId for the instance
func (instance *ThreadSetName) SetPropertyProcessId(value uint32) (err error) { 
    return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *ThreadSetName) GetPropertyProcessId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProcessId")
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
func (instance *ThreadSetName) SetPropertyThreadId(value uint32) (err error) { 
    return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *ThreadSetName) GetPropertyThreadId()(value uint32, err error) { 
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

// SetThreadName sets the value of ThreadName for the instance
func (instance *ThreadSetName) SetPropertyThreadName(value string) (err error) { 
    return instance.SetProperty("ThreadName", (value))
}

// GetThreadName gets the value of ThreadName for the instance
func (instance *ThreadSetName) GetPropertyThreadName()(value string, err error) { 
    retValue, err := instance.GetProperty("ThreadName")
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

