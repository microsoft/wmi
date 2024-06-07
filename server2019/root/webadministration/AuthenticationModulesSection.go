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

// AuthenticationModulesSection struct
type AuthenticationModulesSection struct { 
	*ConfigurationSectionWithCollection

	// 
	AuthenticationModules []AuthenticationModuleElement
}

	func NewAuthenticationModulesSectionEx1(instance *cim.WmiInstance) (newInstance *AuthenticationModulesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &AuthenticationModulesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewAuthenticationModulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AuthenticationModulesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AuthenticationModulesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAuthenticationModules sets the value of AuthenticationModules for the instance
func (instance *AuthenticationModulesSection) SetPropertyAuthenticationModules(value []AuthenticationModuleElement) (err error) { 
    return instance.SetProperty("AuthenticationModules", (value))
}

// GetAuthenticationModules gets the value of AuthenticationModules for the instance
func (instance *AuthenticationModulesSection) GetPropertyAuthenticationModules()(value []AuthenticationModuleElement, err error) { 
    retValue, err := instance.GetProperty("AuthenticationModules")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(AuthenticationModuleElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " AuthenticationModuleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, AuthenticationModuleElement(valuetmp))
    }

    return
}

