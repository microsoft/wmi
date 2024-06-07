// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessAccountingConnectionLocal struct
type RemoteAccessAccountingConnectionLocal struct { 
	*RemoteAccessConnectionLocal

	// 
	SessionId uint64

	// 
	UserName string
}

	func NewRemoteAccessAccountingConnectionLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessAccountingConnectionLocal, err error) {tmp, err := NewRemoteAccessConnectionLocalEx1(instance)
		
	if err != nil { return }
	newInstance = &RemoteAccessAccountingConnectionLocal {
	RemoteAccessConnectionLocal: tmp,
	}
	return
	}
	

	func NewRemoteAccessAccountingConnectionLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessAccountingConnectionLocal, err error) {tmp, err := NewRemoteAccessConnectionLocalEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessAccountingConnectionLocal {
	RemoteAccessConnectionLocal: tmp,
	}
	return
	}
	

// SetSessionId sets the value of SessionId for the instance
func (instance *RemoteAccessAccountingConnectionLocal) SetPropertySessionId(value uint64) (err error) { 
    return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *RemoteAccessAccountingConnectionLocal) GetPropertySessionId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SessionId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetUserName sets the value of UserName for the instance
func (instance *RemoteAccessAccountingConnectionLocal) SetPropertyUserName(value string) (err error) { 
    return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *RemoteAccessAccountingConnectionLocal) GetPropertyUserName()(value string, err error) { 
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

