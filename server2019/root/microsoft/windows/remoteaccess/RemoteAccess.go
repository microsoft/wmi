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

// RemoteAccess struct
type RemoteAccess struct { 
	*RemoteAccessCommon

	// 
	DAConfiguration DirectAccessConfiguration

	// 
	VpnConfiguration VirtualPrivateNetworkConfiguration
}

	func NewRemoteAccessEx1(instance *cim.WmiInstance) (newInstance *RemoteAccess, err error) {tmp, err := NewRemoteAccessCommonEx1(instance)
		
	if err != nil { return }
	newInstance = &RemoteAccess {
	RemoteAccessCommon: tmp,
	}
	return
	}
	

	func NewRemoteAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccess, err error) {tmp, err := NewRemoteAccessCommonEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccess {
	RemoteAccessCommon: tmp,
	}
	return
	}
	

// SetDAConfiguration sets the value of DAConfiguration for the instance
func (instance *RemoteAccess) SetPropertyDAConfiguration(value DirectAccessConfiguration) (err error) { 
    return instance.SetProperty("DAConfiguration", (value))
}

// GetDAConfiguration gets the value of DAConfiguration for the instance
func (instance *RemoteAccess) GetPropertyDAConfiguration()(value DirectAccessConfiguration, err error) { 
    retValue, err := instance.GetProperty("DAConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DirectAccessConfiguration); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DirectAccessConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DirectAccessConfiguration(valuetmp)

    return
}

// SetVpnConfiguration sets the value of VpnConfiguration for the instance
func (instance *RemoteAccess) SetPropertyVpnConfiguration(value VirtualPrivateNetworkConfiguration) (err error) { 
    return instance.SetProperty("VpnConfiguration", (value))
}

// GetVpnConfiguration gets the value of VpnConfiguration for the instance
func (instance *RemoteAccess) GetPropertyVpnConfiguration()(value VirtualPrivateNetworkConfiguration, err error) { 
    retValue, err := instance.GetProperty("VpnConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(VirtualPrivateNetworkConfiguration); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " VirtualPrivateNetworkConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = VirtualPrivateNetworkConfiguration(valuetmp)

    return
}

