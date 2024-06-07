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

// WsdlHelpGeneratorInfo struct
type WsdlHelpGeneratorInfo struct { 
	*EmbeddedObject

	// 
	Href string
}

	func NewWsdlHelpGeneratorInfoEx1(instance *cim.WmiInstance) (newInstance *WsdlHelpGeneratorInfo, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &WsdlHelpGeneratorInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewWsdlHelpGeneratorInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WsdlHelpGeneratorInfo, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WsdlHelpGeneratorInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetHref sets the value of Href for the instance
func (instance *WsdlHelpGeneratorInfo) SetPropertyHref(value string) (err error) { 
    return instance.SetProperty("Href", (value))
}

// GetHref gets the value of Href for the instance
func (instance *WsdlHelpGeneratorInfo) GetPropertyHref()(value string, err error) { 
    retValue, err := instance.GetProperty("Href")
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

