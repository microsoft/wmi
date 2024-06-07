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

// UrlElement struct
type UrlElement struct { 
	*CollectionElement

	// 
	Url string
}

	func NewUrlElementEx1(instance *cim.WmiInstance) (newInstance *UrlElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &UrlElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewUrlElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UrlElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UrlElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetUrl sets the value of Url for the instance
func (instance *UrlElement) SetPropertyUrl(value string) (err error) { 
    return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *UrlElement) GetPropertyUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("Url")
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

