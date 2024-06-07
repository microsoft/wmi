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

// CustomFieldsSettings struct
type CustomFieldsSettings struct { 
	*EmbeddedObject

	// 
	CustomFields []CustomFieldsElement

	// 
	MaxCustomFieldLength uint32
}

	func NewCustomFieldsSettingsEx1(instance *cim.WmiInstance) (newInstance *CustomFieldsSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &CustomFieldsSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewCustomFieldsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CustomFieldsSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CustomFieldsSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetCustomFields sets the value of CustomFields for the instance
func (instance *CustomFieldsSettings) SetPropertyCustomFields(value []CustomFieldsElement) (err error) { 
    return instance.SetProperty("CustomFields", (value))
}

// GetCustomFields gets the value of CustomFields for the instance
func (instance *CustomFieldsSettings) GetPropertyCustomFields()(value []CustomFieldsElement, err error) { 
    retValue, err := instance.GetProperty("CustomFields")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(CustomFieldsElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " CustomFieldsElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, CustomFieldsElement(valuetmp))
    }

    return
}

// SetMaxCustomFieldLength sets the value of MaxCustomFieldLength for the instance
func (instance *CustomFieldsSettings) SetPropertyMaxCustomFieldLength(value uint32) (err error) { 
    return instance.SetProperty("MaxCustomFieldLength", (value))
}

// GetMaxCustomFieldLength gets the value of MaxCustomFieldLength for the instance
func (instance *CustomFieldsSettings) GetPropertyMaxCustomFieldLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxCustomFieldLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

