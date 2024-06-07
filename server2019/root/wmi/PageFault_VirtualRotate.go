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

// PageFault_VirtualRotate struct
type PageFault_VirtualRotate struct { 
	*PageFault_V2

	// 
	BaseAddress uint32

	// 
	SizeInBytes interface{}
}

	func NewPageFault_VirtualRotateEx1(instance *cim.WmiInstance) (newInstance *PageFault_VirtualRotate, err error) {tmp, err := NewPageFault_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &PageFault_VirtualRotate {
	PageFault_V2: tmp,
	}
	return
	}
	

	func NewPageFault_VirtualRotateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PageFault_VirtualRotate, err error) {tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PageFault_VirtualRotate {
	PageFault_V2: tmp,
	}
	return
	}
	

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *PageFault_VirtualRotate) SetPropertyBaseAddress(value uint32) (err error) { 
    return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *PageFault_VirtualRotate) GetPropertyBaseAddress()(value uint32, err error) { 
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

// SetSizeInBytes sets the value of SizeInBytes for the instance
func (instance *PageFault_VirtualRotate) SetPropertySizeInBytes(value interface{}) (err error) { 
    return instance.SetProperty("SizeInBytes", (value))
}

// GetSizeInBytes gets the value of SizeInBytes for the instance
func (instance *PageFault_VirtualRotate) GetPropertySizeInBytes()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("SizeInBytes")
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

