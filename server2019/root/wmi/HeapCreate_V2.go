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

// HeapCreate_V2 struct
type HeapCreate_V2 struct { 
	*HeapTrace_V2

	// 
	HeapFlags uint32

	// 
	HeapHandle uint32
}

	func NewHeapCreate_V2Ex1(instance *cim.WmiInstance) (newInstance *HeapCreate_V2, err error) {tmp, err := NewHeapTrace_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &HeapCreate_V2 {
	HeapTrace_V2: tmp,
	}
	return
	}
	

	func NewHeapCreate_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapCreate_V2, err error) {tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapCreate_V2 {
	HeapTrace_V2: tmp,
	}
	return
	}
	

// SetHeapFlags sets the value of HeapFlags for the instance
func (instance *HeapCreate_V2) SetPropertyHeapFlags(value uint32) (err error) { 
    return instance.SetProperty("HeapFlags", (value))
}

// GetHeapFlags gets the value of HeapFlags for the instance
func (instance *HeapCreate_V2) GetPropertyHeapFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HeapFlags")
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

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapCreate_V2) SetPropertyHeapHandle(value uint32) (err error) { 
    return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapCreate_V2) GetPropertyHeapHandle()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HeapHandle")
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

