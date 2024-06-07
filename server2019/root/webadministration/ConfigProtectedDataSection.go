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

// ConfigProtectedDataSection struct
type ConfigProtectedDataSection struct { 
	*ConfigurationSectionWithCollection

	// 
	DefaultProvider string

	// 
	Providers ProvidersSettings
}

	func NewConfigProtectedDataSectionEx1(instance *cim.WmiInstance) (newInstance *ConfigProtectedDataSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ConfigProtectedDataSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewConfigProtectedDataSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConfigProtectedDataSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConfigProtectedDataSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetDefaultProvider sets the value of DefaultProvider for the instance
func (instance *ConfigProtectedDataSection) SetPropertyDefaultProvider(value string) (err error) { 
    return instance.SetProperty("DefaultProvider", (value))
}

// GetDefaultProvider gets the value of DefaultProvider for the instance
func (instance *ConfigProtectedDataSection) GetPropertyDefaultProvider()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultProvider")
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

// SetProviders sets the value of Providers for the instance
func (instance *ConfigProtectedDataSection) SetPropertyProviders(value ProvidersSettings) (err error) { 
    return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *ConfigProtectedDataSection) GetPropertyProviders()(value ProvidersSettings, err error) { 
    retValue, err := instance.GetProperty("Providers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ProvidersSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ProvidersSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ProvidersSettings(valuetmp)

    return
}

