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

// FormsAuthenticationUser struct
type FormsAuthenticationUser struct { 
	*CollectionElement

	// 
	Name string

	// 
	Password string
}

	func NewFormsAuthenticationUserEx1(instance *cim.WmiInstance) (newInstance *FormsAuthenticationUser, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationUser {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewFormsAuthenticationUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FormsAuthenticationUser, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationUser {
	CollectionElement: tmp,
	}
	return
	}
	

// SetName sets the value of Name for the instance
func (instance *FormsAuthenticationUser) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FormsAuthenticationUser) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetPassword sets the value of Password for the instance
func (instance *FormsAuthenticationUser) SetPropertyPassword(value string) (err error) { 
    return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *FormsAuthenticationUser) GetPropertyPassword()(value string, err error) { 
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

