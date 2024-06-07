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

// Msvm_PciExpress struct
type Msvm_PciExpress struct { 
	*CIM_LogicalDevice

	// 
	DeviceInstancePath string

	// 
	FunctionNumber uint16

	// 
	LocationPath string
}

	func NewMsvm_PciExpressEx1(instance *cim.WmiInstance) (newInstance *Msvm_PciExpress, err error) {tmp, err := NewCIM_LogicalDeviceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_PciExpress {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

	func NewMsvm_PciExpressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_PciExpress, err error) {tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_PciExpress {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

// SetDeviceInstancePath sets the value of DeviceInstancePath for the instance
func (instance *Msvm_PciExpress) SetPropertyDeviceInstancePath(value string) (err error) { 
    return instance.SetProperty("DeviceInstancePath", (value))
}

// GetDeviceInstancePath gets the value of DeviceInstancePath for the instance
func (instance *Msvm_PciExpress) GetPropertyDeviceInstancePath()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceInstancePath")
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

// SetFunctionNumber sets the value of FunctionNumber for the instance
func (instance *Msvm_PciExpress) SetPropertyFunctionNumber(value uint16) (err error) { 
    return instance.SetProperty("FunctionNumber", (value))
}

// GetFunctionNumber gets the value of FunctionNumber for the instance
func (instance *Msvm_PciExpress) GetPropertyFunctionNumber()(value uint16, err error) { 
    retValue, err := instance.GetProperty("FunctionNumber")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetLocationPath sets the value of LocationPath for the instance
func (instance *Msvm_PciExpress) SetPropertyLocationPath(value string) (err error) { 
    return instance.SetProperty("LocationPath", (value))
}

// GetLocationPath gets the value of LocationPath for the instance
func (instance *Msvm_PciExpress) GetPropertyLocationPath()(value string, err error) { 
    retValue, err := instance.GetProperty("LocationPath")
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

