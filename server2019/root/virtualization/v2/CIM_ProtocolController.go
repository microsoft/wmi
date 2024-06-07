// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_ProtocolController struct
type CIM_ProtocolController struct { 
	*CIM_LogicalDevice

	// Maximum number of Units that can be controlled by or accessed through this ProtocolController.
	MaxUnitsControlled uint32
}

	func NewCIM_ProtocolControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_ProtocolController, err error) {tmp, err := NewCIM_LogicalDeviceEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ProtocolController {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

	func NewCIM_ProtocolControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ProtocolController, err error) {tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ProtocolController {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

// SetMaxUnitsControlled sets the value of MaxUnitsControlled for the instance
func (instance *CIM_ProtocolController) SetPropertyMaxUnitsControlled(value uint32) (err error) { 
    return instance.SetProperty("MaxUnitsControlled", (value))
}

// GetMaxUnitsControlled gets the value of MaxUnitsControlled for the instance
func (instance *CIM_ProtocolController) GetPropertyMaxUnitsControlled()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxUnitsControlled")
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

