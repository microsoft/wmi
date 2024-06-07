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

// KerbLogonUser_End struct
type KerbLogonUser_End struct { 
	*KerbLogonUser

	// Logon Domain
	LogonDomain string

	// Logon Type
	LogonType string

	// Status
	Status uint32

	// User Name
	UserName string
}

	func NewKerbLogonUser_EndEx1(instance *cim.WmiInstance) (newInstance *KerbLogonUser_End, err error) {tmp, err := NewKerbLogonUserEx1(instance)
		
	if err != nil { return }
	newInstance = &KerbLogonUser_End {
	KerbLogonUser: tmp,
	}
	return
	}
	

	func NewKerbLogonUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *KerbLogonUser_End, err error) {tmp, err := NewKerbLogonUserEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &KerbLogonUser_End {
	KerbLogonUser: tmp,
	}
	return
	}
	

// SetLogonDomain sets the value of LogonDomain for the instance
func (instance *KerbLogonUser_End) SetPropertyLogonDomain(value string) (err error) { 
    return instance.SetProperty("LogonDomain", (value))
}

// GetLogonDomain gets the value of LogonDomain for the instance
func (instance *KerbLogonUser_End) GetPropertyLogonDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("LogonDomain")
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

// SetLogonType sets the value of LogonType for the instance
func (instance *KerbLogonUser_End) SetPropertyLogonType(value string) (err error) { 
    return instance.SetProperty("LogonType", (value))
}

// GetLogonType gets the value of LogonType for the instance
func (instance *KerbLogonUser_End) GetPropertyLogonType()(value string, err error) { 
    retValue, err := instance.GetProperty("LogonType")
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

// SetStatus sets the value of Status for the instance
func (instance *KerbLogonUser_End) SetPropertyStatus(value uint32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KerbLogonUser_End) GetPropertyStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Status")
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

// SetUserName sets the value of UserName for the instance
func (instance *KerbLogonUser_End) SetPropertyUserName(value string) (err error) { 
    return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *KerbLogonUser_End) GetPropertyUserName()(value string, err error) { 
    retValue, err := instance.GetProperty("UserName")
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

