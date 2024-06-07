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

// Msvm_FlexIoDevice struct
type Msvm_FlexIoDevice struct { 
	*CIM_LogicalDevice

	// 
	EmulatorConfiguration []string

	// 
	EmulatorId string
}

	func NewMsvm_FlexIoDeviceEx1(instance *cim.WmiInstance) (newInstance *Msvm_FlexIoDevice, err error) {tmp, err := NewCIM_LogicalDeviceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_FlexIoDevice {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

	func NewMsvm_FlexIoDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_FlexIoDevice, err error) {tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_FlexIoDevice {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

// SetEmulatorConfiguration sets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDevice) SetPropertyEmulatorConfiguration(value []string) (err error) { 
    return instance.SetProperty("EmulatorConfiguration", (value))
}

// GetEmulatorConfiguration gets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDevice) GetPropertyEmulatorConfiguration()(value []string, err error) { 
    retValue, err := instance.GetProperty("EmulatorConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetEmulatorId sets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDevice) SetPropertyEmulatorId(value string) (err error) { 
    return instance.SetProperty("EmulatorId", (value))
}

// GetEmulatorId gets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDevice) GetPropertyEmulatorId()(value string, err error) { 
    retValue, err := instance.GetProperty("EmulatorId")
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

