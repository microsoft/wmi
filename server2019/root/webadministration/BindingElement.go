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

// BindingElement struct
type BindingElement struct { 
	*CollectionElement

	// 
	BindingInformation string

	// 
	Protocol string

	// 
	SslFlags uint32
}

	func NewBindingElementEx1(instance *cim.WmiInstance) (newInstance *BindingElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &BindingElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BindingElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BindingElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetBindingInformation sets the value of BindingInformation for the instance
func (instance *BindingElement) SetPropertyBindingInformation(value string) (err error) { 
    return instance.SetProperty("BindingInformation", (value))
}

// GetBindingInformation gets the value of BindingInformation for the instance
func (instance *BindingElement) GetPropertyBindingInformation()(value string, err error) { 
    retValue, err := instance.GetProperty("BindingInformation")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *BindingElement) SetPropertyProtocol(value string) (err error) { 
    return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *BindingElement) GetPropertyProtocol()(value string, err error) { 
    retValue, err := instance.GetProperty("Protocol")
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

// SetSslFlags sets the value of SslFlags for the instance
func (instance *BindingElement) SetPropertySslFlags(value uint32) (err error) { 
    return instance.SetProperty("SslFlags", (value))
}

// GetSslFlags gets the value of SslFlags for the instance
func (instance *BindingElement) GetPropertySslFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SslFlags")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

