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

// MSFT_NetTransportFilter struct
type MSFT_NetTransportFilter struct { 
	*CIM_FilterEntryBase

	// 
	DestinationPrefix string

	// 
	LocalPortEnd uint16

	// 
	LocalPortStart uint16

	// 
	Protocol uint16

	// 
	RemotePortEnd uint16

	// 
	RemotePortStart uint16

	// 
	SettingName string
}

	func NewMSFT_NetTransportFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTransportFilter, err error) {tmp, err := NewCIM_FilterEntryBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetTransportFilter {
	CIM_FilterEntryBase: tmp,
	}
	return
	}
	

	func NewMSFT_NetTransportFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetTransportFilter, err error) {tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetTransportFilter {
	CIM_FilterEntryBase: tmp,
	}
	return
	}
	

// SetDestinationPrefix sets the value of DestinationPrefix for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyDestinationPrefix(value string) (err error) { 
    return instance.SetProperty("DestinationPrefix", (value))
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyDestinationPrefix()(value string, err error) { 
    retValue, err := instance.GetProperty("DestinationPrefix")
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

// SetLocalPortEnd sets the value of LocalPortEnd for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyLocalPortEnd(value uint16) (err error) { 
    return instance.SetProperty("LocalPortEnd", (value))
}

// GetLocalPortEnd gets the value of LocalPortEnd for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyLocalPortEnd()(value uint16, err error) { 
    retValue, err := instance.GetProperty("LocalPortEnd")
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

// SetLocalPortStart sets the value of LocalPortStart for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyLocalPortStart(value uint16) (err error) { 
    return instance.SetProperty("LocalPortStart", (value))
}

// GetLocalPortStart gets the value of LocalPortStart for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyLocalPortStart()(value uint16, err error) { 
    retValue, err := instance.GetProperty("LocalPortStart")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyProtocol(value uint16) (err error) { 
    return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyProtocol()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Protocol")
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

// SetRemotePortEnd sets the value of RemotePortEnd for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyRemotePortEnd(value uint16) (err error) { 
    return instance.SetProperty("RemotePortEnd", (value))
}

// GetRemotePortEnd gets the value of RemotePortEnd for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyRemotePortEnd()(value uint16, err error) { 
    retValue, err := instance.GetProperty("RemotePortEnd")
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

// SetRemotePortStart sets the value of RemotePortStart for the instance
func (instance *MSFT_NetTransportFilter) SetPropertyRemotePortStart(value uint16) (err error) { 
    return instance.SetProperty("RemotePortStart", (value))
}

// GetRemotePortStart gets the value of RemotePortStart for the instance
func (instance *MSFT_NetTransportFilter) GetPropertyRemotePortStart()(value uint16, err error) { 
    retValue, err := instance.GetProperty("RemotePortStart")
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

// SetSettingName sets the value of SettingName for the instance
func (instance *MSFT_NetTransportFilter) SetPropertySettingName(value string) (err error) { 
    return instance.SetProperty("SettingName", (value))
}

// GetSettingName gets the value of SettingName for the instance
func (instance *MSFT_NetTransportFilter) GetPropertySettingName()(value string, err error) { 
    retValue, err := instance.GetProperty("SettingName")
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

