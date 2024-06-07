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

// SourceSettings struct
type SourceSettings struct { 
	*EmbeddedObject

	// 
	Sources []SourceElement
}

	func NewSourceSettingsEx1(instance *cim.WmiInstance) (newInstance *SourceSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SourceSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSourceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SourceSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SourceSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetSources sets the value of Sources for the instance
func (instance *SourceSettings) SetPropertySources(value []SourceElement) (err error) { 
    return instance.SetProperty("Sources", (value))
}

// GetSources gets the value of Sources for the instance
func (instance *SourceSettings) GetPropertySources()(value []SourceElement, err error) { 
    retValue, err := instance.GetProperty("Sources")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(SourceElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " SourceElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, SourceElement(valuetmp))
    }

    return
}

