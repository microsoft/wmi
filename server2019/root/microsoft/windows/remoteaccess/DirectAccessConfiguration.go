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

// DirectAccessConfiguration struct
type DirectAccessConfiguration struct { 
	*cim.WmiInstance

	// 
	AppServerPolicy DAAppServer

	// 
	ClientDnsConfiguration DAClientDnsConfiguration

	// 
	ClientPolicy DAClient

	// 
	MgmtServer []string

	// 
	NetworkLocationServerPolicy DANetworkLocationServer

	// 
	ServerConfiguration DAServerConfiguration
}

	func NewDirectAccessConfigurationEx1(instance *cim.WmiInstance) (newInstance *DirectAccessConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DirectAccessConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDirectAccessConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DirectAccessConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DirectAccessConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAppServerPolicy sets the value of AppServerPolicy for the instance
func (instance *DirectAccessConfiguration) SetPropertyAppServerPolicy(value DAAppServer) (err error) { 
    return instance.SetProperty("AppServerPolicy", (value))
}

// GetAppServerPolicy gets the value of AppServerPolicy for the instance
func (instance *DirectAccessConfiguration) GetPropertyAppServerPolicy()(value DAAppServer, err error) { 
    retValue, err := instance.GetProperty("AppServerPolicy")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DAAppServer); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DAAppServer is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DAAppServer(valuetmp)

    return
}

// SetClientDnsConfiguration sets the value of ClientDnsConfiguration for the instance
func (instance *DirectAccessConfiguration) SetPropertyClientDnsConfiguration(value DAClientDnsConfiguration) (err error) { 
    return instance.SetProperty("ClientDnsConfiguration", (value))
}

// GetClientDnsConfiguration gets the value of ClientDnsConfiguration for the instance
func (instance *DirectAccessConfiguration) GetPropertyClientDnsConfiguration()(value DAClientDnsConfiguration, err error) { 
    retValue, err := instance.GetProperty("ClientDnsConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DAClientDnsConfiguration); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DAClientDnsConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DAClientDnsConfiguration(valuetmp)

    return
}

// SetClientPolicy sets the value of ClientPolicy for the instance
func (instance *DirectAccessConfiguration) SetPropertyClientPolicy(value DAClient) (err error) { 
    return instance.SetProperty("ClientPolicy", (value))
}

// GetClientPolicy gets the value of ClientPolicy for the instance
func (instance *DirectAccessConfiguration) GetPropertyClientPolicy()(value DAClient, err error) { 
    retValue, err := instance.GetProperty("ClientPolicy")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DAClient); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DAClient is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DAClient(valuetmp)

    return
}

// SetMgmtServer sets the value of MgmtServer for the instance
func (instance *DirectAccessConfiguration) SetPropertyMgmtServer(value []string) (err error) { 
    return instance.SetProperty("MgmtServer", (value))
}

// GetMgmtServer gets the value of MgmtServer for the instance
func (instance *DirectAccessConfiguration) GetPropertyMgmtServer()(value []string, err error) { 
    retValue, err := instance.GetProperty("MgmtServer")
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

// SetNetworkLocationServerPolicy sets the value of NetworkLocationServerPolicy for the instance
func (instance *DirectAccessConfiguration) SetPropertyNetworkLocationServerPolicy(value DANetworkLocationServer) (err error) { 
    return instance.SetProperty("NetworkLocationServerPolicy", (value))
}

// GetNetworkLocationServerPolicy gets the value of NetworkLocationServerPolicy for the instance
func (instance *DirectAccessConfiguration) GetPropertyNetworkLocationServerPolicy()(value DANetworkLocationServer, err error) { 
    retValue, err := instance.GetProperty("NetworkLocationServerPolicy")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DANetworkLocationServer); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DANetworkLocationServer is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DANetworkLocationServer(valuetmp)

    return
}

// SetServerConfiguration sets the value of ServerConfiguration for the instance
func (instance *DirectAccessConfiguration) SetPropertyServerConfiguration(value DAServerConfiguration) (err error) { 
    return instance.SetProperty("ServerConfiguration", (value))
}

// GetServerConfiguration gets the value of ServerConfiguration for the instance
func (instance *DirectAccessConfiguration) GetPropertyServerConfiguration()(value DAServerConfiguration, err error) { 
    retValue, err := instance.GetProperty("ServerConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DAServerConfiguration); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DAServerConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DAServerConfiguration(valuetmp)

    return
}

