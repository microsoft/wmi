// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DNSClientDohServerAddress struct
type MSFT_DNSClientDohServerAddress struct { 
	*CIM_RemoteServiceAccessPoint

	// 751
	AllowFallbackToUdp bool

	// 752
	AutoUpgrade bool

	// 750
	DohTemplate string

	// 749
	ServerAddress string
}

	func NewMSFT_DNSClientDohServerAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_DNSClientDohServerAddress, err error) {tmp, err := NewCIM_RemoteServiceAccessPointEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_DNSClientDohServerAddress {
	CIM_RemoteServiceAccessPoint: tmp,
	}
	return
	}
	

	func NewMSFT_DNSClientDohServerAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DNSClientDohServerAddress, err error) {tmp, err := NewCIM_RemoteServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DNSClientDohServerAddress {
	CIM_RemoteServiceAccessPoint: tmp,
	}
	return
	}
	

// SetAllowFallbackToUdp sets the value of AllowFallbackToUdp for the instance
func (instance *MSFT_DNSClientDohServerAddress) SetPropertyAllowFallbackToUdp(value bool) (err error) { 
    return instance.SetProperty("AllowFallbackToUdp", (value))
}

// GetAllowFallbackToUdp gets the value of AllowFallbackToUdp for the instance
func (instance *MSFT_DNSClientDohServerAddress) GetPropertyAllowFallbackToUdp()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowFallbackToUdp")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetAutoUpgrade sets the value of AutoUpgrade for the instance
func (instance *MSFT_DNSClientDohServerAddress) SetPropertyAutoUpgrade(value bool) (err error) { 
    return instance.SetProperty("AutoUpgrade", (value))
}

// GetAutoUpgrade gets the value of AutoUpgrade for the instance
func (instance *MSFT_DNSClientDohServerAddress) GetPropertyAutoUpgrade()(value bool, err error) { 
    retValue, err := instance.GetProperty("AutoUpgrade")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDohTemplate sets the value of DohTemplate for the instance
func (instance *MSFT_DNSClientDohServerAddress) SetPropertyDohTemplate(value string) (err error) { 
    return instance.SetProperty("DohTemplate", (value))
}

// GetDohTemplate gets the value of DohTemplate for the instance
func (instance *MSFT_DNSClientDohServerAddress) GetPropertyDohTemplate()(value string, err error) { 
    retValue, err := instance.GetProperty("DohTemplate")
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

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *MSFT_DNSClientDohServerAddress) SetPropertyServerAddress(value string) (err error) { 
    return instance.SetProperty("ServerAddress", (value))
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *MSFT_DNSClientDohServerAddress) GetPropertyServerAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("ServerAddress")
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

