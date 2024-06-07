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

// ads_storage struct
type ads_storage struct { 
	*ds_connectionpoint

	// 
	DS_iconPath []string

	// 
	DS_moniker []Uint8Array

	// 
	DS_monikerDisplayName []string
}

	func Newads_storageEx1(instance *cim.WmiInstance) (newInstance *ads_storage, err error) {tmp, err := Newds_connectionpointEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_storage {
	ds_connectionpoint: tmp,
	}
	return
	}
	

	func Newads_storageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_storage, err error) {tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_storage {
	ds_connectionpoint: tmp,
	}
	return
	}
	

// SetDS_iconPath sets the value of DS_iconPath for the instance
func (instance *ads_storage) SetPropertyDS_iconPath(value []string) (err error) { 
    return instance.SetProperty("DS_iconPath", (value))
}

// GetDS_iconPath gets the value of DS_iconPath for the instance
func (instance *ads_storage) GetPropertyDS_iconPath()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_iconPath")
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

// SetDS_moniker sets the value of DS_moniker for the instance
func (instance *ads_storage) SetPropertyDS_moniker(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_moniker", (value))
}

// GetDS_moniker gets the value of DS_moniker for the instance
func (instance *ads_storage) GetPropertyDS_moniker()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_moniker")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(Uint8Array); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Uint8Array(valuetmp))
    }

    return
}

// SetDS_monikerDisplayName sets the value of DS_monikerDisplayName for the instance
func (instance *ads_storage) SetPropertyDS_monikerDisplayName(value []string) (err error) { 
    return instance.SetProperty("DS_monikerDisplayName", (value))
}

// GetDS_monikerDisplayName gets the value of DS_monikerDisplayName for the instance
func (instance *ads_storage) GetPropertyDS_monikerDisplayName()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_monikerDisplayName")
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

