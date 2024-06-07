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

// ClientTarget struct
type ClientTarget struct { 
	*CollectionElement

	// 
	Alias string

	// 
	UserAgent string
}

	func NewClientTargetEx1(instance *cim.WmiInstance) (newInstance *ClientTarget, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ClientTarget {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewClientTargetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ClientTarget, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ClientTarget {
	CollectionElement: tmp,
	}
	return
	}
	

// SetAlias sets the value of Alias for the instance
func (instance *ClientTarget) SetPropertyAlias(value string) (err error) { 
    return instance.SetProperty("Alias", (value))
}

// GetAlias gets the value of Alias for the instance
func (instance *ClientTarget) GetPropertyAlias()(value string, err error) { 
    retValue, err := instance.GetProperty("Alias")
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

// SetUserAgent sets the value of UserAgent for the instance
func (instance *ClientTarget) SetPropertyUserAgent(value string) (err error) { 
    return instance.SetProperty("UserAgent", (value))
}

// GetUserAgent gets the value of UserAgent for the instance
func (instance *ClientTarget) GetPropertyUserAgent()(value string, err error) { 
    retValue, err := instance.GetProperty("UserAgent")
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

