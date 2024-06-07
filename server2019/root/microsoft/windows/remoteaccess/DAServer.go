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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DAServer struct
type DAServer struct { 
	*RemoteAccessCore

	// 
	ServerConfiguration DAServerConfiguration
}

	func NewDAServerEx1(instance *cim.WmiInstance) (newInstance *DAServer, err error) {tmp, err := NewRemoteAccessCoreEx1(instance)
		
	if err != nil { return }
	newInstance = &DAServer {
	RemoteAccessCore: tmp,
	}
	return
	}
	

	func NewDAServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAServer, err error) {tmp, err := NewRemoteAccessCoreEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAServer {
	RemoteAccessCore: tmp,
	}
	return
	}
	

// SetServerConfiguration sets the value of ServerConfiguration for the instance
func (instance *DAServer) SetPropertyServerConfiguration(value DAServerConfiguration) (err error) { 
    return instance.SetProperty("ServerConfiguration", (value))
}

// GetServerConfiguration gets the value of ServerConfiguration for the instance
func (instance *DAServer) GetPropertyServerConfiguration()(value DAServerConfiguration, err error) { 
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

