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

// FormsAuthenticationCredentials struct
type FormsAuthenticationCredentials struct { 
	*EmbeddedObject

	// 
	Credentials []FormsAuthenticationUser

	// 
	PasswordFormat int32
}

	func NewFormsAuthenticationCredentialsEx1(instance *cim.WmiInstance) (newInstance *FormsAuthenticationCredentials, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationCredentials {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFormsAuthenticationCredentialsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FormsAuthenticationCredentials, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationCredentials {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetCredentials sets the value of Credentials for the instance
func (instance *FormsAuthenticationCredentials) SetPropertyCredentials(value []FormsAuthenticationUser) (err error) { 
    return instance.SetProperty("Credentials", (value))
}

// GetCredentials gets the value of Credentials for the instance
func (instance *FormsAuthenticationCredentials) GetPropertyCredentials()(value []FormsAuthenticationUser, err error) { 
    retValue, err := instance.GetProperty("Credentials")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(FormsAuthenticationUser); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " FormsAuthenticationUser is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, FormsAuthenticationUser(valuetmp))
    }

    return
}

// SetPasswordFormat sets the value of PasswordFormat for the instance
func (instance *FormsAuthenticationCredentials) SetPropertyPasswordFormat(value int32) (err error) { 
    return instance.SetProperty("PasswordFormat", (value))
}

// GetPasswordFormat gets the value of PasswordFormat for the instance
func (instance *FormsAuthenticationCredentials) GetPropertyPasswordFormat()(value int32, err error) { 
    retValue, err := instance.GetProperty("PasswordFormat")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

