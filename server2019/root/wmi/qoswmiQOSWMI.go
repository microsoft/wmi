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

// qoswmiQOSWMI struct
type qoswmiQOSWMI struct { 
	*EventTrace

	// 
	Flags qoswmiQOSWMI_Flags

	// 
	Level qoswmiQOSWMI_Level
}

	func NewqoswmiQOSWMIEx1(instance *cim.WmiInstance) (newInstance *qoswmiQOSWMI, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &qoswmiQOSWMI {
	EventTrace: tmp,
	}
	return
	}
	

	func NewqoswmiQOSWMIEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *qoswmiQOSWMI, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &qoswmiQOSWMI {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *qoswmiQOSWMI) SetPropertyFlags(value qoswmiQOSWMI_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *qoswmiQOSWMI) GetPropertyFlags()(value qoswmiQOSWMI_Flags, err error) { 
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

    value = qoswmiQOSWMI_Flags(valuetmp)

    return
}

// SetLevel sets the value of Level for the instance
func (instance *qoswmiQOSWMI) SetPropertyLevel(value qoswmiQOSWMI_Level) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *qoswmiQOSWMI) GetPropertyLevel()(value qoswmiQOSWMI_Level, err error) { 
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

    value = qoswmiQOSWMI_Level(valuetmp)

    return
}

