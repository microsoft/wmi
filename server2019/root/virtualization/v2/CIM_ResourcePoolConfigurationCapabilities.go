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

// CIM_ResourcePoolConfigurationCapabilities struct
type CIM_ResourcePoolConfigurationCapabilities struct { 
	*CIM_Capabilities

	// This property reflects the methods of the associated service class that are supported that may return a Job.
	AsynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported

	// This property reflects the methods of the associated service class that are supported andblock until completed (no Job is returned.)
	SynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported
}

	func NewCIM_ResourcePoolConfigurationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_ResourcePoolConfigurationCapabilities, err error) {tmp, err := NewCIM_CapabilitiesEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ResourcePoolConfigurationCapabilities {
	CIM_Capabilities: tmp,
	}
	return
	}
	

	func NewCIM_ResourcePoolConfigurationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ResourcePoolConfigurationCapabilities, err error) {tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ResourcePoolConfigurationCapabilities {
	CIM_Capabilities: tmp,
	}
	return
	}
	

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) SetPropertyAsynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported) (err error) { 
    return instance.SetProperty("AsynchronousMethodsSupported", (value))
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) GetPropertyAsynchronousMethodsSupported()(value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported, err error) { 
    retValue, err := instance.GetProperty("AsynchronousMethodsSupported")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported(valuetmp))
    }

    return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) SetPropertySynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported) (err error) { 
    return instance.SetProperty("SynchronousMethodsSupported", (value))
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) GetPropertySynchronousMethodsSupported()(value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported, err error) { 
    retValue, err := instance.GetProperty("SynchronousMethodsSupported")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported(valuetmp))
    }

    return
}

