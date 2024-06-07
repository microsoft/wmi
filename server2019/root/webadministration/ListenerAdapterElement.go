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

// ListenerAdapterElement struct
type ListenerAdapterElement struct { 
	*CollectionElement

	// 
	Identity string

	// 
	Name string

	// 
	ProtocolManagerDll string

	// 
	ProtocolManagerDllInitFunction string
}

	func NewListenerAdapterElementEx1(instance *cim.WmiInstance) (newInstance *ListenerAdapterElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ListenerAdapterElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewListenerAdapterElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ListenerAdapterElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ListenerAdapterElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetIdentity sets the value of Identity for the instance
func (instance *ListenerAdapterElement) SetPropertyIdentity(value string) (err error) { 
    return instance.SetProperty("Identity", (value))
}

// GetIdentity gets the value of Identity for the instance
func (instance *ListenerAdapterElement) GetPropertyIdentity()(value string, err error) { 
    retValue, err := instance.GetProperty("Identity")
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
func (instance *ListenerAdapterElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ListenerAdapterElement) GetPropertyName()(value string, err error) { 
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

// SetProtocolManagerDll sets the value of ProtocolManagerDll for the instance
func (instance *ListenerAdapterElement) SetPropertyProtocolManagerDll(value string) (err error) { 
    return instance.SetProperty("ProtocolManagerDll", (value))
}

// GetProtocolManagerDll gets the value of ProtocolManagerDll for the instance
func (instance *ListenerAdapterElement) GetPropertyProtocolManagerDll()(value string, err error) { 
    retValue, err := instance.GetProperty("ProtocolManagerDll")
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

// SetProtocolManagerDllInitFunction sets the value of ProtocolManagerDllInitFunction for the instance
func (instance *ListenerAdapterElement) SetPropertyProtocolManagerDllInitFunction(value string) (err error) { 
    return instance.SetProperty("ProtocolManagerDllInitFunction", (value))
}

// GetProtocolManagerDllInitFunction gets the value of ProtocolManagerDllInitFunction for the instance
func (instance *ListenerAdapterElement) GetPropertyProtocolManagerDllInitFunction()(value string, err error) { 
    retValue, err := instance.GetProperty("ProtocolManagerDllInitFunction")
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

