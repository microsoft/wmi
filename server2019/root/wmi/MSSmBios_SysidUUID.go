// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSSmBios_SysidUUID struct
type MSSmBios_SysidUUID struct { 
	*MS_SmBios

	// 
	Uuid []uint8
}

	func NewMSSmBios_SysidUUIDEx1(instance *cim.WmiInstance) (newInstance *MSSmBios_SysidUUID, err error) {tmp, err := NewMS_SmBiosEx1(instance)
		
	if err != nil { return }
	newInstance = &MSSmBios_SysidUUID {
	MS_SmBios: tmp,
	}
	return
	}
	

	func NewMSSmBios_SysidUUIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSSmBios_SysidUUID, err error) {tmp, err := NewMS_SmBiosEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSSmBios_SysidUUID {
	MS_SmBios: tmp,
	}
	return
	}
	

// SetUuid sets the value of Uuid for the instance
func (instance *MSSmBios_SysidUUID) SetPropertyUuid(value []uint8) (err error) { 
    return instance.SetProperty("Uuid", (value))
}

// GetUuid gets the value of Uuid for the instance
func (instance *MSSmBios_SysidUUID) GetPropertyUuid()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Uuid")
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

