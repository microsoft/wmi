// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_TerminalSetting struct
type Win32_TerminalSetting struct { 
	*CIM_Setting

	// 
	TerminalName string
}

	func NewWin32_TerminalSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TerminalSetting, err error) {tmp, err := NewCIM_SettingEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_TerminalSetting {
	CIM_Setting: tmp,
	}
	return
	}
	

	func NewWin32_TerminalSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_TerminalSetting, err error) {tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_TerminalSetting {
	CIM_Setting: tmp,
	}
	return
	}
	

// SetTerminalName sets the value of TerminalName for the instance
func (instance *Win32_TerminalSetting) SetPropertyTerminalName(value string) (err error) { 
    return instance.SetProperty("TerminalName", (value))
}

// GetTerminalName gets the value of TerminalName for the instance
func (instance *Win32_TerminalSetting) GetPropertyTerminalName()(value string, err error) { 
    retValue, err := instance.GetProperty("TerminalName")
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

