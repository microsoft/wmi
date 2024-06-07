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

// FastCgiEnvironmentElement struct
type FastCgiEnvironmentElement struct { 
	*CollectionElement

	// 
	Name string

	// 
	Value string
}

	func NewFastCgiEnvironmentElementEx1(instance *cim.WmiInstance) (newInstance *FastCgiEnvironmentElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &FastCgiEnvironmentElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewFastCgiEnvironmentElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FastCgiEnvironmentElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FastCgiEnvironmentElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetName sets the value of Name for the instance
func (instance *FastCgiEnvironmentElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FastCgiEnvironmentElement) GetPropertyName()(value string, err error) { 
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

// SetValue sets the value of Value for the instance
func (instance *FastCgiEnvironmentElement) SetPropertyValue(value string) (err error) { 
    return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *FastCgiEnvironmentElement) GetPropertyValue()(value string, err error) { 
    retValue, err := instance.GetProperty("Value")
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

