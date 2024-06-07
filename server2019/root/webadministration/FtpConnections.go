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

// FtpConnections struct
type FtpConnections struct { 
	*EmbeddedObject

	// 
	ControlChannelTimeout int32

	// 
	DataChannelTimeout int32

	// 
	DisableSocketPooling bool

	// 
	MaxBandwidth uint32

	// 
	MaxConnections uint32

	// 
	MinBytesPerSecond int32

	// 
	ResetOnMaxConnections bool

	// 
	ServerListenBacklog int32

	// 
	UnauthenticatedTimeout int32
}

	func NewFtpConnectionsEx1(instance *cim.WmiInstance) (newInstance *FtpConnections, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpConnections {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpConnectionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpConnections, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpConnections {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetControlChannelTimeout sets the value of ControlChannelTimeout for the instance
func (instance *FtpConnections) SetPropertyControlChannelTimeout(value int32) (err error) { 
    return instance.SetProperty("ControlChannelTimeout", (value))
}

// GetControlChannelTimeout gets the value of ControlChannelTimeout for the instance
func (instance *FtpConnections) GetPropertyControlChannelTimeout()(value int32, err error) { 
    retValue, err := instance.GetProperty("ControlChannelTimeout")
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

// SetDataChannelTimeout sets the value of DataChannelTimeout for the instance
func (instance *FtpConnections) SetPropertyDataChannelTimeout(value int32) (err error) { 
    return instance.SetProperty("DataChannelTimeout", (value))
}

// GetDataChannelTimeout gets the value of DataChannelTimeout for the instance
func (instance *FtpConnections) GetPropertyDataChannelTimeout()(value int32, err error) { 
    retValue, err := instance.GetProperty("DataChannelTimeout")
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

// SetDisableSocketPooling sets the value of DisableSocketPooling for the instance
func (instance *FtpConnections) SetPropertyDisableSocketPooling(value bool) (err error) { 
    return instance.SetProperty("DisableSocketPooling", (value))
}

// GetDisableSocketPooling gets the value of DisableSocketPooling for the instance
func (instance *FtpConnections) GetPropertyDisableSocketPooling()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisableSocketPooling")
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

// SetMaxBandwidth sets the value of MaxBandwidth for the instance
func (instance *FtpConnections) SetPropertyMaxBandwidth(value uint32) (err error) { 
    return instance.SetProperty("MaxBandwidth", (value))
}

// GetMaxBandwidth gets the value of MaxBandwidth for the instance
func (instance *FtpConnections) GetPropertyMaxBandwidth()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxBandwidth")
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

// SetMaxConnections sets the value of MaxConnections for the instance
func (instance *FtpConnections) SetPropertyMaxConnections(value uint32) (err error) { 
    return instance.SetProperty("MaxConnections", (value))
}

// GetMaxConnections gets the value of MaxConnections for the instance
func (instance *FtpConnections) GetPropertyMaxConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxConnections")
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

// SetMinBytesPerSecond sets the value of MinBytesPerSecond for the instance
func (instance *FtpConnections) SetPropertyMinBytesPerSecond(value int32) (err error) { 
    return instance.SetProperty("MinBytesPerSecond", (value))
}

// GetMinBytesPerSecond gets the value of MinBytesPerSecond for the instance
func (instance *FtpConnections) GetPropertyMinBytesPerSecond()(value int32, err error) { 
    retValue, err := instance.GetProperty("MinBytesPerSecond")
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

// SetResetOnMaxConnections sets the value of ResetOnMaxConnections for the instance
func (instance *FtpConnections) SetPropertyResetOnMaxConnections(value bool) (err error) { 
    return instance.SetProperty("ResetOnMaxConnections", (value))
}

// GetResetOnMaxConnections gets the value of ResetOnMaxConnections for the instance
func (instance *FtpConnections) GetPropertyResetOnMaxConnections()(value bool, err error) { 
    retValue, err := instance.GetProperty("ResetOnMaxConnections")
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

// SetServerListenBacklog sets the value of ServerListenBacklog for the instance
func (instance *FtpConnections) SetPropertyServerListenBacklog(value int32) (err error) { 
    return instance.SetProperty("ServerListenBacklog", (value))
}

// GetServerListenBacklog gets the value of ServerListenBacklog for the instance
func (instance *FtpConnections) GetPropertyServerListenBacklog()(value int32, err error) { 
    retValue, err := instance.GetProperty("ServerListenBacklog")
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

// SetUnauthenticatedTimeout sets the value of UnauthenticatedTimeout for the instance
func (instance *FtpConnections) SetPropertyUnauthenticatedTimeout(value int32) (err error) { 
    return instance.SetProperty("UnauthenticatedTimeout", (value))
}

// GetUnauthenticatedTimeout gets the value of UnauthenticatedTimeout for the instance
func (instance *FtpConnections) GetPropertyUnauthenticatedTimeout()(value int32, err error) { 
    retValue, err := instance.GetProperty("UnauthenticatedTimeout")
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

