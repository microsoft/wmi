// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SamCompCreate_Start struct
type SamCompCreate_Start struct { 
	*SamCompCreate

	// Client Network Address
	Client string

	// Signature
	Sam string

	// Client SID
	Sid string

	// SamTraceVersion
	Version uint32
}

	func NewSamCompCreate_StartEx1(instance *cim.WmiInstance) (newInstance *SamCompCreate_Start, err error) {tmp, err := NewSamCompCreateEx1(instance)
		
	if err != nil { return }
	newInstance = &SamCompCreate_Start {
	SamCompCreate: tmp,
	}
	return
	}
	

	func NewSamCompCreate_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamCompCreate_Start, err error) {tmp, err := NewSamCompCreateEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamCompCreate_Start {
	SamCompCreate: tmp,
	}
	return
	}
	

// SetClient sets the value of Client for the instance
func (instance *SamCompCreate_Start) SetPropertyClient(value string) (err error) { 
    return instance.SetProperty("Client", (value))
}

// GetClient gets the value of Client for the instance
func (instance *SamCompCreate_Start) GetPropertyClient()(value string, err error) { 
    retValue, err := instance.GetProperty("Client")
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

// SetSam sets the value of Sam for the instance
func (instance *SamCompCreate_Start) SetPropertySam(value string) (err error) { 
    return instance.SetProperty("Sam", (value))
}

// GetSam gets the value of Sam for the instance
func (instance *SamCompCreate_Start) GetPropertySam()(value string, err error) { 
    retValue, err := instance.GetProperty("Sam")
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

// SetSid sets the value of Sid for the instance
func (instance *SamCompCreate_Start) SetPropertySid(value string) (err error) { 
    return instance.SetProperty("Sid", (value))
}

// GetSid gets the value of Sid for the instance
func (instance *SamCompCreate_Start) GetPropertySid()(value string, err error) { 
    retValue, err := instance.GetProperty("Sid")
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

// SetVersion sets the value of Version for the instance
func (instance *SamCompCreate_Start) SetPropertyVersion(value uint32) (err error) { 
    return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *SamCompCreate_Start) GetPropertyVersion()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Version")
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

