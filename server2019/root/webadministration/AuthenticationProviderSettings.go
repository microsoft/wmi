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

// AuthenticationProviderSettings struct
type AuthenticationProviderSettings struct { 
	*EmbeddedObject

	// 
	Providers []StringElement
}

	func NewAuthenticationProviderSettingsEx1(instance *cim.WmiInstance) (newInstance *AuthenticationProviderSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &AuthenticationProviderSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewAuthenticationProviderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AuthenticationProviderSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AuthenticationProviderSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetProviders sets the value of Providers for the instance
func (instance *AuthenticationProviderSettings) SetPropertyProviders(value []StringElement) (err error) { 
    return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *AuthenticationProviderSettings) GetPropertyProviders()(value []StringElement, err error) { 
    retValue, err := instance.GetProperty("Providers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(StringElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " StringElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, StringElement(valuetmp))
    }

    return
}

