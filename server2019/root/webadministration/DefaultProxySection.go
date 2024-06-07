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

// DefaultProxySection struct
type DefaultProxySection struct { 
	*ConfigurationSectionWithCollection

	// 
	BypassList BypassListSettings

	// 
	Enabled bool

	// 
	Module ModuleSettings

	// 
	Proxy ProxySettings

	// 
	UseDefaultCredentials bool
}

	func NewDefaultProxySectionEx1(instance *cim.WmiInstance) (newInstance *DefaultProxySection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &DefaultProxySection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewDefaultProxySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DefaultProxySection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DefaultProxySection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetBypassList sets the value of BypassList for the instance
func (instance *DefaultProxySection) SetPropertyBypassList(value BypassListSettings) (err error) { 
    return instance.SetProperty("BypassList", (value))
}

// GetBypassList gets the value of BypassList for the instance
func (instance *DefaultProxySection) GetPropertyBypassList()(value BypassListSettings, err error) { 
    retValue, err := instance.GetProperty("BypassList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(BypassListSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " BypassListSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = BypassListSettings(valuetmp)

    return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *DefaultProxySection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *DefaultProxySection) GetPropertyEnabled()(value bool, err error) { 
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

// SetModule sets the value of Module for the instance
func (instance *DefaultProxySection) SetPropertyModule(value ModuleSettings) (err error) { 
    return instance.SetProperty("Module", (value))
}

// GetModule gets the value of Module for the instance
func (instance *DefaultProxySection) GetPropertyModule()(value ModuleSettings, err error) { 
    retValue, err := instance.GetProperty("Module")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ModuleSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ModuleSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ModuleSettings(valuetmp)

    return
}

// SetProxy sets the value of Proxy for the instance
func (instance *DefaultProxySection) SetPropertyProxy(value ProxySettings) (err error) { 
    return instance.SetProperty("Proxy", (value))
}

// GetProxy gets the value of Proxy for the instance
func (instance *DefaultProxySection) GetPropertyProxy()(value ProxySettings, err error) { 
    retValue, err := instance.GetProperty("Proxy")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ProxySettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ProxySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ProxySettings(valuetmp)

    return
}

// SetUseDefaultCredentials sets the value of UseDefaultCredentials for the instance
func (instance *DefaultProxySection) SetPropertyUseDefaultCredentials(value bool) (err error) { 
    return instance.SetProperty("UseDefaultCredentials", (value))
}

// GetUseDefaultCredentials gets the value of UseDefaultCredentials for the instance
func (instance *DefaultProxySection) GetPropertyUseDefaultCredentials()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseDefaultCredentials")
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

