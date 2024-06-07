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

// QueryStringElement struct
type QueryStringElement struct { 
	*CollectionElement

	// 
	QueryString string
}

	func NewQueryStringElementEx1(instance *cim.WmiInstance) (newInstance *QueryStringElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &QueryStringElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewQueryStringElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *QueryStringElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &QueryStringElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetQueryString sets the value of QueryString for the instance
func (instance *QueryStringElement) SetPropertyQueryString(value string) (err error) { 
    return instance.SetProperty("QueryString", (value))
}

// GetQueryString gets the value of QueryString for the instance
func (instance *QueryStringElement) GetPropertyQueryString()(value string, err error) { 
    retValue, err := instance.GetProperty("QueryString")
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

