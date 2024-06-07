// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSCBC80CE2_C2E8_4CBC_8C8C_1C374E56CEFC.Computer
//////////////////////////////////////////////
package computer
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSOP_RegistryKeyBlocked struct
type RSOP_RegistryKeyBlocked struct { 
	*RSOP_SecuritySettingsBlocked

	// 
	Mode RegistryKeyBlocked_Mode

	// 
	Path string

	// 
	SDDLString string
}

	func NewRSOP_RegistryKeyBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_RegistryKeyBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_RegistryKeyBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

	func NewRSOP_RegistryKeyBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_RegistryKeyBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_RegistryKeyBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

// SetMode sets the value of Mode for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertyMode(value RegistryKeyBlocked_Mode) (err error) { 
    return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertyMode()(value RegistryKeyBlocked_Mode, err error) { 
    retValue, err := instance.GetProperty("Mode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RegistryKeyBlocked_Mode(valuetmp)

    return
}

// SetPath sets the value of Path for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertySDDLString(value string) (err error) { 
    return instance.SetProperty("SDDLString", (value))
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertySDDLString()(value string, err error) { 
    retValue, err := instance.GetProperty("SDDLString")
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

