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

// QDat struct
type QDat struct { 
	*cim.WmiInstance

	// qdata
	Bytes []uint8
}

	func NewQDatEx1(instance *cim.WmiInstance) (newInstance *QDat, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &QDat {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewQDatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *QDat, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &QDat {
	WmiInstance: tmp,
	}
	return
	}
	

// SetBytes sets the value of Bytes for the instance
func (instance *QDat) SetPropertyBytes(value []uint8) (err error) { 
    return instance.SetProperty("Bytes", (value))
}

// GetBytes gets the value of Bytes for the instance
func (instance *QDat) GetPropertyBytes()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Bytes")
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

