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

// TraceUrlSettings struct
type TraceUrlSettings struct { 
	*EmbeddedObject

	// 
	TraceUrls []StringElement
}

	func NewTraceUrlSettingsEx1(instance *cim.WmiInstance) (newInstance *TraceUrlSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &TraceUrlSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewTraceUrlSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TraceUrlSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TraceUrlSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetTraceUrls sets the value of TraceUrls for the instance
func (instance *TraceUrlSettings) SetPropertyTraceUrls(value []StringElement) (err error) { 
    return instance.SetProperty("TraceUrls", (value))
}

// GetTraceUrls gets the value of TraceUrls for the instance
func (instance *TraceUrlSettings) GetPropertyTraceUrls()(value []StringElement, err error) { 
    retValue, err := instance.GetProperty("TraceUrls")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(StringElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " StringElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, StringElement(valuetmp))
    }

    return
}

