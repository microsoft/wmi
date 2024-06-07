// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// HNet_ConnectionPortMapping2 struct
type HNet_ConnectionPortMapping2 struct { 
	*cim.WmiInstance

	// 
	Connection HNet_Connection

	// 
	Enabled bool

	// 
	NameActive bool

	// 
	Protocol HNet_PortMappingProtocol

	// 
	TargetIPAddress uint32

	// 
	TargetName string

	// 
	TargetPort uint16
}

	func NewHNet_ConnectionPortMapping2Ex1(instance *cim.WmiInstance) (newInstance *HNet_ConnectionPortMapping2, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &HNet_ConnectionPortMapping2 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewHNet_ConnectionPortMapping2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HNet_ConnectionPortMapping2, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HNet_ConnectionPortMapping2 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyConnection(value HNet_Connection) (err error) { 
    return instance.SetProperty("Connection", (value))
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyConnection()(value HNet_Connection, err error) { 
    retValue, err := instance.GetProperty("Connection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HNet_Connection); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HNet_Connection(valuetmp)

    return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyEnabled()(value bool, err error) { 
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

// SetNameActive sets the value of NameActive for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyNameActive(value bool) (err error) { 
    return instance.SetProperty("NameActive", (value))
}

// GetNameActive gets the value of NameActive for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyNameActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("NameActive")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyProtocol(value HNet_PortMappingProtocol) (err error) { 
    return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyProtocol()(value HNet_PortMappingProtocol, err error) { 
    retValue, err := instance.GetProperty("Protocol")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HNet_PortMappingProtocol); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HNet_PortMappingProtocol is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HNet_PortMappingProtocol(valuetmp)

    return
}

// SetTargetIPAddress sets the value of TargetIPAddress for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetIPAddress(value uint32) (err error) { 
    return instance.SetProperty("TargetIPAddress", (value))
}

// GetTargetIPAddress gets the value of TargetIPAddress for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetIPAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TargetIPAddress")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetName(value string) (err error) { 
    return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetName()(value string, err error) { 
    retValue, err := instance.GetProperty("TargetName")
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

// SetTargetPort sets the value of TargetPort for the instance
func (instance *HNet_ConnectionPortMapping2) SetPropertyTargetPort(value uint16) (err error) { 
    return instance.SetProperty("TargetPort", (value))
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *HNet_ConnectionPortMapping2) GetPropertyTargetPort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("TargetPort")
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

