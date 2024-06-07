// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// VpnConnectionTriggerApplication struct
type VpnConnectionTriggerApplication struct { 
	*cim.WmiInstance

	// 
	ApplicationID []string

	// 
	ConnectionName string
}

	func NewVpnConnectionTriggerApplicationEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionTriggerApplication, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &VpnConnectionTriggerApplication {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewVpnConnectionTriggerApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnConnectionTriggerApplication, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnConnectionTriggerApplication {
	WmiInstance: tmp,
	}
	return
	}
	

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *VpnConnectionTriggerApplication) SetPropertyApplicationID(value []string) (err error) { 
    return instance.SetProperty("ApplicationID", (value))
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *VpnConnectionTriggerApplication) GetPropertyApplicationID()(value []string, err error) { 
    retValue, err := instance.GetProperty("ApplicationID")
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

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerApplication) SetPropertyConnectionName(value string) (err error) { 
    return instance.SetProperty("ConnectionName", (value))
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerApplication) GetPropertyConnectionName()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionName")
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

