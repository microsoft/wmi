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

// UMBus struct
type UMBus struct { 
	*EventTrace

	// Enable Flags
	Flags UMBus_Flags

	// Levels
	Level UMBus_Level
}

	func NewUMBusEx1(instance *cim.WmiInstance) (newInstance *UMBus, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &UMBus {
	EventTrace: tmp,
	}
	return
	}
	

	func NewUMBusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UMBus, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UMBus {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *UMBus) SetPropertyFlags(value UMBus_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *UMBus) GetPropertyFlags()(value UMBus_Flags, err error) { 
    retValue, err := instance.GetProperty("Flags")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = UMBus_Flags(valuetmp)

    return
}

// SetLevel sets the value of Level for the instance
func (instance *UMBus) SetPropertyLevel(value UMBus_Level) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *UMBus) GetPropertyLevel()(value UMBus_Level, err error) { 
    retValue, err := instance.GetProperty("Level")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = UMBus_Level(valuetmp)

    return
}

