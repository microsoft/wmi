// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS69B56437_FE3D_4523_BE69_32394088573F.Computer
//////////////////////////////////////////////
package computer
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSOP_UserPrivilegeRight struct
type RSOP_UserPrivilegeRight struct { 
	*RSOP_SecuritySettings

	// 
	AccountList []string

	// 
	UserRight string
}

	func NewRSOP_UserPrivilegeRightEx1(instance *cim.WmiInstance) (newInstance *RSOP_UserPrivilegeRight, err error) {tmp, err := NewRSOP_SecuritySettingsEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_UserPrivilegeRight {
	RSOP_SecuritySettings: tmp,
	}
	return
	}
	

	func NewRSOP_UserPrivilegeRightEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_UserPrivilegeRight, err error) {tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_UserPrivilegeRight {
	RSOP_SecuritySettings: tmp,
	}
	return
	}
	

// SetAccountList sets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRight) SetPropertyAccountList(value []string) (err error) { 
    return instance.SetProperty("AccountList", (value))
}

// GetAccountList gets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRight) GetPropertyAccountList()(value []string, err error) { 
    retValue, err := instance.GetProperty("AccountList")
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

// SetUserRight sets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRight) SetPropertyUserRight(value string) (err error) { 
    return instance.SetProperty("UserRight", (value))
}

// GetUserRight gets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRight) GetPropertyUserRight()(value string, err error) { 
    retValue, err := instance.GetProperty("UserRight")
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

