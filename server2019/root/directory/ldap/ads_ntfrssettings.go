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

// ads_ntfrssettings struct
type ads_ntfrssettings struct { 
	*ds_applicationsettings

	// 
	DS_fRSExtensions Uint8Array

	// 
	DS_managedBy string
}

	func Newads_ntfrssettingsEx1(instance *cim.WmiInstance) (newInstance *ads_ntfrssettings, err error) {tmp, err := Newds_applicationsettingsEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_ntfrssettings {
	ds_applicationsettings: tmp,
	}
	return
	}
	

	func Newads_ntfrssettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_ntfrssettings, err error) {tmp, err := Newds_applicationsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_ntfrssettings {
	ds_applicationsettings: tmp,
	}
	return
	}
	

// SetDS_fRSExtensions sets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssettings) SetPropertyDS_fRSExtensions(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_fRSExtensions", (value))
}

// GetDS_fRSExtensions gets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssettings) GetPropertyDS_fRSExtensions()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_fRSExtensions")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_ntfrssettings) SetPropertyDS_managedBy(value string) (err error) { 
    return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_ntfrssettings) GetPropertyDS_managedBy()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_managedBy")
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

