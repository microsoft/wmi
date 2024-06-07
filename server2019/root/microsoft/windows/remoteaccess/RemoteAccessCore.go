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

// RemoteAccessCore struct
type RemoteAccessCore struct { 
	*cim.WmiInstance

	// 
	InternalInterface string

	// 
	InternetInterface string

	// 
	SslCertificate []uint8
}

	func NewRemoteAccessCoreEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessCore, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessCore {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessCoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessCore, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessCore {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInternalInterface sets the value of InternalInterface for the instance
func (instance *RemoteAccessCore) SetPropertyInternalInterface(value string) (err error) { 
    return instance.SetProperty("InternalInterface", (value))
}

// GetInternalInterface gets the value of InternalInterface for the instance
func (instance *RemoteAccessCore) GetPropertyInternalInterface()(value string, err error) { 
    retValue, err := instance.GetProperty("InternalInterface")
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

// SetInternetInterface sets the value of InternetInterface for the instance
func (instance *RemoteAccessCore) SetPropertyInternetInterface(value string) (err error) { 
    return instance.SetProperty("InternetInterface", (value))
}

// GetInternetInterface gets the value of InternetInterface for the instance
func (instance *RemoteAccessCore) GetPropertyInternetInterface()(value string, err error) { 
    retValue, err := instance.GetProperty("InternetInterface")
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

// SetSslCertificate sets the value of SslCertificate for the instance
func (instance *RemoteAccessCore) SetPropertySslCertificate(value []uint8) (err error) { 
    return instance.SetProperty("SslCertificate", (value))
}

// GetSslCertificate gets the value of SslCertificate for the instance
func (instance *RemoteAccessCore) GetPropertySslCertificate()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("SslCertificate")
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

