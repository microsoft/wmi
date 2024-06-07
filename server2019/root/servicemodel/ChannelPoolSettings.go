// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ChannelPoolSettings struct
type ChannelPoolSettings struct { 
	*cim.WmiInstance

	// The maximum time the connection can be idle before being disconnected. 
	IdleTimeout string

	// The maximum time for a lease operation to complete before timing out.
	LeaseTimeout string

	// The maximum number of outbound channels per endpoint.
	MaxOutboundChannelsPerEndpoint int32
}

	func NewChannelPoolSettingsEx1(instance *cim.WmiInstance) (newInstance *ChannelPoolSettings, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ChannelPoolSettings {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewChannelPoolSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ChannelPoolSettings, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ChannelPoolSettings {
	WmiInstance: tmp,
	}
	return
	}
	

// SetIdleTimeout sets the value of IdleTimeout for the instance
func (instance *ChannelPoolSettings) SetPropertyIdleTimeout(value string) (err error) { 
    return instance.SetProperty("IdleTimeout", (value))
}

// GetIdleTimeout gets the value of IdleTimeout for the instance
func (instance *ChannelPoolSettings) GetPropertyIdleTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("IdleTimeout")
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

// SetLeaseTimeout sets the value of LeaseTimeout for the instance
func (instance *ChannelPoolSettings) SetPropertyLeaseTimeout(value string) (err error) { 
    return instance.SetProperty("LeaseTimeout", (value))
}

// GetLeaseTimeout gets the value of LeaseTimeout for the instance
func (instance *ChannelPoolSettings) GetPropertyLeaseTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("LeaseTimeout")
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

// SetMaxOutboundChannelsPerEndpoint sets the value of MaxOutboundChannelsPerEndpoint for the instance
func (instance *ChannelPoolSettings) SetPropertyMaxOutboundChannelsPerEndpoint(value int32) (err error) { 
    return instance.SetProperty("MaxOutboundChannelsPerEndpoint", (value))
}

// GetMaxOutboundChannelsPerEndpoint gets the value of MaxOutboundChannelsPerEndpoint for the instance
func (instance *ChannelPoolSettings) GetPropertyMaxOutboundChannelsPerEndpoint()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxOutboundChannelsPerEndpoint")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

