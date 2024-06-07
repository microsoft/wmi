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

// MSKeyboard_ExtendedID struct
type MSKeyboard_ExtendedID struct { 
	*MSKeyboard

	// 
	Active bool

	// 
	InstanceName string

	// 
	Subtype uint32

	// 
	Type uint32
}

	func NewMSKeyboard_ExtendedIDEx1(instance *cim.WmiInstance) (newInstance *MSKeyboard_ExtendedID, err error) {tmp, err := NewMSKeyboardEx1(instance)
		
	if err != nil { return }
	newInstance = &MSKeyboard_ExtendedID {
	MSKeyboard: tmp,
	}
	return
	}
	

	func NewMSKeyboard_ExtendedIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSKeyboard_ExtendedID, err error) {tmp, err := NewMSKeyboardEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSKeyboard_ExtendedID {
	MSKeyboard: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSKeyboard_ExtendedID) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSKeyboard_ExtendedID) GetPropertyActive()(value bool, err error) { 
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSKeyboard_ExtendedID) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSKeyboard_ExtendedID) GetPropertyInstanceName()(value string, err error) { 
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

// SetSubtype sets the value of Subtype for the instance
func (instance *MSKeyboard_ExtendedID) SetPropertySubtype(value uint32) (err error) { 
    return instance.SetProperty("Subtype", (value))
}

// GetSubtype gets the value of Subtype for the instance
func (instance *MSKeyboard_ExtendedID) GetPropertySubtype()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Subtype")
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

// SetType sets the value of Type for the instance
func (instance *MSKeyboard_ExtendedID) SetPropertyType(value uint32) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSKeyboard_ExtendedID) GetPropertyType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Type")
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

