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

// Msvm_ExternalEthernetPortCapabilities struct
type Msvm_ExternalEthernetPortCapabilities struct { 
	*CIM_Capabilities

	// A boolean value which indicates whether IOV is supported by the network adapter.If the value is TRUE, then IOV is supported by the network adapter and IOVSupportReasons will be empty. Otherwise the IOVSupportReasons property will have the reasons why IOV cannot be supported.
	IOVSupport bool

	// An array of strings that indicates the possible reasons why IOV is not supported. If the value of IOVSupport is TRUE this array will be empty. 
	IOVSupportReasons []string
}

	func NewMsvm_ExternalEthernetPortCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_ExternalEthernetPortCapabilities, err error) {tmp, err := NewCIM_CapabilitiesEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ExternalEthernetPortCapabilities {
	CIM_Capabilities: tmp,
	}
	return
	}
	

	func NewMsvm_ExternalEthernetPortCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ExternalEthernetPortCapabilities, err error) {tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ExternalEthernetPortCapabilities {
	CIM_Capabilities: tmp,
	}
	return
	}
	

// SetIOVSupport sets the value of IOVSupport for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) SetPropertyIOVSupport(value bool) (err error) { 
    return instance.SetProperty("IOVSupport", (value))
}

// GetIOVSupport gets the value of IOVSupport for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) GetPropertyIOVSupport()(value bool, err error) { 
    retValue, err := instance.GetProperty("IOVSupport")
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

// SetIOVSupportReasons sets the value of IOVSupportReasons for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) SetPropertyIOVSupportReasons(value []string) (err error) { 
    return instance.SetProperty("IOVSupportReasons", (value))
}

// GetIOVSupportReasons gets the value of IOVSupportReasons for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) GetPropertyIOVSupportReasons()(value []string, err error) { 
    retValue, err := instance.GetProperty("IOVSupportReasons")
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
func  (instance* Msvm_ExternalEthernetPortCapabilities) GetRelatedExternalEthernetPort() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ExternalEthernetPort"); 
	}
	

