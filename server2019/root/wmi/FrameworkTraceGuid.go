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

// FrameworkTraceGuid struct
type FrameworkTraceGuid struct { 
	*EventTrace

	// 
	Flags uint32

	// 
	Level uint32
}

	func NewFrameworkTraceGuidEx1(instance *cim.WmiInstance) (newInstance *FrameworkTraceGuid, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &FrameworkTraceGuid {
	EventTrace: tmp,
	}
	return
	}
	

	func NewFrameworkTraceGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FrameworkTraceGuid, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FrameworkTraceGuid {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *FrameworkTraceGuid) SetPropertyFlags(value uint32) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *FrameworkTraceGuid) GetPropertyFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Flags")
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

// SetLevel sets the value of Level for the instance
func (instance *FrameworkTraceGuid) SetPropertyLevel(value uint32) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *FrameworkTraceGuid) GetPropertyLevel()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Level")
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

