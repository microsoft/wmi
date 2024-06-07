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

// ProviderSettings struct
type ProviderSettings struct { 
	*EmbeddedObject

	// 
	Providers []ProviderElement
}

	func NewProviderSettingsEx1(instance *cim.WmiInstance) (newInstance *ProviderSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ProviderSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewProviderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProviderSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProviderSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetProviders sets the value of Providers for the instance
func (instance *ProviderSettings) SetPropertyProviders(value []ProviderElement) (err error) { 
    return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *ProviderSettings) GetPropertyProviders()(value []ProviderElement, err error) { 
    retValue, err := instance.GetProperty("Providers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ProviderElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ProviderElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ProviderElement(valuetmp))
    }

    return
}

