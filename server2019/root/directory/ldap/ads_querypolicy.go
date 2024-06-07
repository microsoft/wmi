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

// ads_querypolicy struct
type ads_querypolicy struct { 
	*ds_top

	// 
	DS_lDAPAdminLimits []string

	// 
	DS_lDAPIPDenyList []Uint8Array
}

	func Newads_querypolicyEx1(instance *cim.WmiInstance) (newInstance *ads_querypolicy, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_querypolicy {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_querypolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_querypolicy, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_querypolicy {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_lDAPAdminLimits sets the value of DS_lDAPAdminLimits for the instance
func (instance *ads_querypolicy) SetPropertyDS_lDAPAdminLimits(value []string) (err error) { 
    return instance.SetProperty("DS_lDAPAdminLimits", (value))
}

// GetDS_lDAPAdminLimits gets the value of DS_lDAPAdminLimits for the instance
func (instance *ads_querypolicy) GetPropertyDS_lDAPAdminLimits()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_lDAPAdminLimits")
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

// SetDS_lDAPIPDenyList sets the value of DS_lDAPIPDenyList for the instance
func (instance *ads_querypolicy) SetPropertyDS_lDAPIPDenyList(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_lDAPIPDenyList", (value))
}

// GetDS_lDAPIPDenyList gets the value of DS_lDAPIPDenyList for the instance
func (instance *ads_querypolicy) GetPropertyDS_lDAPIPDenyList()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_lDAPIPDenyList")
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

