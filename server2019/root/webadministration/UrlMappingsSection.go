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

// UrlMappingsSection struct
type UrlMappingsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	Enabled bool

	// 
	UrlMappings []UrlMapping
}

	func NewUrlMappingsSectionEx1(instance *cim.WmiInstance) (newInstance *UrlMappingsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &UrlMappingsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewUrlMappingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UrlMappingsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UrlMappingsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *UrlMappingsSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *UrlMappingsSection) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetUrlMappings sets the value of UrlMappings for the instance
func (instance *UrlMappingsSection) SetPropertyUrlMappings(value []UrlMapping) (err error) { 
    return instance.SetProperty("UrlMappings", (value))
}

// GetUrlMappings gets the value of UrlMappings for the instance
func (instance *UrlMappingsSection) GetPropertyUrlMappings()(value []UrlMapping, err error) { 
    retValue, err := instance.GetProperty("UrlMappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(UrlMapping); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " UrlMapping is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, UrlMapping(valuetmp))
    }

    return
}

