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

// MSWmi_GuidRegistrationInfo struct
type MSWmi_GuidRegistrationInfo struct { 
	*WMIEvent

	// 
	Active bool

	// 
	GuidCount uint32

	// 
	GuidList []MSWmi_Guid

	// 
	InstanceName string

	// 
	Operation uint32
}

	func NewMSWmi_GuidRegistrationInfoEx1(instance *cim.WmiInstance) (newInstance *MSWmi_GuidRegistrationInfo, err error) {tmp, err := NewWMIEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSWmi_GuidRegistrationInfo {
	WMIEvent: tmp,
	}
	return
	}
	

	func NewMSWmi_GuidRegistrationInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSWmi_GuidRegistrationInfo, err error) {tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSWmi_GuidRegistrationInfo {
	WMIEvent: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSWmi_GuidRegistrationInfo) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSWmi_GuidRegistrationInfo) GetPropertyActive()(value bool, err error) { 
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

// SetGuidCount sets the value of GuidCount for the instance
func (instance *MSWmi_GuidRegistrationInfo) SetPropertyGuidCount(value uint32) (err error) { 
    return instance.SetProperty("GuidCount", (value))
}

// GetGuidCount gets the value of GuidCount for the instance
func (instance *MSWmi_GuidRegistrationInfo) GetPropertyGuidCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GuidCount")
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

// SetGuidList sets the value of GuidList for the instance
func (instance *MSWmi_GuidRegistrationInfo) SetPropertyGuidList(value []MSWmi_Guid) (err error) { 
    return instance.SetProperty("GuidList", (value))
}

// GetGuidList gets the value of GuidList for the instance
func (instance *MSWmi_GuidRegistrationInfo) GetPropertyGuidList()(value []MSWmi_Guid, err error) { 
    retValue, err := instance.GetProperty("GuidList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSWmi_Guid); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSWmi_Guid is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSWmi_Guid(valuetmp))
    }

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSWmi_GuidRegistrationInfo) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSWmi_GuidRegistrationInfo) GetPropertyInstanceName()(value string, err error) { 
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

// SetOperation sets the value of Operation for the instance
func (instance *MSWmi_GuidRegistrationInfo) SetPropertyOperation(value uint32) (err error) { 
    return instance.SetProperty("Operation", (value))
}

// GetOperation gets the value of Operation for the instance
func (instance *MSWmi_GuidRegistrationInfo) GetPropertyOperation()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Operation")
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

