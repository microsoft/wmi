// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_InstModification struct
type CIM_InstModification struct { 
	*CIM_InstIndication

	// A copy of the 'previous' instance whose change generated the Indication. PreviousInstance contains 'older' values of an instance's properties (as compared to SourceInstance), selected by the IndicationFilter's Query.
	PreviousInstance interface{}
}

	func NewCIM_InstModificationEx1(instance *cim.WmiInstance) (newInstance *CIM_InstModification, err error) {tmp, err := NewCIM_InstIndicationEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_InstModification {
	CIM_InstIndication: tmp,
	}
	return
	}
	

	func NewCIM_InstModificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_InstModification, err error) {tmp, err := NewCIM_InstIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_InstModification {
	CIM_InstIndication: tmp,
	}
	return
	}
	

// SetPreviousInstance sets the value of PreviousInstance for the instance
func (instance *CIM_InstModification) SetPropertyPreviousInstance(value interface{}) (err error) { 
    return instance.SetProperty("PreviousInstance", (value))
}

// GetPreviousInstance gets the value of PreviousInstance for the instance
func (instance *CIM_InstModification) GetPropertyPreviousInstance()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("PreviousInstance")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(interface{}); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = interface{}(valuetmp)

    return
}

