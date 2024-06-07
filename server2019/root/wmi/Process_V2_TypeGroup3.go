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

// Process_V2_TypeGroup3 struct
type Process_V2_TypeGroup3 struct { 
	*Process_V2

	// 
	DirectoryTableBase uint32

	// 
	ProcessId uint32
}

	func NewProcess_V2_TypeGroup3Ex1(instance *cim.WmiInstance) (newInstance *Process_V2_TypeGroup3, err error) {tmp, err := NewProcess_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &Process_V2_TypeGroup3 {
	Process_V2: tmp,
	}
	return
	}
	

	func NewProcess_V2_TypeGroup3Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Process_V2_TypeGroup3, err error) {tmp, err := NewProcess_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Process_V2_TypeGroup3 {
	Process_V2: tmp,
	}
	return
	}
	

// SetDirectoryTableBase sets the value of DirectoryTableBase for the instance
func (instance *Process_V2_TypeGroup3) SetPropertyDirectoryTableBase(value uint32) (err error) { 
    return instance.SetProperty("DirectoryTableBase", (value))
}

// GetDirectoryTableBase gets the value of DirectoryTableBase for the instance
func (instance *Process_V2_TypeGroup3) GetPropertyDirectoryTableBase()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DirectoryTableBase")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup3) SetPropertyProcessId(value uint32) (err error) { 
    return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup3) GetPropertyProcessId()(value uint32, err error) { 
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

