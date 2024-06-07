// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessAccounting struct
type RemoteAccessAccounting struct { 
	*cim.WmiInstance

	// 
	InboxAccounting RemoteAccessInboxAccounting

	// 
	RadiusAccounting RemoteAccessRadiusAccounting
}

	func NewRemoteAccessAccountingEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessAccounting, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessAccounting {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessAccounting, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessAccounting {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInboxAccounting sets the value of InboxAccounting for the instance
func (instance *RemoteAccessAccounting) SetPropertyInboxAccounting(value RemoteAccessInboxAccounting) (err error) { 
    return instance.SetProperty("InboxAccounting", (value))
}

// GetInboxAccounting gets the value of InboxAccounting for the instance
func (instance *RemoteAccessAccounting) GetPropertyInboxAccounting()(value RemoteAccessInboxAccounting, err error) { 
    retValue, err := instance.GetProperty("InboxAccounting")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(RemoteAccessInboxAccounting); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " RemoteAccessInboxAccounting is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RemoteAccessInboxAccounting(valuetmp)

    return
}

// SetRadiusAccounting sets the value of RadiusAccounting for the instance
func (instance *RemoteAccessAccounting) SetPropertyRadiusAccounting(value RemoteAccessRadiusAccounting) (err error) { 
    return instance.SetProperty("RadiusAccounting", (value))
}

// GetRadiusAccounting gets the value of RadiusAccounting for the instance
func (instance *RemoteAccessAccounting) GetPropertyRadiusAccounting()(value RemoteAccessRadiusAccounting, err error) { 
    retValue, err := instance.GetProperty("RadiusAccounting")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(RemoteAccessRadiusAccounting); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " RemoteAccessRadiusAccounting is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RemoteAccessRadiusAccounting(valuetmp)

    return
}

