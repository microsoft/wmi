// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MsmqIntegrationBindingElement struct
type MsmqIntegrationBindingElement struct { 
	*MsmqBindingElementBase

	// The format the binding uses to serialize messages.
	SerializationFormat string
}

	func NewMsmqIntegrationBindingElementEx1(instance *cim.WmiInstance) (newInstance *MsmqIntegrationBindingElement, err error) {tmp, err := NewMsmqBindingElementBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MsmqIntegrationBindingElement {
	MsmqBindingElementBase: tmp,
	}
	return
	}
	

	func NewMsmqIntegrationBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MsmqIntegrationBindingElement, err error) {tmp, err := NewMsmqBindingElementBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MsmqIntegrationBindingElement {
	MsmqBindingElementBase: tmp,
	}
	return
	}
	

// SetSerializationFormat sets the value of SerializationFormat for the instance
func (instance *MsmqIntegrationBindingElement) SetPropertySerializationFormat(value string) (err error) { 
    return instance.SetProperty("SerializationFormat", (value))
}

// GetSerializationFormat gets the value of SerializationFormat for the instance
func (instance *MsmqIntegrationBindingElement) GetPropertySerializationFormat()(value string, err error) { 
    retValue, err := instance.GetProperty("SerializationFormat")
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

