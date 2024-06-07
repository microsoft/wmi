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

// ServicePointManagerSettings struct
type ServicePointManagerSettings struct { 
	*EmbeddedObject

	// 
	CheckCertificateName bool

	// 
	CheckCertificateRevocationList bool

	// 
	DnsRefreshTimeout int32

	// 
	EnableDnsRoundRobin bool

	// 
	Expect100Continue bool

	// 
	UseNagleAlgorithm bool
}

	func NewServicePointManagerSettingsEx1(instance *cim.WmiInstance) (newInstance *ServicePointManagerSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ServicePointManagerSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewServicePointManagerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServicePointManagerSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServicePointManagerSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetCheckCertificateName sets the value of CheckCertificateName for the instance
func (instance *ServicePointManagerSettings) SetPropertyCheckCertificateName(value bool) (err error) { 
    return instance.SetProperty("CheckCertificateName", (value))
}

// GetCheckCertificateName gets the value of CheckCertificateName for the instance
func (instance *ServicePointManagerSettings) GetPropertyCheckCertificateName()(value bool, err error) { 
    retValue, err := instance.GetProperty("CheckCertificateName")
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

// SetCheckCertificateRevocationList sets the value of CheckCertificateRevocationList for the instance
func (instance *ServicePointManagerSettings) SetPropertyCheckCertificateRevocationList(value bool) (err error) { 
    return instance.SetProperty("CheckCertificateRevocationList", (value))
}

// GetCheckCertificateRevocationList gets the value of CheckCertificateRevocationList for the instance
func (instance *ServicePointManagerSettings) GetPropertyCheckCertificateRevocationList()(value bool, err error) { 
    retValue, err := instance.GetProperty("CheckCertificateRevocationList")
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

// SetDnsRefreshTimeout sets the value of DnsRefreshTimeout for the instance
func (instance *ServicePointManagerSettings) SetPropertyDnsRefreshTimeout(value int32) (err error) { 
    return instance.SetProperty("DnsRefreshTimeout", (value))
}

// GetDnsRefreshTimeout gets the value of DnsRefreshTimeout for the instance
func (instance *ServicePointManagerSettings) GetPropertyDnsRefreshTimeout()(value int32, err error) { 
    retValue, err := instance.GetProperty("DnsRefreshTimeout")
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

// SetEnableDnsRoundRobin sets the value of EnableDnsRoundRobin for the instance
func (instance *ServicePointManagerSettings) SetPropertyEnableDnsRoundRobin(value bool) (err error) { 
    return instance.SetProperty("EnableDnsRoundRobin", (value))
}

// GetEnableDnsRoundRobin gets the value of EnableDnsRoundRobin for the instance
func (instance *ServicePointManagerSettings) GetPropertyEnableDnsRoundRobin()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableDnsRoundRobin")
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

// SetExpect100Continue sets the value of Expect100Continue for the instance
func (instance *ServicePointManagerSettings) SetPropertyExpect100Continue(value bool) (err error) { 
    return instance.SetProperty("Expect100Continue", (value))
}

// GetExpect100Continue gets the value of Expect100Continue for the instance
func (instance *ServicePointManagerSettings) GetPropertyExpect100Continue()(value bool, err error) { 
    retValue, err := instance.GetProperty("Expect100Continue")
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

// SetUseNagleAlgorithm sets the value of UseNagleAlgorithm for the instance
func (instance *ServicePointManagerSettings) SetPropertyUseNagleAlgorithm(value bool) (err error) { 
    return instance.SetProperty("UseNagleAlgorithm", (value))
}

// GetUseNagleAlgorithm gets the value of UseNagleAlgorithm for the instance
func (instance *ServicePointManagerSettings) GetPropertyUseNagleAlgorithm()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseNagleAlgorithm")
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

