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

// RemoteAccessLoadBalancerNode struct
type RemoteAccessLoadBalancerNode struct { 
	*cim.WmiInstance

	// 
	HostName string

	// 
	StatusCode string

	// 
	VpnIPAddressAssignmentSetting VpnIPAddressAssignment
}

	func NewRemoteAccessLoadBalancerNodeEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessLoadBalancerNode, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessLoadBalancerNode {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessLoadBalancerNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessLoadBalancerNode, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessLoadBalancerNode {
	WmiInstance: tmp,
	}
	return
	}
	

// SetHostName sets the value of HostName for the instance
func (instance *RemoteAccessLoadBalancerNode) SetPropertyHostName(value string) (err error) { 
    return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *RemoteAccessLoadBalancerNode) GetPropertyHostName()(value string, err error) { 
    retValue, err := instance.GetProperty("HostName")
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

// SetStatusCode sets the value of StatusCode for the instance
func (instance *RemoteAccessLoadBalancerNode) SetPropertyStatusCode(value string) (err error) { 
    return instance.SetProperty("StatusCode", (value))
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *RemoteAccessLoadBalancerNode) GetPropertyStatusCode()(value string, err error) { 
    retValue, err := instance.GetProperty("StatusCode")
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

// SetVpnIPAddressAssignmentSetting sets the value of VpnIPAddressAssignmentSetting for the instance
func (instance *RemoteAccessLoadBalancerNode) SetPropertyVpnIPAddressAssignmentSetting(value VpnIPAddressAssignment) (err error) { 
    return instance.SetProperty("VpnIPAddressAssignmentSetting", (value))
}

// GetVpnIPAddressAssignmentSetting gets the value of VpnIPAddressAssignmentSetting for the instance
func (instance *RemoteAccessLoadBalancerNode) GetPropertyVpnIPAddressAssignmentSetting()(value VpnIPAddressAssignment, err error) { 
    retValue, err := instance.GetProperty("VpnIPAddressAssignmentSetting")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(VpnIPAddressAssignment); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " VpnIPAddressAssignment is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = VpnIPAddressAssignment(valuetmp)

    return
}

