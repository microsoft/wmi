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

// Application struct
type Application struct { 
	*ConfiguredObject

	// 
	ApplicationPool string

	// 
	EnabledProtocols string

	// 
	Path string

	// 
	PreloadEnabled bool

	// 
	ServiceAutoStartEnabled bool

	// 
	ServiceAutoStartProvider string

	// 
	SiteName string

	// 
	VirtualDirectoryDefaults VirtualDirectoryElementDefaults
}

	func NewApplicationEx1(instance *cim.WmiInstance) (newInstance *Application, err error) {tmp, err := NewConfiguredObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &Application {
	ConfiguredObject: tmp,
	}
	return
	}
	

	func NewApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Application, err error) {tmp, err := NewConfiguredObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Application {
	ConfiguredObject: tmp,
	}
	return
	}
	

// SetApplicationPool sets the value of ApplicationPool for the instance
func (instance *Application) SetPropertyApplicationPool(value string) (err error) { 
    return instance.SetProperty("ApplicationPool", (value))
}

// GetApplicationPool gets the value of ApplicationPool for the instance
func (instance *Application) GetPropertyApplicationPool()(value string, err error) { 
    retValue, err := instance.GetProperty("ApplicationPool")
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

// SetEnabledProtocols sets the value of EnabledProtocols for the instance
func (instance *Application) SetPropertyEnabledProtocols(value string) (err error) { 
    return instance.SetProperty("EnabledProtocols", (value))
}

// GetEnabledProtocols gets the value of EnabledProtocols for the instance
func (instance *Application) GetPropertyEnabledProtocols()(value string, err error) { 
    retValue, err := instance.GetProperty("EnabledProtocols")
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

// SetPath sets the value of Path for the instance
func (instance *Application) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Application) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

// SetPreloadEnabled sets the value of PreloadEnabled for the instance
func (instance *Application) SetPropertyPreloadEnabled(value bool) (err error) { 
    return instance.SetProperty("PreloadEnabled", (value))
}

// GetPreloadEnabled gets the value of PreloadEnabled for the instance
func (instance *Application) GetPropertyPreloadEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("PreloadEnabled")
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

// SetServiceAutoStartEnabled sets the value of ServiceAutoStartEnabled for the instance
func (instance *Application) SetPropertyServiceAutoStartEnabled(value bool) (err error) { 
    return instance.SetProperty("ServiceAutoStartEnabled", (value))
}

// GetServiceAutoStartEnabled gets the value of ServiceAutoStartEnabled for the instance
func (instance *Application) GetPropertyServiceAutoStartEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("ServiceAutoStartEnabled")
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

// SetServiceAutoStartProvider sets the value of ServiceAutoStartProvider for the instance
func (instance *Application) SetPropertyServiceAutoStartProvider(value string) (err error) { 
    return instance.SetProperty("ServiceAutoStartProvider", (value))
}

// GetServiceAutoStartProvider gets the value of ServiceAutoStartProvider for the instance
func (instance *Application) GetPropertyServiceAutoStartProvider()(value string, err error) { 
    retValue, err := instance.GetProperty("ServiceAutoStartProvider")
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

// SetSiteName sets the value of SiteName for the instance
func (instance *Application) SetPropertySiteName(value string) (err error) { 
    return instance.SetProperty("SiteName", (value))
}

// GetSiteName gets the value of SiteName for the instance
func (instance *Application) GetPropertySiteName()(value string, err error) { 
    retValue, err := instance.GetProperty("SiteName")
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

// SetVirtualDirectoryDefaults sets the value of VirtualDirectoryDefaults for the instance
func (instance *Application) SetPropertyVirtualDirectoryDefaults(value VirtualDirectoryElementDefaults) (err error) { 
    return instance.SetProperty("VirtualDirectoryDefaults", (value))
}

// GetVirtualDirectoryDefaults gets the value of VirtualDirectoryDefaults for the instance
func (instance *Application) GetPropertyVirtualDirectoryDefaults()(value VirtualDirectoryElementDefaults, err error) { 
    retValue, err := instance.GetProperty("VirtualDirectoryDefaults")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(VirtualDirectoryElementDefaults); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " VirtualDirectoryElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = VirtualDirectoryElementDefaults(valuetmp)

    return
}

// 

// <param name="ApplicationPath" type="string "></param>
// <param name="PhysicalPath" type="string "></param>
// <param name="SiteName" type="string "></param>
func (instance *Application) Create( /* IN */ ApplicationPath string,
 /* IN */ SiteName string,
 /* OPTIONAL IN */ PhysicalPath string) (err error) {_, err = instance.InvokeMethodWithReturn("Create" , ApplicationPath, SiteName, PhysicalPath);
	if err != nil { return }
	return
	
}


// 

// <param name="PropertyName" type="string "></param>
func (instance *Application) RevertToParent( /* OPTIONAL IN */ PropertyName string) (err error) {_, err = instance.InvokeMethodWithReturn("RevertToParent" , PropertyName);
	if err != nil { return }
	return
	
}


