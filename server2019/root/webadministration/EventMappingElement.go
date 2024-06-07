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

// EventMappingElement struct
type EventMappingElement struct { 
	*CollectionElement

	// 
	EndEventCode int32

	// 
	Name string

	// 
	StartEventCode int32

	// 
	Type string
}

	func NewEventMappingElementEx1(instance *cim.WmiInstance) (newInstance *EventMappingElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &EventMappingElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewEventMappingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EventMappingElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EventMappingElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetEndEventCode sets the value of EndEventCode for the instance
func (instance *EventMappingElement) SetPropertyEndEventCode(value int32) (err error) { 
    return instance.SetProperty("EndEventCode", (value))
}

// GetEndEventCode gets the value of EndEventCode for the instance
func (instance *EventMappingElement) GetPropertyEndEventCode()(value int32, err error) { 
    retValue, err := instance.GetProperty("EndEventCode")
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

// SetName sets the value of Name for the instance
func (instance *EventMappingElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *EventMappingElement) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetStartEventCode sets the value of StartEventCode for the instance
func (instance *EventMappingElement) SetPropertyStartEventCode(value int32) (err error) { 
    return instance.SetProperty("StartEventCode", (value))
}

// GetStartEventCode gets the value of StartEventCode for the instance
func (instance *EventMappingElement) GetPropertyStartEventCode()(value int32, err error) { 
    retValue, err := instance.GetProperty("StartEventCode")
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

// SetType sets the value of Type for the instance
func (instance *EventMappingElement) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *EventMappingElement) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

