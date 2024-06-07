// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSD5239530_48B3_486F_9FC3_5701A2D5A27D.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSoP_PolicySettingStatus struct
type RSoP_PolicySettingStatus struct { 
	*cim.WmiInstance

	// 
	errorCode uint32

	// 
	eventID uint32

	// 
	eventLogName string

	// 
	eventSource string

	// 
	eventTime string

	// 
	id string

	// 
	status int32
}

	func NewRSoP_PolicySettingStatusEx1(instance *cim.WmiInstance) (newInstance *RSoP_PolicySettingStatus, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RSoP_PolicySettingStatus {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRSoP_PolicySettingStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSoP_PolicySettingStatus, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSoP_PolicySettingStatus {
	WmiInstance: tmp,
	}
	return
	}
	

// SeterrorCode sets the value of errorCode for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyerrorCode(value uint32) (err error) { 
    return instance.SetProperty("errorCode", (value))
}

// GeterrorCode gets the value of errorCode for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyerrorCode()(value uint32, err error) { 
    retValue, err := instance.GetProperty("errorCode")
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

// SeteventID sets the value of eventID for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventID(value uint32) (err error) { 
    return instance.SetProperty("eventID", (value))
}

// GeteventID gets the value of eventID for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("eventID")
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

// SeteventLogName sets the value of eventLogName for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventLogName(value string) (err error) { 
    return instance.SetProperty("eventLogName", (value))
}

// GeteventLogName gets the value of eventLogName for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventLogName()(value string, err error) { 
    retValue, err := instance.GetProperty("eventLogName")
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

// SeteventSource sets the value of eventSource for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventSource(value string) (err error) { 
    return instance.SetProperty("eventSource", (value))
}

// GeteventSource gets the value of eventSource for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventSource()(value string, err error) { 
    retValue, err := instance.GetProperty("eventSource")
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

// SeteventTime sets the value of eventTime for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyeventTime(value string) (err error) { 
    return instance.SetProperty("eventTime", (value))
}

// GeteventTime gets the value of eventTime for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyeventTime()(value string, err error) { 
    retValue, err := instance.GetProperty("eventTime")
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

// Setid sets the value of id for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertyid(value string) (err error) { 
    return instance.SetProperty("id", (value))
}

// Getid gets the value of id for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertyid()(value string, err error) { 
    retValue, err := instance.GetProperty("id")
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

// Setstatus sets the value of status for the instance
func (instance *RSoP_PolicySettingStatus) SetPropertystatus(value int32) (err error) { 
    return instance.SetProperty("status", (value))
}

// Getstatus gets the value of status for the instance
func (instance *RSoP_PolicySettingStatus) GetPropertystatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("status")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

