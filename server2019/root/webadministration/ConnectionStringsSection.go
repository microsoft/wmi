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

// ConnectionStringsSection struct
type ConnectionStringsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ConnectionStrings []ConnectionStringSettings
}

	func NewConnectionStringsSectionEx1(instance *cim.WmiInstance) (newInstance *ConnectionStringsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ConnectionStringsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewConnectionStringsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConnectionStringsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConnectionStringsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetConnectionStrings sets the value of ConnectionStrings for the instance
func (instance *ConnectionStringsSection) SetPropertyConnectionStrings(value []ConnectionStringSettings) (err error) { 
    return instance.SetProperty("ConnectionStrings", (value))
}

// GetConnectionStrings gets the value of ConnectionStrings for the instance
func (instance *ConnectionStringsSection) GetPropertyConnectionStrings()(value []ConnectionStringSettings, err error) { 
    retValue, err := instance.GetProperty("ConnectionStrings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ConnectionStringSettings); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ConnectionStringSettings is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ConnectionStringSettings(valuetmp))
    }

    return
}

