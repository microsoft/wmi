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

// QueryStringSequenceSettings struct
type QueryStringSequenceSettings struct { 
	*EmbeddedObject

	// 
	Sequence []SequenceElement
}

	func NewQueryStringSequenceSettingsEx1(instance *cim.WmiInstance) (newInstance *QueryStringSequenceSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &QueryStringSequenceSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewQueryStringSequenceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *QueryStringSequenceSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &QueryStringSequenceSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetSequence sets the value of Sequence for the instance
func (instance *QueryStringSequenceSettings) SetPropertySequence(value []SequenceElement) (err error) { 
    return instance.SetProperty("Sequence", (value))
}

// GetSequence gets the value of Sequence for the instance
func (instance *QueryStringSequenceSettings) GetPropertySequence()(value []SequenceElement, err error) { 
    retValue, err := instance.GetProperty("Sequence")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(SequenceElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " SequenceElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, SequenceElement(valuetmp))
    }

    return
}

