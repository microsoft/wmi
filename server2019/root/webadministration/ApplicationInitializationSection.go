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

// ApplicationInitializationSection struct
type ApplicationInitializationSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ApplicationInitialization []AppInitApps

	// 
	DoAppInitAfterRestart bool

	// 
	RemapManagedRequestsTo string

	// 
	SkipManagedModules bool
}

	func NewApplicationInitializationSectionEx1(instance *cim.WmiInstance) (newInstance *ApplicationInitializationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ApplicationInitializationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewApplicationInitializationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ApplicationInitializationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ApplicationInitializationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetApplicationInitialization sets the value of ApplicationInitialization for the instance
func (instance *ApplicationInitializationSection) SetPropertyApplicationInitialization(value []AppInitApps) (err error) { 
    return instance.SetProperty("ApplicationInitialization", (value))
}

// GetApplicationInitialization gets the value of ApplicationInitialization for the instance
func (instance *ApplicationInitializationSection) GetPropertyApplicationInitialization()(value []AppInitApps, err error) { 
    retValue, err := instance.GetProperty("ApplicationInitialization")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(AppInitApps); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " AppInitApps is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, AppInitApps(valuetmp))
    }

    return
}

// SetDoAppInitAfterRestart sets the value of DoAppInitAfterRestart for the instance
func (instance *ApplicationInitializationSection) SetPropertyDoAppInitAfterRestart(value bool) (err error) { 
    return instance.SetProperty("DoAppInitAfterRestart", (value))
}

// GetDoAppInitAfterRestart gets the value of DoAppInitAfterRestart for the instance
func (instance *ApplicationInitializationSection) GetPropertyDoAppInitAfterRestart()(value bool, err error) { 
    retValue, err := instance.GetProperty("DoAppInitAfterRestart")
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

// SetRemapManagedRequestsTo sets the value of RemapManagedRequestsTo for the instance
func (instance *ApplicationInitializationSection) SetPropertyRemapManagedRequestsTo(value string) (err error) { 
    return instance.SetProperty("RemapManagedRequestsTo", (value))
}

// GetRemapManagedRequestsTo gets the value of RemapManagedRequestsTo for the instance
func (instance *ApplicationInitializationSection) GetPropertyRemapManagedRequestsTo()(value string, err error) { 
    retValue, err := instance.GetProperty("RemapManagedRequestsTo")
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

// SetSkipManagedModules sets the value of SkipManagedModules for the instance
func (instance *ApplicationInitializationSection) SetPropertySkipManagedModules(value bool) (err error) { 
    return instance.SetProperty("SkipManagedModules", (value))
}

// GetSkipManagedModules gets the value of SkipManagedModules for the instance
func (instance *ApplicationInitializationSection) GetPropertySkipManagedModules()(value bool, err error) { 
    retValue, err := instance.GetProperty("SkipManagedModules")
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

