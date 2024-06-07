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

// Msvm_ResourceAllocationSettingData struct
type Msvm_ResourceAllocationSettingData struct { 
	*CIM_ResourceAllocationSettingData

	// 
	TargetVtl uint8

	// 
	VirtualSystemIdentifiers []string
}

	func NewMsvm_ResourceAllocationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourceAllocationSettingData, err error) {tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ResourceAllocationSettingData {
	CIM_ResourceAllocationSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_ResourceAllocationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ResourceAllocationSettingData, err error) {tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ResourceAllocationSettingData {
	CIM_ResourceAllocationSettingData: tmp,
	}
	return
	}
	

// SetTargetVtl sets the value of TargetVtl for the instance
func (instance *Msvm_ResourceAllocationSettingData) SetPropertyTargetVtl(value uint8) (err error) { 
    return instance.SetProperty("TargetVtl", (value))
}

// GetTargetVtl gets the value of TargetVtl for the instance
func (instance *Msvm_ResourceAllocationSettingData) GetPropertyTargetVtl()(value uint8, err error) { 
    retValue, err := instance.GetProperty("TargetVtl")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_ResourceAllocationSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) { 
    return instance.SetProperty("VirtualSystemIdentifiers", (value))
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_ResourceAllocationSettingData) GetPropertyVirtualSystemIdentifiers()(value []string, err error) { 
    retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
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
func  (instance* Msvm_ResourceAllocationSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_AllocationCapabilities"); 
	}
	

