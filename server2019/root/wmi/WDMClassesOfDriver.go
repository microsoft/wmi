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

// WDMClassesOfDriver struct
type WDMClassesOfDriver struct { 
	*cim.WmiInstance

	// 
	ClassName string

	// 
	Driver string

	// 
	HighDateTime uint32

	// 
	LowDateTime uint32
}

	func NewWDMClassesOfDriverEx1(instance *cim.WmiInstance) (newInstance *WDMClassesOfDriver, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &WDMClassesOfDriver {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewWDMClassesOfDriverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WDMClassesOfDriver, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WDMClassesOfDriver {
	WmiInstance: tmp,
	}
	return
	}
	

// SetClassName sets the value of ClassName for the instance
func (instance *WDMClassesOfDriver) SetPropertyClassName(value string) (err error) { 
    return instance.SetProperty("ClassName", (value))
}

// GetClassName gets the value of ClassName for the instance
func (instance *WDMClassesOfDriver) GetPropertyClassName()(value string, err error) { 
    retValue, err := instance.GetProperty("ClassName")
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

// SetDriver sets the value of Driver for the instance
func (instance *WDMClassesOfDriver) SetPropertyDriver(value string) (err error) { 
    return instance.SetProperty("Driver", (value))
}

// GetDriver gets the value of Driver for the instance
func (instance *WDMClassesOfDriver) GetPropertyDriver()(value string, err error) { 
    retValue, err := instance.GetProperty("Driver")
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

// SetHighDateTime sets the value of HighDateTime for the instance
func (instance *WDMClassesOfDriver) SetPropertyHighDateTime(value uint32) (err error) { 
    return instance.SetProperty("HighDateTime", (value))
}

// GetHighDateTime gets the value of HighDateTime for the instance
func (instance *WDMClassesOfDriver) GetPropertyHighDateTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HighDateTime")
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

// SetLowDateTime sets the value of LowDateTime for the instance
func (instance *WDMClassesOfDriver) SetPropertyLowDateTime(value uint32) (err error) { 
    return instance.SetProperty("LowDateTime", (value))
}

// GetLowDateTime gets the value of LowDateTime for the instance
func (instance *WDMClassesOfDriver) GetPropertyLowDateTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LowDateTime")
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

