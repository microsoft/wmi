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

// WmiMonitorBrightnessMethods struct
type WmiMonitorBrightnessMethods struct { 
	*MSMonitorClass

	// 
	Active bool

	// 
	InstanceName string
}

	func NewWmiMonitorBrightnessMethodsEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorBrightnessMethods, err error) {tmp, err := NewMSMonitorClassEx1(instance)
		
	if err != nil { return }
	newInstance = &WmiMonitorBrightnessMethods {
	MSMonitorClass: tmp,
	}
	return
	}
	

	func NewWmiMonitorBrightnessMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WmiMonitorBrightnessMethods, err error) {tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WmiMonitorBrightnessMethods {
	MSMonitorClass: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorBrightnessMethods) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorBrightnessMethods) GetPropertyActive()(value bool, err error) { 
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
func (instance *WmiMonitorBrightnessMethods) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorBrightnessMethods) GetPropertyInstanceName()(value string, err error) { 
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

// 

// <param name="Brightness" type="uint8 "></param>
// <param name="Timeout" type="uint32 "></param>
func (instance *WmiMonitorBrightnessMethods) WmiSetBrightness( /* IN */ Timeout uint32,
 /* IN */ Brightness uint8) (err error) {_, err = instance.InvokeMethodWithReturn("WmiSetBrightness" , Timeout, Brightness);
	if err != nil { return }
	return
	
}


// 
func (instance *WmiMonitorBrightnessMethods) WmiRevertToPolicyBrightness() (err error) {_, err = instance.InvokeMethodWithReturn("WmiRevertToPolicyBrightness" );
	if err != nil { return }
	return
	
}


// 

// <param name="State" type="bool "></param>
func (instance *WmiMonitorBrightnessMethods) WmiSetALSBrightnessState( /* IN */ State bool) (err error) {_, err = instance.InvokeMethodWithReturn("WmiSetALSBrightnessState" , State);
	if err != nil { return }
	return
	
}


// 

// <param name="Brightness" type="uint8 "></param>
func (instance *WmiMonitorBrightnessMethods) WmiSetALSBrightness( /* IN */ Brightness uint8) (err error) {_, err = instance.InvokeMethodWithReturn("WmiSetALSBrightness" , Brightness);
	if err != nil { return }
	return
	
}


