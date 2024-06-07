// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpFirewallSupport struct
type FtpFirewallSupport struct { 
	*EmbeddedObject

	// 
	ExternalIp4Address string
}

	func NewFtpFirewallSupportEx1(instance *cim.WmiInstance) (newInstance *FtpFirewallSupport, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpFirewallSupport {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpFirewallSupportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpFirewallSupport, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpFirewallSupport {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetExternalIp4Address sets the value of ExternalIp4Address for the instance
func (instance *FtpFirewallSupport) SetPropertyExternalIp4Address(value string) (err error) { 
    return instance.SetProperty("ExternalIp4Address", (value))
}

// GetExternalIp4Address gets the value of ExternalIp4Address for the instance
func (instance *FtpFirewallSupport) GetPropertyExternalIp4Address()(value string, err error) { 
    retValue, err := instance.GetProperty("ExternalIp4Address")
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

