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

// DPC_V1 struct
type DPC_V1 struct { 
	*PerfInfo_V1

	// 
	InitialTime interface{}

	// 
	Routine uint32
}

	func NewDPC_V1Ex1(instance *cim.WmiInstance) (newInstance *DPC_V1, err error) {tmp, err := NewPerfInfo_V1Ex1(instance)
		
	if err != nil { return }
	newInstance = &DPC_V1 {
	PerfInfo_V1: tmp,
	}
	return
	}
	

	func NewDPC_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DPC_V1, err error) {tmp, err := NewPerfInfo_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DPC_V1 {
	PerfInfo_V1: tmp,
	}
	return
	}
	

// SetInitialTime sets the value of InitialTime for the instance
func (instance *DPC_V1) SetPropertyInitialTime(value interface{}) (err error) { 
    return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *DPC_V1) GetPropertyInitialTime()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("InitialTime")
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

// SetRoutine sets the value of Routine for the instance
func (instance *DPC_V1) SetPropertyRoutine(value uint32) (err error) { 
    return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *DPC_V1) GetPropertyRoutine()(value uint32, err error) { 
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

