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

// PassportAuthentication struct
type PassportAuthentication struct { 
	*EmbeddedObject

	// 
	RedirectUrl string
}

	func NewPassportAuthenticationEx1(instance *cim.WmiInstance) (newInstance *PassportAuthentication, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &PassportAuthentication {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewPassportAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PassportAuthentication, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PassportAuthentication {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetRedirectUrl sets the value of RedirectUrl for the instance
func (instance *PassportAuthentication) SetPropertyRedirectUrl(value string) (err error) { 
    return instance.SetProperty("RedirectUrl", (value))
}

// GetRedirectUrl gets the value of RedirectUrl for the instance
func (instance *PassportAuthentication) GetPropertyRedirectUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("RedirectUrl")
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

