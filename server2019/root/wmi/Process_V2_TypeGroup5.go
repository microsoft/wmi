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

// Process_V2_TypeGroup5 struct
type Process_V2_TypeGroup5 struct { 
	*Process_V2

	// 
	Object uint32
}

	func NewProcess_V2_TypeGroup5Ex1(instance *cim.WmiInstance) (newInstance *Process_V2_TypeGroup5, err error) {tmp, err := NewProcess_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &Process_V2_TypeGroup5 {
	Process_V2: tmp,
	}
	return
	}
	

	func NewProcess_V2_TypeGroup5Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Process_V2_TypeGroup5, err error) {tmp, err := NewProcess_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Process_V2_TypeGroup5 {
	Process_V2: tmp,
	}
	return
	}
	

// SetObject sets the value of Object for the instance
func (instance *Process_V2_TypeGroup5) SetPropertyObject(value uint32) (err error) { 
    return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *Process_V2_TypeGroup5) GetPropertyObject()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Object")
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

