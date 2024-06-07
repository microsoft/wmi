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

// CustomBindingElement struct
type CustomBindingElement struct { 
	*BindingElement

	// A string that contains the configuration name of the binding. This value is a user-defined string that acts as the identification string for the custom binding.
	Name string
}

	func NewCustomBindingElementEx1(instance *cim.WmiInstance) (newInstance *CustomBindingElement, err error) {tmp, err := NewBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &CustomBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

	func NewCustomBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CustomBindingElement, err error) {tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CustomBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

// SetName sets the value of Name for the instance
func (instance *CustomBindingElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *CustomBindingElement) GetPropertyName()(value string, err error) { 
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

