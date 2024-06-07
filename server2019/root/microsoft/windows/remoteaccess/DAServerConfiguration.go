// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DAServerConfiguration struct
type DAServerConfiguration struct { 
	*cim.WmiInstance

	// 
	ClientIPv6Prefix string

	// 
	ComputerCertAuthentication string

	// 
	ConnectToAddress string

	// 
	DAInstallType string

	// 
	GpoName string

	// 
	IntermediateRootCertificate bool

	// 
	InternalIPv6Prefix []string

	// 
	IPsecRootCertificate []uint8

	// 
	IsNatDeployed bool

	// 
	IsSingleNic bool

	// 
	TeredoState string

	// 
	UserAuthentication string
}

	func NewDAServerConfigurationEx1(instance *cim.WmiInstance) (newInstance *DAServerConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DAServerConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDAServerConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAServerConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAServerConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetClientIPv6Prefix sets the value of ClientIPv6Prefix for the instance
func (instance *DAServerConfiguration) SetPropertyClientIPv6Prefix(value string) (err error) { 
    return instance.SetProperty("ClientIPv6Prefix", (value))
}

// GetClientIPv6Prefix gets the value of ClientIPv6Prefix for the instance
func (instance *DAServerConfiguration) GetPropertyClientIPv6Prefix()(value string, err error) { 
    retValue, err := instance.GetProperty("ClientIPv6Prefix")
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

// SetComputerCertAuthentication sets the value of ComputerCertAuthentication for the instance
func (instance *DAServerConfiguration) SetPropertyComputerCertAuthentication(value string) (err error) { 
    return instance.SetProperty("ComputerCertAuthentication", (value))
}

// GetComputerCertAuthentication gets the value of ComputerCertAuthentication for the instance
func (instance *DAServerConfiguration) GetPropertyComputerCertAuthentication()(value string, err error) { 
    retValue, err := instance.GetProperty("ComputerCertAuthentication")
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

// SetConnectToAddress sets the value of ConnectToAddress for the instance
func (instance *DAServerConfiguration) SetPropertyConnectToAddress(value string) (err error) { 
    return instance.SetProperty("ConnectToAddress", (value))
}

// GetConnectToAddress gets the value of ConnectToAddress for the instance
func (instance *DAServerConfiguration) GetPropertyConnectToAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectToAddress")
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

// SetDAInstallType sets the value of DAInstallType for the instance
func (instance *DAServerConfiguration) SetPropertyDAInstallType(value string) (err error) { 
    return instance.SetProperty("DAInstallType", (value))
}

// GetDAInstallType gets the value of DAInstallType for the instance
func (instance *DAServerConfiguration) GetPropertyDAInstallType()(value string, err error) { 
    retValue, err := instance.GetProperty("DAInstallType")
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

// SetGpoName sets the value of GpoName for the instance
func (instance *DAServerConfiguration) SetPropertyGpoName(value string) (err error) { 
    return instance.SetProperty("GpoName", (value))
}

// GetGpoName gets the value of GpoName for the instance
func (instance *DAServerConfiguration) GetPropertyGpoName()(value string, err error) { 
    retValue, err := instance.GetProperty("GpoName")
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

// SetIntermediateRootCertificate sets the value of IntermediateRootCertificate for the instance
func (instance *DAServerConfiguration) SetPropertyIntermediateRootCertificate(value bool) (err error) { 
    return instance.SetProperty("IntermediateRootCertificate", (value))
}

// GetIntermediateRootCertificate gets the value of IntermediateRootCertificate for the instance
func (instance *DAServerConfiguration) GetPropertyIntermediateRootCertificate()(value bool, err error) { 
    retValue, err := instance.GetProperty("IntermediateRootCertificate")
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

// SetInternalIPv6Prefix sets the value of InternalIPv6Prefix for the instance
func (instance *DAServerConfiguration) SetPropertyInternalIPv6Prefix(value []string) (err error) { 
    return instance.SetProperty("InternalIPv6Prefix", (value))
}

// GetInternalIPv6Prefix gets the value of InternalIPv6Prefix for the instance
func (instance *DAServerConfiguration) GetPropertyInternalIPv6Prefix()(value []string, err error) { 
    retValue, err := instance.GetProperty("InternalIPv6Prefix")
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

// SetIPsecRootCertificate sets the value of IPsecRootCertificate for the instance
func (instance *DAServerConfiguration) SetPropertyIPsecRootCertificate(value []uint8) (err error) { 
    return instance.SetProperty("IPsecRootCertificate", (value))
}

// GetIPsecRootCertificate gets the value of IPsecRootCertificate for the instance
func (instance *DAServerConfiguration) GetPropertyIPsecRootCertificate()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("IPsecRootCertificate")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetIsNatDeployed sets the value of IsNatDeployed for the instance
func (instance *DAServerConfiguration) SetPropertyIsNatDeployed(value bool) (err error) { 
    return instance.SetProperty("IsNatDeployed", (value))
}

// GetIsNatDeployed gets the value of IsNatDeployed for the instance
func (instance *DAServerConfiguration) GetPropertyIsNatDeployed()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsNatDeployed")
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

// SetIsSingleNic sets the value of IsSingleNic for the instance
func (instance *DAServerConfiguration) SetPropertyIsSingleNic(value bool) (err error) { 
    return instance.SetProperty("IsSingleNic", (value))
}

// GetIsSingleNic gets the value of IsSingleNic for the instance
func (instance *DAServerConfiguration) GetPropertyIsSingleNic()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsSingleNic")
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

// SetTeredoState sets the value of TeredoState for the instance
func (instance *DAServerConfiguration) SetPropertyTeredoState(value string) (err error) { 
    return instance.SetProperty("TeredoState", (value))
}

// GetTeredoState gets the value of TeredoState for the instance
func (instance *DAServerConfiguration) GetPropertyTeredoState()(value string, err error) { 
    retValue, err := instance.GetProperty("TeredoState")
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

// SetUserAuthentication sets the value of UserAuthentication for the instance
func (instance *DAServerConfiguration) SetPropertyUserAuthentication(value string) (err error) { 
    return instance.SetProperty("UserAuthentication", (value))
}

// GetUserAuthentication gets the value of UserAuthentication for the instance
func (instance *DAServerConfiguration) GetPropertyUserAuthentication()(value string, err error) { 
    retValue, err := instance.GetProperty("UserAuthentication")
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

