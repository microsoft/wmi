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

// ads_groupofnames struct
type ads_groupofnames struct { 
	*ds_top

	// 
	DS_businessCategory []string

	// 
	DS_member []string

	// 
	DS_o []string

	// 
	DS_ou []string

	// 
	DS_owner string

	// 
	DS_seeAlso []string
}

	func Newads_groupofnamesEx1(instance *cim.WmiInstance) (newInstance *ads_groupofnames, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_groupofnames {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_groupofnamesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_groupofnames, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_groupofnames {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_businessCategory sets the value of DS_businessCategory for the instance
func (instance *ads_groupofnames) SetPropertyDS_businessCategory(value []string) (err error) { 
    return instance.SetProperty("DS_businessCategory", (value))
}

// GetDS_businessCategory gets the value of DS_businessCategory for the instance
func (instance *ads_groupofnames) GetPropertyDS_businessCategory()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_businessCategory")
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

// SetDS_member sets the value of DS_member for the instance
func (instance *ads_groupofnames) SetPropertyDS_member(value []string) (err error) { 
    return instance.SetProperty("DS_member", (value))
}

// GetDS_member gets the value of DS_member for the instance
func (instance *ads_groupofnames) GetPropertyDS_member()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_member")
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_groupofnames) SetPropertyDS_o(value []string) (err error) { 
    return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_groupofnames) GetPropertyDS_o()(value []string, err error) { 
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
func (instance *ads_groupofnames) SetPropertyDS_ou(value []string) (err error) { 
    return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_groupofnames) GetPropertyDS_ou()(value []string, err error) { 
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

// SetDS_owner sets the value of DS_owner for the instance
func (instance *ads_groupofnames) SetPropertyDS_owner(value string) (err error) { 
    return instance.SetProperty("DS_owner", (value))
}

// GetDS_owner gets the value of DS_owner for the instance
func (instance *ads_groupofnames) GetPropertyDS_owner()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_owner")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_groupofnames) SetPropertyDS_seeAlso(value []string) (err error) { 
    return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_groupofnames) GetPropertyDS_seeAlso()(value []string, err error) { 
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

