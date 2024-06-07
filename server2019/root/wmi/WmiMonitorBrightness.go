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

// WmiMonitorBrightness struct
type WmiMonitorBrightness struct { 
	*MSMonitorClass

	// 
	Active bool

	// 
	CurrentBrightness uint8

	// 
	InstanceName string

	// 
	Level []uint8

	// 
	Levels uint32
}

	func NewWmiMonitorBrightnessEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorBrightness, err error) {tmp, err := NewMSMonitorClassEx1(instance)
		
	if err != nil { return }
	newInstance = &WmiMonitorBrightness {
	MSMonitorClass: tmp,
	}
	return
	}
	

	func NewWmiMonitorBrightnessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WmiMonitorBrightness, err error) {tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WmiMonitorBrightness {
	MSMonitorClass: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorBrightness) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorBrightness) GetPropertyActive()(value bool, err error) { 
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

// SetCurrentBrightness sets the value of CurrentBrightness for the instance
func (instance *WmiMonitorBrightness) SetPropertyCurrentBrightness(value uint8) (err error) { 
    return instance.SetProperty("CurrentBrightness", (value))
}

// GetCurrentBrightness gets the value of CurrentBrightness for the instance
func (instance *WmiMonitorBrightness) GetPropertyCurrentBrightness()(value uint8, err error) { 
    retValue, err := instance.GetProperty("CurrentBrightness")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *WmiMonitorBrightness) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorBrightness) GetPropertyInstanceName()(value string, err error) { 
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

// SetLevel sets the value of Level for the instance
func (instance *WmiMonitorBrightness) SetPropertyLevel(value []uint8) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *WmiMonitorBrightness) GetPropertyLevel()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Level")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetLevels sets the value of Levels for the instance
func (instance *WmiMonitorBrightness) SetPropertyLevels(value uint32) (err error) { 
    return instance.SetProperty("Levels", (value))
}

// GetLevels gets the value of Levels for the instance
func (instance *WmiMonitorBrightness) GetPropertyLevels()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Levels")
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

