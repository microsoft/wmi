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

// ds_connectionpoint struct
type ds_connectionpoint struct { 
	*ds_leaf

	// 
	DS_keywords []string

	// 
	DS_managedBy string

	// 
	DS_msDS_Settings []string
}

	func Newds_connectionpointEx1(instance *cim.WmiInstance) (newInstance *ds_connectionpoint, err error) {tmp, err := Newds_leafEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_connectionpoint {
	ds_leaf: tmp,
	}
	return
	}
	

	func Newds_connectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_connectionpoint, err error) {tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_connectionpoint {
	ds_leaf: tmp,
	}
	return
	}
	

// SetDS_keywords sets the value of DS_keywords for the instance
func (instance *ds_connectionpoint) SetPropertyDS_keywords(value []string) (err error) { 
    return instance.SetProperty("DS_keywords", (value))
}

// GetDS_keywords gets the value of DS_keywords for the instance
func (instance *ds_connectionpoint) GetPropertyDS_keywords()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_keywords")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ds_connectionpoint) SetPropertyDS_managedBy(value string) (err error) { 
    return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ds_connectionpoint) GetPropertyDS_managedBy()(value string, err error) { 
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

// SetDS_msDS_Settings sets the value of DS_msDS_Settings for the instance
func (instance *ds_connectionpoint) SetPropertyDS_msDS_Settings(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_Settings", (value))
}

// GetDS_msDS_Settings gets the value of DS_msDS_Settings for the instance
func (instance *ds_connectionpoint) GetPropertyDS_msDS_Settings()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_Settings")
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

