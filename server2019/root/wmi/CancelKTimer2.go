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

// CancelKTimer2 struct
type CancelKTimer2 struct { 
	*PerfInfo_V2

	// 
	Timer uint32
}

	func NewCancelKTimer2Ex1(instance *cim.WmiInstance) (newInstance *CancelKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &CancelKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewCancelKTimer2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CancelKTimer2, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CancelKTimer2 {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetTimer sets the value of Timer for the instance
func (instance *CancelKTimer2) SetPropertyTimer(value uint32) (err error) { 
    return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *CancelKTimer2) GetPropertyTimer()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Timer")
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

