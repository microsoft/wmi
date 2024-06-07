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

// ads_serviceinstance struct
type ads_serviceinstance struct { 
	*ds_connectionpoint

	// 
	DS_serviceClassID Uint8Array

	// 
	DS_serviceInstanceVersion Uint8Array

	// 
	DS_winsockAddresses []Uint8Array
}

	func Newads_serviceinstanceEx1(instance *cim.WmiInstance) (newInstance *ads_serviceinstance, err error) {tmp, err := Newds_connectionpointEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_serviceinstance {
	ds_connectionpoint: tmp,
	}
	return
	}
	

	func Newads_serviceinstanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_serviceinstance, err error) {tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_serviceinstance {
	ds_connectionpoint: tmp,
	}
	return
	}
	

// SetDS_serviceClassID sets the value of DS_serviceClassID for the instance
func (instance *ads_serviceinstance) SetPropertyDS_serviceClassID(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_serviceClassID", (value))
}

// GetDS_serviceClassID gets the value of DS_serviceClassID for the instance
func (instance *ads_serviceinstance) GetPropertyDS_serviceClassID()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_serviceClassID")
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

// SetDS_serviceInstanceVersion sets the value of DS_serviceInstanceVersion for the instance
func (instance *ads_serviceinstance) SetPropertyDS_serviceInstanceVersion(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_serviceInstanceVersion", (value))
}

// GetDS_serviceInstanceVersion gets the value of DS_serviceInstanceVersion for the instance
func (instance *ads_serviceinstance) GetPropertyDS_serviceInstanceVersion()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_serviceInstanceVersion")
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

// SetDS_winsockAddresses sets the value of DS_winsockAddresses for the instance
func (instance *ads_serviceinstance) SetPropertyDS_winsockAddresses(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_winsockAddresses", (value))
}

// GetDS_winsockAddresses gets the value of DS_winsockAddresses for the instance
func (instance *ads_serviceinstance) GetPropertyDS_winsockAddresses()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_winsockAddresses")
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

