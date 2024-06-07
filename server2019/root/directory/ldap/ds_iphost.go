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

// ds_iphost struct
type ds_iphost struct { 
	*ds_top

	// 
	DS_ipHostNumber []string

	// 
	DS_l string

	// 
	DS_manager string

	// 
	DS_uid []string
}

	func Newds_iphostEx1(instance *cim.WmiInstance) (newInstance *ds_iphost, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_iphost {
	ds_top: tmp,
	}
	return
	}
	

	func Newds_iphostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_iphost, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_iphost {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_ipHostNumber sets the value of DS_ipHostNumber for the instance
func (instance *ds_iphost) SetPropertyDS_ipHostNumber(value []string) (err error) { 
    return instance.SetProperty("DS_ipHostNumber", (value))
}

// GetDS_ipHostNumber gets the value of DS_ipHostNumber for the instance
func (instance *ds_iphost) GetPropertyDS_ipHostNumber()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_ipHostNumber")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ds_iphost) SetPropertyDS_l(value string) (err error) { 
    return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ds_iphost) GetPropertyDS_l()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_l")
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

// SetDS_manager sets the value of DS_manager for the instance
func (instance *ds_iphost) SetPropertyDS_manager(value string) (err error) { 
    return instance.SetProperty("DS_manager", (value))
}

// GetDS_manager gets the value of DS_manager for the instance
func (instance *ds_iphost) GetPropertyDS_manager()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_manager")
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ds_iphost) SetPropertyDS_uid(value []string) (err error) { 
    return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ds_iphost) GetPropertyDS_uid()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_uid")
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

