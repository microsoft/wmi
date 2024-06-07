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

// DAClientDnsConfiguration struct
type DAClientDnsConfiguration struct { 
	*cim.WmiInstance

	// 
	NrptEntry []DnsClientNrptRule

	// 
	NrptGlobalSettings DnsClientNrptGlobal
}

	func NewDAClientDnsConfigurationEx1(instance *cim.WmiInstance) (newInstance *DAClientDnsConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DAClientDnsConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDAClientDnsConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAClientDnsConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAClientDnsConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetNrptEntry sets the value of NrptEntry for the instance
func (instance *DAClientDnsConfiguration) SetPropertyNrptEntry(value []DnsClientNrptRule) (err error) { 
    return instance.SetProperty("NrptEntry", (value))
}

// GetNrptEntry gets the value of NrptEntry for the instance
func (instance *DAClientDnsConfiguration) GetPropertyNrptEntry()(value []DnsClientNrptRule, err error) { 
    retValue, err := instance.GetProperty("NrptEntry")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(DnsClientNrptRule); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " DnsClientNrptRule is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, DnsClientNrptRule(valuetmp))
    }

    return
}

// SetNrptGlobalSettings sets the value of NrptGlobalSettings for the instance
func (instance *DAClientDnsConfiguration) SetPropertyNrptGlobalSettings(value DnsClientNrptGlobal) (err error) { 
    return instance.SetProperty("NrptGlobalSettings", (value))
}

// GetNrptGlobalSettings gets the value of NrptGlobalSettings for the instance
func (instance *DAClientDnsConfiguration) GetPropertyNrptGlobalSettings()(value DnsClientNrptGlobal, err error) { 
    retValue, err := instance.GetProperty("NrptGlobalSettings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DnsClientNrptGlobal); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DnsClientNrptGlobal is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DnsClientNrptGlobal(valuetmp)

    return
}

