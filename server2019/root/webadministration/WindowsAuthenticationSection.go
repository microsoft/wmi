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

// WindowsAuthenticationSection struct
type WindowsAuthenticationSection struct { 
	*ConfigurationSectionWithCollection

	// 
	AuthPersistNonNTLM bool

	// 
	AuthPersistSingleRequest bool

	// 
	Enabled bool

	// 
	ExtendedProtection ExtendedProtectionSettings

	// 
	Providers AuthenticationProviderSettings

	// 
	UseAppPoolCredentials bool

	// 
	UseKernelMode bool
}

	func NewWindowsAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *WindowsAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &WindowsAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewWindowsAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WindowsAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WindowsAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAuthPersistNonNTLM sets the value of AuthPersistNonNTLM for the instance
func (instance *WindowsAuthenticationSection) SetPropertyAuthPersistNonNTLM(value bool) (err error) { 
    return instance.SetProperty("AuthPersistNonNTLM", (value))
}

// GetAuthPersistNonNTLM gets the value of AuthPersistNonNTLM for the instance
func (instance *WindowsAuthenticationSection) GetPropertyAuthPersistNonNTLM()(value bool, err error) { 
    retValue, err := instance.GetProperty("AuthPersistNonNTLM")
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

// SetAuthPersistSingleRequest sets the value of AuthPersistSingleRequest for the instance
func (instance *WindowsAuthenticationSection) SetPropertyAuthPersistSingleRequest(value bool) (err error) { 
    return instance.SetProperty("AuthPersistSingleRequest", (value))
}

// GetAuthPersistSingleRequest gets the value of AuthPersistSingleRequest for the instance
func (instance *WindowsAuthenticationSection) GetPropertyAuthPersistSingleRequest()(value bool, err error) { 
    retValue, err := instance.GetProperty("AuthPersistSingleRequest")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *WindowsAuthenticationSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *WindowsAuthenticationSection) GetPropertyEnabled()(value bool, err error) { 
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

// SetExtendedProtection sets the value of ExtendedProtection for the instance
func (instance *WindowsAuthenticationSection) SetPropertyExtendedProtection(value ExtendedProtectionSettings) (err error) { 
    return instance.SetProperty("ExtendedProtection", (value))
}

// GetExtendedProtection gets the value of ExtendedProtection for the instance
func (instance *WindowsAuthenticationSection) GetPropertyExtendedProtection()(value ExtendedProtectionSettings, err error) { 
    retValue, err := instance.GetProperty("ExtendedProtection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ExtendedProtectionSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ExtendedProtectionSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ExtendedProtectionSettings(valuetmp)

    return
}

// SetProviders sets the value of Providers for the instance
func (instance *WindowsAuthenticationSection) SetPropertyProviders(value AuthenticationProviderSettings) (err error) { 
    return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *WindowsAuthenticationSection) GetPropertyProviders()(value AuthenticationProviderSettings, err error) { 
    retValue, err := instance.GetProperty("Providers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(AuthenticationProviderSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " AuthenticationProviderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = AuthenticationProviderSettings(valuetmp)

    return
}

// SetUseAppPoolCredentials sets the value of UseAppPoolCredentials for the instance
func (instance *WindowsAuthenticationSection) SetPropertyUseAppPoolCredentials(value bool) (err error) { 
    return instance.SetProperty("UseAppPoolCredentials", (value))
}

// GetUseAppPoolCredentials gets the value of UseAppPoolCredentials for the instance
func (instance *WindowsAuthenticationSection) GetPropertyUseAppPoolCredentials()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseAppPoolCredentials")
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

// SetUseKernelMode sets the value of UseKernelMode for the instance
func (instance *WindowsAuthenticationSection) SetPropertyUseKernelMode(value bool) (err error) { 
    return instance.SetProperty("UseKernelMode", (value))
}

// GetUseKernelMode gets the value of UseKernelMode for the instance
func (instance *WindowsAuthenticationSection) GetPropertyUseKernelMode()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseKernelMode")
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

