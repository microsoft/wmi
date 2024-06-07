// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DN_With_String struct
type DN_With_String struct { 
	*cim.WmiInstance

	// 
	dnString string

	// 
	value string
}

	func NewDN_With_StringEx1(instance *cim.WmiInstance) (newInstance *DN_With_String, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DN_With_String {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDN_With_StringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DN_With_String, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DN_With_String {
	WmiInstance: tmp,
	}
	return
	}
	

// SetdnString sets the value of dnString for the instance
func (instance *DN_With_String) SetPropertydnString(value string) (err error) { 
    return instance.SetProperty("dnString", (value))
}

// GetdnString gets the value of dnString for the instance
func (instance *DN_With_String) GetPropertydnString()(value string, err error) { 
    retValue, err := instance.GetProperty("dnString")
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

// Setvalue sets the value of value for the instance
func (instance *DN_With_String) SetPropertyvalue(value string) (err error) { 
    return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *DN_With_String) GetPropertyvalue()(value string, err error) { 
    retValue, err := instance.GetProperty("value")
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

