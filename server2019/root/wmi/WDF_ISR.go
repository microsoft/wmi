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

// WDF_ISR struct
type WDF_ISR struct { 
	*PerfInfo_V2

	// 
	Routine uint32
}

	func NewWDF_ISREx1(instance *cim.WmiInstance) (newInstance *WDF_ISR, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &WDF_ISR {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewWDF_ISREx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WDF_ISR, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WDF_ISR {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetRoutine sets the value of Routine for the instance
func (instance *WDF_ISR) SetPropertyRoutine(value uint32) (err error) { 
    return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *WDF_ISR) GetPropertyRoutine()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Routine")
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

