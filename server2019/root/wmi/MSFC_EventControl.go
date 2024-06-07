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

// MSFC_EventControl struct
type MSFC_EventControl struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMSFC_EventControlEx1(instance *cim.WmiInstance) (newInstance *MSFC_EventControl, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFC_EventControl {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFC_EventControlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFC_EventControl, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFC_EventControl {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSFC_EventControl) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_EventControl) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSFC_EventControl) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_EventControl) GetPropertyInstanceName()(value string, err error) { 
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

// <param name="AllTargets" type="uint32 "></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) AddTarget( /* IN */ HbaPortWWN []uint8,
 /* IN */ DiscoveredPortWWN []uint8,
 /* IN */ AllTargets uint32,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("AddTarget" , HbaPortWWN, DiscoveredPortWWN, AllTargets)
	if err != nil { return }
	return
	
}


// 

// <param name="AllTargets" type="uint32 "></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) RemoveTarget( /* IN */ HbaPortWWN []uint8,
 /* IN */ DiscoveredPortWWN []uint8,
 /* IN */ AllTargets uint32,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("RemoveTarget" , HbaPortWWN, DiscoveredPortWWN, AllTargets)
	if err != nil { return }
	return
	
}


// 

// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) AddPort( /* IN */ PortWWN []uint8,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("AddPort" , PortWWN)
	if err != nil { return }
	return
	
}


// 

// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) RemovePort( /* IN */ PortWWN []uint8,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("RemovePort" , PortWWN)
	if err != nil { return }
	return
	
}


// 

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) AddLink( /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("AddLink" )
	if err != nil { return }
	return
	
}


// 

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_EventControl) RemoveLink( /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("RemoveLink" )
	if err != nil { return }
	return
	
}


