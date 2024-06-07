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

// HeapAffinityManagerEnable struct
type HeapAffinityManagerEnable struct { 
	*HeapTrace_V2

	// 
	BucketIndex uint32

	// 
	HeapHandle uint32
}

	func NewHeapAffinityManagerEnableEx1(instance *cim.WmiInstance) (newInstance *HeapAffinityManagerEnable, err error) {tmp, err := NewHeapTrace_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &HeapAffinityManagerEnable {
	HeapTrace_V2: tmp,
	}
	return
	}
	

	func NewHeapAffinityManagerEnableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapAffinityManagerEnable, err error) {tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapAffinityManagerEnable {
	HeapTrace_V2: tmp,
	}
	return
	}
	

// SetBucketIndex sets the value of BucketIndex for the instance
func (instance *HeapAffinityManagerEnable) SetPropertyBucketIndex(value uint32) (err error) { 
    return instance.SetProperty("BucketIndex", (value))
}

// GetBucketIndex gets the value of BucketIndex for the instance
func (instance *HeapAffinityManagerEnable) GetPropertyBucketIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BucketIndex")
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
func (instance *HeapAffinityManagerEnable) SetPropertyHeapHandle(value uint32) (err error) { 
    return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapAffinityManagerEnable) GetPropertyHeapHandle()(value uint32, err error) { 
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

