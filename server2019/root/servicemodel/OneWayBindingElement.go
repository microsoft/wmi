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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// OneWayBindingElement struct
type OneWayBindingElement struct { 
	*BindingElement

	// The channel pool settings.
	ChannelPoolSettings ChannelPoolSettings

	// The maximum number of accepted channels.
	MaxAcceptedChannels int32

	// Whether the packet is routable.
	PacketRoutable bool
}

	func NewOneWayBindingElementEx1(instance *cim.WmiInstance) (newInstance *OneWayBindingElement, err error) {tmp, err := NewBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &OneWayBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

	func NewOneWayBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OneWayBindingElement, err error) {tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OneWayBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

// SetChannelPoolSettings sets the value of ChannelPoolSettings for the instance
func (instance *OneWayBindingElement) SetPropertyChannelPoolSettings(value ChannelPoolSettings) (err error) { 
    return instance.SetProperty("ChannelPoolSettings", (value))
}

// GetChannelPoolSettings gets the value of ChannelPoolSettings for the instance
func (instance *OneWayBindingElement) GetPropertyChannelPoolSettings()(value ChannelPoolSettings, err error) { 
    retValue, err := instance.GetProperty("ChannelPoolSettings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ChannelPoolSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ChannelPoolSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ChannelPoolSettings(valuetmp)

    return
}

// SetMaxAcceptedChannels sets the value of MaxAcceptedChannels for the instance
func (instance *OneWayBindingElement) SetPropertyMaxAcceptedChannels(value int32) (err error) { 
    return instance.SetProperty("MaxAcceptedChannels", (value))
}

// GetMaxAcceptedChannels gets the value of MaxAcceptedChannels for the instance
func (instance *OneWayBindingElement) GetPropertyMaxAcceptedChannels()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxAcceptedChannels")
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

// SetPacketRoutable sets the value of PacketRoutable for the instance
func (instance *OneWayBindingElement) SetPropertyPacketRoutable(value bool) (err error) { 
    return instance.SetProperty("PacketRoutable", (value))
}

// GetPacketRoutable gets the value of PacketRoutable for the instance
func (instance *OneWayBindingElement) GetPropertyPacketRoutable()(value bool, err error) { 
    retValue, err := instance.GetProperty("PacketRoutable")
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

