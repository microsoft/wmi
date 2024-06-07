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

// MessageEncodingBindingElement struct
type MessageEncodingBindingElement struct { 
	*BindingElement

	// The SOAP version of the messages sent using the binding.
	MessageVersion string
}

	func NewMessageEncodingBindingElementEx1(instance *cim.WmiInstance) (newInstance *MessageEncodingBindingElement, err error) {tmp, err := NewBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &MessageEncodingBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

	func NewMessageEncodingBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MessageEncodingBindingElement, err error) {tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MessageEncodingBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

// SetMessageVersion sets the value of MessageVersion for the instance
func (instance *MessageEncodingBindingElement) SetPropertyMessageVersion(value string) (err error) { 
    return instance.SetProperty("MessageVersion", (value))
}

// GetMessageVersion gets the value of MessageVersion for the instance
func (instance *MessageEncodingBindingElement) GetPropertyMessageVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("MessageVersion")
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

