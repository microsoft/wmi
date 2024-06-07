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

// ObObjectEvent struct
type ObObjectEvent struct { 
	*ObTrace

	// 
	Object uint32

	// 
	ObjectType uint16
}

	func NewObObjectEventEx1(instance *cim.WmiInstance) (newInstance *ObObjectEvent, err error) {tmp, err := NewObTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &ObObjectEvent {
	ObTrace: tmp,
	}
	return
	}
	

	func NewObObjectEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ObObjectEvent, err error) {tmp, err := NewObTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ObObjectEvent {
	ObTrace: tmp,
	}
	return
	}
	

// SetObject sets the value of Object for the instance
func (instance *ObObjectEvent) SetPropertyObject(value uint32) (err error) { 
    return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *ObObjectEvent) GetPropertyObject()(value uint32, err error) { 
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

// SetObjectType sets the value of ObjectType for the instance
func (instance *ObObjectEvent) SetPropertyObjectType(value uint16) (err error) { 
    return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *ObObjectEvent) GetPropertyObjectType()(value uint16, err error) { 
    retValue, err := instance.GetProperty("ObjectType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

