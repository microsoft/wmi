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

// VpnAuthProtocol struct
type VpnAuthProtocol struct { 
	*cim.WmiInstance

	// 
	CertificateAdvertised []uint8

	// 
	CertificateEKUsToAccept []string

	// 
	RootCertificateNameToAccept []uint8

	// 
	TunnelAuthProtocolsAdvertised string

	// 
	UserAuthProtocolAccepted []string
}

	func NewVpnAuthProtocolEx1(instance *cim.WmiInstance) (newInstance *VpnAuthProtocol, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &VpnAuthProtocol {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewVpnAuthProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnAuthProtocol, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnAuthProtocol {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCertificateAdvertised sets the value of CertificateAdvertised for the instance
func (instance *VpnAuthProtocol) SetPropertyCertificateAdvertised(value []uint8) (err error) { 
    return instance.SetProperty("CertificateAdvertised", (value))
}

// GetCertificateAdvertised gets the value of CertificateAdvertised for the instance
func (instance *VpnAuthProtocol) GetPropertyCertificateAdvertised()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("CertificateAdvertised")
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

// SetCertificateEKUsToAccept sets the value of CertificateEKUsToAccept for the instance
func (instance *VpnAuthProtocol) SetPropertyCertificateEKUsToAccept(value []string) (err error) { 
    return instance.SetProperty("CertificateEKUsToAccept", (value))
}

// GetCertificateEKUsToAccept gets the value of CertificateEKUsToAccept for the instance
func (instance *VpnAuthProtocol) GetPropertyCertificateEKUsToAccept()(value []string, err error) { 
    retValue, err := instance.GetProperty("CertificateEKUsToAccept")
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

// SetRootCertificateNameToAccept sets the value of RootCertificateNameToAccept for the instance
func (instance *VpnAuthProtocol) SetPropertyRootCertificateNameToAccept(value []uint8) (err error) { 
    return instance.SetProperty("RootCertificateNameToAccept", (value))
}

// GetRootCertificateNameToAccept gets the value of RootCertificateNameToAccept for the instance
func (instance *VpnAuthProtocol) GetPropertyRootCertificateNameToAccept()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("RootCertificateNameToAccept")
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

// SetTunnelAuthProtocolsAdvertised sets the value of TunnelAuthProtocolsAdvertised for the instance
func (instance *VpnAuthProtocol) SetPropertyTunnelAuthProtocolsAdvertised(value string) (err error) { 
    return instance.SetProperty("TunnelAuthProtocolsAdvertised", (value))
}

// GetTunnelAuthProtocolsAdvertised gets the value of TunnelAuthProtocolsAdvertised for the instance
func (instance *VpnAuthProtocol) GetPropertyTunnelAuthProtocolsAdvertised()(value string, err error) { 
    retValue, err := instance.GetProperty("TunnelAuthProtocolsAdvertised")
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

// SetUserAuthProtocolAccepted sets the value of UserAuthProtocolAccepted for the instance
func (instance *VpnAuthProtocol) SetPropertyUserAuthProtocolAccepted(value []string) (err error) { 
    return instance.SetProperty("UserAuthProtocolAccepted", (value))
}

// GetUserAuthProtocolAccepted gets the value of UserAuthProtocolAccepted for the instance
func (instance *VpnAuthProtocol) GetPropertyUserAuthProtocolAccepted()(value []string, err error) { 
    retValue, err := instance.GetProperty("UserAuthProtocolAccepted")
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

