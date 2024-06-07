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

// DenyStringElement struct
type DenyStringElement struct { 
	*CollectionElement

	// 
	String string
}

	func NewDenyStringElementEx1(instance *cim.WmiInstance) (newInstance *DenyStringElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &DenyStringElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewDenyStringElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DenyStringElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DenyStringElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetString sets the value of String for the instance
func (instance *DenyStringElement) SetPropertyString(value string) (err error) { 
    return instance.SetProperty("String", (value))
}

// GetString gets the value of String for the instance
func (instance *DenyStringElement) GetPropertyString()(value string, err error) { 
    retValue, err := instance.GetProperty("String")
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

