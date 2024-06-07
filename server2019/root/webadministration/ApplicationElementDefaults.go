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

// ApplicationElementDefaults struct
type ApplicationElementDefaults struct { 
	*EmbeddedObject

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
}

	func NewApplicationElementDefaultsEx1(instance *cim.WmiInstance) (newInstance *ApplicationElementDefaults, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ApplicationElementDefaults {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewApplicationElementDefaultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ApplicationElementDefaults, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ApplicationElementDefaults {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetApplicationPool sets the value of ApplicationPool for the instance
func (instance *ApplicationElementDefaults) SetPropertyApplicationPool(value string) (err error) { 
    return instance.SetProperty("ApplicationPool", (value))
}

// GetApplicationPool gets the value of ApplicationPool for the instance
func (instance *ApplicationElementDefaults) GetPropertyApplicationPool()(value string, err error) { 
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
func (instance *ApplicationElementDefaults) SetPropertyEnabledProtocols(value string) (err error) { 
    return instance.SetProperty("EnabledProtocols", (value))
}

// GetEnabledProtocols gets the value of EnabledProtocols for the instance
func (instance *ApplicationElementDefaults) GetPropertyEnabledProtocols()(value string, err error) { 
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
func (instance *ApplicationElementDefaults) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *ApplicationElementDefaults) GetPropertyPath()(value string, err error) { 
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
func (instance *ApplicationElementDefaults) SetPropertyPreloadEnabled(value bool) (err error) { 
    return instance.SetProperty("PreloadEnabled", (value))
}

// GetPreloadEnabled gets the value of PreloadEnabled for the instance
func (instance *ApplicationElementDefaults) GetPropertyPreloadEnabled()(value bool, err error) { 
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
func (instance *ApplicationElementDefaults) SetPropertyServiceAutoStartEnabled(value bool) (err error) { 
    return instance.SetProperty("ServiceAutoStartEnabled", (value))
}

// GetServiceAutoStartEnabled gets the value of ServiceAutoStartEnabled for the instance
func (instance *ApplicationElementDefaults) GetPropertyServiceAutoStartEnabled()(value bool, err error) { 
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
func (instance *ApplicationElementDefaults) SetPropertyServiceAutoStartProvider(value string) (err error) { 
    return instance.SetProperty("ServiceAutoStartProvider", (value))
}

// GetServiceAutoStartProvider gets the value of ServiceAutoStartProvider for the instance
func (instance *ApplicationElementDefaults) GetPropertyServiceAutoStartProvider()(value string, err error) { 
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

