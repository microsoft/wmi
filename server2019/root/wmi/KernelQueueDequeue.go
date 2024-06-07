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

// KernelQueueDequeue struct
type KernelQueueDequeue struct { 
	*Thread_V2

	// 
	Entries []uint32

	// 
	EntryCount uint32

	// 
	ThreadId uint32
}

	func NewKernelQueueDequeueEx1(instance *cim.WmiInstance) (newInstance *KernelQueueDequeue, err error) {tmp, err := NewThread_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &KernelQueueDequeue {
	Thread_V2: tmp,
	}
	return
	}
	

	func NewKernelQueueDequeueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *KernelQueueDequeue, err error) {tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &KernelQueueDequeue {
	Thread_V2: tmp,
	}
	return
	}
	

// SetEntries sets the value of Entries for the instance
func (instance *KernelQueueDequeue) SetPropertyEntries(value []uint32) (err error) { 
    return instance.SetProperty("Entries", (value))
}

// GetEntries gets the value of Entries for the instance
func (instance *KernelQueueDequeue) GetPropertyEntries()(value []uint32, err error) { 
    retValue, err := instance.GetProperty("Entries")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint32(valuetmp))
    }

    return
}

// SetEntryCount sets the value of EntryCount for the instance
func (instance *KernelQueueDequeue) SetPropertyEntryCount(value uint32) (err error) { 
    return instance.SetProperty("EntryCount", (value))
}

// GetEntryCount gets the value of EntryCount for the instance
func (instance *KernelQueueDequeue) GetPropertyEntryCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EntryCount")
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
func (instance *KernelQueueDequeue) SetPropertyThreadId(value uint32) (err error) { 
    return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *KernelQueueDequeue) GetPropertyThreadId()(value uint32, err error) { 
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

