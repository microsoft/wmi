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

// RemoteAccessS2SConfiguration struct
type RemoteAccessS2SConfiguration struct { 
	*cim.WmiInstance

	// 
	Ipv4TransportInfo []uint8

	// 
	Ipv6TransportInfo []uint8

	// 
	Password string

	// 
	S2SInterfaceConfiguration interface{}

	// 
	SharedSecret string
}

	func NewRemoteAccessS2SConfigurationEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessS2SConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessS2SConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessS2SConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessS2SConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessS2SConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetIpv4TransportInfo sets the value of Ipv4TransportInfo for the instance
func (instance *RemoteAccessS2SConfiguration) SetPropertyIpv4TransportInfo(value []uint8) (err error) { 
    return instance.SetProperty("Ipv4TransportInfo", (value))
}

// GetIpv4TransportInfo gets the value of Ipv4TransportInfo for the instance
func (instance *RemoteAccessS2SConfiguration) GetPropertyIpv4TransportInfo()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Ipv4TransportInfo")
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

// SetIpv6TransportInfo sets the value of Ipv6TransportInfo for the instance
func (instance *RemoteAccessS2SConfiguration) SetPropertyIpv6TransportInfo(value []uint8) (err error) { 
    return instance.SetProperty("Ipv6TransportInfo", (value))
}

// GetIpv6TransportInfo gets the value of Ipv6TransportInfo for the instance
func (instance *RemoteAccessS2SConfiguration) GetPropertyIpv6TransportInfo()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Ipv6TransportInfo")
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

// SetPassword sets the value of Password for the instance
func (instance *RemoteAccessS2SConfiguration) SetPropertyPassword(value string) (err error) { 
    return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *RemoteAccessS2SConfiguration) GetPropertyPassword()(value string, err error) { 
    retValue, err := instance.GetProperty("Password")
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

// SetS2SInterfaceConfiguration sets the value of S2SInterfaceConfiguration for the instance
func (instance *RemoteAccessS2SConfiguration) SetPropertyS2SInterfaceConfiguration(value interface{}) (err error) { 
    return instance.SetProperty("S2SInterfaceConfiguration", (value))
}

// GetS2SInterfaceConfiguration gets the value of S2SInterfaceConfiguration for the instance
func (instance *RemoteAccessS2SConfiguration) GetPropertyS2SInterfaceConfiguration()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("S2SInterfaceConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(interface{}); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = interface{}(valuetmp)

    return
}

// SetSharedSecret sets the value of SharedSecret for the instance
func (instance *RemoteAccessS2SConfiguration) SetPropertySharedSecret(value string) (err error) { 
    return instance.SetProperty("SharedSecret", (value))
}

// GetSharedSecret gets the value of SharedSecret for the instance
func (instance *RemoteAccessS2SConfiguration) GetPropertySharedSecret()(value string, err error) { 
    retValue, err := instance.GetProperty("SharedSecret")
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

