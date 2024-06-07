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

// IisClientCertificateMappingAuthenticationSection struct
type IisClientCertificateMappingAuthenticationSection struct { 
	*ConfigurationSectionWithCollection

	// 
	DefaultLogonDomain string

	// 
	Enabled bool

	// 
	LogonMethod int32

	// 
	ManyToOneCertificateMappingsEnabled bool

	// 
	ManyToOneMappings ManyToOneSettings

	// 
	OneToOneCertificateMappingsEnabled bool

	// 
	OneToOneMappings OneToOneSettings
}

	func NewIisClientCertificateMappingAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *IisClientCertificateMappingAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &IisClientCertificateMappingAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewIisClientCertificateMappingAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IisClientCertificateMappingAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IisClientCertificateMappingAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetDefaultLogonDomain sets the value of DefaultLogonDomain for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyDefaultLogonDomain(value string) (err error) { 
    return instance.SetProperty("DefaultLogonDomain", (value))
}

// GetDefaultLogonDomain gets the value of DefaultLogonDomain for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyDefaultLogonDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultLogonDomain")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyEnabled()(value bool, err error) { 
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

// SetLogonMethod sets the value of LogonMethod for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyLogonMethod(value int32) (err error) { 
    return instance.SetProperty("LogonMethod", (value))
}

// GetLogonMethod gets the value of LogonMethod for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyLogonMethod()(value int32, err error) { 
    retValue, err := instance.GetProperty("LogonMethod")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetManyToOneCertificateMappingsEnabled sets the value of ManyToOneCertificateMappingsEnabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyManyToOneCertificateMappingsEnabled(value bool) (err error) { 
    return instance.SetProperty("ManyToOneCertificateMappingsEnabled", (value))
}

// GetManyToOneCertificateMappingsEnabled gets the value of ManyToOneCertificateMappingsEnabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyManyToOneCertificateMappingsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("ManyToOneCertificateMappingsEnabled")
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

// SetManyToOneMappings sets the value of ManyToOneMappings for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyManyToOneMappings(value ManyToOneSettings) (err error) { 
    return instance.SetProperty("ManyToOneMappings", (value))
}

// GetManyToOneMappings gets the value of ManyToOneMappings for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyManyToOneMappings()(value ManyToOneSettings, err error) { 
    retValue, err := instance.GetProperty("ManyToOneMappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ManyToOneSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ManyToOneSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ManyToOneSettings(valuetmp)

    return
}

// SetOneToOneCertificateMappingsEnabled sets the value of OneToOneCertificateMappingsEnabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyOneToOneCertificateMappingsEnabled(value bool) (err error) { 
    return instance.SetProperty("OneToOneCertificateMappingsEnabled", (value))
}

// GetOneToOneCertificateMappingsEnabled gets the value of OneToOneCertificateMappingsEnabled for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyOneToOneCertificateMappingsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("OneToOneCertificateMappingsEnabled")
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

// SetOneToOneMappings sets the value of OneToOneMappings for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) SetPropertyOneToOneMappings(value OneToOneSettings) (err error) { 
    return instance.SetProperty("OneToOneMappings", (value))
}

// GetOneToOneMappings gets the value of OneToOneMappings for the instance
func (instance *IisClientCertificateMappingAuthenticationSection) GetPropertyOneToOneMappings()(value OneToOneSettings, err error) { 
    retValue, err := instance.GetProperty("OneToOneMappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(OneToOneSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " OneToOneSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = OneToOneSettings(valuetmp)

    return
}

