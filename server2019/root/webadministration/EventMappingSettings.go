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

// EventMappingSettings struct
type EventMappingSettings struct { 
	*EmbeddedObject

	// 
	EventMappings []EventMappingElement
}

	func NewEventMappingSettingsEx1(instance *cim.WmiInstance) (newInstance *EventMappingSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &EventMappingSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewEventMappingSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EventMappingSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EventMappingSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetEventMappings sets the value of EventMappings for the instance
func (instance *EventMappingSettings) SetPropertyEventMappings(value []EventMappingElement) (err error) { 
    return instance.SetProperty("EventMappings", (value))
}

// GetEventMappings gets the value of EventMappings for the instance
func (instance *EventMappingSettings) GetPropertyEventMappings()(value []EventMappingElement, err error) { 
    retValue, err := instance.GetProperty("EventMappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(EventMappingElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " EventMappingElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, EventMappingElement(valuetmp))
    }

    return
}

