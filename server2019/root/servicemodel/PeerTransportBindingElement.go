// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// PeerTransportBindingElement struct
type PeerTransportBindingElement struct { 
	*TransportBindingElement

	// The IP address on which the peer node will listen for messages.
	ListenIPAddress string

	// The network interface port on which this binding will process peer channel messages.
	Port int32

	// Peer transport security settings.
	Security PeerSecuritySettings
}

	func NewPeerTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *PeerTransportBindingElement, err error) {tmp, err := NewTransportBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &PeerTransportBindingElement {
	TransportBindingElement: tmp,
	}
	return
	}
	

	func NewPeerTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PeerTransportBindingElement, err error) {tmp, err := NewTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PeerTransportBindingElement {
	TransportBindingElement: tmp,
	}
	return
	}
	

// SetListenIPAddress sets the value of ListenIPAddress for the instance
func (instance *PeerTransportBindingElement) SetPropertyListenIPAddress(value string) (err error) { 
    return instance.SetProperty("ListenIPAddress", (value))
}

// GetListenIPAddress gets the value of ListenIPAddress for the instance
func (instance *PeerTransportBindingElement) GetPropertyListenIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("ListenIPAddress")
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

// SetPort sets the value of Port for the instance
func (instance *PeerTransportBindingElement) SetPropertyPort(value int32) (err error) { 
    return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *PeerTransportBindingElement) GetPropertyPort()(value int32, err error) { 
    retValue, err := instance.GetProperty("Port")
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

// SetSecurity sets the value of Security for the instance
func (instance *PeerTransportBindingElement) SetPropertySecurity(value PeerSecuritySettings) (err error) { 
    return instance.SetProperty("Security", (value))
}

// GetSecurity gets the value of Security for the instance
func (instance *PeerTransportBindingElement) GetPropertySecurity()(value PeerSecuritySettings, err error) { 
    retValue, err := instance.GetProperty("Security")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(PeerSecuritySettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " PeerSecuritySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PeerSecuritySettings(valuetmp)

    return
}

