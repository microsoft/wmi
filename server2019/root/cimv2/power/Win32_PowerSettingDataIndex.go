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

// Win32_PowerSettingDataIndex struct
type Win32_PowerSettingDataIndex struct { 
	*CIM_SettingData

	// 
	SettingIndexValue uint32
}

	func NewWin32_PowerSettingDataIndexEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDataIndex, err error) {tmp, err := NewCIM_SettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDataIndex {
	CIM_SettingData: tmp,
	}
	return
	}
	

	func NewWin32_PowerSettingDataIndexEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PowerSettingDataIndex, err error) {tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDataIndex {
	CIM_SettingData: tmp,
	}
	return
	}
	

// SetSettingIndexValue sets the value of SettingIndexValue for the instance
func (instance *Win32_PowerSettingDataIndex) SetPropertySettingIndexValue(value uint32) (err error) { 
    return instance.SetProperty("SettingIndexValue", (value))
}

// GetSettingIndexValue gets the value of SettingIndexValue for the instance
func (instance *Win32_PowerSettingDataIndex) GetPropertySettingIndexValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SettingIndexValue")
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

