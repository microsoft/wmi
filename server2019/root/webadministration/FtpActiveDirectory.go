// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpActiveDirectory struct
type FtpActiveDirectory struct { 
	*EmbeddedObject

	// 
	AdCacheRefresh string

	// 
	AdPassword string

	// 
	AdUserName string
}

	func NewFtpActiveDirectoryEx1(instance *cim.WmiInstance) (newInstance *FtpActiveDirectory, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpActiveDirectory {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpActiveDirectoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpActiveDirectory, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpActiveDirectory {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAdCacheRefresh sets the value of AdCacheRefresh for the instance
func (instance *FtpActiveDirectory) SetPropertyAdCacheRefresh(value string) (err error) { 
    return instance.SetProperty("AdCacheRefresh", (value))
}

// GetAdCacheRefresh gets the value of AdCacheRefresh for the instance
func (instance *FtpActiveDirectory) GetPropertyAdCacheRefresh()(value string, err error) { 
    retValue, err := instance.GetProperty("AdCacheRefresh")
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

// SetAdPassword sets the value of AdPassword for the instance
func (instance *FtpActiveDirectory) SetPropertyAdPassword(value string) (err error) { 
    return instance.SetProperty("AdPassword", (value))
}

// GetAdPassword gets the value of AdPassword for the instance
func (instance *FtpActiveDirectory) GetPropertyAdPassword()(value string, err error) { 
    retValue, err := instance.GetProperty("AdPassword")
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

// SetAdUserName sets the value of AdUserName for the instance
func (instance *FtpActiveDirectory) SetPropertyAdUserName(value string) (err error) { 
    return instance.SetProperty("AdUserName", (value))
}

// GetAdUserName gets the value of AdUserName for the instance
func (instance *FtpActiveDirectory) GetPropertyAdUserName()(value string, err error) { 
    retValue, err := instance.GetProperty("AdUserName")
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

