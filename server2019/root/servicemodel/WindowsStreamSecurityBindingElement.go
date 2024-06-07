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

// WindowsStreamSecurityBindingElement struct
type WindowsStreamSecurityBindingElement struct { 
	*BindingElement

	// The ProtectionLevel for the TCP stream.
	ProtectionLevel string
}

	func NewWindowsStreamSecurityBindingElementEx1(instance *cim.WmiInstance) (newInstance *WindowsStreamSecurityBindingElement, err error) {tmp, err := NewBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &WindowsStreamSecurityBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

	func NewWindowsStreamSecurityBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WindowsStreamSecurityBindingElement, err error) {tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WindowsStreamSecurityBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

// SetProtectionLevel sets the value of ProtectionLevel for the instance
func (instance *WindowsStreamSecurityBindingElement) SetPropertyProtectionLevel(value string) (err error) { 
    return instance.SetProperty("ProtectionLevel", (value))
}

// GetProtectionLevel gets the value of ProtectionLevel for the instance
func (instance *WindowsStreamSecurityBindingElement) GetPropertyProtectionLevel()(value string, err error) { 
    retValue, err := instance.GetProperty("ProtectionLevel")
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

