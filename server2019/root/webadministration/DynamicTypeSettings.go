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

// DynamicTypeSettings struct
type DynamicTypeSettings struct { 
	*EmbeddedObject

	// 
	DynamicTypes []HttpCompressionMimeTypeElement
}

	func NewDynamicTypeSettingsEx1(instance *cim.WmiInstance) (newInstance *DynamicTypeSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &DynamicTypeSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewDynamicTypeSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DynamicTypeSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DynamicTypeSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetDynamicTypes sets the value of DynamicTypes for the instance
func (instance *DynamicTypeSettings) SetPropertyDynamicTypes(value []HttpCompressionMimeTypeElement) (err error) { 
    return instance.SetProperty("DynamicTypes", (value))
}

// GetDynamicTypes gets the value of DynamicTypes for the instance
func (instance *DynamicTypeSettings) GetPropertyDynamicTypes()(value []HttpCompressionMimeTypeElement, err error) { 
    retValue, err := instance.GetProperty("DynamicTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(HttpCompressionMimeTypeElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " HttpCompressionMimeTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, HttpCompressionMimeTypeElement(valuetmp))
    }

    return
}

