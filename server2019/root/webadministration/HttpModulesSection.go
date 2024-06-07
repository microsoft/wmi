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

// HttpModulesSection struct
type HttpModulesSection struct { 
	*ConfigurationSectionWithCollection

	// 
	HttpModules []HttpModuleAction
}

	func NewHttpModulesSectionEx1(instance *cim.WmiInstance) (newInstance *HttpModulesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpModulesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewHttpModulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpModulesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpModulesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetHttpModules sets the value of HttpModules for the instance
func (instance *HttpModulesSection) SetPropertyHttpModules(value []HttpModuleAction) (err error) { 
    return instance.SetProperty("HttpModules", (value))
}

// GetHttpModules gets the value of HttpModules for the instance
func (instance *HttpModulesSection) GetPropertyHttpModules()(value []HttpModuleAction, err error) { 
    retValue, err := instance.GetProperty("HttpModules")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(HttpModuleAction); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " HttpModuleAction is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, HttpModuleAction(valuetmp))
    }

    return
}

