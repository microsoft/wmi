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

// ads_rrasadministrationdictionary struct
type ads_rrasadministrationdictionary struct { 
	*ds_top

	// 
	DS_msRRASVendorAttributeEntry []string
}

	func Newads_rrasadministrationdictionaryEx1(instance *cim.WmiInstance) (newInstance *ads_rrasadministrationdictionary, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_rrasadministrationdictionary {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_rrasadministrationdictionaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_rrasadministrationdictionary, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_rrasadministrationdictionary {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msRRASVendorAttributeEntry sets the value of DS_msRRASVendorAttributeEntry for the instance
func (instance *ads_rrasadministrationdictionary) SetPropertyDS_msRRASVendorAttributeEntry(value []string) (err error) { 
    return instance.SetProperty("DS_msRRASVendorAttributeEntry", (value))
}

// GetDS_msRRASVendorAttributeEntry gets the value of DS_msRRASVendorAttributeEntry for the instance
func (instance *ads_rrasadministrationdictionary) GetPropertyDS_msRRASVendorAttributeEntry()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msRRASVendorAttributeEntry")
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

