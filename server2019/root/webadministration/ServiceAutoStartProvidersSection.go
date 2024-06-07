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

// ServiceAutoStartProvidersSection struct
type ServiceAutoStartProvidersSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ServiceAutoStartProviders []NameTypeElement
}

	func NewServiceAutoStartProvidersSectionEx1(instance *cim.WmiInstance) (newInstance *ServiceAutoStartProvidersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ServiceAutoStartProvidersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewServiceAutoStartProvidersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServiceAutoStartProvidersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServiceAutoStartProvidersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetServiceAutoStartProviders sets the value of ServiceAutoStartProviders for the instance
func (instance *ServiceAutoStartProvidersSection) SetPropertyServiceAutoStartProviders(value []NameTypeElement) (err error) { 
    return instance.SetProperty("ServiceAutoStartProviders", (value))
}

// GetServiceAutoStartProviders gets the value of ServiceAutoStartProviders for the instance
func (instance *ServiceAutoStartProvidersSection) GetPropertyServiceAutoStartProviders()(value []NameTypeElement, err error) { 
    retValue, err := instance.GetProperty("ServiceAutoStartProviders")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(NameTypeElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " NameTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, NameTypeElement(valuetmp))
    }

    return
}

