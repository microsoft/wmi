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

// RSOP_RegistryValueBlocked struct
type RSOP_RegistryValueBlocked struct { 
	*RSOP_SecuritySettingsBlocked

	// 
	Data string

	// 
	Path string

	// 
	Type RegistryValueBlocked_Type
}

	func NewRSOP_RegistryValueBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_RegistryValueBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_RegistryValueBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

	func NewRSOP_RegistryValueBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_RegistryValueBlocked, err error) {tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_RegistryValueBlocked {
	RSOP_SecuritySettingsBlocked: tmp,
	}
	return
	}
	

// SetData sets the value of Data for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyData(value string) (err error) { 
    return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyData()(value string, err error) { 
    retValue, err := instance.GetProperty("Data")
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

// SetPath sets the value of Path for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyPath()(value string, err error) { 
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

// SetType sets the value of Type for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyType(value RegistryValueBlocked_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyType()(value RegistryValueBlocked_Type, err error) { 
    retValue, err := instance.GetProperty("Type")
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

    value = RegistryValueBlocked_Type(valuetmp)

    return
}

