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

// AuthorizationSection struct
type AuthorizationSection struct { 
	*ConfigurationSectionWithCollection

	// 
	Authorization []AuthorizationRule

	// 
	BypassLoginPages bool
}

	func NewAuthorizationSectionEx1(instance *cim.WmiInstance) (newInstance *AuthorizationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &AuthorizationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewAuthorizationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AuthorizationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AuthorizationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAuthorization sets the value of Authorization for the instance
func (instance *AuthorizationSection) SetPropertyAuthorization(value []AuthorizationRule) (err error) { 
    return instance.SetProperty("Authorization", (value))
}

// GetAuthorization gets the value of Authorization for the instance
func (instance *AuthorizationSection) GetPropertyAuthorization()(value []AuthorizationRule, err error) { 
    retValue, err := instance.GetProperty("Authorization")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(AuthorizationRule); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " AuthorizationRule is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, AuthorizationRule(valuetmp))
    }

    return
}

// SetBypassLoginPages sets the value of BypassLoginPages for the instance
func (instance *AuthorizationSection) SetPropertyBypassLoginPages(value bool) (err error) { 
    return instance.SetProperty("BypassLoginPages", (value))
}

// GetBypassLoginPages gets the value of BypassLoginPages for the instance
func (instance *AuthorizationSection) GetPropertyBypassLoginPages()(value bool, err error) { 
    retValue, err := instance.GetProperty("BypassLoginPages")
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

