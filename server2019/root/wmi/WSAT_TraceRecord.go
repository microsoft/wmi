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

// WSAT_TraceRecord struct
type WSAT_TraceRecord struct { 
	*WSAT_TraceEvent

	// Activity ID
	ActivityID interface{}

	// EventID
	EventID int32

	// Trace Record
	TraceRecord string
}

	func NewWSAT_TraceRecordEx1(instance *cim.WmiInstance) (newInstance *WSAT_TraceRecord, err error) {tmp, err := NewWSAT_TraceEventEx1(instance)
		
	if err != nil { return }
	newInstance = &WSAT_TraceRecord {
	WSAT_TraceEvent: tmp,
	}
	return
	}
	

	func NewWSAT_TraceRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WSAT_TraceRecord, err error) {tmp, err := NewWSAT_TraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WSAT_TraceRecord {
	WSAT_TraceEvent: tmp,
	}
	return
	}
	

// SetActivityID sets the value of ActivityID for the instance
func (instance *WSAT_TraceRecord) SetPropertyActivityID(value interface{}) (err error) { 
    return instance.SetProperty("ActivityID", (value))
}

// GetActivityID gets the value of ActivityID for the instance
func (instance *WSAT_TraceRecord) GetPropertyActivityID()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("ActivityID")
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

// SetEventID sets the value of EventID for the instance
func (instance *WSAT_TraceRecord) SetPropertyEventID(value int32) (err error) { 
    return instance.SetProperty("EventID", (value))
}

// GetEventID gets the value of EventID for the instance
func (instance *WSAT_TraceRecord) GetPropertyEventID()(value int32, err error) { 
    retValue, err := instance.GetProperty("EventID")
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

    value = int32(valuetmp)

    return
}

// SetTraceRecord sets the value of TraceRecord for the instance
func (instance *WSAT_TraceRecord) SetPropertyTraceRecord(value string) (err error) { 
    return instance.SetProperty("TraceRecord", (value))
}

// GetTraceRecord gets the value of TraceRecord for the instance
func (instance *WSAT_TraceRecord) GetPropertyTraceRecord()(value string, err error) { 
    retValue, err := instance.GetProperty("TraceRecord")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

