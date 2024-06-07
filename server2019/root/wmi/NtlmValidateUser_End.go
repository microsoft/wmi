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

// NtlmValidateUser_End struct
type NtlmValidateUser_End struct { 
	*NtlmValidateUser

	// Domain Name
	LogonDomain string

	// Logon Server
	LogonServer string

	// Success Bitmask
	Success uint32

	// User Name
	UserName string

	// Logon Workstation
	Workstation string
}

	func NewNtlmValidateUser_EndEx1(instance *cim.WmiInstance) (newInstance *NtlmValidateUser_End, err error) {tmp, err := NewNtlmValidateUserEx1(instance)
		
	if err != nil { return }
	newInstance = &NtlmValidateUser_End {
	NtlmValidateUser: tmp,
	}
	return
	}
	

	func NewNtlmValidateUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *NtlmValidateUser_End, err error) {tmp, err := NewNtlmValidateUserEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &NtlmValidateUser_End {
	NtlmValidateUser: tmp,
	}
	return
	}
	

// SetLogonDomain sets the value of LogonDomain for the instance
func (instance *NtlmValidateUser_End) SetPropertyLogonDomain(value string) (err error) { 
    return instance.SetProperty("LogonDomain", (value))
}

// GetLogonDomain gets the value of LogonDomain for the instance
func (instance *NtlmValidateUser_End) GetPropertyLogonDomain()(value string, err error) { 
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

// SetLogonServer sets the value of LogonServer for the instance
func (instance *NtlmValidateUser_End) SetPropertyLogonServer(value string) (err error) { 
    return instance.SetProperty("LogonServer", (value))
}

// GetLogonServer gets the value of LogonServer for the instance
func (instance *NtlmValidateUser_End) GetPropertyLogonServer()(value string, err error) { 
    retValue, err := instance.GetProperty("LogonServer")
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

// SetSuccess sets the value of Success for the instance
func (instance *NtlmValidateUser_End) SetPropertySuccess(value uint32) (err error) { 
    return instance.SetProperty("Success", (value))
}

// GetSuccess gets the value of Success for the instance
func (instance *NtlmValidateUser_End) GetPropertySuccess()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Success")
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
func (instance *NtlmValidateUser_End) SetPropertyUserName(value string) (err error) { 
    return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *NtlmValidateUser_End) GetPropertyUserName()(value string, err error) { 
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

// SetWorkstation sets the value of Workstation for the instance
func (instance *NtlmValidateUser_End) SetPropertyWorkstation(value string) (err error) { 
    return instance.SetProperty("Workstation", (value))
}

// GetWorkstation gets the value of Workstation for the instance
func (instance *NtlmValidateUser_End) GetPropertyWorkstation()(value string, err error) { 
    retValue, err := instance.GetProperty("Workstation")
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

