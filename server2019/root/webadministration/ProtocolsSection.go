// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ProtocolsSection struct
type ProtocolsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	Protocols []ProtocolElement
}

	func NewProtocolsSectionEx1(instance *cim.WmiInstance) (newInstance *ProtocolsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ProtocolsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewProtocolsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProtocolsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProtocolsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetProtocols sets the value of Protocols for the instance
func (instance *ProtocolsSection) SetPropertyProtocols(value []ProtocolElement) (err error) { 
    return instance.SetProperty("Protocols", (value))
}

// GetProtocols gets the value of Protocols for the instance
func (instance *ProtocolsSection) GetPropertyProtocols()(value []ProtocolElement, err error) { 
    retValue, err := instance.GetProperty("Protocols")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ProtocolElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ProtocolElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ProtocolElement(valuetmp))
    }

    return
}

