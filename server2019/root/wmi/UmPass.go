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

// UmPass struct
type UmPass struct { 
	*EventTrace

	// Enable Flags
	Flags UmPass_Flags

	// Levels
	Level UmPass_Level
}

	func NewUmPassEx1(instance *cim.WmiInstance) (newInstance *UmPass, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &UmPass {
	EventTrace: tmp,
	}
	return
	}
	

	func NewUmPassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UmPass, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UmPass {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *UmPass) SetPropertyFlags(value UmPass_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *UmPass) GetPropertyFlags()(value UmPass_Flags, err error) { 
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

    value = UmPass_Flags(valuetmp)

    return
}

// SetLevel sets the value of Level for the instance
func (instance *UmPass) SetPropertyLevel(value UmPass_Level) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *UmPass) GetPropertyLevel()(value UmPass_Level, err error) { 
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

    value = UmPass_Level(valuetmp)

    return
}

