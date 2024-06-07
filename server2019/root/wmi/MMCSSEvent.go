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

// MMCSSEvent struct
type MMCSSEvent struct { 
	*MMCSSTrace

	// 
	ScheduledPID uint32

	// 
	ScheduledTID uint32

	// 
	SchedulingPriority uint32

	// 
	TaskIndex uint32
}

	func NewMMCSSEventEx1(instance *cim.WmiInstance) (newInstance *MMCSSEvent, err error) {tmp, err := NewMMCSSTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &MMCSSEvent {
	MMCSSTrace: tmp,
	}
	return
	}
	

	func NewMMCSSEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MMCSSEvent, err error) {tmp, err := NewMMCSSTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MMCSSEvent {
	MMCSSTrace: tmp,
	}
	return
	}
	

// SetScheduledPID sets the value of ScheduledPID for the instance
func (instance *MMCSSEvent) SetPropertyScheduledPID(value uint32) (err error) { 
    return instance.SetProperty("ScheduledPID", (value))
}

// GetScheduledPID gets the value of ScheduledPID for the instance
func (instance *MMCSSEvent) GetPropertyScheduledPID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ScheduledPID")
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

// SetScheduledTID sets the value of ScheduledTID for the instance
func (instance *MMCSSEvent) SetPropertyScheduledTID(value uint32) (err error) { 
    return instance.SetProperty("ScheduledTID", (value))
}

// GetScheduledTID gets the value of ScheduledTID for the instance
func (instance *MMCSSEvent) GetPropertyScheduledTID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ScheduledTID")
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

// SetSchedulingPriority sets the value of SchedulingPriority for the instance
func (instance *MMCSSEvent) SetPropertySchedulingPriority(value uint32) (err error) { 
    return instance.SetProperty("SchedulingPriority", (value))
}

// GetSchedulingPriority gets the value of SchedulingPriority for the instance
func (instance *MMCSSEvent) GetPropertySchedulingPriority()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SchedulingPriority")
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

// SetTaskIndex sets the value of TaskIndex for the instance
func (instance *MMCSSEvent) SetPropertyTaskIndex(value uint32) (err error) { 
    return instance.SetProperty("TaskIndex", (value))
}

// GetTaskIndex gets the value of TaskIndex for the instance
func (instance *MMCSSEvent) GetPropertyTaskIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TaskIndex")
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

