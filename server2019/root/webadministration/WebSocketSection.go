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

// WebSocketSection struct
type WebSocketSection struct { 
	*ConfigurationSection

	// 
	Enabled bool

	// 
	PingInterval string

	// 
	ReadBufferBlockSize uint32

	// 
	ReceiveBufferLimit uint32
}

	func NewWebSocketSectionEx1(instance *cim.WmiInstance) (newInstance *WebSocketSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &WebSocketSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewWebSocketSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WebSocketSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WebSocketSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *WebSocketSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *WebSocketSection) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetPingInterval sets the value of PingInterval for the instance
func (instance *WebSocketSection) SetPropertyPingInterval(value string) (err error) { 
    return instance.SetProperty("PingInterval", (value))
}

// GetPingInterval gets the value of PingInterval for the instance
func (instance *WebSocketSection) GetPropertyPingInterval()(value string, err error) { 
    retValue, err := instance.GetProperty("PingInterval")
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

// SetReadBufferBlockSize sets the value of ReadBufferBlockSize for the instance
func (instance *WebSocketSection) SetPropertyReadBufferBlockSize(value uint32) (err error) { 
    return instance.SetProperty("ReadBufferBlockSize", (value))
}

// GetReadBufferBlockSize gets the value of ReadBufferBlockSize for the instance
func (instance *WebSocketSection) GetPropertyReadBufferBlockSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadBufferBlockSize")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReceiveBufferLimit sets the value of ReceiveBufferLimit for the instance
func (instance *WebSocketSection) SetPropertyReceiveBufferLimit(value uint32) (err error) { 
    return instance.SetProperty("ReceiveBufferLimit", (value))
}

// GetReceiveBufferLimit gets the value of ReceiveBufferLimit for the instance
func (instance *WebSocketSection) GetPropertyReceiveBufferLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReceiveBufferLimit")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

