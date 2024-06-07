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

// WebDavPropertyStoresSettings struct
type WebDavPropertyStoresSettings struct { 
	*EmbeddedObject

	// 
	PropertyStores []NameImageElement
}

	func NewWebDavPropertyStoresSettingsEx1(instance *cim.WmiInstance) (newInstance *WebDavPropertyStoresSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &WebDavPropertyStoresSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewWebDavPropertyStoresSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WebDavPropertyStoresSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WebDavPropertyStoresSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetPropertyStores sets the value of PropertyStores for the instance
func (instance *WebDavPropertyStoresSettings) SetPropertyPropertyStores(value []NameImageElement) (err error) { 
    return instance.SetProperty("PropertyStores", (value))
}

// GetPropertyStores gets the value of PropertyStores for the instance
func (instance *WebDavPropertyStoresSettings) GetPropertyPropertyStores()(value []NameImageElement, err error) { 
    retValue, err := instance.GetProperty("PropertyStores")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(NameImageElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " NameImageElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, NameImageElement(valuetmp))
    }

    return
}

