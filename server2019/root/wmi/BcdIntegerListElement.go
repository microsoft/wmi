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

// BcdIntegerListElement struct
type BcdIntegerListElement struct { 
	*BcdElement

	// This is the array of integer values of the element.
	Integers []uint64
}

	func NewBcdIntegerListElementEx1(instance *cim.WmiInstance) (newInstance *BcdIntegerListElement, err error) {tmp, err := NewBcdElementEx1(instance)
		
	if err != nil { return }
	newInstance = &BcdIntegerListElement {
	BcdElement: tmp,
	}
	return
	}
	

	func NewBcdIntegerListElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BcdIntegerListElement, err error) {tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BcdIntegerListElement {
	BcdElement: tmp,
	}
	return
	}
	

// SetIntegers sets the value of Integers for the instance
func (instance *BcdIntegerListElement) SetPropertyIntegers(value []uint64) (err error) { 
    return instance.SetProperty("Integers", (value))
}

// GetIntegers gets the value of Integers for the instance
func (instance *BcdIntegerListElement) GetPropertyIntegers()(value []uint64, err error) { 
    retValue, err := instance.GetProperty("Integers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint64); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint64(valuetmp))
    }

    return
}

