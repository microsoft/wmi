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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_foreignsecurityprincipal struct
type ads_foreignsecurityprincipal struct { 
	*ds_top

	// 
	DS_foreignIdentifier Uint8Array

	// 
	DS_objectSid Uint8Array
}

	func Newads_foreignsecurityprincipalEx1(instance *cim.WmiInstance) (newInstance *ads_foreignsecurityprincipal, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_foreignsecurityprincipal {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_foreignsecurityprincipalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_foreignsecurityprincipal, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_foreignsecurityprincipal {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_foreignIdentifier sets the value of DS_foreignIdentifier for the instance
func (instance *ads_foreignsecurityprincipal) SetPropertyDS_foreignIdentifier(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_foreignIdentifier", (value))
}

// GetDS_foreignIdentifier gets the value of DS_foreignIdentifier for the instance
func (instance *ads_foreignsecurityprincipal) GetPropertyDS_foreignIdentifier()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_foreignIdentifier")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_objectSid sets the value of DS_objectSid for the instance
func (instance *ads_foreignsecurityprincipal) SetPropertyDS_objectSid(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ads_foreignsecurityprincipal) GetPropertyDS_objectSid()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_objectSid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

