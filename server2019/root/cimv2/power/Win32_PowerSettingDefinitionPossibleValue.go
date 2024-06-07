// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PowerSettingDefinitionPossibleValue struct
type Win32_PowerSettingDefinitionPossibleValue struct { 
	*CIM_SettingData

	// 
	BinaryValue []uint8

	// 
	SettingIndex uint32

	// 
	StringValue string

	// 
	UInt32Value uint32

	// 
	UInt64Value uint64
}

	func NewWin32_PowerSettingDefinitionPossibleValueEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDefinitionPossibleValue, err error) {tmp, err := NewCIM_SettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDefinitionPossibleValue {
	CIM_SettingData: tmp,
	}
	return
	}
	

	func NewWin32_PowerSettingDefinitionPossibleValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PowerSettingDefinitionPossibleValue, err error) {tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDefinitionPossibleValue {
	CIM_SettingData: tmp,
	}
	return
	}
	

// SetBinaryValue sets the value of BinaryValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyBinaryValue(value []uint8) (err error) { 
    return instance.SetProperty("BinaryValue", (value))
}

// GetBinaryValue gets the value of BinaryValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyBinaryValue()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("BinaryValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetSettingIndex sets the value of SettingIndex for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertySettingIndex(value uint32) (err error) { 
    return instance.SetProperty("SettingIndex", (value))
}

// GetSettingIndex gets the value of SettingIndex for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertySettingIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SettingIndex")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetStringValue sets the value of StringValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyStringValue(value string) (err error) { 
    return instance.SetProperty("StringValue", (value))
}

// GetStringValue gets the value of StringValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyStringValue()(value string, err error) { 
    retValue, err := instance.GetProperty("StringValue")
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

// SetUInt32Value sets the value of UInt32Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyUInt32Value(value uint32) (err error) { 
    return instance.SetProperty("UInt32Value", (value))
}

// GetUInt32Value gets the value of UInt32Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyUInt32Value()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UInt32Value")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetUInt64Value sets the value of UInt64Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyUInt64Value(value uint64) (err error) { 
    return instance.SetProperty("UInt64Value", (value))
}

// GetUInt64Value gets the value of UInt64Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyUInt64Value()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UInt64Value")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

