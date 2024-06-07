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

// MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn struct
type MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn struct { 
	*MSFT_NetAdapterPowerManagement_WakePattern

	// 
	DestinationAddress string

	// 
	DestinationPort uint16

	// 
	SourceAddress string

	// 
	SourcePort uint16
}

	func NewMSFT_NetAdapterPowerManagement_WakePattern_TcpSynEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn, err error) {tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn {
	MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
	}
	

	func NewMSFT_NetAdapterPowerManagement_WakePattern_TcpSynEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn, err error) {tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn {
	MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
	}
	

// SetDestinationAddress sets the value of DestinationAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertyDestinationAddress(value string) (err error) { 
    return instance.SetProperty("DestinationAddress", (value))
}

// GetDestinationAddress gets the value of DestinationAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertyDestinationAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("DestinationAddress")
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

// SetDestinationPort sets the value of DestinationPort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertyDestinationPort(value uint16) (err error) { 
    return instance.SetProperty("DestinationPort", (value))
}

// GetDestinationPort gets the value of DestinationPort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertyDestinationPort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("DestinationPort")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetSourceAddress sets the value of SourceAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertySourceAddress(value string) (err error) { 
    return instance.SetProperty("SourceAddress", (value))
}

// GetSourceAddress gets the value of SourceAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertySourceAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("SourceAddress")
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

// SetSourcePort sets the value of SourcePort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) SetPropertySourcePort(value uint16) (err error) { 
    return instance.SetProperty("SourcePort", (value))
}

// GetSourcePort gets the value of SourcePort for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_TcpSyn) GetPropertySourcePort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("SourcePort")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

