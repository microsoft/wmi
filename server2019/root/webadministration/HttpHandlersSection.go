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

// HttpHandlersSection struct
type HttpHandlersSection struct { 
	*ConfigurationSectionWithCollection

	// 
	HttpHandlers []HttpHandlerAction
}

	func NewHttpHandlersSectionEx1(instance *cim.WmiInstance) (newInstance *HttpHandlersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpHandlersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewHttpHandlersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpHandlersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpHandlersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetHttpHandlers sets the value of HttpHandlers for the instance
func (instance *HttpHandlersSection) SetPropertyHttpHandlers(value []HttpHandlerAction) (err error) { 
    return instance.SetProperty("HttpHandlers", (value))
}

// GetHttpHandlers gets the value of HttpHandlers for the instance
func (instance *HttpHandlersSection) GetPropertyHttpHandlers()(value []HttpHandlerAction, err error) { 
    retValue, err := instance.GetProperty("HttpHandlers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(HttpHandlerAction); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " HttpHandlerAction is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, HttpHandlerAction(valuetmp))
    }

    return
}

