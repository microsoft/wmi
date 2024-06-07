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

// DriverCompleteRequest struct
type DriverCompleteRequest struct { 
	*DiskIo_V2

	// 
	Irp uint32

	// 
	RoutineAddr uint32

	// 
	UniqMatchId uint32
}

	func NewDriverCompleteRequestEx1(instance *cim.WmiInstance) (newInstance *DriverCompleteRequest, err error) {tmp, err := NewDiskIo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &DriverCompleteRequest {
	DiskIo_V2: tmp,
	}
	return
	}
	

	func NewDriverCompleteRequestEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DriverCompleteRequest, err error) {tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DriverCompleteRequest {
	DiskIo_V2: tmp,
	}
	return
	}
	

// SetIrp sets the value of Irp for the instance
func (instance *DriverCompleteRequest) SetPropertyIrp(value uint32) (err error) { 
    return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DriverCompleteRequest) GetPropertyIrp()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Irp")
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

// SetRoutineAddr sets the value of RoutineAddr for the instance
func (instance *DriverCompleteRequest) SetPropertyRoutineAddr(value uint32) (err error) { 
    return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *DriverCompleteRequest) GetPropertyRoutineAddr()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoutineAddr")
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

// SetUniqMatchId sets the value of UniqMatchId for the instance
func (instance *DriverCompleteRequest) SetPropertyUniqMatchId(value uint32) (err error) { 
    return instance.SetProperty("UniqMatchId", (value))
}

// GetUniqMatchId gets the value of UniqMatchId for the instance
func (instance *DriverCompleteRequest) GetPropertyUniqMatchId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UniqMatchId")
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

