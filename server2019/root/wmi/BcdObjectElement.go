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

// BcdObjectElement struct
type BcdObjectElement struct { 
	*BcdElement

	// This is the guid id of the object this element refers to.
	Id string
}

	func NewBcdObjectElementEx1(instance *cim.WmiInstance) (newInstance *BcdObjectElement, err error) {tmp, err := NewBcdElementEx1(instance)
		
	if err != nil { return }
	newInstance = &BcdObjectElement {
	BcdElement: tmp,
	}
	return
	}
	

	func NewBcdObjectElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BcdObjectElement, err error) {tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BcdObjectElement {
	BcdElement: tmp,
	}
	return
	}
	

// SetId sets the value of Id for the instance
func (instance *BcdObjectElement) SetPropertyId(value string) (err error) { 
    return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *BcdObjectElement) GetPropertyId()(value string, err error) { 
    retValue, err := instance.GetProperty("Id")
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

