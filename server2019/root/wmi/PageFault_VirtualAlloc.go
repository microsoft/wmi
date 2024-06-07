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

// PageFault_VirtualAlloc struct
type PageFault_VirtualAlloc struct { 
	*PageFault_V2

	// 
	BaseAddress uint32

	// 
	ProcessId uint32

	// 
	RegionSize interface{}
}

	func NewPageFault_VirtualAllocEx1(instance *cim.WmiInstance) (newInstance *PageFault_VirtualAlloc, err error) {tmp, err := NewPageFault_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &PageFault_VirtualAlloc {
	PageFault_V2: tmp,
	}
	return
	}
	

	func NewPageFault_VirtualAllocEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PageFault_VirtualAlloc, err error) {tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PageFault_VirtualAlloc {
	PageFault_V2: tmp,
	}
	return
	}
	

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *PageFault_VirtualAlloc) SetPropertyBaseAddress(value uint32) (err error) { 
    return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *PageFault_VirtualAlloc) GetPropertyBaseAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BaseAddress")
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
func (instance *PageFault_VirtualAlloc) SetPropertyProcessId(value uint32) (err error) { 
    return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *PageFault_VirtualAlloc) GetPropertyProcessId()(value uint32, err error) { 
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

// SetRegionSize sets the value of RegionSize for the instance
func (instance *PageFault_VirtualAlloc) SetPropertyRegionSize(value interface{}) (err error) { 
    return instance.SetProperty("RegionSize", (value))
}

// GetRegionSize gets the value of RegionSize for the instance
func (instance *PageFault_VirtualAlloc) GetPropertyRegionSize()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("RegionSize")
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

