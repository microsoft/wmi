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

// MSFT_NetEventVmSwitchProvider struct
type MSFT_NetEventVmSwitchProvider struct { 
	*MSFT_NetEventProviderBase

	// 
	PortIds []uint32

	// 
	SwitchName string
}

	func NewMSFT_NetEventVmSwitchProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventVmSwitchProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventVmSwitchProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

	func NewMSFT_NetEventVmSwitchProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetEventVmSwitchProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventVmSwitchProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

// SetPortIds sets the value of PortIds for the instance
func (instance *MSFT_NetEventVmSwitchProvider) SetPropertyPortIds(value []uint32) (err error) { 
    return instance.SetProperty("PortIds", (value))
}

// GetPortIds gets the value of PortIds for the instance
func (instance *MSFT_NetEventVmSwitchProvider) GetPropertyPortIds()(value []uint32, err error) { 
    retValue, err := instance.GetProperty("PortIds")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint32(valuetmp))
    }

    return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmSwitchProvider) SetPropertySwitchName(value string) (err error) { 
    return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmSwitchProvider) GetPropertySwitchName()(value string, err error) { 
    retValue, err := instance.GetProperty("SwitchName")
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

