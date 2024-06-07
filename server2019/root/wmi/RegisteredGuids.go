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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RegisteredGuids struct
type RegisteredGuids struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	EnableFlags uint32

	// 
	EnableLevel uint32

	// 
	GuidType uint32

	// 
	InstanceName string

	// 
	IsEnabled bool

	// 
	LoggerId uint32
}

	func NewRegisteredGuidsEx1(instance *cim.WmiInstance) (newInstance *RegisteredGuids, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RegisteredGuids {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRegisteredGuidsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RegisteredGuids, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RegisteredGuids {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *RegisteredGuids) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *RegisteredGuids) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
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

// SetEnableFlags sets the value of EnableFlags for the instance
func (instance *RegisteredGuids) SetPropertyEnableFlags(value uint32) (err error) { 
    return instance.SetProperty("EnableFlags", (value))
}

// GetEnableFlags gets the value of EnableFlags for the instance
func (instance *RegisteredGuids) GetPropertyEnableFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EnableFlags")
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

// SetEnableLevel sets the value of EnableLevel for the instance
func (instance *RegisteredGuids) SetPropertyEnableLevel(value uint32) (err error) { 
    return instance.SetProperty("EnableLevel", (value))
}

// GetEnableLevel gets the value of EnableLevel for the instance
func (instance *RegisteredGuids) GetPropertyEnableLevel()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EnableLevel")
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

// SetGuidType sets the value of GuidType for the instance
func (instance *RegisteredGuids) SetPropertyGuidType(value uint32) (err error) { 
    return instance.SetProperty("GuidType", (value))
}

// GetGuidType gets the value of GuidType for the instance
func (instance *RegisteredGuids) GetPropertyGuidType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GuidType")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *RegisteredGuids) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *RegisteredGuids) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *RegisteredGuids) SetPropertyIsEnabled(value bool) (err error) { 
    return instance.SetProperty("IsEnabled", (value))
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *RegisteredGuids) GetPropertyIsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsEnabled")
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

// SetLoggerId sets the value of LoggerId for the instance
func (instance *RegisteredGuids) SetPropertyLoggerId(value uint32) (err error) { 
    return instance.SetProperty("LoggerId", (value))
}

// GetLoggerId gets the value of LoggerId for the instance
func (instance *RegisteredGuids) GetPropertyLoggerId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoggerId")
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

