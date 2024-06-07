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

// ads_msmq_group struct
type ads_msmq_group struct { 
	*ds_top

	// 
	DS_member []string
}

	func Newads_msmq_groupEx1(instance *cim.WmiInstance) (newInstance *ads_msmq_group, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msmq_group {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msmq_groupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msmq_group, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msmq_group {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_member sets the value of DS_member for the instance
func (instance *ads_msmq_group) SetPropertyDS_member(value []string) (err error) { 
    return instance.SetProperty("DS_member", (value))
}

// GetDS_member gets the value of DS_member for the instance
func (instance *ads_msmq_group) GetPropertyDS_member()(value []string, err error) { 
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

