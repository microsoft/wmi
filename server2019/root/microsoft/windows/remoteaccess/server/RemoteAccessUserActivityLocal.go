// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessUserActivityLocal struct
type RemoteAccessUserActivityLocal struct { 
	*cim.WmiInstance

	// 
	ProtocolID uint32

	// 
	ServerIPAddress string

	// 
	ServerPort uint32
}

	func NewRemoteAccessUserActivityLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessUserActivityLocal, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessUserActivityLocal {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessUserActivityLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessUserActivityLocal, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessUserActivityLocal {
	WmiInstance: tmp,
	}
	return
	}
	

// SetProtocolID sets the value of ProtocolID for the instance
func (instance *RemoteAccessUserActivityLocal) SetPropertyProtocolID(value uint32) (err error) { 
    return instance.SetProperty("ProtocolID", (value))
}

// GetProtocolID gets the value of ProtocolID for the instance
func (instance *RemoteAccessUserActivityLocal) GetPropertyProtocolID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProtocolID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetServerIPAddress sets the value of ServerIPAddress for the instance
func (instance *RemoteAccessUserActivityLocal) SetPropertyServerIPAddress(value string) (err error) { 
    return instance.SetProperty("ServerIPAddress", (value))
}

// GetServerIPAddress gets the value of ServerIPAddress for the instance
func (instance *RemoteAccessUserActivityLocal) GetPropertyServerIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("ServerIPAddress")
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

// SetServerPort sets the value of ServerPort for the instance
func (instance *RemoteAccessUserActivityLocal) SetPropertyServerPort(value uint32) (err error) { 
    return instance.SetProperty("ServerPort", (value))
}

// GetServerPort gets the value of ServerPort for the instance
func (instance *RemoteAccessUserActivityLocal) GetPropertyServerPort()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ServerPort")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

