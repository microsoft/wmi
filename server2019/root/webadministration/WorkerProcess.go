// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// WorkerProcess struct
type WorkerProcess struct { 
	*Object

	// 
	AppPoolName string

	// 
	Guid string

	// 
	ProcessId uint32
}

	func NewWorkerProcessEx1(instance *cim.WmiInstance) (newInstance *WorkerProcess, err error) {tmp, err := NewObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &WorkerProcess {
	Object: tmp,
	}
	return
	}
	

	func NewWorkerProcessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WorkerProcess, err error) {tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WorkerProcess {
	Object: tmp,
	}
	return
	}
	

// SetAppPoolName sets the value of AppPoolName for the instance
func (instance *WorkerProcess) SetPropertyAppPoolName(value string) (err error) { 
    return instance.SetProperty("AppPoolName", (value))
}

// GetAppPoolName gets the value of AppPoolName for the instance
func (instance *WorkerProcess) GetPropertyAppPoolName()(value string, err error) { 
    retValue, err := instance.GetProperty("AppPoolName")
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

// SetGuid sets the value of Guid for the instance
func (instance *WorkerProcess) SetPropertyGuid(value string) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *WorkerProcess) GetPropertyGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("Guid")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *WorkerProcess) SetPropertyProcessId(value uint32) (err error) { 
    return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *WorkerProcess) GetPropertyProcessId()(value uint32, err error) { 
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

// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WorkerProcess) GetState() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("GetState" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="OutputElement" type="HttpRequest []"></param>
func (instance *WorkerProcess) GetExecutingRequests( /* OUT */ OutputElement []HttpRequest) (err error) {_, err = instance.InvokeMethod("GetExecutingRequests" )
	if err != nil { return }
	return
	
}


