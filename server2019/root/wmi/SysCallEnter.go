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

// SysCallEnter struct
type SysCallEnter struct { 
	*PerfInfo_V2

	// 
	SysCallAddress uint32
}

	func NewSysCallEnterEx1(instance *cim.WmiInstance) (newInstance *SysCallEnter, err error) {tmp, err := NewPerfInfo_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SysCallEnter {
	PerfInfo_V2: tmp,
	}
	return
	}
	

	func NewSysCallEnterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SysCallEnter, err error) {tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SysCallEnter {
	PerfInfo_V2: tmp,
	}
	return
	}
	

// SetSysCallAddress sets the value of SysCallAddress for the instance
func (instance *SysCallEnter) SetPropertySysCallAddress(value uint32) (err error) { 
    return instance.SetProperty("SysCallAddress", (value))
}

// GetSysCallAddress gets the value of SysCallAddress for the instance
func (instance *SysCallEnter) GetPropertySysCallAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SysCallAddress")
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

