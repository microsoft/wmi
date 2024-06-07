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

// DAAppServer struct
type DAAppServer struct { 
	*cim.WmiInstance

	// 
	AppServerConnection DAAppServerConnection

	// 
	GpoName []string

	// 
	SecurityGroupNameList []string
}

	func NewDAAppServerEx1(instance *cim.WmiInstance) (newInstance *DAAppServer, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DAAppServer {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDAAppServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAAppServer, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAAppServer {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAppServerConnection sets the value of AppServerConnection for the instance
func (instance *DAAppServer) SetPropertyAppServerConnection(value DAAppServerConnection) (err error) { 
    return instance.SetProperty("AppServerConnection", (value))
}

// GetAppServerConnection gets the value of AppServerConnection for the instance
func (instance *DAAppServer) GetPropertyAppServerConnection()(value DAAppServerConnection, err error) { 
    retValue, err := instance.GetProperty("AppServerConnection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DAAppServerConnection); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DAAppServerConnection is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DAAppServerConnection(valuetmp)

    return
}

// SetGpoName sets the value of GpoName for the instance
func (instance *DAAppServer) SetPropertyGpoName(value []string) (err error) { 
    return instance.SetProperty("GpoName", (value))
}

// GetGpoName gets the value of GpoName for the instance
func (instance *DAAppServer) GetPropertyGpoName()(value []string, err error) { 
    retValue, err := instance.GetProperty("GpoName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetSecurityGroupNameList sets the value of SecurityGroupNameList for the instance
func (instance *DAAppServer) SetPropertySecurityGroupNameList(value []string) (err error) { 
    return instance.SetProperty("SecurityGroupNameList", (value))
}

// GetSecurityGroupNameList gets the value of SecurityGroupNameList for the instance
func (instance *DAAppServer) GetPropertySecurityGroupNameList()(value []string, err error) { 
    retValue, err := instance.GetProperty("SecurityGroupNameList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

