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

// BypassListSettings struct
type BypassListSettings struct { 
	*EmbeddedObject

	// 
	BypassList []BypassElement
}

	func NewBypassListSettingsEx1(instance *cim.WmiInstance) (newInstance *BypassListSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &BypassListSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewBypassListSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BypassListSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BypassListSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetBypassList sets the value of BypassList for the instance
func (instance *BypassListSettings) SetPropertyBypassList(value []BypassElement) (err error) { 
    return instance.SetProperty("BypassList", (value))
}

// GetBypassList gets the value of BypassList for the instance
func (instance *BypassListSettings) GetPropertyBypassList()(value []BypassElement, err error) { 
    retValue, err := instance.GetProperty("BypassList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(BypassElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " BypassElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, BypassElement(valuetmp))
    }

    return
}

