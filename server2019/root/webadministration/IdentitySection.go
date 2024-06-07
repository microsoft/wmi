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

// IdentitySection struct
type IdentitySection struct { 
	*ConfigurationSection

	// 
	Impersonate bool

	// 
	Password string

	// 
	UserName string
}

	func NewIdentitySectionEx1(instance *cim.WmiInstance) (newInstance *IdentitySection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &IdentitySection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewIdentitySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IdentitySection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IdentitySection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetImpersonate sets the value of Impersonate for the instance
func (instance *IdentitySection) SetPropertyImpersonate(value bool) (err error) { 
    return instance.SetProperty("Impersonate", (value))
}

// GetImpersonate gets the value of Impersonate for the instance
func (instance *IdentitySection) GetPropertyImpersonate()(value bool, err error) { 
    retValue, err := instance.GetProperty("Impersonate")
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

// SetPassword sets the value of Password for the instance
func (instance *IdentitySection) SetPropertyPassword(value string) (err error) { 
    return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *IdentitySection) GetPropertyPassword()(value string, err error) { 
    retValue, err := instance.GetProperty("Password")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetUserName sets the value of UserName for the instance
func (instance *IdentitySection) SetPropertyUserName(value string) (err error) { 
    return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *IdentitySection) GetPropertyUserName()(value string, err error) { 
    retValue, err := instance.GetProperty("UserName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

