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

// SoapServerProtocolFactory struct
type SoapServerProtocolFactory struct { 
	*EmbeddedObject

	// 
	Type string
}

	func NewSoapServerProtocolFactoryEx1(instance *cim.WmiInstance) (newInstance *SoapServerProtocolFactory, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SoapServerProtocolFactory {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSoapServerProtocolFactoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SoapServerProtocolFactory, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SoapServerProtocolFactory {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetType sets the value of Type for the instance
func (instance *SoapServerProtocolFactory) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *SoapServerProtocolFactory) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

