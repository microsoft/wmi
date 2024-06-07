// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS39FB71E5_E3CB_4676_B31D_9E3B74C5A6D5.Computer
//////////////////////////////////////////////
package computer
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSOP_SecuritySettingStringBlocked struct
type RSOP_SecuritySettingStringBlocked struct { 
	*RSOP_SecuritySettingsBlocked

	// 
	KeyName string

	// 
	Setting string
}

	func NewRSOP_SecuritySettingStringBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_SecuritySettingStringBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_SecuritySettingStringBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

	func NewRSOP_SecuritySettingStringBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_SecuritySettingStringBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_SecuritySettingStringBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingStringBlocked) SetPropertyKeyName(value string) (err error) { 
    return instance.SetProperty("KeyName", (value))
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingStringBlocked) GetPropertyKeyName()(value string, err error) { 
    retValue, err := instance.GetProperty("KeyName")
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

// SetSetting sets the value of Setting for the instance
func (instance *RSOP_SecuritySettingStringBlocked) SetPropertySetting(value string) (err error) { 
    return instance.SetProperty("Setting", (value))
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecuritySettingStringBlocked) GetPropertySetting()(value string, err error) { 
    retValue, err := instance.GetProperty("Setting")
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

