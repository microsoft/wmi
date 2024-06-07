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

// CtlGuid struct
type CtlGuid struct { 
	*EventTrace

	// Enable Flags
	Flags CtlGuid_Flags

	// Levels
	Level CtlGuid_Level
}

	func NewCtlGuidEx1(instance *cim.WmiInstance) (newInstance *CtlGuid, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &CtlGuid {
	EventTrace: tmp,
	}
	return
	}
	

	func NewCtlGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CtlGuid, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CtlGuid {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *CtlGuid) SetPropertyFlags(value CtlGuid_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *CtlGuid) GetPropertyFlags()(value CtlGuid_Flags, err error) { 
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

    value = CtlGuid_Flags(valuetmp)

    return
}

// SetLevel sets the value of Level for the instance
func (instance *CtlGuid) SetPropertyLevel(value CtlGuid_Level) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *CtlGuid) GetPropertyLevel()(value CtlGuid_Level, err error) { 
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

    value = CtlGuid_Level(valuetmp)

    return
}

