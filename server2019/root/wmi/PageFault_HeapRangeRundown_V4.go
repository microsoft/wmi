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

// PageFault_HeapRangeRundown_V4 struct
type PageFault_HeapRangeRundown_V4 struct { 
	*PageFault

	// 
	HeapHandle uint32

	// 
	HRFlags uint32

	// 
	HRHeapTag uint64

	// 
	HRPid uint32

	// 
	HRRangeCount uint32
}

	func NewPageFault_HeapRangeRundown_V4Ex1(instance *cim.WmiInstance) (newInstance *PageFault_HeapRangeRundown_V4, err error) {tmp, err := NewPageFaultEx1(instance)
		
	if err != nil { return }
	newInstance = &PageFault_HeapRangeRundown_V4 {
	PageFault: tmp,
	}
	return
	}
	

	func NewPageFault_HeapRangeRundown_V4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PageFault_HeapRangeRundown_V4, err error) {tmp, err := NewPageFaultEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PageFault_HeapRangeRundown_V4 {
	PageFault: tmp,
	}
	return
	}
	

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *PageFault_HeapRangeRundown_V4) SetPropertyHeapHandle(value uint32) (err error) { 
    return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *PageFault_HeapRangeRundown_V4) GetPropertyHeapHandle()(value uint32, err error) { 
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

// SetHRFlags sets the value of HRFlags for the instance
func (instance *PageFault_HeapRangeRundown_V4) SetPropertyHRFlags(value uint32) (err error) { 
    return instance.SetProperty("HRFlags", (value))
}

// GetHRFlags gets the value of HRFlags for the instance
func (instance *PageFault_HeapRangeRundown_V4) GetPropertyHRFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HRFlags")
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

// SetHRHeapTag sets the value of HRHeapTag for the instance
func (instance *PageFault_HeapRangeRundown_V4) SetPropertyHRHeapTag(value uint64) (err error) { 
    return instance.SetProperty("HRHeapTag", (value))
}

// GetHRHeapTag gets the value of HRHeapTag for the instance
func (instance *PageFault_HeapRangeRundown_V4) GetPropertyHRHeapTag()(value uint64, err error) { 
    retValue, err := instance.GetProperty("HRHeapTag")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetHRPid sets the value of HRPid for the instance
func (instance *PageFault_HeapRangeRundown_V4) SetPropertyHRPid(value uint32) (err error) { 
    return instance.SetProperty("HRPid", (value))
}

// GetHRPid gets the value of HRPid for the instance
func (instance *PageFault_HeapRangeRundown_V4) GetPropertyHRPid()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HRPid")
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

// SetHRRangeCount sets the value of HRRangeCount for the instance
func (instance *PageFault_HeapRangeRundown_V4) SetPropertyHRRangeCount(value uint32) (err error) { 
    return instance.SetProperty("HRRangeCount", (value))
}

// GetHRRangeCount gets the value of HRRangeCount for the instance
func (instance *PageFault_HeapRangeRundown_V4) GetPropertyHRRangeCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HRRangeCount")
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

