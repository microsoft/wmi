// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_ExpressionValue_sint8 struct
type MSFT_ExpressionValue_sint8 struct { 
	*MSFT_ExpressionValue

	// 
	value int8
}

	func NewMSFT_ExpressionValue_sint8Ex1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionValue_sint8, err error) {tmp, err := NewMSFT_ExpressionValueEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_ExpressionValue_sint8 {
	MSFT_ExpressionValue: tmp,
	}
	return
	}
	

	func NewMSFT_ExpressionValue_sint8Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_ExpressionValue_sint8, err error) {tmp, err := NewMSFT_ExpressionValueEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_ExpressionValue_sint8 {
	MSFT_ExpressionValue: tmp,
	}
	return
	}
	

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint8) SetPropertyvalue(value int8) (err error) { 
    return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint8) GetPropertyvalue()(value int8, err error) { 
    retValue, err := instance.GetProperty("value")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int8(valuetmp)

    return
}

