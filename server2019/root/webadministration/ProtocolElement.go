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

// ProtocolElement struct
type ProtocolElement struct { 
	*CollectionElement

	// 
	AppDomainHandlerType string

	// 
	Name string

	// 
	ProcessHandlerType string

	// 
	Validate bool
}

	func NewProtocolElementEx1(instance *cim.WmiInstance) (newInstance *ProtocolElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ProtocolElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewProtocolElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProtocolElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProtocolElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetAppDomainHandlerType sets the value of AppDomainHandlerType for the instance
func (instance *ProtocolElement) SetPropertyAppDomainHandlerType(value string) (err error) { 
    return instance.SetProperty("AppDomainHandlerType", (value))
}

// GetAppDomainHandlerType gets the value of AppDomainHandlerType for the instance
func (instance *ProtocolElement) GetPropertyAppDomainHandlerType()(value string, err error) { 
    retValue, err := instance.GetProperty("AppDomainHandlerType")
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

// SetName sets the value of Name for the instance
func (instance *ProtocolElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ProtocolElement) GetPropertyName()(value string, err error) { 
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

// SetProcessHandlerType sets the value of ProcessHandlerType for the instance
func (instance *ProtocolElement) SetPropertyProcessHandlerType(value string) (err error) { 
    return instance.SetProperty("ProcessHandlerType", (value))
}

// GetProcessHandlerType gets the value of ProcessHandlerType for the instance
func (instance *ProtocolElement) GetPropertyProcessHandlerType()(value string, err error) { 
    retValue, err := instance.GetProperty("ProcessHandlerType")
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

// SetValidate sets the value of Validate for the instance
func (instance *ProtocolElement) SetPropertyValidate(value bool) (err error) { 
    return instance.SetProperty("Validate", (value))
}

// GetValidate gets the value of Validate for the instance
func (instance *ProtocolElement) GetPropertyValidate()(value bool, err error) { 
    retValue, err := instance.GetProperty("Validate")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

