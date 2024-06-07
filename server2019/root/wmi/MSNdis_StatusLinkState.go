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

// MSNdis_StatusLinkState struct
type MSNdis_StatusLinkState struct { 
	*WMIEvent

	// 
	Active bool

	// 
	InstanceName string

	// 
	NdisLinkStateStatusIndication []uint8

	// 
	NumberElements uint32
}

	func NewMSNdis_StatusLinkStateEx1(instance *cim.WmiInstance) (newInstance *MSNdis_StatusLinkState, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_StatusLinkState {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewMSNdis_StatusLinkStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_StatusLinkState, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_StatusLinkState {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_StatusLinkState) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_StatusLinkState) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_StatusLinkState) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_StatusLinkState) GetPropertyInstanceName()(value string, err error) { 
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

// SetNdisLinkStateStatusIndication sets the value of NdisLinkStateStatusIndication for the instance
func (instance *MSNdis_StatusLinkState) SetPropertyNdisLinkStateStatusIndication(value []uint8) (err error) { 
    return instance.SetProperty("NdisLinkStateStatusIndication", (value))
}

// GetNdisLinkStateStatusIndication gets the value of NdisLinkStateStatusIndication for the instance
func (instance *MSNdis_StatusLinkState) GetPropertyNdisLinkStateStatusIndication()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("NdisLinkStateStatusIndication")
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

// SetNumberElements sets the value of NumberElements for the instance
func (instance *MSNdis_StatusLinkState) SetPropertyNumberElements(value uint32) (err error) { 
    return instance.SetProperty("NumberElements", (value))
}

// GetNumberElements gets the value of NumberElements for the instance
func (instance *MSNdis_StatusLinkState) GetPropertyNumberElements()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberElements")
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

