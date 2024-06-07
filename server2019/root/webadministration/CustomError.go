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

// CustomError struct
type CustomError struct { 
	*CollectionElement

	// 
	Redirect string

	// 
	StatusCode int32
}

	func NewCustomErrorEx1(instance *cim.WmiInstance) (newInstance *CustomError, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &CustomError {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewCustomErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CustomError, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CustomError {
	CollectionElement: tmp,
	}
	return
	}
	

// SetRedirect sets the value of Redirect for the instance
func (instance *CustomError) SetPropertyRedirect(value string) (err error) { 
    return instance.SetProperty("Redirect", (value))
}

// GetRedirect gets the value of Redirect for the instance
func (instance *CustomError) GetPropertyRedirect()(value string, err error) { 
    retValue, err := instance.GetProperty("Redirect")
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

// SetStatusCode sets the value of StatusCode for the instance
func (instance *CustomError) SetPropertyStatusCode(value int32) (err error) { 
    return instance.SetProperty("StatusCode", (value))
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *CustomError) GetPropertyStatusCode()(value int32, err error) { 
    retValue, err := instance.GetProperty("StatusCode")
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

