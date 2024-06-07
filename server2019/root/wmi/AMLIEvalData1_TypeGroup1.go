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

// AMLIEvalData1_TypeGroup1 struct
type AMLIEvalData1_TypeGroup1 struct { 
	*AMLIEvalData1

	// 
	DataString string
}

	func NewAMLIEvalData1_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *AMLIEvalData1_TypeGroup1, err error) {tmp, err := NewAMLIEvalData1Ex1(instance)
		
	if err != nil { return }
	newInstance = &AMLIEvalData1_TypeGroup1 {
	AMLIEvalData1: tmp,
	}
	return
	}
	

	func NewAMLIEvalData1_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AMLIEvalData1_TypeGroup1, err error) {tmp, err := NewAMLIEvalData1Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AMLIEvalData1_TypeGroup1 {
	AMLIEvalData1: tmp,
	}
	return
	}
	

// SetDataString sets the value of DataString for the instance
func (instance *AMLIEvalData1_TypeGroup1) SetPropertyDataString(value string) (err error) { 
    return instance.SetProperty("DataString", (value))
}

// GetDataString gets the value of DataString for the instance
func (instance *AMLIEvalData1_TypeGroup1) GetPropertyDataString()(value string, err error) { 
    retValue, err := instance.GetProperty("DataString")
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

