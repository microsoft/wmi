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

// RemoteAccessBgpPeerConfig struct
type RemoteAccessBgpPeerConfig struct { 
	*cim.WmiInstance

	// 
	Peer []string

	// 
	PeerStaus []bool

	// 
	RoutingDomain string
}

	func NewRemoteAccessBgpPeerConfigEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessBgpPeerConfig, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessBgpPeerConfig {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessBgpPeerConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessBgpPeerConfig, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessBgpPeerConfig {
	WmiInstance: tmp,
	}
	return
	}
	

// SetPeer sets the value of Peer for the instance
func (instance *RemoteAccessBgpPeerConfig) SetPropertyPeer(value []string) (err error) { 
    return instance.SetProperty("Peer", (value))
}

// GetPeer gets the value of Peer for the instance
func (instance *RemoteAccessBgpPeerConfig) GetPropertyPeer()(value []string, err error) { 
    retValue, err := instance.GetProperty("Peer")
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

// SetPeerStaus sets the value of PeerStaus for the instance
func (instance *RemoteAccessBgpPeerConfig) SetPropertyPeerStaus(value []bool) (err error) { 
    return instance.SetProperty("PeerStaus", (value))
}

// GetPeerStaus gets the value of PeerStaus for the instance
func (instance *RemoteAccessBgpPeerConfig) GetPropertyPeerStaus()(value []bool, err error) { 
    retValue, err := instance.GetProperty("PeerStaus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(bool); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, bool(valuetmp))
    }

    return
}

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *RemoteAccessBgpPeerConfig) SetPropertyRoutingDomain(value string) (err error) { 
    return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *RemoteAccessBgpPeerConfig) GetPropertyRoutingDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomain")
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

