// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_UserDevice struct
type CIM_UserDevice struct { 
	*CIM_LogicalDevice

	// An indication of whether the Device is locked, preventing user input or output.
	IsLocked bool
}

	func NewCIM_UserDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_UserDevice, err error) {tmp, err := NewCIM_LogicalDeviceEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_UserDevice {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

	func NewCIM_UserDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_UserDevice, err error) {tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_UserDevice {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

// SetIsLocked sets the value of IsLocked for the instance
func (instance *CIM_UserDevice) SetPropertyIsLocked(value bool) (err error) { 
    return instance.SetProperty("IsLocked", (value))
}

// GetIsLocked gets the value of IsLocked for the instance
func (instance *CIM_UserDevice) GetPropertyIsLocked()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsLocked")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

