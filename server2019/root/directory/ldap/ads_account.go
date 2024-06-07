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

// ads_account struct
type ads_account struct { 
	*ds_top

	// 
	DS_host []string

	// 
	DS_l string

	// 
	DS_o []string

	// 
	DS_ou []string

	// 
	DS_seeAlso []string

	// 
	DS_uid []string
}

	func Newads_accountEx1(instance *cim.WmiInstance) (newInstance *ads_account, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_account {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_accountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_account, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_account {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_host sets the value of DS_host for the instance
func (instance *ads_account) SetPropertyDS_host(value []string) (err error) { 
    return instance.SetProperty("DS_host", (value))
}

// GetDS_host gets the value of DS_host for the instance
func (instance *ads_account) GetPropertyDS_host()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_host")
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
func (instance *ads_account) SetPropertyDS_l(value string) (err error) { 
    return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_account) GetPropertyDS_l()(value string, err error) { 
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_account) SetPropertyDS_o(value []string) (err error) { 
    return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_account) GetPropertyDS_o()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_o")
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

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_account) SetPropertyDS_ou(value []string) (err error) { 
    return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_account) GetPropertyDS_ou()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_ou")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_account) SetPropertyDS_seeAlso(value []string) (err error) { 
    return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_account) GetPropertyDS_seeAlso()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_seeAlso")
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ads_account) SetPropertyDS_uid(value []string) (err error) { 
    return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ads_account) GetPropertyDS_uid()(value []string, err error) { 
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

