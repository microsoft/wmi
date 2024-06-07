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

// Win32_PowerSettingElementSettingDataIndex struct
type Win32_PowerSettingElementSettingDataIndex struct { 
	*CIM_ElementSettingData

	// 
	IsACSetting uint16
}

	func NewWin32_PowerSettingElementSettingDataIndexEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingElementSettingDataIndex, err error) {tmp, err := NewCIM_ElementSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingElementSettingDataIndex {
	CIM_ElementSettingData: tmp,
	}
	return
	}
	

	func NewWin32_PowerSettingElementSettingDataIndexEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PowerSettingElementSettingDataIndex, err error) {tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingElementSettingDataIndex {
	CIM_ElementSettingData: tmp,
	}
	return
	}
	

// SetIsACSetting sets the value of IsACSetting for the instance
func (instance *Win32_PowerSettingElementSettingDataIndex) SetPropertyIsACSetting(value uint16) (err error) { 
    return instance.SetProperty("IsACSetting", (value))
}

// GetIsACSetting gets the value of IsACSetting for the instance
func (instance *Win32_PowerSettingElementSettingDataIndex) GetPropertyIsACSetting()(value uint16, err error) { 
    retValue, err := instance.GetProperty("IsACSetting")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

