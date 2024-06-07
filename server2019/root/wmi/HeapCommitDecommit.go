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

// HeapCommitDecommit struct
type HeapCommitDecommit struct { 
	*HeapTrace_V2

	// 
	Block uint32

	// 
	Caller uint32

	// 
	HeapHandle uint32

	// 
	Size interface{}
}

	func NewHeapCommitDecommitEx1(instance *cim.WmiInstance) (newInstance *HeapCommitDecommit, err error) {tmp, err := NewHeapTrace_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &HeapCommitDecommit {
	HeapTrace_V2: tmp,
	}
	return
	}
	

	func NewHeapCommitDecommitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapCommitDecommit, err error) {tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapCommitDecommit {
	HeapTrace_V2: tmp,
	}
	return
	}
	

// SetBlock sets the value of Block for the instance
func (instance *HeapCommitDecommit) SetPropertyBlock(value uint32) (err error) { 
    return instance.SetProperty("Block", (value))
}

// GetBlock gets the value of Block for the instance
func (instance *HeapCommitDecommit) GetPropertyBlock()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Block")
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

// SetCaller sets the value of Caller for the instance
func (instance *HeapCommitDecommit) SetPropertyCaller(value uint32) (err error) { 
    return instance.SetProperty("Caller", (value))
}

// GetCaller gets the value of Caller for the instance
func (instance *HeapCommitDecommit) GetPropertyCaller()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Caller")
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
func (instance *HeapCommitDecommit) SetPropertyHeapHandle(value uint32) (err error) { 
    return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapCommitDecommit) GetPropertyHeapHandle()(value uint32, err error) { 
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

// SetSize sets the value of Size for the instance
func (instance *HeapCommitDecommit) SetPropertySize(value interface{}) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *HeapCommitDecommit) GetPropertySize()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("Size")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(interface{}); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = interface{}(valuetmp)

    return
}

