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

// RemoteAccessRadiusServerPurpose struct
type RemoteAccessRadiusServerPurpose struct { 
	*cim.WmiInstance

	// 
	Purpose string

	// 
	ServerName string
}

	func NewRemoteAccessRadiusServerPurposeEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessRadiusServerPurpose, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessRadiusServerPurpose {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessRadiusServerPurposeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessRadiusServerPurpose, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessRadiusServerPurpose {
	WmiInstance: tmp,
	}
	return
	}
	

// SetPurpose sets the value of Purpose for the instance
func (instance *RemoteAccessRadiusServerPurpose) SetPropertyPurpose(value string) (err error) { 
    return instance.SetProperty("Purpose", (value))
}

// GetPurpose gets the value of Purpose for the instance
func (instance *RemoteAccessRadiusServerPurpose) GetPropertyPurpose()(value string, err error) { 
    retValue, err := instance.GetProperty("Purpose")
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

// SetServerName sets the value of ServerName for the instance
func (instance *RemoteAccessRadiusServerPurpose) SetPropertyServerName(value string) (err error) { 
    return instance.SetProperty("ServerName", (value))
}

// GetServerName gets the value of ServerName for the instance
func (instance *RemoteAccessRadiusServerPurpose) GetPropertyServerName()(value string, err error) { 
    retValue, err := instance.GetProperty("ServerName")
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

