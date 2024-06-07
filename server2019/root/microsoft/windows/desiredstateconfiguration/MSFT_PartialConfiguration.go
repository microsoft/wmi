// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_PartialConfiguration struct
type MSFT_PartialConfiguration struct { 
	*OMI_MetaConfigurationResource

	// 
	ConfigurationSource []string

	// 
	DependsOn []string

	// 
	Description string

	// 
	ExclusiveResources []string

	// 
	RefreshMode string

	// 
	ResourceModuleSource []string
}

	func NewMSFT_PartialConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_PartialConfiguration, err error) {tmp, err := NewOMI_MetaConfigurationResourceEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_PartialConfiguration {
	OMI_MetaConfigurationResource: tmp,
	}
	return
	}
	

	func NewMSFT_PartialConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_PartialConfiguration, err error) {tmp, err := NewOMI_MetaConfigurationResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_PartialConfiguration {
	OMI_MetaConfigurationResource: tmp,
	}
	return
	}
	

// SetConfigurationSource sets the value of ConfigurationSource for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyConfigurationSource(value []string) (err error) { 
    return instance.SetProperty("ConfigurationSource", (value))
}

// GetConfigurationSource gets the value of ConfigurationSource for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyConfigurationSource()(value []string, err error) { 
    retValue, err := instance.GetProperty("ConfigurationSource")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDependsOn sets the value of DependsOn for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyDependsOn(value []string) (err error) { 
    return instance.SetProperty("DependsOn", (value))
}

// GetDependsOn gets the value of DependsOn for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyDependsOn()(value []string, err error) { 
    retValue, err := instance.GetProperty("DependsOn")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyDescription(value string) (err error) { 
    return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("Description")
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

// SetExclusiveResources sets the value of ExclusiveResources for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyExclusiveResources(value []string) (err error) { 
    return instance.SetProperty("ExclusiveResources", (value))
}

// GetExclusiveResources gets the value of ExclusiveResources for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyExclusiveResources()(value []string, err error) { 
    retValue, err := instance.GetProperty("ExclusiveResources")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetRefreshMode sets the value of RefreshMode for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyRefreshMode(value string) (err error) { 
    return instance.SetProperty("RefreshMode", (value))
}

// GetRefreshMode gets the value of RefreshMode for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyRefreshMode()(value string, err error) { 
    retValue, err := instance.GetProperty("RefreshMode")
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

// SetResourceModuleSource sets the value of ResourceModuleSource for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyResourceModuleSource(value []string) (err error) { 
    return instance.SetProperty("ResourceModuleSource", (value))
}

// GetResourceModuleSource gets the value of ResourceModuleSource for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyResourceModuleSource()(value []string, err error) { 
    retValue, err := instance.GetProperty("ResourceModuleSource")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

