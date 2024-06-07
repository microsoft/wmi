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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ThirdPartyVpnConnection struct
type ThirdPartyVpnConnection struct { 
	*VpnCommonConfig

	// 
	CustomConfiguration string

	// 
	PlugInApplicationID string

	// 
	VpnConfigurationXml string
}

	func NewThirdPartyVpnConnectionEx1(instance *cim.WmiInstance) (newInstance *ThirdPartyVpnConnection, err error) {tmp, err := NewVpnCommonConfigEx1(instance)
		
	if err != nil { return }
	newInstance = &ThirdPartyVpnConnection {
	VpnCommonConfig: tmp,
	}
	return
	}
	

	func NewThirdPartyVpnConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ThirdPartyVpnConnection, err error) {tmp, err := NewVpnCommonConfigEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ThirdPartyVpnConnection {
	VpnCommonConfig: tmp,
	}
	return
	}
	

// SetCustomConfiguration sets the value of CustomConfiguration for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyCustomConfiguration(value string) (err error) { 
    return instance.SetProperty("CustomConfiguration", (value))
}

// GetCustomConfiguration gets the value of CustomConfiguration for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyCustomConfiguration()(value string, err error) { 
    retValue, err := instance.GetProperty("CustomConfiguration")
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

// SetPlugInApplicationID sets the value of PlugInApplicationID for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyPlugInApplicationID(value string) (err error) { 
    return instance.SetProperty("PlugInApplicationID", (value))
}

// GetPlugInApplicationID gets the value of PlugInApplicationID for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyPlugInApplicationID()(value string, err error) { 
    retValue, err := instance.GetProperty("PlugInApplicationID")
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

// SetVpnConfigurationXml sets the value of VpnConfigurationXml for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyVpnConfigurationXml(value string) (err error) { 
    return instance.SetProperty("VpnConfigurationXml", (value))
}

// GetVpnConfigurationXml gets the value of VpnConfigurationXml for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyVpnConfigurationXml()(value string, err error) { 
    retValue, err := instance.GetProperty("VpnConfigurationXml")
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

