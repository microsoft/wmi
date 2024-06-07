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

// TraceAreaSettings struct
type TraceAreaSettings struct { 
	*EmbeddedObject

	// 
	Areas []TraceAreaDefinition
}

	func NewTraceAreaSettingsEx1(instance *cim.WmiInstance) (newInstance *TraceAreaSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &TraceAreaSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewTraceAreaSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TraceAreaSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TraceAreaSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAreas sets the value of Areas for the instance
func (instance *TraceAreaSettings) SetPropertyAreas(value []TraceAreaDefinition) (err error) { 
    return instance.SetProperty("Areas", (value))
}

// GetAreas gets the value of Areas for the instance
func (instance *TraceAreaSettings) GetPropertyAreas()(value []TraceAreaDefinition, err error) { 
    retValue, err := instance.GetProperty("Areas")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(TraceAreaDefinition); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " TraceAreaDefinition is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, TraceAreaDefinition(valuetmp))
    }

    return
}

