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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_VirtualMachineToDisks struct
type Msvm_VirtualMachineToDisks struct { 
	*cim.WmiInstance

	// 
	DisksToExport []string

	// 
	VirtualMachineId string
}

	func NewMsvm_VirtualMachineToDisksEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualMachineToDisks, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &Msvm_VirtualMachineToDisks {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMsvm_VirtualMachineToDisksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_VirtualMachineToDisks, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_VirtualMachineToDisks {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDisksToExport sets the value of DisksToExport for the instance
func (instance *Msvm_VirtualMachineToDisks) SetPropertyDisksToExport(value []string) (err error) { 
    return instance.SetProperty("DisksToExport", (value))
}

// GetDisksToExport gets the value of DisksToExport for the instance
func (instance *Msvm_VirtualMachineToDisks) GetPropertyDisksToExport()(value []string, err error) { 
    retValue, err := instance.GetProperty("DisksToExport")
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

// SetVirtualMachineId sets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualMachineToDisks) SetPropertyVirtualMachineId(value string) (err error) { 
    return instance.SetProperty("VirtualMachineId", (value))
}

// GetVirtualMachineId gets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualMachineToDisks) GetPropertyVirtualMachineId()(value string, err error) { 
    retValue, err := instance.GetProperty("VirtualMachineId")
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

