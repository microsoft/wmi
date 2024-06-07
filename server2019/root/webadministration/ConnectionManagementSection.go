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

// ConnectionManagementSection struct
type ConnectionManagementSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ConnectionManagement []ConnectionManagementElement
}

	func NewConnectionManagementSectionEx1(instance *cim.WmiInstance) (newInstance *ConnectionManagementSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ConnectionManagementSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewConnectionManagementSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConnectionManagementSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConnectionManagementSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetConnectionManagement sets the value of ConnectionManagement for the instance
func (instance *ConnectionManagementSection) SetPropertyConnectionManagement(value []ConnectionManagementElement) (err error) { 
    return instance.SetProperty("ConnectionManagement", (value))
}

// GetConnectionManagement gets the value of ConnectionManagement for the instance
func (instance *ConnectionManagementSection) GetPropertyConnectionManagement()(value []ConnectionManagementElement, err error) { 
    retValue, err := instance.GetProperty("ConnectionManagement")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ConnectionManagementElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ConnectionManagementElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ConnectionManagementElement(valuetmp))
    }

    return
}

