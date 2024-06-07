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

// StaticTypeSettings struct
type StaticTypeSettings struct { 
	*EmbeddedObject

	// 
	StaticTypes []HttpCompressionMimeTypeElement
}

	func NewStaticTypeSettingsEx1(instance *cim.WmiInstance) (newInstance *StaticTypeSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &StaticTypeSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewStaticTypeSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *StaticTypeSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &StaticTypeSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetStaticTypes sets the value of StaticTypes for the instance
func (instance *StaticTypeSettings) SetPropertyStaticTypes(value []HttpCompressionMimeTypeElement) (err error) { 
    return instance.SetProperty("StaticTypes", (value))
}

// GetStaticTypes gets the value of StaticTypes for the instance
func (instance *StaticTypeSettings) GetPropertyStaticTypes()(value []HttpCompressionMimeTypeElement, err error) { 
    retValue, err := instance.GetProperty("StaticTypes")
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

