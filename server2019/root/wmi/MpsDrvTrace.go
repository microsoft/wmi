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

// MpsDrvTrace struct
type MpsDrvTrace struct { 
	*EventTrace

	// Enable Flags
	Flags MpsDrvTrace_Flags
}

	func NewMpsDrvTraceEx1(instance *cim.WmiInstance) (newInstance *MpsDrvTrace, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &MpsDrvTrace {
	EventTrace: tmp,
	}
	return
	}
	

	func NewMpsDrvTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MpsDrvTrace, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MpsDrvTrace {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *MpsDrvTrace) SetPropertyFlags(value MpsDrvTrace_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MpsDrvTrace) GetPropertyFlags()(value MpsDrvTrace_Flags, err error) { 
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

    value = MpsDrvTrace_Flags(valuetmp)

    return
}

