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

// HeapAffinitySlotAssigned struct
type HeapAffinitySlotAssigned struct { 
	*HeapTrace_V2

	// 
	SlotIndex uint32
}

	func NewHeapAffinitySlotAssignedEx1(instance *cim.WmiInstance) (newInstance *HeapAffinitySlotAssigned, err error) {tmp, err := NewHeapTrace_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &HeapAffinitySlotAssigned {
	HeapTrace_V2: tmp,
	}
	return
	}
	

	func NewHeapAffinitySlotAssignedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapAffinitySlotAssigned, err error) {tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapAffinitySlotAssigned {
	HeapTrace_V2: tmp,
	}
	return
	}
	

// SetSlotIndex sets the value of SlotIndex for the instance
func (instance *HeapAffinitySlotAssigned) SetPropertySlotIndex(value uint32) (err error) { 
    return instance.SetProperty("SlotIndex", (value))
}

// GetSlotIndex gets the value of SlotIndex for the instance
func (instance *HeapAffinitySlotAssigned) GetPropertySlotIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SlotIndex")
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

