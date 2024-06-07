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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSWmi_Guid struct
type MSWmi_Guid struct { 
	*cim.WmiInstance

	// 
	Guid []uint8
}

	func NewMSWmi_GuidEx1(instance *cim.WmiInstance) (newInstance *MSWmi_Guid, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSWmi_Guid {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSWmi_GuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSWmi_Guid, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSWmi_Guid {
	WmiInstance: tmp,
	}
	return
	}
	

// SetGuid sets the value of Guid for the instance
func (instance *MSWmi_Guid) SetPropertyGuid(value []uint8) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSWmi_Guid) GetPropertyGuid()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Guid")
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

